package nlpcloud

import (
	"net/http"
)

// AdGenerationParams wraps all the parameters for the "ad-generation" endpoint.
type AdGenerationParams struct {
	Keywords []string `json:"keywords"`
}

// AdGeneration generates a product description or an ad by contacting the API.
func (c *Client) AdGeneration(params AdGenerationParams) (*AdGeneration, error) {
	adGeneration := &AdGeneration{}
	err := c.issueRequest(http.MethodPost, "ad-generation", params, adGeneration)
	if err != nil {
		return nil, err
	}
	return adGeneration, nil
}

// ClassificationParams wraps all the parameters for the "classification" endpoint.
type ClassificationParams struct {
	Text       string    `json:"text"`
	Labels     *[]string `json:"labels,omitempty"`
	MultiClass *bool     `json:"multi_class,omitempty"`
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

// DependenciesParams wraps all the parameters for the "dependencies" endpoint.
type DependenciesParams struct {
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

// EntitiesParams wraps all the parameters for the "entities" endpoint.
type EntitiesParams struct {
	Text           string  `json:"text"`
	SearchedEntity *string `json:"searched_entity,omitempty"`
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

// GenerationParams wraps all the parameters for the "generation" endpoint.
type GenerationParams struct {
	Text               string    `json:"text"`
	MinLength          *int      `json:"min_length,omitempty"`
	MaxLength          *int      `json:"max_length,omitempty"`
	LengthNoInput      *bool     `json:"length_no_input,omitempty"`
	EndSequence        *string   `json:"end_sequence,omitempty"`
	RemoveInput        *bool     `json:"remove_input,omitempty"`
	DoSample           *bool     `json:"do_sample,omitempty"`
	NumBeams           *int      `json:"num_beams,omitempty"`
	EarlyStopping      *bool     `json:"early_stopping,omitempty"`
	NoRepeatNgramSize  *int      `json:"no_repeat_ngram_size,omitempty"`
	NumReturnSequences *int      `json:"num_return_sequences,omitempty"`
	TopK               *int      `json:"top_k,omitempty"`
	TopP               *float64  `json:"top_p,omitempty"`
	Temperature        *float64  `json:"temperature,omitempty"`
	RepetitionPenalty  *float64  `json:"repetition_penalty,omitempty"`
	LengthPenalty      *float64  `json:"length_penalty,omitempty"`
	BadWords           *[]string `json:"bad_words,omitempty"`
	RemoveEndSequence  *bool     `json:"remove_end_sequence,omitempty"`
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

// GSCorrectionParams wraps all the parameters for the "gs-correction" endpoint.
type GSCorrectionParams struct {
	Text string `json:"text"`
}

// GSCorrection corrects the grammar and spelling by contacting the API.
func (c *Client) GSCorrection(params GSCorrectionParams) (*GSCorrection, error) {
	gSCorrection := &GSCorrection{}
	err := c.issueRequest(http.MethodPost, "gs-correction", params, gSCorrection)
	if err != nil {
		return nil, err
	}
	return gSCorrection, nil
}

// IntentClassificationParams wraps all the parameters for the "intent-classification" endpoint.
type IntentClassificationParams struct {
	Text string `json:"text"`
}

// IntentClassification classifies intent from a block of text by contacting the API.
func (c *Client) IntentClassification(params IntentClassificationParams) (*IntentClassification, error) {
	intentClassification := &IntentClassification{}
	err := c.issueRequest(http.MethodPost, "intent-classification", params, intentClassification)
	if err != nil {
		return nil, err
	}
	return intentClassification, nil
}

// KwKpExtractionParams wraps all the parameters for the "kw-kp-extraction" endpoint.
type KwKpExtractionParams struct {
	Text string `json:"text"`
}

// AdGeneration generates a product description or an ad by contacting the API.
func (c *Client) KwKpExtraction(params KwKpExtractionParams) (*KwKpExtraction, error) {
	kwKpExtraction := &KwKpExtraction{}
	err := c.issueRequest(http.MethodPost, "kw-kp-extraction", params, kwKpExtraction)
	if err != nil {
		return nil, err
	}
	return kwKpExtraction, nil
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

// LibVersions returns the versions used with the model by calling the API.
func (c *Client) LibVersions() (*LibVersions, error) {
	libVersions := &LibVersions{}
	err := c.issueRequest(http.MethodGet, "versions", nil, libVersions)
	if err != nil {
		return nil, err
	}
	return libVersions, nil
}

// ParaphrasingParams wraps all the parameters for the "paraphrasing" endpoint.
type ParaphrasingParams struct {
	Text string `json:"text"`
}

// Paraphrasing paraphrases a block of text by contacting the API.
func (c *Client) Paraphrasing(params ParaphrasingParams) (*Paraphrasing, error) {
	paraphrasing := &Paraphrasing{}
	err := c.issueRequest(http.MethodPost, "paraphrasing", params, paraphrasing)
	if err != nil {
		return nil, err
	}
	return paraphrasing, nil
}

// QuestionParams wraps all the parameters for the "question" endpoint.
type QuestionParams struct {
	Question string  `json:"question"`
	Context  *string `json:"context,omitempty"`
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

// SummarizationParams wraps all the parameters for the "summarization" endpoint.
type SummarizationParams struct {
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

// BatchSummarizationParams wraps all the parameters for the "batch-summarization" endpoint.
type BatchSummarizationParams struct {
	Texts []string `json:"texts"`
}

// BatchSummarization summarizes a batch of blocks of text by contacting the API.
func (c *Client) BatchSummarization(params BatchSummarizationParams) (*BatchSummarization, error) {
	batchSummarization := &BatchSummarization{}
	err := c.issueRequest(http.MethodPost, "batch-summarization", params, batchSummarization)
	if err != nil {
		return nil, err
	}
	return batchSummarization, nil
}

// TokensParams wraps all the parameters for the "tokens" endpoint.
type TokensParams struct {
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

// AdGeneration holds the generated product description or ad returned by the API.
type AdGeneration struct {
	GeneratedText string `json:"generated_text"`
}

// Classification holds the text classification returned by the API.
type Classification struct {
	Labels []string  `json:"labels"`
	Scores []float64 `json:"scores"`
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

// Generation holds a generated text returned by the API.
type Generation struct {
	GeneratedText     string `json:"generated_text"`
	NbGeneratedTokens int    `json:"nb_generated_tokens"`
	NbInputTokens     int    `json:"nb_input_tokens"`
}

// GSCorrection holds the corrected text returned by the API.
type GSCorrection struct {
	Correction string `json:"correction"`
}

// IntentClassification holds the classified intent returned by the API.
type IntentClassification struct {
	Intent string `json:"intent"`
}

// KwKpExtraction holds the extracted keywords and keyphrases returned by the API.
type KwKpExtraction struct {
	KeywordsAndKeyphrases string `json:"keywords_and_keyphrases"`
}

// LangDetection holds the languages of a text returned by the API.
type LangDetection struct {
	Languages []map[string]float64 `json:"languages"`
}

// LibVersions holds the versions of the libraries used behind the hood with the model.
type LibVersions map[string]string

// Paraphrasing holds a paraphrased text returned by the API.
type Paraphrasing struct {
	ParaphrasedText string `json:"paraphrased_text"`
}

// Question holds the answer to a question by the API.
type Question struct {
	Answer string  `json:"answer"`
	Score  float64 `json:"score"`
	Start  int     `json:"start"`
	End    int     `json:"end"`
}

// ScoredLabel holds a label and its score for sentiment analysis.
type ScoredLabel struct {
	Label string  `json:"label"`
	Score float64 `json:"score"`
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

// Sentiment holds the sentiment of a text returned by the API.
type Sentiment struct {
	ScoredLabels []ScoredLabel `json:"scored_labels"`
}

// Summarization holds a summarized text returned by the API.
type Summarization struct {
	SummaryText string `json:"summary_text"`
}

// BatchSummarization holds a batch of summarized texts returned by the API.
type BatchSummarization struct {
	SummaryTexts []string `json:"summary_texts"`
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

// Translation holds a translated text returned by the API.
type Translation struct {
	TranslationText string `json:"translation_text"`
}
