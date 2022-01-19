package nlpcloud_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/nlpcloud/nlpcloud-go"
	"github.com/stretchr/testify/assert"
)

func TestClientEntities(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client           nlpcloud.HTTPClient
		Params           nlpcloud.EntitiesParams
		ExpectedEntities *nlpcloud.Entities
		ExpectedErr      error
	}{
		"nil-client": {
			Client:           nil,
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:           newFakeHTTPClient(``, 0, errFake),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      errFake,
		},
		"bad-request": {
			Client:           newFakeHTTPClient(``, 400, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:           newFakeHTTPClient(``, 401, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:           newFakeHTTPClient(``, 402, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:           newFakeHTTPClient(``, 403, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:           newFakeHTTPClient(``, 404, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:           newFakeHTTPClient(``, 405, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:           newFakeHTTPClient(``, 406, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:           newFakeHTTPClient(``, 413, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:           newFakeHTTPClient(``, 422, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:           newFakeHTTPClient(``, 429, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:           newFakeHTTPClient(``, 500, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:           newFakeHTTPClient(``, 503, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr:      nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:           newFakeHTTPClient(``, 0, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:           newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:           nlpcloud.EntitiesParams{},
			ExpectedEntities: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"entities":[{"start":0,"end":8,"type":"PERSON","text":"John Doe"},{"start":30,"end":39,"type":"ORG","text":"Microsoft"},{"start":43,"end":50,"type":"GPE","text":"Seattle"},{"start":57,"end":61,"type":"DATE","text":"1999"}]}`, http.StatusOK, nil),
			Params: nlpcloud.EntitiesParams{
				Text: "John Doe has been working for Microsoft in Seattle since 1999.",
			},
			ExpectedEntities: &nlpcloud.Entities{
				Entities: []nlpcloud.Entity{
					{
						Start: 0,
						End:   8,
						Type:  "PERSON",
						Text:  "John Doe",
					}, {
						Start: 30,
						End:   39,
						Type:  "ORG",
						Text:  "Microsoft",
					}, {
						Start: 43,
						End:   50,
						Type:  "GPE",
						Text:  "Seattle",
					}, {
						Start: 57,
						End:   61,
						Type:  "DATE",
						Text:  "1999",
					},
				},
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			entities, err := client.Entities(tt.Params)

			assert.Equal(tt.ExpectedEntities, entities)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientClassification(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client                 nlpcloud.HTTPClient
		Params                 nlpcloud.ClassificationParams
		ExpectedClassification *nlpcloud.Classification
		ExpectedErr            error
	}{
		"nil-client": {
			Client:                 nil,
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:                 newFakeHTTPClient(``, 0, errFake),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            errFake,
		},
		"bad-request": {
			Client:                 newFakeHTTPClient(``, 400, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:                 newFakeHTTPClient(``, 401, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:                 newFakeHTTPClient(``, 402, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:                 newFakeHTTPClient(``, 403, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:                 newFakeHTTPClient(``, 404, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:                 newFakeHTTPClient(``, 405, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:                 newFakeHTTPClient(``, 406, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:                 newFakeHTTPClient(``, 413, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:                 newFakeHTTPClient(``, 422, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:                 newFakeHTTPClient(``, 429, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:                 newFakeHTTPClient(``, 500, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:                 newFakeHTTPClient(``, 503, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr:            nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:                 newFakeHTTPClient(``, 0, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:                 newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:                 nlpcloud.ClassificationParams{},
			ExpectedClassification: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"labels":["job","space","nature"],"scores":[0.9258800745010376,0.1938474327325821,0.010988450609147549]}`, http.StatusOK, nil),
			Params: nlpcloud.ClassificationParams{
				Text: "`John Doe is a Go Developer at Google. He has been working there for 10 years and has been awarded employee of the year.",
				Labels: []string{
					"job",
					"nature",
					"space",
				},
				MultiClass: b(true),
			},
			ExpectedClassification: &nlpcloud.Classification{
				Labels: []string{
					"job",
					"space",
					"nature",
				},
				Scores: []float64{
					0.9258800745010376,
					0.1938474327325821,
					0.010988450609147549,
				},
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			classification, err := client.Classification(tt.Params)

			assert.Equal(tt.ExpectedClassification, classification)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientSentiment(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client            nlpcloud.HTTPClient
		Params            nlpcloud.SentimentParams
		ExpectedSentiment *nlpcloud.Sentiment
		ExpectedErr       error
	}{
		"nil-client": {
			Client:            nil,
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:            newFakeHTTPClient(``, 0, errFake),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       errFake,
		},
		"bad-request": {
			Client:            newFakeHTTPClient(``, 400, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:            newFakeHTTPClient(``, 401, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:            newFakeHTTPClient(``, 402, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:            newFakeHTTPClient(``, 403, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:            newFakeHTTPClient(``, 404, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:            newFakeHTTPClient(``, 405, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:            newFakeHTTPClient(``, 406, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:            newFakeHTTPClient(``, 413, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:            newFakeHTTPClient(``, 422, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:            newFakeHTTPClient(``, 429, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:            newFakeHTTPClient(``, 500, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:            newFakeHTTPClient(``, 503, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr:       nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:            newFakeHTTPClient(``, 0, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:            newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:            nlpcloud.SentimentParams{},
			ExpectedSentiment: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"scored_labels":[{"label":"POSITIVE","score":0.9996881484985352}]}`, http.StatusOK, nil),
			Params: nlpcloud.SentimentParams{
				Text: "NLP Cloud proposes an amazing service!",
			},
			ExpectedSentiment: &nlpcloud.Sentiment{
				ScoredLabels: []nlpcloud.ScoredLabel{
					{
						Label: "POSITIVE",
						Score: 0.9996881484985352,
					},
				},
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			sentiment, err := client.Sentiment(tt.Params)

			assert.Equal(tt.ExpectedSentiment, sentiment)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientQuestion(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client           nlpcloud.HTTPClient
		Params           nlpcloud.QuestionParams
		ExpectedQuestion *nlpcloud.Question
		ExpectedErr      error
	}{
		"nil-client": {
			Client:           nil,
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:           newFakeHTTPClient(``, 0, errFake),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      errFake,
		},
		"bad-request": {
			Client:           newFakeHTTPClient(``, 400, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:           newFakeHTTPClient(``, 401, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:           newFakeHTTPClient(``, 402, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:           newFakeHTTPClient(``, 403, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:           newFakeHTTPClient(``, 404, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:           newFakeHTTPClient(``, 405, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:           newFakeHTTPClient(``, 406, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:           newFakeHTTPClient(``, 413, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:           newFakeHTTPClient(``, 422, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:           newFakeHTTPClient(``, 429, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:           newFakeHTTPClient(``, 500, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:           newFakeHTTPClient(``, 503, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr:      nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:           newFakeHTTPClient(``, 0, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:           newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:           nlpcloud.QuestionParams{},
			ExpectedQuestion: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"answer":"Emmanuel Macron","score":0.9595934152603149,"start":17,"end":32}`, http.StatusOK, nil),
			Params: nlpcloud.QuestionParams{
				Context: "French president Emmanuel Macron said the country was at war with an invisible, elusive enemy, and the measures were unprecedented,	but circumstances demanded them.",
				Question: "Who is the French president?",
			},
			ExpectedQuestion: &nlpcloud.Question{
				Answer: "Emmanuel Macron",
				Score:  0.9595934152603149,
				Start:  17,
				End:    32,
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			question, err := client.Question(tt.Params)

			assert.Equal(tt.ExpectedQuestion, question)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientSummarization(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client                nlpcloud.HTTPClient
		Params                nlpcloud.SummarizationParams
		ExpectedSummarization *nlpcloud.Summarization
		ExpectedErr           error
	}{
		"nil-client": {
			Client:                nil,
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:                newFakeHTTPClient(``, 0, errFake),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           errFake,
		},
		"bad-request": {
			Client:                newFakeHTTPClient(``, 400, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:                newFakeHTTPClient(``, 401, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:                newFakeHTTPClient(``, 402, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:                newFakeHTTPClient(``, 403, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:                newFakeHTTPClient(``, 404, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:                newFakeHTTPClient(``, 405, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:                newFakeHTTPClient(``, 406, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:                newFakeHTTPClient(``, 413, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:                newFakeHTTPClient(``, 422, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:                newFakeHTTPClient(``, 429, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:                newFakeHTTPClient(``, 500, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:                newFakeHTTPClient(``, 503, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr:           nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:                newFakeHTTPClient(``, 0, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:                newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:                nlpcloud.SummarizationParams{},
			ExpectedSummarization: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"summary_text":"Over 951,000 doses were given in the past 24 hours. That's the largest number of shots given in one day since the rollout began. That number is likely to jump quickly after the federal government gave states the OK to vaccinate anyone over 65. A number of states have now opened mass vaccination sites."}`, http.StatusOK, nil),
			Params: nlpcloud.SummarizationParams{
				Text: "One month after the United States began what has become a " +
					"troubled rollout of a national COVID vaccination campaign, the effort is finally " +
					"gathering real steam. Close to a million doses -- over 951,000, to be more exact -- " +
					"made their way into the arms of Americans in the past 24 hours, the U.S. Centers " +
					"for Disease Control and Prevention reported Wednesday. That s the largest number " +
					"of shots given in one day since the rollout began and a big jump from the " +
					"previous day, when just under 340,000 doses were given, CBS News reported. " +
					"That number is likely to jump quickly after the federal government on Tuesday " +
					"gave states the OK to vaccinate anyone over 65 and said it would release all " +
					"the doses of vaccine it has available for distribution. Meanwhile, a number " +
					"of states have now opened mass vaccination sites in an effort to get larger " +
					"numbers of people inoculated, CBS News reported.",
			},
			ExpectedSummarization: &nlpcloud.Summarization{
				SummaryText: "Over 951,000 doses were given in the past 24 hours. That's the largest number of shots given in one day since the rollout began. That number is likely to jump quickly after the federal government gave states the OK to vaccinate anyone over 65. A number of states have now opened mass vaccination sites.",
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			summarization, err := client.Summarization(tt.Params)

			assert.Equal(tt.ExpectedSummarization, summarization)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientGeneration(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client             nlpcloud.HTTPClient
		Params             nlpcloud.GenerationParams
		ExpectedGeneration *nlpcloud.Generation
		ExpectedErr        error
	}{
		"nil-client": {
			Client:             nil,
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:             newFakeHTTPClient(``, 0, errFake),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        errFake,
		},
		"bad-request": {
			Client:             newFakeHTTPClient(``, 400, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:             newFakeHTTPClient(``, 401, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:             newFakeHTTPClient(``, 402, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:             newFakeHTTPClient(``, 403, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:             newFakeHTTPClient(``, 404, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:             newFakeHTTPClient(``, 405, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:             newFakeHTTPClient(``, 406, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:             newFakeHTTPClient(``, 413, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:             newFakeHTTPClient(``, 422, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:             newFakeHTTPClient(``, 429, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:             newFakeHTTPClient(``, 500, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:             newFakeHTTPClient(``, 503, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr:        nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:             newFakeHTTPClient(``, 0, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:             newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:             nlpcloud.GenerationParams{},
			ExpectedGeneration: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"generated_text":"GPT-J is a powerful NLP model for text generation. This is the open-source version of GPT-3 by OpenAI. It is the most advanced NLP model created as of today.","nb_generated_tokens":33}`, http.StatusOK, nil),
			Params: nlpcloud.GenerationParams{
				Text:      "GPT-J is a powerful NLP model",
				MinLength: i(10),
				MaxLength: i(50),
			},
			ExpectedGeneration: &nlpcloud.Generation{
				GeneratedText:     "GPT-J is a powerful NLP model for text generation. This is the open-source version of GPT-3 by OpenAI. It is the most advanced NLP model created as of today.",
				NbGeneratedTokens: 33,
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			generation, err := client.Generation(tt.Params)

			assert.Equal(tt.ExpectedGeneration, generation)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientTranslation(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client              nlpcloud.HTTPClient
		Params              nlpcloud.TranslationParams
		ExpectedTranslation *nlpcloud.Translation
		ExpectedErr         error
	}{
		"nil-client": {
			Client:              nil,
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:              newFakeHTTPClient(``, 0, errFake),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         errFake,
		},
		"bad-request": {
			Client:              newFakeHTTPClient(``, 400, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:              newFakeHTTPClient(``, 401, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:              newFakeHTTPClient(``, 402, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:              newFakeHTTPClient(``, 403, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:              newFakeHTTPClient(``, 404, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:              newFakeHTTPClient(``, 405, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:              newFakeHTTPClient(``, 406, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:              newFakeHTTPClient(``, 413, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:              newFakeHTTPClient(``, 422, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:              newFakeHTTPClient(``, 429, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:              newFakeHTTPClient(``, 500, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:              newFakeHTTPClient(``, 503, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr:         nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:              newFakeHTTPClient(``, 0, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:              newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:              nlpcloud.TranslationParams{},
			ExpectedTranslation: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"translation_text":"John Doe travaille pour Microsoft à Seattle depuis 1999."}`, http.StatusOK, nil),
			Params: nlpcloud.TranslationParams{
				Text: "John Doe has been working for Microsoft in Seattle since 1999.",
			},
			ExpectedTranslation: &nlpcloud.Translation{
				TranslationText: "John Doe travaille pour Microsoft à Seattle depuis 1999.",
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			translation, err := client.Translation(tt.Params)

			assert.Equal(tt.ExpectedTranslation, translation)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientLangDetection(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client                nlpcloud.HTTPClient
		Params                nlpcloud.LangDetectionParams
		ExpectedLangDetection *nlpcloud.LangDetection
		ExpectedErr           error
	}{
		"nil-client": {
			Client:                nil,
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:                newFakeHTTPClient(``, 0, errFake),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           errFake,
		},
		"bad-request": {
			Client:                newFakeHTTPClient(``, 400, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:                newFakeHTTPClient(``, 401, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:                newFakeHTTPClient(``, 402, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:                newFakeHTTPClient(``, 403, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:                newFakeHTTPClient(``, 404, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:                newFakeHTTPClient(``, 405, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:                newFakeHTTPClient(``, 406, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:                newFakeHTTPClient(``, 413, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:                newFakeHTTPClient(``, 422, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:                newFakeHTTPClient(``, 429, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:                newFakeHTTPClient(``, 500, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:                newFakeHTTPClient(``, 503, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr:           nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:                newFakeHTTPClient(``, 0, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:                newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:                nlpcloud.LangDetectionParams{},
			ExpectedLangDetection: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"languages":[{"en":0.7142834369645996},{"fr":0.28571521669868466}]}`, http.StatusOK, nil),
			Params: nlpcloud.LangDetectionParams{
				Text: "John Doe has been working for Microsoft in Seattle since 1999. Et il parle aussi un peu français.",
			},
			ExpectedLangDetection: &nlpcloud.LangDetection{
				Languages: []map[string]float64{
					{
						"en": 0.7142834369645996,
					}, {
						"fr": 0.28571521669868466,
					},
				},
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			langDetection, err := client.LangDetection(tt.Params)

			assert.Equal(tt.ExpectedLangDetection, langDetection)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientDependencies(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client               nlpcloud.HTTPClient
		Params               nlpcloud.DependenciesParams
		ExpectedDependencies *nlpcloud.Dependencies
		ExpectedErr          error
	}{
		"nil-client": {
			Client:               nil,
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:               newFakeHTTPClient(``, 0, errFake),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          errFake,
		},
		"bad-request": {
			Client:               newFakeHTTPClient(``, 400, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:               newFakeHTTPClient(``, 401, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:               newFakeHTTPClient(``, 402, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:               newFakeHTTPClient(``, 403, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:               newFakeHTTPClient(``, 404, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:               newFakeHTTPClient(``, 405, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:               newFakeHTTPClient(``, 406, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:               newFakeHTTPClient(``, 413, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:               newFakeHTTPClient(``, 422, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:               newFakeHTTPClient(``, 429, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:               newFakeHTTPClient(``, 500, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:               newFakeHTTPClient(``, 503, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr:          nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:               newFakeHTTPClient(``, 0, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:               newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:               nlpcloud.DependenciesParams{},
			ExpectedDependencies: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"words":[{"text":"John","tag":"NNP"},{"text":"Doe","tag":"NNP"},{"text":"is","tag":"VBZ"},{"text":"a","tag":"DT"},{"text":"Go","tag":"NNP"},{"text":"Developer","tag":"NN"},{"text":"at","tag":"IN"},{"text":"Google","tag":"NNP"}],"arcs":[{"start":0,"end":1,"label":"compound","text":"John","dir":"left"},{"start":1,"end":2,"label":"nsubj","text":"Doe","dir":"left"},{"start":3,"end":5,"label":"det","text":"a","dir":"left"},{"start":4,"end":5,"label":"compound","text":"Go","dir":"left"},{"start":2,"end":5,"label":"attr","text":"Developer","dir":"right"},{"start":5,"end":6,"label":"prep","text":"at","dir":"right"},{"start":6,"end":7,"label":"pobj","text":"Google","dir":"right"}]}`, http.StatusOK, nil),
			Params: nlpcloud.DependenciesParams{
				Text: "John Doe is a Go Developer at Google",
			},
			ExpectedDependencies: &nlpcloud.Dependencies{
				Words: []nlpcloud.Word{
					{
						Text: "John",
						Tag:  "NNP",
					}, {
						Text: "Doe",
						Tag:  "NNP",
					}, {
						Text: "is",
						Tag:  "VBZ",
					}, {
						Text: "a",
						Tag:  "DT",
					}, {
						Text: "Go",
						Tag:  "NNP",
					}, {
						Text: "Developer",
						Tag:  "NN",
					}, {
						Text: "at",
						Tag:  "IN",
					}, {
						Text: "Google",
						Tag:  "NNP",
					},
				},
				Arcs: []nlpcloud.Arc{
					{
						Start: 0,
						End:   1,
						Label: "compound",
						Text:  "John",
						Dir:   "left",
					}, {
						Start: 1,
						End:   2,
						Label: "nsubj",
						Text:  "Doe",
						Dir:   "left",
					}, {
						Start: 3,
						End:   5,
						Label: "det",
						Text:  "a",
						Dir:   "left",
					}, {
						Start: 4,
						End:   5,
						Label: "compound",
						Text:  "Go",
						Dir:   "left",
					}, {
						Start: 2,
						End:   5,
						Label: "attr",
						Text:  "Developer",
						Dir:   "right",
					}, {
						Start: 5,
						End:   6,
						Label: "prep",
						Text:  "at",
						Dir:   "right",
					}, {
						Start: 6,
						End:   7,
						Label: "pobj",
						Text:  "Google",
						Dir:   "right",
					},
				},
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			dependencies, err := client.Dependencies(tt.Params)

			assert.Equal(tt.ExpectedDependencies, dependencies)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientSentenceDependencies(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client                       nlpcloud.HTTPClient
		Params                       nlpcloud.SentenceDependenciesParams
		ExpectedSentenceDependencies *nlpcloud.SentenceDependencies
		ExpectedErr                  error
	}{
		"nil-client": {
			Client:                       nil,
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:                       newFakeHTTPClient(``, 0, errFake),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  errFake,
		},
		"bad-request": {
			Client:                       newFakeHTTPClient(``, 400, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:                       newFakeHTTPClient(``, 401, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:                       newFakeHTTPClient(``, 402, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:                       newFakeHTTPClient(``, 403, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:                       newFakeHTTPClient(``, 404, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:                       newFakeHTTPClient(``, 405, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:                       newFakeHTTPClient(``, 406, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:                       newFakeHTTPClient(``, 413, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:                       newFakeHTTPClient(``, 422, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:                       newFakeHTTPClient(``, 429, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:                       newFakeHTTPClient(``, 500, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:                       newFakeHTTPClient(``, 503, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr:                  nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:                       newFakeHTTPClient(``, 0, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:                       newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:                       nlpcloud.SentenceDependenciesParams{},
			ExpectedSentenceDependencies: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"sentence_dependencies":[{"sentence":"John Doe is a Go Developer at Google.","dependencies":{"words":[{"text":"John","tag":"NNP"},{"text":"Doe","tag":"NNP"},{"text":"is","tag":"VBZ"},{"text":"a","tag":"DT"},{"text":"Go","tag":"NNP"},{"text":"Developer","tag":"NN"},{"text":"at","tag":"IN"},{"text":"Google","tag":"NNP"},{"text":".","tag":"."}],"arcs":[{"start":0,"end":1,"label":"compound","text":"John","dir":"left"},{"start":1,"end":2,"label":"nsubj","text":"Doe","dir":"left"},{"start":3,"end":5,"label":"det","text":"a","dir":"left"},{"start":4,"end":5,"label":"compound","text":"Go","dir":"left"},{"start":2,"end":5,"label":"attr","text":"Developer","dir":"right"},{"start":5,"end":6,"label":"prep","text":"at","dir":"right"},{"start":6,"end":7,"label":"pobj","text":"Google","dir":"right"},{"start":2,"end":8,"label":"punct","text":".","dir":"right"}]}},{"sentence":"Before that, he worked at Microsoft.","dependencies":{"words":[{"text":"Before","tag":"IN"},{"text":"that","tag":"DT"},{"text":",","tag":","},{"text":"he","tag":"PRP"},{"text":"worked","tag":"VBD"},{"text":"at","tag":"IN"},{"text":"Microsoft","tag":"NNP"},{"text":".","tag":"."}],"arcs":[{"start":9,"end":13,"label":"prep","text":"Before","dir":"left"},{"start":9,"end":10,"label":"pobj","text":"that","dir":"right"},{"start":11,"end":13,"label":"punct","text":",","dir":"left"},{"start":12,"end":13,"label":"nsubj","text":"he","dir":"left"},{"start":13,"end":14,"label":"prep","text":"at","dir":"right"},{"start":14,"end":15,"label":"pobj","text":"Microsoft","dir":"right"},{"start":13,"end":16,"label":"punct","text":".","dir":"right"}]}}]}`, http.StatusOK, nil),
			Params: nlpcloud.SentenceDependenciesParams{
				Text: "John Doe is a Go Developer at Google. Before that, he worked at Microsoft.",
			},
			ExpectedSentenceDependencies: &nlpcloud.SentenceDependencies{
				SentenceDependencies: []nlpcloud.SentenceDependency{
					{
						Sentence: "John Doe is a Go Developer at Google.",
						Dependencies: nlpcloud.Dependencies{
							Words: []nlpcloud.Word{
								{
									Text: "John",
									Tag:  "NNP",
								}, {
									Text: "Doe",
									Tag:  "NNP",
								}, {
									Text: "is",
									Tag:  "VBZ",
								}, {
									Text: "a",
									Tag:  "DT",
								}, {
									Text: "Go",
									Tag:  "NNP",
								}, {
									Text: "Developer",
									Tag:  "NN",
								}, {
									Text: "at",
									Tag:  "IN",
								}, {
									Text: "Google",
									Tag:  "NNP",
								}, {
									Text: ".",
									Tag:  ".",
								},
							},
							Arcs: []nlpcloud.Arc{
								{
									Start: 0,
									End:   1,
									Label: "compound",
									Text:  "John",
									Dir:   "left",
								}, {
									Start: 1,
									End:   2,
									Label: "nsubj",
									Text:  "Doe",
									Dir:   "left",
								}, {
									Start: 3,
									End:   5,
									Label: "det",
									Text:  "a",
									Dir:   "left",
								}, {
									Start: 4,
									End:   5,
									Label: "compound",
									Text:  "Go",
									Dir:   "left",
								}, {
									Start: 2,
									End:   5,
									Label: "attr",
									Text:  "Developer",
									Dir:   "right",
								}, {
									Start: 5,
									End:   6,
									Label: "prep",
									Text:  "at",
									Dir:   "right",
								}, {
									Start: 6,
									End:   7,
									Label: "pobj",
									Text:  "Google",
									Dir:   "right",
								}, {
									Start: 2,
									End:   8,
									Label: "punct",
									Text:  ".",
									Dir:   "right",
								},
							},
						},
					}, {
						Sentence: "Before that, he worked at Microsoft.",
						Dependencies: nlpcloud.Dependencies{
							Words: []nlpcloud.Word{
								{
									Text: "Before",
									Tag:  "IN",
								}, {
									Text: "that",
									Tag:  "DT",
								}, {
									Text: ",",
									Tag:  ",",
								}, {
									Text: "he",
									Tag:  "PRP",
								}, {
									Text: "worked",
									Tag:  "VBD",
								}, {
									Text: "at",
									Tag:  "IN",
								}, {
									Text: "Microsoft",
									Tag:  "NNP",
								}, {
									Text: ".",
									Tag:  ".",
								},
							},
							Arcs: []nlpcloud.Arc{
								{
									Start: 9,
									End:   13,
									Label: "prep",
									Text:  "Before",
									Dir:   "left",
								}, {
									Start: 9,
									End:   10,
									Label: "pobj",
									Text:  "that",
									Dir:   "right",
								}, {
									Start: 11,
									End:   13,
									Label: "punct",
									Text:  ",",
									Dir:   "left",
								}, {
									Start: 12,
									End:   13,
									Label: "nsubj",
									Text:  "he",
									Dir:   "left",
								}, {
									Start: 13,
									End:   14,
									Label: "prep",
									Text:  "at",
									Dir:   "right",
								}, {
									Start: 14,
									End:   15,
									Label: "pobj",
									Text:  "Microsoft",
									Dir:   "right",
								}, {
									Start: 13,
									End:   16,
									Label: "punct",
									Text:  ".",
									Dir:   "right",
								},
							},
						},
					},
				},
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			sentenceDependencies, err := client.SentenceDependencies(tt.Params)

			assert.Equal(tt.ExpectedSentenceDependencies, sentenceDependencies)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientTokens(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client         nlpcloud.HTTPClient
		Params         nlpcloud.TokensParams
		ExpectedTokens *nlpcloud.Tokens
		ExpectedErr    error
	}{
		"nil-client": {
			Client:         nil,
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:         newFakeHTTPClient(``, 0, errFake),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    errFake,
		},
		"bad-request": {
			Client:         newFakeHTTPClient(``, 400, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:         newFakeHTTPClient(``, 401, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:         newFakeHTTPClient(``, 402, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:         newFakeHTTPClient(``, 403, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:         newFakeHTTPClient(``, 404, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:         newFakeHTTPClient(``, 405, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:         newFakeHTTPClient(``, 406, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:         newFakeHTTPClient(``, 413, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:         newFakeHTTPClient(``, 422, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:         newFakeHTTPClient(``, 429, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:         newFakeHTTPClient(``, 500, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:         newFakeHTTPClient(``, 503, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr:    nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:         newFakeHTTPClient(``, 0, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:         newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:         nlpcloud.TokensParams{},
			ExpectedTokens: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"tokens":[{"start":0,"end":4,"index":1,"text":"John","lemma":"John","ws_after":true},{"start":5,"end":7,"index":2,"text":"is","lemma":"be","ws_after":true},{"start":8,"end":9,"index":3,"text":"a","lemma":"a","ws_after":true},{"start":10,"end":12,"index":4,"text":"Go","lemma":"Go","ws_after":true},{"start":13,"end":22,"index":5,"text":"Developer","lemma":"developer","ws_after":true},{"start":23,"end":25,"index":6,"text":"at","lemma":"at","ws_after":true},{"start":26,"end":32,"index":7,"text":"Google","lemma":"Google","ws_after":false},{"start":32,"end":33,"index":8,"text":".","lemma":".","ws_after":false}]}`, http.StatusOK, nil),
			Params: nlpcloud.TokensParams{
				Text: "John is a Go Developer at Google.",
			},
			ExpectedTokens: &nlpcloud.Tokens{
				Tokens: []nlpcloud.Token{
					{
						Start:   0,
						End:     4,
						Index:   1,
						Text:    "John",
						Lemma:   "John",
						WSAfter: true,
					}, {
						Start:   5,
						End:     7,
						Index:   2,
						Text:    "is",
						Lemma:   "be",
						WSAfter: true,
					}, {
						Start:   8,
						End:     9,
						Index:   3,
						Text:    "a",
						Lemma:   "a",
						WSAfter: true,
					}, {
						Start:   10,
						End:     12,
						Index:   4,
						Text:    "Go",
						Lemma:   "Go",
						WSAfter: true,
					}, {
						Start:   13,
						End:     22,
						Index:   5,
						Text:    "Developer",
						Lemma:   "developer",
						WSAfter: true,
					}, {
						Start:   23,
						End:     25,
						Index:   6,
						Text:    "at",
						Lemma:   "at",
						WSAfter: true,
					}, {
						Start:   26,
						End:     32,
						Index:   7,
						Text:    "Google",
						Lemma:   "Google",
						WSAfter: false,
					}, {
						Start:   32,
						End:     33,
						Index:   8,
						Text:    ".",
						Lemma:   ".",
						WSAfter: false,
					},
				},
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			Tokens, err := client.Tokens(tt.Params)

			assert.Equal(tt.ExpectedTokens, Tokens)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func TestClientLibVersions(t *testing.T) {
	t.Parallel()

	var tests = map[string]struct {
		Client              nlpcloud.HTTPClient
		ExpectedLibVersions *nlpcloud.LibVersions
		ExpectedErr         error
	}{
		"nil-client": {
			Client:              nil,
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrNilClient,
		},
		"failing-client": {
			Client:              newFakeHTTPClient(``, 0, errFake),
			ExpectedLibVersions: nil,
			ExpectedErr:         errFake,
		},
		"bad-request": {
			Client:              newFakeHTTPClient(``, 400, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrBadRequest,
		},
		"unauthorized": {
			Client:              newFakeHTTPClient(``, 401, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrUnauthorized,
		},
		"payment-required": {
			Client:              newFakeHTTPClient(``, 402, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrPaymentRequired,
		},
		"forbidden": {
			Client:              newFakeHTTPClient(``, 403, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrForbidden,
		},
		"not-found": {
			Client:              newFakeHTTPClient(``, 404, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrNotFound,
		},
		"method-not-allowed": {
			Client:              newFakeHTTPClient(``, 405, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrMethodNotAllowed,
		},
		"not-acceptable": {
			Client:              newFakeHTTPClient(``, 406, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrNotAcceptable,
		},
		"request-entity-too-large": {
			Client:              newFakeHTTPClient(``, 413, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrRequestEntityTooLarge,
		},
		"unprocessable-entity": {
			Client:              newFakeHTTPClient(``, 422, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrUnprocessableEntity,
		},
		"too-many-requests": {
			Client:              newFakeHTTPClient(``, 429, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrTooManyRequests,
		},
		"internal-server-error": {
			Client:              newFakeHTTPClient(``, 500, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrInternalServerError,
		},
		"service-unavailable": {
			Client:              newFakeHTTPClient(``, 503, nil),
			ExpectedLibVersions: nil,
			ExpectedErr:         nlpcloud.ErrServiceUnavailable,
		},
		"unexpected-statuscode": {
			Client:              newFakeHTTPClient(``, 0, nil),
			ExpectedLibVersions: nil,
			ExpectedErr: &nlpcloud.ErrUnexpectedStatus{
				Body:       []byte(``),
				StatusCode: 0,
			},
		},
		"failing-unmarshal": {
			Client:              newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			ExpectedLibVersions: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"pytorch":"1.7.1","transformers":"4.3.2"}`, http.StatusOK, nil),
			ExpectedLibVersions: &nlpcloud.LibVersions{
				"pytorch":      "1.7.1",
				"transformers": "4.3.2",
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			client := nlpcloud.NewClient(tt.Client, "fake-model", "fake-token", true, "")
			assert := assert.New(t)

			libVersions, err := client.LibVersions()

			assert.Equal(tt.ExpectedLibVersions, libVersions)
			checkErr(tt.ExpectedErr, err, assert)
		})
	}
}

func b(b bool) *bool {
	return &b
}

func i(i int) *int {
	return &i
}
