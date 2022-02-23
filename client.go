package nlpcloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// HTTPError is an error type returned when the HTTP request
// is failing.
type HTTPError struct {
	Detail string
	Status int
}

func (h HTTPError) Error() string {
	return fmt.Sprintf("http error with status %d: %v", h.Status, h.Detail)
}

func (h HTTPError) GetDetail() string {
	return h.Detail
}

func (h HTTPError) GetStatusCode() int {
	return h.Status
}

var _ error = (*HTTPError)(nil)

// HTTPClient defines what a HTTP client have to implement in order to get
// used by the Client.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Makes sure the *http.Client works with the HTTPClient.
var _ HTTPClient = (*http.Client)(nil)

// Client holds the necessary information to connect to API.
type Client struct {
	client  HTTPClient
	rootURL string
	token   string
}

// NewClient initializes a new Client.
func NewClient(client HTTPClient, model, token string, gpu bool, lang string) *Client {
	rootUrl := "https://api.nlpcloud.io/v1/"
	if gpu {
		rootUrl += "gpu/"
	}
	if lang != "" {
		rootUrl += lang + "/"
	}
	rootUrl += model

	return &Client{
		client:  client,
		rootURL: rootUrl,
		token:   token,
	}
}

func (c *Client) issueRequest(method, endpoint string, params, dst interface{}) error {
	// Check the client is properly defined
	if c.client == nil {
		return errors.New("client is nil")
	}

	// Marshal the request body if needed (in most cases, for POST)
	var buf io.Reader = nil
	if params != nil {
		j, err := json.Marshal(params)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(j)
	}

	// Create the request backbone
	req, err := http.NewRequest(method, c.rootURL+"/"+endpoint, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Token "+c.token)
	req.Header.Set("User-Agent", "nlpcloud-go-client")

	// Issue the request
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for request failure
	if resp.StatusCode != http.StatusOK {
		return &HTTPError{
			Detail: string(body),
			Status: resp.StatusCode,
		}
	}

	// Unmarshal response
	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}

	return nil
}
