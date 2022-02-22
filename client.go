package nlpcloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	// ErrNilClient is an error returned when the sub-Client is nil but
	// need to get used.
	ErrNilClient = errors.New("client is nil")

	// ErrBadRequest is an error returned on status code 400.
	ErrBadRequest = errors.New("your request is invalid")

	// ErrUnauthorized is an error returned on status code 401.
	ErrUnauthorized = errors.New("your API token is wrong")

	// ErrPaymentRequired is an error returned on status code 402.
	ErrPaymentRequired = errors.New("you are trying to access a ressource that is only accessible after payment")

	// ErrForbidden is an error returned on status code 403.
	ErrForbidden = errors.New("you don't have the sufficient rights to access the resource. Please make sure you subscribed to the proper plan that grants you access to this resource")

	// ErrNotFound is an error returned on status code 404.
	ErrNotFound = errors.New("the specified resource could not be found")

	// ErrMethodNotAllowed is an error returned on status code 405.
	ErrMethodNotAllowed = errors.New("you tried to access a resource with an invalid method")

	// ErrNotAcceptable is an error returned on status code 406.
	ErrNotAcceptable = errors.New("you requested a format that isn't json")

	// ErrRequestEntityTooLarge is an error returned on status code 413.
	ErrRequestEntityTooLarge = errors.New("the piece of text that you are sending is too large. Please see the maximum sizes in the documentation")

	// ErrUnprocessableEntity is an error returned on status code 422.
	ErrUnprocessableEntity = errors.New("your request is not properly formatted. Happens for example if your JSON payload is not correctly formatted, or if you omit the \"Content-Type: application/json\" header")

	// ErrTooManyRequests is an error returned on status code 429.
	ErrTooManyRequests = errors.New("you made too many requests in a short while, please slow down")

	// ErrInternalServerError is an error returned on status code 500.
	ErrInternalServerError = errors.New("sorry, we had a problem with our server. Please try again later")

	// ErrBadGateway is an error returned on status code 502.
	ErrBadGateway = errors.New("sorry, our reverse proxy was not able to contact the model you're requesting. Please try again later")

	// ErrServiceUnavailable is an error returned on status code 503.
	ErrServiceUnavailable = errors.New("sorry, the model you are requesting had a temporary issue. Please try again later. The error is returned together with a \"Retry-After\" header, mentioning the number of seconds you should wait before trying again")

	// ErrGatewayTimeout is an error returned on status code 504.
	ErrGatewayTimeout = errors.New("sorry, the model you are requesting is temporarily overloaded. Please try again later")
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
		return ErrNilClient
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
		switch resp.StatusCode {
		case 400:
			return ErrBadRequest
		case 401:
			return ErrUnauthorized
		case 402:
			return ErrPaymentRequired
		case 403:
			return ErrForbidden
		case 404:
			return ErrNotFound
		case 405:
			return ErrMethodNotAllowed
		case 406:
			return ErrNotAcceptable
		case 413:
			return ErrRequestEntityTooLarge
		case 422:
			return ErrUnprocessableEntity
		case 429:
			return ErrTooManyRequests
		case 500:
			return ErrInternalServerError
		case 502:
			return ErrBadGateway
		case 503:
			return ErrServiceUnavailable
		case 504:
			return ErrGatewayTimeout
		default:
			return &ErrUnexpectedStatus{
				Body:       body,
				StatusCode: resp.StatusCode,
			}
		}
	}

	// Unmarshal response
	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}

	return nil
}
