package nlpcloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

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
func NewClient(client HTTPClient, model, token string, gpu bool) *Client {
	rootUrl := "https://api.nlpcloud.io/v1/"
	if gpu {
		rootUrl += "gpu/"
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
		return ErrNilClient
	}

	// Marshal the request body if needed (in most cases, for POST)
	var buf io.Reader = nil
	if params != nil {
		j, _ := json.Marshal(params)
		buf = bytes.NewBuffer(j)
	}

	// Create the request backbone
	req, _ := http.NewRequest(method, c.rootURL+"/"+endpoint, buf)
	req.Header.Set("Authorization", "Token "+c.token)
	req.Header.Set("User-Agent", "nlpcloud-go-client")

	// Issue the request
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body) // assume the client builds Body properly

	// Check for request failure
	if resp.StatusCode != http.StatusOK {
		return &ErrUnexpectedStatus{
			Body:       body,
			StatusCode: resp.StatusCode,
		}
	}

	// Unmarshal response
	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}

	return nil
}

var (
	// ErrNilClient is an error returned when the sub-Client is nil but
	// need to get used.
	ErrNilClient = errors.New("client is nil")
)

// ErrUnexpectedStatus is an error type returned when the HTTP request
// returned with an unexpected status code, meaning something failed.
type ErrUnexpectedStatus struct {
	Body       []byte
	StatusCode int
}

func (e ErrUnexpectedStatus) Error() string {
	return fmt.Sprintf("unexpected status %d with body %v", e.StatusCode, e.Body)
}

var _ error = (*ErrUnexpectedStatus)(nil)
