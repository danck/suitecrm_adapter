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
}

type AuthData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginStr struct {
	AuthData AuthData `json:"user_auth"`
	AppName  string   `json:"application_name"`
}

func CreateConnection(addr string, user string, pwd string) (Connection, error) {
	h := md5.New()
	io.WriteString(h, pwd)
	pwdHashStr := fmt.Sprintf("%x", h.Sum(nil))
	restDataBytes, err := json.Marshal(LoginStr{
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
		log.Printf("Failed to set up connection to SuiteCRM: %s", err)
		return Connection{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error while getting session ID: %s", err)
		return Connection{}, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Error while getting session ID: %s", string(body[:]))
		return Connection{}, err
	}

	sid, ok := data["id"].(string)
	if !ok {
		log.Printf("Error while getting session ID: %s", err)
		return Connection{}, err
	}
	return Connection{SessionId: sid, Client: cl}, nil
}

func send() {}
