package nlpcloud

import (
	"net/http"
)

// EntitiesParams wraps all the parameters for the "entities" endpoint.
type EntitiesParams struct {
	Text string `json:"text"`
}

// Entities extracts entities from a block of text by contacting the API.
func (c *Client) Entities(params EntitiesParams) (*Entities, error) {
	entities := &Entities{}
	err := c.issueRequest(http.MethodPost, "entities", params, entities)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// ClassificationParams wraps all the parameters for the "classification" endpoint.
type ClassificationParams struct {
	Text       string   `json:"text"`
	Labels     []string `json:"labels"`
	MultiClass *bool    `json:"multi_class,omitempty"`
}

// Classification applies scored labels to a block of text by contacting the API.
func (c *Client) Classification(params ClassificationParams) (*Classification, error) {
	classification := &Classification{}
	err := c.issueRequest(http.MethodPost, "classification", params, classification)
	if err != nil {
		return nil, err
	}
	return classification, nil
}

// SentimentParams wraps all the parameters for the "sentiment" endpoint.
type SentimentParams struct {
	Text string `json:"text"`
}

// Sentiment defines the sentime of a block of text by contacting the API.
func (c *Client) Sentiment(params SentimentParams) (*Sentiment, error) {
	sentiment := &Sentiment{}
	err := c.issueRequest(http.MethodPost, "sentiment", params, sentiment)
	if err != nil {
		return nil, err
	}
	return sentiment, nil
}

// QuestionParams wraps all the parameters for the "question" endpoint.
type QuestionParams struct {
	Context  string `json:"context"`
	Question string `json:"question"`
}

// Question answers a question with a context by contacting the API.
func (c *Client) Question(params QuestionParams) (*Question, error) {
	ques := &Question{}
	err := c.issueRequest(http.MethodPost, "question", params, ques)
	if err != nil {
		return nil, err
	}
	return ques, nil
}

// SummarizationParams wraps all the parameters for the "summarization" endpoint.
type SummarizationParams struct {
	// Text should not exceed 1024 words.
	Text string `json:"text"`
}

// Summarization summarizes a block of text by contacting the API.
func (c *Client) Summarization(params SummarizationParams) (*Summarization, error) {
	summarization := &Summarization{}
	err := c.issueRequest(http.MethodPost, "summarization", params, summarization)
	if err != nil {
		return nil, err
	}
	return summarization, nil
}

// GenerationParams wraps all the parameters for the "generation" endpoint.
type GenerationParams struct {
	// Text should not exceed 1200 tokens.
	Text               string    `json:"text"`
	MinLength          *int      `json:"min_length,omitempty"`
	MaxLength          *int      `json:"max_length,omitempty"`
	LengthNoInput      *bool     `json:"length_no_input,omitempty"`
	EndSequence        *string   `json:"end_sequence,omitempty"`
	RemoveInput        *bool     `json:"remove_input,omitempty"`
	DoSample           *bool     `json:"do_sample,omitempty"`
	NumBeans           *int      `json:"num_beams,omitempty"`
	EarlyStopping      *bool     `json:"early_stopping,omitempty"`
	NoRepeatNgramSize  *int      `json:"no_repeat_ngram_size,omitempty"`
	NumReturnSequences *int      `json:"num_return_sequences,omitempty"`
	TopK               *int      `json:"top_k,omitempty"`
	TopP               *float64  `json:"top_p,omitempty"`
	Temperature        *float64  `json:"temperature,omitempty"`
	RepetitionPenalty  *float64  `json:"repetition_penalty,omitempty"`
	LengthPenalty      *float64  `json:"length_penalty,omitempty"`
	BadWords           *[]string `json:"bad_words,omitempty"`
}

// Generation generates a block of text by contacting the API.
func (c *Client) Generation(params GenerationParams) (*Generation, error) {
	generation := &Generation{}
	err := c.issueRequest(http.MethodPost, "generation", params, generation)
	if err != nil {
		return nil, err
	}
	return generation, nil
}

// TranslationParams wraps all the parameters for the "translation" endpoint.
type TranslationParams struct {
	Text string `json:"text"`
}

// Translation translates a block of text by contacting the API.
func (c *Client) Translation(params TranslationParams) (*Translation, error) {
	translation := &Translation{}
	err := c.issueRequest(http.MethodPost, "translation", params, translation)
	if err != nil {
		return nil, err
	}
	return translation, nil
}

// LangDetectionParams wraps all the parameters for the "langdetection" endpoint.
type LangDetectionParams struct {
	// Text should not exceed 100.000 characters.
	Text string `json:"text"`
}

// LangDetection returns an estimation of the text language by contacting the API.
func (c *Client) LangDetection(params LangDetectionParams) (*LangDetection, error) {
	langDetection := &LangDetection{}
	err := c.issueRequest(http.MethodPost, "langdetection", params, langDetection)
	if err != nil {
		return nil, err
	}
	return langDetection, nil
}

// DependenciesParams wraps all the parameters for the "dependencies" endpoint.
type DependenciesParams struct {
	// Text should not exceed 1.000 characters.
	Text string `json:"text"`
}

// Dependencies gets POS dependencies from a block of text by contacting the API.
func (c *Client) Dependencies(params DependenciesParams) (*Dependencies, error) {
	dependencies := &Dependencies{}
	err := c.issueRequest(http.MethodPost, "dependencies", params, dependencies)
	if err != nil {
		return nil, err
	}
	return dependencies, nil
}

// SentenceDependenciesParams wraps all the parameters for the "sentence-dependencies" endpoint.
type SentenceDependenciesParams struct {
	Text string `json:"text"`
}

// SentenceDependencies gets POS dependencies with arcs from a block of text by contacting the API.
func (c *Client) SentenceDependencies(params SentenceDependenciesParams) (*SentenceDependencies, error) {
	sentenceDependencies := &SentenceDependencies{}
	err := c.issueRequest(http.MethodPost, "sentence-dependencies", params, sentenceDependencies)
	if err != nil {
		return nil, err
	}
	return sentenceDependencies, nil
}

// TokensParams wraps all the parameters for the "tokens" endpoint.
type TokensParams struct {
	// Text should not exceed 1.000 characters.
	Text string `json:"text"`
}

// Tokens tokenize and lemmatize text by calling the API.
func (c *Client) Tokens(params TokensParams) (*Tokens, error) {
	tokens := &Tokens{}
	err := c.issueRequest(http.MethodPost, "tokens", params, tokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

// LibVersions returns the spaCy versions used with the model by calling the API.
func (c *Client) LibVersions() (*LibVersions, error) {
	libVersions := &LibVersions{}
	err := c.issueRequest(http.MethodGet, "versions", nil, libVersions)
	if err != nil {
		return nil, err
	}
	return libVersions, nil
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
	Scores []float64 `json:"scores"`
}

// ScoredLabel holds a label and its score for sentiment analysis.
type ScoredLabel struct {
	Label string  `json:"label"`
	Score float64 `json:"score"`
}

// Sentiment holds the sentiment of a text returned by the API.
type Sentiment struct {
	ScoredLabels []ScoredLabel `json:"scored_labels"`
}

// Question holds the answer to a question by the API.
type Question struct {
	Answer string  `json:"answer"`
	Score  float64 `json:"score"`
	Start  int     `json:"start"`
	End    int     `json:"end"`
}

// Summarization holds a summarized text returned by the API.
type Summarization struct {
	SummaryText string `json:"summary_text"`
}

// Generation holds a generated text returned by the API.
type Generation struct {
	GeneratedText     string `json:"generated_text"`
	NbGeneratedTokens int    `json:"nb_generated_tokens"`
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

// LangDetection holds the languages of a text returned by the API.
type LangDetection struct {
	Languages []map[string]float64 `json:"languages"`
}

// Dependencies holds a list of POS dependencies returned by the API.
type Dependencies struct {
	Words []Word `json:"words"`
	Arcs  []Arc  `json:"arcs"`
}

// SentenceDependency holds a POS dependency for one sentence
// returned by the API.
type SentenceDependency struct {
	Sentence     string       `json:"sentence"`
	Dependencies Dependencies `json:"dependencies"`
}

// SentenceDependencies holds a list of POS dependencies for several sentences
// returned by the API.
type SentenceDependencies struct {
	SentenceDependencies []SentenceDependency `json:"sentence_dependencies"`
}

// Tokens holds a list of Token returned by the API.
type Tokens struct {
	Tokens []Token `json:"tokens"`
}

// Token holds a token value from Tokens.
type Token struct {
	Text    string `json:"text"`
	Lemma   string `json:"lemma"`
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Index   int    `json:"index"`
	WSAfter bool   `json:"ws_after"`
}

// LibVersions holds the versions of the libraries used behind the hood with the model.
type LibVersions map[string]string
