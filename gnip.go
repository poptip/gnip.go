package gnip

import (
	bytelib "bytes"
	"encoding/json"
	"fmt"
  "io/ioutil"
	"net/http"
)

const (
	streamBaseUrl = "https://stream.gnip.com:443/accounts"
	streamSuffix  = "publishers/twitter/streams/track/Production.json"
	rulesBaseUrl  = "https://api.gnip.com:443/accounts"
	rulesSuffix   = "publishers/twitter/streams/track/Production/rules.json"
	bufferSize    = 33554432
)

type Rules struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Value string `json:"value"`
}

type Client struct {
	username    string
	password    string
	HttpClient  http.Client
  streamUrl   string
  rulesUrl    string
}

func NewClient(un, pw, account string) *Client {
  c := &Client{}
  c.username = un;
  c.password = pw;
  c.streamUrl = fmt.Sprintf("%s/%s/%s",
    streamBaseUrl, account, streamSuffix)
  c.rulesUrl = fmt.Sprintf("%s/%s/%s",
    rulesBaseUrl, account, rulesSuffix)
	return c
}

func (c *Client) Connect() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.streamUrl, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.password)
	return c.HttpClient.Do(req)
}

func (c *Client) GetAllActiveRules() ([]Rule, error) {
	req, err := http.NewRequest("GET", c.rulesUrl, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.password)
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return nil, err
  }

  var rules Rules
  if err = json.Unmarshal(body, &rules); err != nil {
    return nil, err
  }
	return rules.Rules, nil
}

func (c *Client) AddRules(rules []Rule) error {
	payload := Rules{Rules: rules}
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.rulesUrl, bytelib.NewBuffer(bytes))
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.username, c.password)
	_, err = c.HttpClient.Do(req)
  return err
}

func (c *Client) RemoveRules(rules []Rule) error {
	payload := &Rules{Rules: rules}
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", c.rulesUrl, bytelib.NewBuffer(bytes))
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.username, c.password)
	_, err = c.HttpClient.Do(req)
  return err
}
