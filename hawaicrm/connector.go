package hawaicrm

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	Addr      string
	SessionID string
	Client    *http.Client
)

// Transfer type for SuiteCRM get/set requests
// A map doesn't work here because order seems to be important
type RestData struct {
	Session       string         `json:"session"`
	ModuleName    string         `json:"module_name"`
	NameValueList []KeyValuePair `json:"name_value_list"`
}

// Transfer type for SuiteCRM relationship setting
type RestRelationData struct {
	Session       string         `json:"session"`
	ModuleName    string         `json:"module_name"`
	ModuleID      string         `json:"module_id"`
	LinkFieldName string         `json:"link_field_name"`
	RelatedIDs    []string       `json:"related_ids"`
	NameValueList []KeyValuePair `json:"name_value_list"`
}

type KeyValuePair struct {
	Key   string `json:"name"`
	Value string `json:"value"`
}

// Transfer type for SuiteCRM session requests
type LoginData struct {
	AuthData AuthData `json:"user_auth"`
	AppName  string   `json:"application_name"`
}

type AuthData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// Connect creates a connection to SuiteCRM
func Connect(addr string, user string, pwd string) error {
	Addr = addr

	// Set up login data
	h := md5.New()
	io.WriteString(h, pwd)
	pwdHashStr := fmt.Sprintf("%x", h.Sum(nil))
	restDataBytes, err := json.Marshal(LoginData{
		AppName: "SuiteCRM Adapter",
		AuthData: AuthData{
			UserName: user,
			Password: pwdHashStr,
		},
	})
	if err != nil {
		return errors.New("illegal credential format: " + err.Error())
	}
	restData := string(restDataBytes[:])

	// Set up connection request
	URL, err := url.Parse(addr)
	if err != nil {
		return errors.New("Illegal URL: " + err.Error())
	}
	URL.Scheme = "http"
	q := URL.Query()
	q.Set("method", "login")
	q.Set("input_type", "json")
	q.Set("response_type", "json")
	q.Set("rest_data", restData)
	URL.RawQuery = q.Encode()

	// Connect to SuiteCRM instance
	tr := &http.Transport{}
	Client = &http.Client{Transport: tr}
	resp, err := Client.Get(URL.String())
	if err != nil {
		err = errors.New("Can't connect to SuiteCRM: " + err.Error())
		return err
	}
	defer resp.Body.Close()
	log.Printf("Connection Query: %s", URL.String())

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("Can't read response by SuiteCRM: " + err.Error())
		return err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		err = errors.New("Can't read response by SuiteCRM: " + err.Error())
		return err
	}

	sid, ok := data["id"].(string)
	if !ok {
		err = errors.New("Can't read session id: " + err.Error())
		return err
	}
	SessionID = sid
	log.Printf("Connection established. ID %s", SessionID)
	return nil
}

// crmSetEntry takes a module and a list of key-value-pairs to create an entry
func crmSetEntry(module string, nameValueList []KeyValuePair) ([]byte, error) {
	r := RestData{SessionID, module, nameValueList}
	restDataJSON, err := json.Marshal(r)
	if err != nil {
		err := errors.New("illegal JSON format: " + err.Error())
		return nil, err
	}
	restDataJSONString := string(restDataJSON[:])

	return send(module, "set_entry", restDataJSONString)
}

// crmGetEntry takes a module and a list of key-value-pairs to retrieve an entry
func crmGetEntry(module string, nameValueList []KeyValuePair) ([]byte, error) {
	r := RestData{SessionID, module, nameValueList}
	restDataJSON, err := json.Marshal(r)
	if err != nil {
		err := errors.New("illegal JSON format: " + err.Error())
		return nil, err
	}
	restDataJSONString := string(restDataJSON[:])

	return send(module, "get_entry", restDataJSONString)
}

// crmSetRelationship takes two modules and IDs to set a relationship between
// the two or more entities
func crmSetRelationship(
	module string,
	moduleID string,
	linkFieldName string,
	relatedIDs []string,
	nameValueList []KeyValuePair) ([]byte, error) {

	r := RestRelationData{SessionID, module, moduleID, linkFieldName, relatedIDs, nameValueList}
	restDataJSON, err := json.Marshal(r)
	if err != nil {
		err := errors.New("illegal JSON format: " + err.Error())
		return nil, err
	}
	restDataJSONString := string(restDataJSON[:])

	return send(module, "set_relationship", restDataJSONString)
}

func send(module string, method string, restData string) ([]byte, error) {
	if Client == nil {
		err := errors.New("Connection not available: HTTP Client is nil")
		return nil, err
	}

	URL, err := url.Parse(Addr)
	if err != nil {
		err := errors.New("Illegal URL: " + err.Error())
		return nil, err
	}

	URL.Scheme = "http"

	// Assemble query string
	q := URL.Query()
	q.Set("method", method)
	q.Set("input_type", "json")
	q.Set("response_type", "json")
	q.Set("rest_data", restData)
	URL.RawQuery = q.Encode()

	// Create and send request
	req, err := http.NewRequest("POST", URL.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("Sent: %s", URL.String())
	log.Printf("Received: %s", body)
	return body, nil
}
