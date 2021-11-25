//go:build integration
// +build integration

package integration_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/nlpcloud/nlpcloud-go"
	"github.com/stretchr/testify/assert"
)

func TestClientEntities(t *testing.T) {
	assert := assert.New(t)

	mdwClient := &MdwClient{}
	client := nlpcloud.NewClient(mdwClient, "en_core_web_lg", APIToken, false)
	resp, err := client.Entities(nlpcloud.EntitiesParams{
		Text: "John Doe has been working for Microsoft in Seattle since 1999.",
	})

	// Ensure no error
	if !assert.Nil(err) {
		t.Errorf("Last body [%s]\n", mdwClient.LastBody)
	}

	// Reencode to JSON
	buf := &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(resp)

	// Decode both to interfaces
	var expected interface{}
	var actual interface{}
	_ = json.Unmarshal(mdwClient.LastBody, &expected)
	_ = json.Unmarshal(buf.Bytes(), &actual)

	// Compares both to check valid API (and not nil)
	assert.NotNil(expected)
	assert.Equal(expected, actual)
}

func TestClientClassification(t *testing.T) {
	assert := assert.New(t)

	mdwClient := &MdwClient{}
	client := nlpcloud.NewClient(mdwClient, "bart-large-mnli", APIToken, false)
	resp, err := client.Classification(nlpcloud.ClassificationParams{
		Text:       "John Doe is a Go Developer at Google. He has been working there for 10 years and has been awarded employee of the year.",
		Labels:     []string{"job", "nature", "space"},
		MultiClass: b(true),
	})

	// Ensure no error
	if !assert.Nil(err) {
		t.Errorf("Last body [%s]\n", mdwClient.LastBody)
	}

	// Reencode to JSON
	buf := &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(resp)

	// Decode both to interfaces
	var expected interface{}
	var actual interface{}
	_ = json.Unmarshal(mdwClient.LastBody, &expected)
	_ = json.Unmarshal(buf.Bytes(), &actual)

	// Compares both to check valid API (and not nil)
	assert.NotNil(expected)
	assert.Equal(expected, actual)
}

func TestClientSentiment(t *testing.T) {
	assert := assert.New(t)

	mdwClient := &MdwClient{}
	client := nlpcloud.NewClient(mdwClient, "distilbert-base-uncased-finetuned-sst-2-english", APIToken, false)
	resp, err := client.Sentiment(nlpcloud.SentimentParams{
		Text: "NLP Cloud proposes an amazing service!",
	})

	// Ensure no error
	if !assert.Nil(err) {
		t.Errorf("Last body [%s]\n", mdwClient.LastBody)
	}

	// Reencode to JSON
	buf := &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(resp)

	// Decode both to interfaces
	var expected interface{}
	var actual interface{}
	_ = json.Unmarshal(mdwClient.LastBody, &expected)
	_ = json.Unmarshal(buf.Bytes(), &actual)

	// Compares both to check valid API (and not nil)
	assert.NotNil(expected)
	assert.Equal(expected, actual)
}

func TestClientQuestion(t *testing.T) {
	assert := assert.New(t)

	mdwClient := &MdwClient{}
	client := nlpcloud.NewClient(mdwClient, "roberta-base-squad2", APIToken, false)
	resp, err := client.Question(nlpcloud.QuestionParams{
		Context:  "French president Emmanuel Macron said the country was at war with an invisible, elusive enemy, and the measures were unprecedented, but circumstances demanded them.",
		Question: "Who is the French president?",
	})

	// Ensure no error
	if !assert.Nil(err) {
		t.Errorf("Last body [%s]\n", mdwClient.LastBody)
	}

	// Reencode to JSON
	buf := &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(resp)

	// Decode both to interfaces
	var expected interface{}
	var actual interface{}
	_ = json.Unmarshal(mdwClient.LastBody, &expected)
	_ = json.Unmarshal(buf.Bytes(), &actual)

	// Compares both to check valid API (and not nil)
	assert.NotNil(expected)
	assert.Equal(expected, actual)
}

