package nlpcloud

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// AdGenerationParams wraps all the parameters for the "ad-generation" endpoint.
type AdGenerationParams struct {
	Keywords []string `json:"keywords"`
}

// AdGeneration generates a product description or an ad by contacting the API.
func (c *Client) AdGeneration(params AdGenerationParams, opts ...Option) (*AdGeneration, error) {
	adGeneration := &AdGeneration{}
	err := c.issueRequest(http.MethodPost, "ad-generation", params, adGeneration, opts...)
	if err != nil {
		return nil, err
	}
	return adGeneration, nil
}

// ASRParams wraps all the parameters for the "asr" endpoint.
type ASRParams struct {
	URL           *string `json:"url"`
	EncodedFile   *string `json:"encoded_file"`
	InputLanguage *string `json:"input_language"`
}

// ASR extracts text from an audio file by contacting the API.
func (c *Client) ASR(params ASRParams, opts ...Option) (*ASR, error) {
	asr := &ASR{}
	err := c.issueRequest(http.MethodPost, "asr", params, asr, opts...)
	if err != nil {
		return nil, err
	}
	return asr, nil
}

// AsyncResultParams wraps all the parameters for the "async-result" endpoint.
type AsyncResultParams struct {
	URL string `json:"url"`
}

// AsyncResult extracts gets an async result by contacting the API.
func (c *Client) AsyncResult(params AsyncResultParams, opts ...Option) (*AsyncResult, error) {
	asyncResult := &AsyncResult{}

	resp, err := http.Get(params.URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return nil, &HTTPError{
			Detail: string(body),
			Status: resp.StatusCode,
		}
	}

	if err = json.Unmarshal(body, asyncResult); err != nil {
		return nil, err
	}

	return asyncResult, nil
}

type Exchange struct {
	Input    string `json:"input"`
	Response string `json:"response"`
}

// ChatbotParams wraps all the parameters for the "chatbot" endpoint.
type ChatbotParams struct {
	Input   string      `json:"input"`
	History *[]Exchange `json:"history,omitempty"`
	Context *string     `json:"context,omitempty"`
}

// Chatbot responds as a human by contacting the API.
func (c *Client) Chatbot(params ChatbotParams, opts ...Option) (*Chatbot, error) {
	chatbot := &Chatbot{}
	err := c.issueRequest(http.MethodPost, "chatbot", params, chatbot, opts...)
	if err != nil {
		return nil, err
	}
	return chatbot, nil
}

// StreamingChatbot responds as a human by contacting the API, and returns a stream.
func (c *Client) StreamingChatbot(params ChatbotParams, opts ...Option) (io.ReadCloser, error) {
	streamBody, err := c.issueStreamingRequest(http.MethodPost, "chatbot", params)
	if err != nil {
		return nil, err
	}

	return streamBody, nil
}

// ClassificationParams wraps all the parameters for the "classification" endpoint.
type ClassificationParams struct {
	Text       string    `json:"text"`
	Labels     *[]string `json:"labels,omitempty"`
	MultiClass *bool     `json:"multi_class,omitempty"`
}

// Classification applies scored labels to a block of text by contacting the API.
func (c *Client) Classification(params ClassificationParams, opts ...Option) (*Classification, error) {
	classification := &Classification{}
	err := c.issueRequest(http.MethodPost, "classification", params, classification, opts...)
	if err != nil {
		return nil, err
	}
	return classification, nil
}

// BatchClassificationParams wraps all the parameters for the "batch-classification" endpoint.
type BatchClassificationParams struct {
	Texts  []string `json:"texts"`
	Labels []string `json:"labels"`
}

// BatchClassification classifies a batch of blocks of text by contacting the API.
func (c *Client) BatchClassification(params BatchClassificationParams, opts ...Option) (*BatchClassification, error) {
	batchClassification := &BatchClassification{}
	err := c.issueRequest(http.MethodPost, "batch-classification", params, batchClassification, opts...)
	if err != nil {
		return nil, err
	}
	return batchClassification, nil
}

