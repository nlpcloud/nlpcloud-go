package spacycloud

import (
	"fmt"
	"net/http"
)

const (
	baseURL    = "https://api.spacycloud.io"
	apiVersion = "v1"
)

type Client struct {
	rootURL string
	token   string
	client  *http.Client
}

func NewClient(model, token string) Client {
	return Client{
		rootURL: fmt.Sprintf("%v/%v/%v", baseURL, apiVersion, model),
		token:   token,
		client:  &http.Client{},
	}
}

func (c *Client) Entities(text string) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/entities", c.rootURL), nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Token %v", c.token))
	if err != nil {
		return
	}
	resp, err := c.client.Do(req)

}
