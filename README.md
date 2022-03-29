# Go Client For NLP Cloud

[![reference](https://godoc.org/github.com/nlpcloud/nlpcloud-go/v5?status.svg=)](https://pkg.go.dev/github.com/nlpcloud/nlpcloud-go)
[![go report](https://goreportcard.com/badge/github.com/nlpcloud/nlpcloud-go)](https://goreportcard.com/report/github.com/nlpcloud/nlpcloud-go)

This is a Go client for the [NLP Cloud](https://nlpcloud.io) API. See the [documentation](https://docs.nlpcloud.io) for more details.

NLP Cloud serves high performance pre-trained or custom models for NER, sentiment-analysis, classification, summarization, paraphrasing, intent classification, product description and ad generation, chatbot, grammar and spelling correction, keywords and keyphrases extraction, text generation, question answering, machine translation, language detection, semantic similarity, tokenization, POS tagging, embeddings, and dependency parsing. It is ready for production, served through a REST API.

You can either use the NLP Cloud pre-trained models, fine-tune your own models, or deploy your own models.

If you face an issue, don't hesitate to raise it as a Github issue. Thanks!

## Installation

Install using `go install`.

```shell
go install github.com/nlpcloud/nlpcloud-go
```

## Examples

Here is a full example that summarizes a text using Facebook's Bart Large CNN model, with a fake token:

```go
package main

import (
    "net/http"
    
    "github.com/nlpcloud/nlpcloud-go"
)

func main() {
    client := nlpcloud.NewClient(&http.Client{}, "bart-large-cnn", "4eC39HqLyjWDarjtT1zdp7dc", false, "")
    summarization, err := client.Summarization(nlpcloud.SummarizationParams{Text: `One month after
    the United States began what has become a troubled rollout of a national COVID vaccination
    campaign, the effort is finally gathering real steam. Close to a million doses -- over 951,000, to be more exact -- 
    made their way into the arms of Americans in the past 24 hours, the U.S. Centers 
    for Disease Control and Prevention reported Wednesday. That s the largest number 
    of shots given in one day since the rollout began and a big jump from the 
    previous day, when just under 340,000 doses were given, CBS News reported. 
    That number is likely to jump quickly after the federal government on Tuesday 
    gave states the OK to vaccinate anyone over 65 and said it would release all 
    the doses of vaccine it has available for distribution. Meanwhile, a number 
    of states have now opened mass vaccination sites in an effort to get larger 
    numbers of people inoculated, CBS News reported.`})
    ...
}
```

Here is a full example that does the same thing, but on a GPU:

```go
package main

import (
    "net/http"
    
    "github.com/nlpcloud/nlpcloud-go"
)

func main() {
    client := nlpcloud.NewClient(&http.Client{}, "bart-large-cnn", "4eC39HqLyjWDarjtT1zdp7dc", true, "")
    summarization, err := client.Summarization(nlpcloud.SummarizationParams{Text: `One month after
    the United States began what has become a troubled rollout of a national COVID vaccination
    campaign, the effort is finally gathering real steam. Close to a million doses -- over 951,000, to be more exact -- 
    made their way into the arms of Americans in the past 24 hours, the U.S. Centers 
    for Disease Control and Prevention reported Wednesday. That s the largest number 
    of shots given in one day since the rollout began and a big jump from the 
    previous day, when just under 340,000 doses were given, CBS News reported. 
    That number is likely to jump quickly after the federal government on Tuesday 
    gave states the OK to vaccinate anyone over 65 and said it would release all 
    the doses of vaccine it has available for distribution. Meanwhile, a number 
    of states have now opened mass vaccination sites in an effort to get larger 
    numbers of people inoculated, CBS News reported.`})
    ...
}
```

Here is a full example that does the same thing, but on a French text:

```go
package main

import (
    "net/http"
    
    "github.com/nlpcloud/nlpcloud-go"
)

func main() {
    client := nlpcloud.NewClient(&http.Client{}, "bart-large-cnn", "4eC39HqLyjWDarjtT1zdp7dc", true, "fr")
    summarization, err := client.Summarization(nlpcloud.SummarizationParams{Text: `Sur des images aériennes, 
    prises la veille par un vol de surveillance de la Nouvelle-Zélande, la côte d’une île est bordée 
    d’arbres passés du vert au gris sous l’effet des retombées volcaniques. On y voit aussi des immeubles
    endommagés côtoyer des bâtiments intacts. « D’après le peu d’informations
    dont nous disposons, l’échelle de la dévastation pourrait être immense, 
    spécialement pour les îles les plus isolées », avait déclaré plus tôt 
    Katie Greenwood, de la Fédération internationale des sociétés de la Croix-Rouge.
    Selon l’Organisation mondiale de la santé (OMS), une centaine de maisons ont
    été endommagées, dont cinquante ont été détruites sur l’île principale de
    Tonga, Tongatapu. La police locale, citée par les autorités néo-zélandaises,
    a également fait état de deux morts, dont une Britannique âgée de 50 ans,
    Angela Glover, emportée par le tsunami après avoir essayé de sauver les chiens
    de son refuge, selon sa famille.`})
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
