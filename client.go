package nlpcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
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

// ClientParams wraps all the parameters for the client initialization.
type ClientParams struct {
	Model string
	Token string
	GPU   bool
	Lang  string
	Async bool
}

// NewClient initializes a new Client.
func NewClient(client HTTPClient, clientParams ClientParams) *Client {
	rootUrl := "https://api.nlpcloud.io/v1/"
	if clientParams.Lang == "en" {
		clientParams.Lang = ""
	}
	if clientParams.Lang == "eng_Latn" {
		clientParams.Lang = ""
	}
	if clientParams.GPU {
		rootUrl += "gpu/"
	}
	if clientParams.Async {
		rootUrl += "async/"
	}
	if clientParams.Lang != "" {
		rootUrl += clientParams.Lang + "/"
	}
	rootUrl += clientParams.Model

	return &Client{
		client:  client,
		rootURL: rootUrl,
		token:   clientParams.Token,
	}
}

func (c *Client) issueRequest(method, endpoint string, params, dst interface{}, opts ...Option) error {
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

	// Apply the options
	options := &options{
		Ctx: context.Background(),
	}
	for _, opt := range opts {
		opt.apply(options)
	}

	// Create the request backbone
	req, err := http.NewRequestWithContext(options.Ctx, method, c.rootURL+"/"+endpoint, buf)
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
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return &HTTPError{
			Detail: string(body),
			Status: resp.StatusCode,
		}
	}

	// Unmarshal response
	if err = json.Unmarshal(body, dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) issueStreamingRequest(method, endpoint string, params interface{}, opts ...Option) (io.ReadCloser, error) {
	// Check the client is properly defined
	if c.client == nil {
		return nil, errors.New("client is nil")
	}

	// Marshal the request body if needed (in most cases, for POST)
	var buf io.Reader = nil
	if params != nil {
		j, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		streamingPayload := strings.TrimSuffix(string(j), "}") + `,"stream":true}`
		buf = bytes.NewBuffer([]byte(streamingPayload))
	}

	// Apply the options
	options := &options{
		Ctx: context.Background(),
	}
	for _, opt := range opts {
		opt.apply(options)
	}

	// Create the request backbone
	req, err := http.NewRequestWithContext(options.Ctx, method, c.rootURL+"/"+endpoint, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Token "+c.token)
	req.Header.Set("User-Agent", "nlpcloud-go-client")

	// Issue the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check for request failure
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return nil, &HTTPError{
			Status: resp.StatusCode,
		}
	}

	return resp.Body, nil
}

type Option interface {
	apply(*options)
}

type options struct {
	Ctx context.Context
}

type ctxOpt struct {
	ctx context.Context
}

func (opt ctxOpt) apply(opts *options) {
	opts.Ctx = opt.ctx
}

// WithContext returns an Option that defines the context.Context
// to use with issuing a request.
// Default is context.Background.
func WithContext(ctx context.Context) Option {
	return &ctxOpt{
		ctx: ctx,
	}
}
