# Go Client For NLP Cloud

This is a Go client for the NLP Cloud API: https://docs.nlpcloud.io

NLP Cloud serves high performance pre-trained for NER, sentiment-analysis, classification, summarization, text generation, question answering, machine translation, language detection, tokenization, POS tagging, and dependency parsing. It is ready for production, served through a REST API.

Pre-trained models are the spaCy models and some transformers-based models from Hugging Face. You can also deploy your own transformers-based models, or spaCy models.

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
    "http"
    
    "github.com/nlpcloud/nlpcloud-go"
)

func main() {
    client := nlpcloud.NewClient(&http.Client{}, "en_core_web_lg", "fake-token", false)
    entities, err := client.Entities(nlpcloud.EntitiesParams{Text: "John Doe is a Go Developer at Google"})
    ...
}
```

And a full example that uses your own custom model `7894`:

```go
package main

import (
    "http"

    "github.com/nlpcloud/nlpcloud-go"
)

func main() {
    client := nlpcloud.NewClient(&http.Client{}, "custom_model/7894", "fake-token", false)
    entities, err := client.Entities(nlpcloud.EntitiesParams{Text: "John Doe is a Go Developer at Google"})
    ...
}
```

## Usage

### Client Initialization

While it uses a HTTP REST API, you'll have to pass an instance that implements interface `HTTPClient`.
It works with a `*http.Client`.

Pass the model you want to use and the NLP Cloud token to the client during initialization.

The model can either be a pretrained model like `en_core_web_lg`, `bart-large-mnli`... but also one of your custom transformers-based models, or spaCy models, using `custom_model/<model id>` (e.g. `custom_model/2568`).

Your token can be retrieved from your [NLP Cloud dashboard](https://nlpcloud.io/home/token).

```go
package main

import (
    "http"
    
    "github.com/nlpcloud/nlpcloud-go"
)

func main() {
    client := nlpcloud.NewClient(&http.Client, "<model>", "<token>", false)
    ...
}
```

If you want to use a GPU, set the 4th parameter as `true`.

### API endpoint

Depending on the API endpoint, it may have parameters (only `LibVersions` does not follow this rule).
In the case it has some, you can call an enpoint using the following.

```go
res, err := nlpcloud.TheAPIEndpoint(params TheAPIEndpointParams)
```

## Tests

This Go wrapper aims to support the NLP Cloud API, being automatically full-tested and validated.
It's achieved through unit and integration tests that you can run to make sure everything works as expected.

### Unit tests

To run them, you can execute the following.
```bash
go test ./... -v -count=1 -cover
```

### Integration tests

To run them, you can execute the following.
```bash
API_TOKEN=<you_api_token> go test -tags=integration -v ./internal/integration/...
```

Notice it needs a valid API token that has access to the models and functionalities documented.