// CodeGenerationParams wraps all the parameters for the "code-generation" endpoint.
type CodeGenerationParams struct {
	Intruction string `json:"instruction"`
}

// CodeGeneration generates source code by contacting the API.
func (c *Client) CodeGeneration(params CodeGenerationParams, opts ...Option) (*CodeGeneration, error) {
	codeGeneration := &CodeGeneration{}
	err := c.issueRequest(http.MethodPost, "code-generation", params, codeGeneration, opts...)
	if err != nil {
		return nil, err
	}
	return codeGeneration, nil
}

// DependenciesParams wraps all the parameters for the "dependencies" endpoint.
type DependenciesParams struct {
	Text string `json:"text"`
}

// Dependencies gets POS dependencies from a block of text by contacting the API.
func (c *Client) Dependencies(params DependenciesParams, opts ...Option) (*Dependencies, error) {
	dependencies := &Dependencies{}
	err := c.issueRequest(http.MethodPost, "dependencies", params, dependencies, opts...)
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
func (c *Client) Entities(params EntitiesParams, opts ...Option) (*Entities, error) {
	entities := &Entities{}
	err := c.issueRequest(http.MethodPost, "entities", params, entities, opts...)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// EmbeddingsParams wraps all the parameters for the "embeddings" endpoint.
type EmbeddingsParams struct {
	Sentences []string `json:"sentences"`
}

// Embeddings extracts embeddings from a list of sentences by contacting the API.
func (c *Client) Embeddings(params EmbeddingsParams, opts ...Option) (*Embeddings, error) {
	embeddings := &Embeddings{}
	err := c.issueRequest(http.MethodPost, "embeddings", params, embeddings, opts...)
	if err != nil {
		return nil, err
	}
	return embeddings, nil
}

// GenerationParams wraps all the parameters for the "generation" endpoint.
type GenerationParams struct {
	Text               string    `json:"text"`
	MaxLength          *int      `json:"max_length,omitempty"`
	LengthNoInput      *bool     `json:"length_no_input,omitempty"`
	EndSequence        *string   `json:"end_sequence,omitempty"`
	RemoveInput        *bool     `json:"remove_input,omitempty"`
	NumBeams           *int      `json:"num_beams,omitempty"`
	NumReturnSequences *int      `json:"num_return_sequences,omitempty"`
	TopK               *int      `json:"top_k,omitempty"`
	TopP               *float64  `json:"top_p,omitempty"`
	Temperature        *float64  `json:"temperature,omitempty"`
	RepetitionPenalty  *float64  `json:"repetition_penalty,omitempty"`
	BadWords           *[]string `json:"bad_words,omitempty"`
	RemoveEndSequence  *bool     `json:"remove_end_sequence,omitempty"`
	// NoChatPrompt is experimental. It is not suited for production.
	NoChatPrompt *bool `json:"no_chat_prompt,omitempty"`
}

// Generation generates a block of text by contacting the API.
func (c *Client) Generation(params GenerationParams, opts ...Option) (*Generation, error) {
	generation := &Generation{}
	err := c.issueRequest(http.MethodPost, "generation", params, generation, opts...)
	if err != nil {
		return nil, err
	}
	return generation, nil
}

// StreamingGeneration generates a block of text by contacting the API, and returns a stream.
func (c *Client) StreamingGeneration(params GenerationParams, opts ...Option) (io.ReadCloser, error) {
	streamBody, err := c.issueStreamingRequest(http.MethodPost, "generation", params)
	if err != nil {
		return nil, err
	}

	return streamBody, nil
}

// BatchGenerationParams wraps all the parameters for the "batch-generation" endpoint.
type BatchGenerationParams struct {
	Texts []string `json:"texts"`
}

// BatchGeneration generates a batch of blocks of text by contacting the API.
func (c *Client) BatchGeneration(params BatchGenerationParams) (*BatchGeneration, error) {
	batchGeneration := &BatchGeneration{}
	err := c.issueRequest(http.MethodPost, "batch-generation", params, batchGeneration)
	if err != nil {
		return nil, err
	}
	return batchGeneration, nil
}

// GSCorrectionParams wraps all the parameters for the "gs-correction" endpoint.
type GSCorrectionParams struct {
	Text string `json:"text"`
}

// GSCorrection corrects the grammar and spelling by contacting the API.
func (c *Client) GSCorrection(params GSCorrectionParams, opts ...Option) (*GSCorrection, error) {
	gSCorrection := &GSCorrection{}
	err := c.issueRequest(http.MethodPost, "gs-correction", params, gSCorrection, opts...)
	if err != nil {
		return nil, err
	}
	return gSCorrection, nil
}

// ImageGenerationParams wraps all the parameters for the "image-generation" endpoint.
type ImageGenerationParams struct {
	Text string `json:"text"`
}

// ImageGeneration generates an image out of a text instruction by contacting the API.
func (c *Client) ImageGeneration(params ImageGenerationParams, opts ...Option) (*ImageGeneration, error) {
	imageGeneration := &ImageGeneration{}
	err := c.issueRequest(http.MethodPost, "image-generation", params, imageGeneration, opts...)
	if err != nil {
		return nil, err
	}
	return imageGeneration, nil
}

// IntentClassificationParams wraps all the parameters for the "intent-classification" endpoint.
type IntentClassificationParams struct {
	Text string `json:"text"`
}

// IntentClassification classifies intent from a block of text by contacting the API.
func (c *Client) IntentClassification(params IntentClassificationParams, opts ...Option) (*IntentClassification, error) {
	intentClassification := &IntentClassification{}
	err := c.issueRequest(http.MethodPost, "intent-classification", params, intentClassification, opts...)
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
func (c *Client) KwKpExtraction(params KwKpExtractionParams, opts ...Option) (*KwKpExtraction, error) {
	kwKpExtraction := &KwKpExtraction{}
	err := c.issueRequest(http.MethodPost, "kw-kp-extraction", params, kwKpExtraction, opts...)
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
func (c *Client) LangDetection(params LangDetectionParams, opts ...Option) (*LangDetection, error) {
	langDetection := &LangDetection{}
	err := c.issueRequest(http.MethodPost, "langdetection", params, langDetection, opts...)
	if err != nil {
		return nil, err
	}
	return langDetection, nil
}

// ParaphrasingParams wraps all the parameters for the "paraphrasing" endpoint.
type ParaphrasingParams struct {
	Text string `json:"text"`
}

// Paraphrasing paraphrases a block of text by contacting the API.
func (c *Client) Paraphrasing(params ParaphrasingParams, opts ...Option) (*Paraphrasing, error) {
	paraphrasing := &Paraphrasing{}
	err := c.issueRequest(http.MethodPost, "paraphrasing", params, paraphrasing, opts...)
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
func (c *Client) Question(params QuestionParams, opts ...Option) (*Question, error) {
	ques := &Question{}
	err := c.issueRequest(http.MethodPost, "question", params, ques, opts...)
	if err != nil {
		return nil, err
	}
	return ques, nil
}

// SemanticSearchParams wraps all the parameters for the "semantic-search" endpoint.
type SemanticSearchParams struct {
	Text       string `json:"text"`
	NumResults int    `json:"num_results"`
}

// SemanticSearch performs semantic search on custom data contacting the API.
func (c *Client) SemanticSearch(params SemanticSearchParams, opts ...Option) (*SemanticSearch, error) {
	semanticSearch := &SemanticSearch{}
	err := c.issueRequest(http.MethodPost, "semantic-similarity", params, semanticSearch, opts...)
	if err != nil {
		return nil, err
	}
	return semanticSearch, nil
}

// SemanticSimilarityParams wraps all the parameters for the "semantic-similarity" endpoint.
type SemanticSimilarityParams struct {
	Sentences [2]string `json:"sentences"`
}

// SemanticSimilarity calculates a semantic similarity score out of 2 sentences by contacting the API.
func (c *Client) SemanticSimilarity(params SemanticSimilarityParams, opts ...Option) (*SemanticSimilarity, error) {
	semanticSimilarity := &SemanticSimilarity{}
	err := c.issueRequest(http.MethodPost, "semantic-similarity", params, semanticSimilarity, opts...)
	if err != nil {
		return nil, err
	}
	return semanticSimilarity, nil
}

// SentenceDependenciesParams wraps all the parameters for the "sentence-dependencies" endpoint.
type SentenceDependenciesParams struct {
	Text string `json:"text"`
}

// SentenceDependencies gets POS dependencies with arcs from a block of text by contacting the API.
func (c *Client) SentenceDependencies(params SentenceDependenciesParams, opts ...Option) (*SentenceDependencies, error) {
	sentenceDependencies := &SentenceDependencies{}
	err := c.issueRequest(http.MethodPost, "sentence-dependencies", params, sentenceDependencies, opts...)
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
func (c *Client) Sentiment(params SentimentParams, opts ...Option) (*Sentiment, error) {
	sentiment := &Sentiment{}
	err := c.issueRequest(http.MethodPost, "sentiment", params, sentiment, opts...)
	if err != nil {
		return nil, err
	}
	return sentiment, nil
}

// SpeechSynthesisParams wraps all the parameters for the "speech-synthesis" endpoint.
type SpeechSynthesisParams struct {
	Text  string  `json:"text"`
	Voice *string `json:"voice"`
}

// SpeechSynthesis generates audio out of a text by contacting the API.
func (c *Client) SpeechSynthesis(params SpeechSynthesisParams, opts ...Option) (*SpeechSynthesis, error) {
	speechSynthesis := &SpeechSynthesis{}
	err := c.issueRequest(http.MethodPost, "speech-synthesis", params, speechSynthesis, opts...)
	if err != nil {
		return nil, err
	}
	return speechSynthesis, nil
}

// SummarizationParams wraps all the parameters for the "summarization" endpoint.
type SummarizationParams struct {
	Text string  `json:"text"`
	Size *string `json:"size"`
}

// Summarization summarizes a block of text by contacting the API.
func (c *Client) Summarization(params SummarizationParams, opts ...Option) (*Summarization, error) {
	summarization := &Summarization{}
	err := c.issueRequest(http.MethodPost, "summarization", params, summarization, opts...)
	if err != nil {
		return nil, err
	}
	return summarization, nil
}

// BatchSummarizationParams wraps all the parameters for the "batch-summarization" endpoint.
type BatchSummarizationParams struct {
	Texts []string `json:"texts"`
	Size  string   `json:"size"`
}

// BatchSummarization summarizes a batch of blocks of text by contacting the API.
func (c *Client) BatchSummarization(params BatchSummarizationParams, opts ...Option) (*BatchSummarization, error) {
	batchSummarization := &BatchSummarization{}
	err := c.issueRequest(http.MethodPost, "batch-summarization", params, batchSummarization, opts...)
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
func (c *Client) Tokens(params TokensParams, opts ...Option) (*Tokens, error) {
	tokens := &Tokens{}
	err := c.issueRequest(http.MethodPost, "tokens", params, tokens, opts...)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

// TranslationParams wraps all the parameters for the "translation" endpoint.
type TranslationParams struct {
	Text   string  `json:"text"`
	Source *string `json:"source,omitempty"`
	Target *string `json:"target,omitempty"`
}

// Translation translates a block of text by contacting the API.
func (c *Client) Translation(params TranslationParams, opts ...Option) (*Translation, error) {
	translation := &Translation{}
	err := c.issueRequest(http.MethodPost, "translation", params, translation, opts...)
	if err != nil {
		return nil, err
	}
	return translation, nil
}

// BatchTranslationParams wraps all the parameters for the "batch-translation" endpoint.
type BatchTranslationParams struct {
	Texts   []string  `json:"texts"`
	Sources *[]string `json:"sources,omitempty"`
	Targets *[]string `json:"targets,omitempty"`
}

// BatchTranslation translates a batch of blocks of text by contacting the API.
func (c *Client) BatchTranslation(params BatchTranslationParams, opts ...Option) (*BatchTranslation, error) {
	batchTranslation := &BatchTranslation{}
	err := c.issueRequest(http.MethodPost, "batch-translation", params, batchTranslation, opts...)
	if err != nil {
		return nil, err
	}
	return batchTranslation, nil
}

// AdGeneration holds the generated product description or ad returned by the API.
type AdGeneration struct {
	GeneratedText string `json:"generated_text"`
}

// Async holds the information returned when making an async request.
type Async struct {
	URL string `json:"url"`
}

// AsyncResult holds the information returned when getting an async result.

type AsyncResult struct {
	CreatedOn   time.Time `json:"created_on"`
	FinishedOn  time.Time `json:"finished_on"`
	RequestBody string    `json:"request_body"`
	HTTPCode    int       `json:"http_code"`
	ErrorDetail string    `json:"error_detail"`
	Content     string    `json:"content"`
}

// Segment holds an ASR segment.
type Segment struct {
	ID      int     `json:"id"`
	Starter float64 `json:"start"`
	End     float64 `json:"end"`
	Text    string  `json:"text"`
}

// Word holds an ASR word.
type ASRWord struct {
	ID      int     `json:"id"`
	Starter float64 `json:"start"`
	End     float64 `json:"end"`
	Text    string  `json:"text"`
	Prob    float64 `json:"prob"`
}

// ASR holds the extracted text returned by the API.
type ASR struct {
	Text     string    `json:"text"`
	Duration int       `json:"duration"`
	Language string    `json:"language"`
	Segments []Segment `json:"segments"`
	Words    []ASRWord `json:"words"`
}

// Chatbot holds the chatbot response returned by the API.
type Chatbot struct {
	Response string     `json:"response"`
	History  []Exchange `json:"history"`
}

// Classification holds the text classification returned by the API.
type Classification struct {
	Labels []string  `json:"labels"`
	Scores []float64 `json:"scores"`
}

// BatchClassification holds a batch of scores returned by the API.
type BatchClassification struct {
	Scores []float64 `json:"scores"`
}

// CodeGeneration holds the generated code returned by the API.
type CodeGeneration struct {
	GeneratedCode string `json:"generated_code"`
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

// Embeddings holds text embeddings returned by the API.

type Embeddings struct {
	Embeddings [][]float64 `json:"embeddings"`
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

// BatchGeneration holds a batch of generations returned by the API.
type BatchGeneration struct {
	Generations []Generation `json:"generations"`
}

// GSCorrection holds the corrected text returned by the API.
type GSCorrection struct {
	Correction string `json:"correction"`
}

// ImageGeneration holds the generated image url returned by the API.
type ImageGeneration struct {
	URL string `json:"url"`
}

// IntentClassification holds the classified intent returned by the API.
type IntentClassification struct {
	Intent string `json:"intent"`
}

// KwKpExtraction holds the extracted keywords and keyphrases returned by the API.
type KwKpExtraction struct {
	KeywordsAndKeyphrases []string `json:"keywords_and_keyphrases"`
}

// LangDetection holds the languages of a text returned by the API.
type LangDetection struct {
	Languages []map[string]float64 `json:"languages"`
}

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

// SearchResult holds a search result from semantic search.
type SearchResult struct {
	Score float64 `json:"score"`
	Text  string  `json:"text"`
}

// SemanticSearch holds semantic search results returned by the API.
type SemanticSearch struct {
	SearchResults []SearchResult `json:"search_results"`
}

// SemanticSimilarity holds semantic similarity score returned by the API.
type SemanticSimilarity struct {
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

// SpeechSynthesis holds the generated audio file url returned by the API.
type SpeechSynthesis struct {
	URL string `json:"url"`
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

// BatchTranslation holds a batch of translated texts returned by the API.
type BatchTranslation struct {
	TranslationTexts []string `json:"translation_texts"`
}
