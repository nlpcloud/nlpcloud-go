package nlpcloud_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/nlpcloud/nlpcloud-go"
	"github.com/stretchr/testify/assert"
)

var (
	errStrTypeOf = reflect.TypeOf(errors.New(""))
	errFake      = errors.New("this is a fake error")
)

func checkErr(expErr, err error, assert *assert.Assertions) {
	// Check err type
	typeErr := reflect.TypeOf(err)
	typeExpErr := reflect.TypeOf(expErr)
	if !assert.Equal(typeExpErr, typeErr) {
		return
	}

	// Check Error content is not empty
	if err != nil && err.Error() == "" {
		assert.FailNow("Error should not have an empty content")
	}

	// Check if the error is generated using errors.New
	if typeErr == errStrTypeOf {
		if err.Error() != expErr.Error() {
			assert.Fail("Error message differs: got \"%s\" instead of \"%s\".", err.Error(), expErr.Error())
		}
		return
	}

	switch err.(type) {
	case *url.Error:
		castedErr := err.(*url.Error)
		castedExpErr := expErr.(*url.Error)

		assert.Equal(castedExpErr.Op, castedErr.Op)
		assert.Equal(castedExpErr.URL, castedErr.URL)

	case *json.SyntaxError:
		castedErr := err.(*json.SyntaxError)
		castedExpErr := expErr.(*json.SyntaxError)

		assert.Equal(castedExpErr.Offset, castedErr.Offset)

	case *nlpcloud.ErrUnexpectedStatus:
		castedErr := err.(*nlpcloud.ErrUnexpectedStatus)
		castedExpErr := expErr.(*nlpcloud.ErrUnexpectedStatus)

		assert.Equal(castedExpErr.Body, castedErr.Body)
		assert.Equal(castedExpErr.StatusCode, castedErr.StatusCode)

	case nil:
		return

	default:
		assert.Fail("\033[31mcheckErr Unsupported type: %s\033[0m\n", typeErr)
	}
}

// FakeHTTPClient is an implementation of HTTPClient that
// does nothing expect returning what you said it to.
type fakeHTTPClient struct {
	Response *http.Response
	Err      error
}

func (f fakeHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return f.Response, f.Err
}

var _ nlpcloud.HTTPClient = (*fakeHTTPClient)(nil)

func newFakeHTTPClient(body string, statusCode int, err error) nlpcloud.HTTPClient {
	return &fakeHTTPClient{
		Response: &http.Response{
			StatusCode: statusCode,
			Body:       newFakeReadCloser(body),
		},
		Err: err,
	}
}

// FakeReadCloser mocks an io.ReadCloser.
type fakeReadCloser struct {
	data      []byte
	readIndex int64
}

func (f *fakeReadCloser) Read(p []byte) (n int, err error) {
	if f.readIndex >= int64(len(f.data)) {
		err = io.EOF
		return
	}

	n = copy(p, f.data[f.readIndex:])
	f.readIndex += int64(n)
	return
}

func (f *fakeReadCloser) Close() error {
	return nil
}

var _ io.ReadCloser = (*fakeReadCloser)(nil)

func newFakeReadCloser(str string) *fakeReadCloser {
	return &fakeReadCloser{
		data: []byte(str),
	}
}
