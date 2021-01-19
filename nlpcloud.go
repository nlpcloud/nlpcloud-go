package nlpcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL    = "https://api.nlpcloud.io"
	apiVersion = "v1"
)

// Client holds the necessary information to connect to API.
type Client struct {
	rootURL string
	token   string
	client  *http.Client
}

// Entity holds an NER entity returned by the API.
type Entity struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Type  string `json:"type"`
	Text  string `json:"text"`
}

// Entities holds a list of NER entities returned by the API.
type Entities struct {
	Entities []Entity `json:"entities"`
}

// Word holds POS tag for a word.
type Word struct {
	Text string `json:"text"`
	Tag  string `json:"tag"`
}

// Arc holds information related to POS direction.
type Arc struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Label string `json:"label"`
	Text  string `json:"text"`
	Dir   string `json:"dir"`
}

// Dependencies holds a list of POS dependencies returned by the API.
type Dependencies struct {
	Words []Word `json:"words"`
	Arcs  []Arc  `json:"arcs"`
}

// SentenceDependency holds a POS dependency for one sentence
// returned by the API.
type SentenceDependency struct {
	Sentence     string `json:"sentence"`
	Dependencies `json:"dependencies"`
}

// SentenceDependencies holds a list of POS dependencies for several sentences
// returned by the API.
type SentenceDependencies struct {
	SentenceDependencies []SentenceDependency `json:"sentence_dependencies"`
}

type LibVersion struct {
	Spacy string `json:"spacy"`
}

type userInput struct {
	Text string `json:"text"`
}

// Entities extracts entities from a block of text by contacting the API.
func (c *Client) Entities(text string) (entities Entities, err error) {
	body, err := c.apiPost("entities", text)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &entities)
	if err != nil {
		return
	}

	return
}

// Dependencies gets POS dependencies from a block of text by contacting the API.
func (c *Client) Dependencies(text string) (dependencies Dependencies, err error) {
	body, err := c.apiPost("dependencies", text)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &dependencies)
	if err != nil {
		return
	}

	return
}

// SentenceDependencies gets POS dependencies with arcs from a block of text by contacting the API.
func (c *Client) SentenceDependencies(text string) (sentenceDependencies SentenceDependencies, err error) {
	body, err := c.apiPost("sentence-dependencies", text)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &sentenceDependencies)
	if err != nil {
		return
	}

	return
}

// LibVersions returns the spaCy versions used with the model.
// Only showing the spaCy version is temporary. More lib versions
// will be added soon.
func (c *Client) LibVersions() (libVersion LibVersion, err error) {
	body, err := c.apiGet("version")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &libVersion)
	if err != nil {
		return
	}
	return
}

// NewClient initializes a new Client.
func NewClient(model, token string) Client {
	return Client{
		rootURL: fmt.Sprintf("%v/%v/%v", baseURL, apiVersion, model),
		token:   token,
		client:  &http.Client{},
	}
}

func (c *Client) apiPost(endpoint, text string) (body []byte, err error) {
	data := new(bytes.Buffer)

	err = json.NewEncoder(data).Encode(userInput{Text: text})
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/%v", c.rootURL, endpoint), data)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %v", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ = ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("status code: %v, message: %v", resp.StatusCode, string(body))
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func (c *Client) apiGet(endpoint string) (body []byte, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", c.rootURL, endpoint), nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %v", c.token))

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
