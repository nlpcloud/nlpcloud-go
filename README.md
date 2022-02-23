# Go Client For NLP Cloud

[![reference](https://godoc.org/github.com/nlpcloud/nlpcloud-go/v5?status.svg=)](https://pkg.go.dev/github.com/nlpcloud/nlpcloud-go)
[![go report](https://goreportcard.com/badge/github.com/nlpcloud/nlpcloud-go)](https://goreportcard.com/report/github.com/nlpcloud/nlpcloud-go)

This is a Go client for the [NLP Cloud](https://nlpcloud.io) API. See the [documentation](https://docs.nlpcloud.io) for more details.

NLP Cloud serves high performance pre-trained for NER, sentiment-analysis, classification, summarization, text generation, question answering, machine translation, language detection, tokenization, lemmatization, POS tagging, and dependency parsing. It is ready for production, served through a REST API.

You can either use the NLP Cloud pre-trained models, fine-tune your own models, or deploy your own models.

If you face an issue, don't hesitate to raise it as a Github issue. Thanks!

## Installation

Install using `go install`.

```shell
go install github.com/nlpcloud/nlpcloud-go
```

## Examples

Here is a full example that performs Named Entity Recognition (NER) using spaCy's `en_core_web_lg` model, with a fake token:

```go
package main

import (
    "net/http"
    
    "github.com/nlpcloud/nlpcloud-go"
)

func main() {
    client := nlpcloud.NewClient(&http.Client{}, "en_core_web_lg", "<your token>", false, "")
    entities, err := client.Entities(nlpcloud.EntitiesParams{Text: "John Doe is a Go Developer at Google"})
    ...
}
```

And a full example that uses your own custom model `7894`:

```go
package main

import (
    "net/http"

    "github.com/nlpcloud/nlpcloud-go"
)

func main() {
    client := nlpcloud.NewClient(&http.Client{}, "custom_model/7894", "<your token>", false, "")
    entities, err := client.Entities(nlpcloud.EntitiesParams{Text: "John Doe is a Go Developer at Google"})
    ...
}
```

## Usage

### Client Initialization

While it uses a HTTP REST API, you'll have to pass an instance that implements interface `HTTPClient`.
It works with a `*http.Client`.

Pass the model you want to use and the NLP Cloud token to the client during initialization.

The model can either be a pre-trained model like `en_core_web_lg`, `bart-large-mnli`, ... but also one of your custom models using `custom_model/<model id>` (e.g. `custom_model/2568`).

Your token can be retrieved from your [NLP Cloud dashboard](https://nlpcloud.io/home/token).

```go
package main

import (
    "net/http"
    
    "github.com/nlpcloud/nlpcloud-go"
)

func main() {
    client := nlpcloud.NewClient(&http.Client, "<model>", "<token>", false, "<language>")
    ...
}
```

If you want to use a GPU, set the 4th parameter as `true`.

If you want to use the multilingual add-on in order to process non-English texts, set your language code in the 5th parameter. For example, if you want to process French text, you should set the 5th parameter as `"fr"`.

### API endpoint

Depending on the API endpoint, it may have parameters (only `LibVersions` does not follow this rule).
In case it has parameters, you can call an endpoint using the following:

```go
res, err := nlpcloud.TheAPIEndpoint(params TheAPIEndpointParams)
```

## Tests

This Go wrapper aims to support the NLP Cloud API, being automatically full-tested and validated.
It's achieved through unit and integration tests that you can run to make sure everything works as expected.

### Unit tests

To run them, you can execute the following:

```bash
go test ./... -v -count=1 -cover
```

### Integration tests

To run them, you can execute the following:

```bash
API_TOKEN=<you_api_token> go test -tags=integration -v ./internal/integration/...
```

Notice it needs a valid API token that has access to all the NLP Cloud models.
