package zapi

import (
	"bytes"
	"io/ioutil"
	//	"log"
	"net/http"
	//	"strings"
)

type Client interface {
	Do(http.Request) (http.Response, error)
}

type PasswordClient struct {
	Email    string
	Password string
	Testing  *testing.T   `json:"-"`
	client   *http.Client `json:"-"`
}

func NewPasswordClient(username, password string) (client Client) {
	return PasswordClient{
		Email:    username,
		Password: password,
		client:   &http.Client{Transport: &http.Transport{}},
	}
}

func (c *PasswordClient) Do(req http.Request) (http.Response, error) {
	req.SetBasicAuth(c.Username, c.Password)
	return c.client.Do(req)
}

func (client *Client) Request() {
	req, err := http.NewRequest(method, url, bytes.NewBufferString(params))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
}
