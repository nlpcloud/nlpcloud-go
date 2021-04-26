# Go Client For NLP Cloud

This is a Go client for the NLP Cloud API: https://docs.nlpcloud.io

NLP Cloud serves high performance pre-trained models for NER, sentiment-analysis, classification, summarization, question answering, and POS tagging, ready for production, served through a REST API. 

Pre-trained models are the spaCy models and some transformers-based models from Hugging Face. You can also deploy your own transformers-based models, or spaCy models.

If you face an issue, don't hesitate to raise it as a Github issue. Thanks!

## Installation

Install using `go get`.

```shell
go get -u github.com/nlpcloud/nlpcloud-go
```

## Examples

Here is a full example that performs Named Entity Recognition (NER) using spaCy's `en_core_web_lg` model, with a fake token:

```go
package main

import "github.com/nlpcloud/nlpcloud-go"

func main() {
    client := nlpcloud.NewClient("en_core_web_lg", "4eC39HqLyjWDarjtT1zdp7dc")
    client.Entities("John Doe is a Go Developer at Google")
}
```

And a full example that uses your own custom model `7894`:

```go
package main

import "github.com/nlpcloud/nlpcloud-go"

func main() {
    client := nlpcloud.NewClient("custom_model/7894", "4eC39HqLyjWDarjtT1zdp7dc")
    client.Entities("John Doe is a Go Developer at Google")
}
```

An `Entities` struct is returned.

## Usage

### Client Initialization

Pass the model you want to use and the NLP Cloud token to the client during initialization.

The model can either be a pretrained model like `en_core_web_lg`, `bart-large-mnli`... but also one of your custom transformers-based models, or spaCy models, using `custom_model/<model id>` (e.g. `custom_model/2568`).

Your token can be retrieved from your [NLP Cloud dashboard](https://nlpcloud.io/home/token).

```go
package main

import "github.com/nlpcloud/nlpcloud-go"

func main() {
    client := nlpcloud.NewClient("<model>", "<token>")
}
```

### Entities Endpoint

Call the `Entities()` method and pass the text you want to perform named entity recognition (NER) on.

```go
client.Entities("<Your block of text>")
```

The above command returns an `Entities` struct.

### Classification Endpoint

Call the `Classification()` method and pass 3 arguments:

1. The text you want to classify, as a string
1. The candidate labels for your text, as a slice of strings
1. Whether the classification should be multi-class or not, as a boolean

```go
client.Classification("<Your block of text>", []string{"label 1", "label 2", "..."}, true|false)
```

The above command returns a `Classification` struct.

### Sentiment Analysis Endpoint

Call the `Sentiment()` method and pass the text you want to analyze the sentiment of:

```go
client.Sentiment("<Your block of text>")
```

The above command returns a `Sentiment` struct.

### Question Answering Endpoint

Call the `Question()` method and pass the following:

1. A context that the model will use to try to answer your question
1. Your question

```go
client.Question("<Your context>", "<Your question>")
```

The above command returns an `Question` struct.

### Summarization Endpoint

Call the `Summarization()` method and pass the text you want to summarize.

**Note that your block of text should not exceed 1024 words, otherwise you will get an error. Also note that this model works best for blocks of text between 56 and 142 words.**

```go
client.Summarization("<Your text to summarize>")
```

The above command returns a `Summarization` struct.

### Dependencies Endpoint

Call the `Dependencies()` method and pass the text you want to perform part of speech tagging (POS) + arcs on.

```go
client.Dependencies("<Your block of text>")
```

The above command returns a `Dependencies` struct.

### Sentence Dependencies Endpoint

Call the `DentenceDependencies()` method and pass a block of text made up of several sentencies you want to perform POS + arcs on.

```go
client.SentenceDependencies("<Your block of text>")
```

The above command returns a `SentenceDependencies` struct.

### Library Versions Endpoint

Call the `LibVersions()` method to know the versions of the libraries used behind the hood with the model (for example the PyTorch, TensorFlow, or spaCy version used).

```go
client.LibVersions()
```

The above command returns a `LibVersion` struct.
