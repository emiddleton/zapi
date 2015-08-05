package zapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	//"log"
	"encoding/json"
	"net/http"
	"net/url"
	//"strings"
	"time"
)

type Client struct {
	Url        string       `json:"url"`
	Username   string       `json:"username"`
	Password   string       `json:"password"`
	Token      string       `json:"token"`
	httpClient *http.Client `json:"-"`
}

func NewClientFromFile(path string) (client Client, err error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return client, err
	}

	if err := json.Unmarshal(data, &client); err != nil {
		return client, err
	}

	client.httpClient = &http.Client{}

	return client, err
}

func NewPasswordClient(url, username, password string) Client {
	return Client{
		Username:   username,
		Password:   password,
		httpClient: &http.Client{Transport: &http.Transport{}},
	}
}
func NewTokenClient(url, username, token string) Client {
	return Client{
		Username:   username,
		Token:      token,
		httpClient: &http.Client{Transport: &http.Transport{}},
	}
}

type Filter struct {
	Key   string
	Value string
}

type Filters []Filter

func (fs *Filters) toParams() (vals url.Values) {
	for _, filter := range fs {
		vals.Add(filter.Key, filter.Value)
	}
	return vals
}

type Date time.Time

func (d *Date) MarshalJSON() ([]byte, error) {
	t := time.Time(*d).Format(fmt.Sprintf("\"%s\"", time.RFC3339))
	return []byte(t), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(fmt.Sprintf("\"%s\"", time.RFC3339), string(b))
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (c *Client) Do(req *http.Request) ([]byte, error) {
	req.Header.Add("Content-Type", "application/json")
	if c.Token == "" {
		req.SetBasicAuth(c.Username, c.Password)
	} else {
		req.SetBasicAuth(fmt.Sprintf("%s/token", c.Username), c.Token)
	}
	fmt.Printf("%#v\n", req.URL)
	fmt.Printf("%#v\n", req)
	resp, err := c.httpClient.Do(req)
	fmt.Printf("%#v\n", resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (client *Client) Get(path string, params *url.Values) ([]byte, error) {
	urlPath := fmt.Sprintf("%s%s", client.Url, path)
	urlRaw, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		urlRaw.RawQuery = params.Encode()
	}
	fmt.Printf("%s\n", urlRaw.String())
	req, err := http.NewRequest("GET", urlRaw.String(), nil)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func (client *Client) Put(path string, params *url.Values, requestBody []byte) ([]byte, error) {
	urlPath := fmt.Sprintf("%s%s", client.Url, path)
	urlRaw, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		urlRaw.RawQuery = params.Encode()
	}
	fmt.Printf("%s\n", urlRaw.String())
	req, err := http.NewRequest("PUT", urlRaw.String(), bytes.NewBufferString(string(requestBody)))
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func (client *Client) Post(path string, params *url.Values, requestBody []byte) ([]byte, error) {
	urlPath := fmt.Sprintf("%s%s", client.Url, path)
	urlRaw, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	if params != nil {
		urlRaw.RawQuery = params.Encode()
	}
	fmt.Printf("%s\n", urlRaw.String())
	req, err := http.NewRequest("POST", urlRaw.String(), bytes.NewBufferString(string(requestBody)))
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
