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

// Classification holds the text classification returned by the API.
type Classification struct {
	Labels []string  `json:"labels"`
	Scores []float32 `json:"scores"`
}

// ScoredLabel holds a label and its score for sentiment analysis.
type ScoredLabel struct {
	Label string  `json:"label"`
	Score float32 `json:"score"`
}

// Sentiment holds the sentiment of a text returned by the API.
type Sentiment struct {
	ScoredLabels []ScoredLabel `json:"scored_labels"`
}

// Question holds the answer to a question by the API.
type Question struct {
	Answer string  `json:"string"`
	Score  float32 `json:"score"`
	Start  int     `json:"start"`
	End    int     `json:"end"`
}

// Summarization holds a summarized text returned by the API.
type Summarization struct {
	SummaryText string `json:"summary_text"`
}

// Translation holds a translated text returned by the API.
type Translation struct {
	TranslationText string `json:"translation_text"`
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

// LibVersions returns the versions behind the spaCy model.
type LibVersions struct {
	Spacy string `json:"spacy"`
}

type textInput struct {
	Text string `json:"text"`
}

type classificationInput struct {
	Text       string   `json:"text"`
	Labels     []string `json:"labels"`
	MultiClass bool     `json:"multi_class"`
}

type questionInput struct {
	Context  string `json:"context"`
	Question string `json:"question"`
}

// Entities extracts entities from a block of text by contacting the API.
func (c *Client) Entities(text string) (entities Entities, err error) {
	data := new(bytes.Buffer)

	err = json.NewEncoder(data).Encode(textInput{Text: text})
	if err != nil {
		return
	}

	body, err := c.apiPost("entities", data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &entities)
	if err != nil {
		return
	}

	return
}

// Classification applies scored labels to a block of text by contacting the API.
func (c *Client) Classification(text string, labels []string, multiClass bool) (classification Classification, err error) {
	data := new(bytes.Buffer)

	err = json.NewEncoder(data).Encode(classificationInput{
		Text:       text,
		Labels:     labels,
		MultiClass: multiClass,
	})
	if err != nil {
		return
	}

	body, err := c.apiPost("classification", data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &classification)
	if err != nil {
		return
	}

	return
}

// Sentiment defines the sentime of a block of text by contacting the API.
func (c *Client) Sentiment(text string) (sentiment Sentiment, err error) {
	data := new(bytes.Buffer)

	err = json.NewEncoder(data).Encode(textInput{Text: text})
	if err != nil {
		return
	}

	body, err := c.apiPost("sentiment", data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &sentiment)
	if err != nil {
		return
	}

	return
}

// Question answers a question with a context by contacting the API.
func (c *Client) Question(context, question string) (questionResponse Question, err error) {
	data := new(bytes.Buffer)

	err = json.NewEncoder(data).Encode(questionInput{
		Context:  context,
		Question: question,
	})
	if err != nil {
		return
	}

	body, err := c.apiPost("question", data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &questionResponse)
	if err != nil {
		return
	}

	return
}

// Summarization summarizes a block of text by contacting the API.
// Text should not exceed 1024 words.
func (c *Client) Summarization(text string) (summarization Summarization, err error) {
	data := new(bytes.Buffer)

	err = json.NewEncoder(data).Encode(textInput{Text: text})
	if err != nil {
		return
	}

	body, err := c.apiPost("summarization", data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &summarization)
	if err != nil {
		return
	}

	return
}

// Translation translates a block of text by contacting the API.
func (c *Client) Translation(text string) (translation Translation, err error) {
	data := new(bytes.Buffer)

	err = json.NewEncoder(data).Encode(textInput{Text: text})
	if err != nil {
		return
	}

	body, err := c.apiPost("translation", data)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &translation)
	if err != nil {
		return
	}

	return
}

// Dependencies gets POS dependencies from a block of text by contacting the API.
func (c *Client) Dependencies(text string) (dependencies Dependencies, err error) {
	data := new(bytes.Buffer)

	err = json.NewEncoder(data).Encode(textInput{Text: text})
	if err != nil {
		return
	}

	body, err := c.apiPost("dependencies", data)
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
	data := new(bytes.Buffer)

	err = json.NewEncoder(data).Encode(textInput{Text: text})
	if err != nil {
		return
	}

	body, err := c.apiPost("sentence-dependencies", data)
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
func (c *Client) LibVersions() (libVersions LibVersions, err error) {
	body, err := c.apiGet("versions")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &libVersions)
	if err != nil {
		return
	}
	return
}

// NewClient initializes a new Client.
func NewClient(model, token string, gpu bool) Client {
	var rootURL string

	if gpu {
		rootURL = fmt.Sprintf("%v/%v/gpu/%v", baseURL, apiVersion, model)
	} else {
		rootURL = fmt.Sprintf("%v/%v/%v", baseURL, apiVersion, model)
	}

	return Client{
		rootURL: rootURL,
		token:   token,
		client:  &http.Client{},
	}
}

func (c *Client) apiPost(endpoint string, data *bytes.Buffer) (body []byte, err error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/%v", c.rootURL, endpoint), data)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %v", c.token))
	req.Header.Set("User-Agent", "nlpcloud-go-client")

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
	req.Header.Set("User-Agent", "nlpcloud-go-client")

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
