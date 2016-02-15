package main

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

type Connection struct {
	SessionId string
	Client    *http.Client
	Retries   int
}

type AuthData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginData struct {
	AuthData AuthData `json:"user_auth"`
	AppName  string   `json:"application_name"`
}

func CreateConnection(addr string, user string, pwd string) (*Connection, error) {
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
		log.Fatalf("illegal credential format: %s", err)
	}
	restData := string(restDataBytes[:])

	url, err := url.Parse(addr)
	if err != nil {
		log.Fatalf("Illegal URL: %s", err)
	}
	url.Scheme = "http"
	q := url.Query()
	q.Set("method", "login")
	q.Set("input_type", "json")
	q.Set("response_type", "json")
	q.Set("rest_data", restData)
	url.RawQuery = q.Encode()

	tr := &http.Transport{}
	cl := &http.Client{Transport: tr}
	resp, err := cl.Get(url.String())
	if err != nil {
		err = errors.New("Can't connect to SuiteCRM: " + err.Error())
		return &Connection{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("Can't read response by SuiteCRM: " + err.Error())
		return &Connection{}, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		err = errors.New("Can't read response by SuiteCRM: " + err.Error())
		return &Connection{}, err
	}

	sid, ok := data["id"].(string)
	if !ok {
		err = errors.New("Can't read session id: " + err.Error())
		return &Connection{}, err
	}
	return &Connection{SessionId: sid, Client: cl}, nil
}

// Use reconnect() if a connection needs to be reastablished.
// Each unsuccessful attempt will increment the number of retries stored in the
// 'Connection'-structure.
// A successful connection will reset the counter.
func (c *Connection) Reconnect() {
	// TODO(danck):
	//	con, err := c.CreateConnection(c.)
	//	if err != nil {
	//		c.Retries = c.Retries + 1
	//		return c, err
	//	}
	//	return con, nil
}

func (c *Connection) Send() {}