func TestClientSummarization(t *testing.T) {
	assert := assert.New(t)

	mdwClient := &MdwClient{}
	client := nlpcloud.NewClient(mdwClient, "roberta-base-squad2", APIToken, false)
	resp, err := client.Summarization(nlpcloud.SummarizationParams{
		Text: "One month after the United States began what has become a troubled rollout of a national COVID vaccination campaign, the effort is finally gathering real steam. Close to a million doses -- over 951,000, to be more exact -- made their way into the arms of Americans in the past 24 hours, the U.S. Centers for Disease Control and Prevention reported Wednesday. That is the largest number of shots given in one day since the rollout began and a big jump from the previous day, when just under 340,000 doses were given, CBS News reported. That number is likely to jump quickly after the federal government on Tuesday gave states the OK to vaccinate anyone over 65 and said it would release all the doses of vaccine it has available for distribution. Meanwhile, a number of states have now opened mass vaccination sites in an effort to get larger numbers of people inoculated, CBS News reported.",
	})

	// Ensure no error
	if !assert.Nil(err) {
		t.Errorf("Last body [%s]\n", mdwClient.LastBody)
	}

	// Reencode to JSON
	buf := &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(resp)

	// Decode both to interfaces
	var expected interface{}
	var actual interface{}
	_ = json.Unmarshal(mdwClient.LastBody, &expected)
	_ = json.Unmarshal(buf.Bytes(), &actual)

	// Compares both to check valid API (and not nil)
	assert.NotNil(expected)
	assert.Equal(expected, actual)
}

func TestClientGeneration(t *testing.T) {
	assert := assert.New(t)

	mdwClient := &MdwClient{}
	client := nlpcloud.NewClient(mdwClient, "gpt-j", APIToken, false)
	resp, err := client.Generation(nlpcloud.GenerationParams{
		Text:      "GPT Neo is a powerful NLP model",
		MinLength: i(10),
		MaxLength: i(50),
	})

	// Ensure no error
	if !assert.Nil(err) {
		t.Errorf("Last body [%s]\n", mdwClient.LastBody)
	}

	// Reencode to JSON
	buf := &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(resp)

	// Decode both to interfaces
	var expected interface{}
	var actual interface{}
	_ = json.Unmarshal(mdwClient.LastBody, &expected)
	_ = json.Unmarshal(buf.Bytes(), &actual)

	// Compares both to check valid API (and not nil)
	assert.NotNil(expected)
	assert.Equal(expected, actual)
}

func TestClientTranslation(t *testing.T) {
	assert := assert.New(t)

	mdwClient := &MdwClient{}
	client := nlpcloud.NewClient(mdwClient, "opus-mt-en-fr", APIToken, false)
	resp, err := client.Translation(nlpcloud.TranslationParams{
		Text: "John Doe has been working for Microsoft in Seattle since 1999.",
	})

	// Ensure no error
	if !assert.Nil(err) {
		t.Errorf("Last body [%s]\n", mdwClient.LastBody)
	}

	// Reencode to JSON
	buf := &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(resp)

	// Decode both to interfaces
	var expected interface{}
	var actual interface{}
	_ = json.Unmarshal(mdwClient.LastBody, &expected)
	_ = json.Unmarshal(buf.Bytes(), &actual)

	// Compares both to check valid API (and not nil)
	assert.NotNil(expected)
	assert.Equal(expected, actual)
}

func TestClientLangDetection(t *testing.T) {
	assert := assert.New(t)

	mdwClient := &MdwClient{}
	client := nlpcloud.NewClient(mdwClient, "python-langdetect", APIToken, false)
	resp, err := client.LangDetection(nlpcloud.LangDetectionParams{
		Text: "John Doe has been working for Microsoft in Seattle since 1999. Il parle aussi un peu français.",
	})

	// Ensure no error
	if !assert.Nil(err) {
		t.Errorf("Last body [%s]\n", mdwClient.LastBody)
	}

	// Reencode to JSON
	buf := &bytes.Buffer{}
	_ = json.NewEncoder(buf).Encode(resp)

	// Decode both to interfaces
	var expected interface{}
	var actual interface{}
	_ = json.Unmarshal(mdwClient.LastBody, &expected)
	_ = json.Unmarshal(buf.Bytes(), &actual)

	// Compares both to check valid API (and not nil)
	assert.NotNil(expected)
	assert.Equal(expected, actual)
}

func b(b bool) *bool {
	return &b
}

func i(i int) *int {
	return &i
}
