# Go Client For spaCy Cloud

This is a Go client for the spaCy Cloud API: https://docs.spacycloud.io

spaCy Cloud serves all the spaCy pre-trained models, and your own custom models, through a RESTful API, so it's easy for you to use them in production.

If you face an issue, don't hesitate to raise it as a Github issue. Thanks!

## Installation

Install using Go `get`.

```shell
go get -u github.com/spacycloud/spacycloud-go
```

## Examples

Here is a full example that uses the `en_core_web_sm` model, with a fake token:

```go
package main

import "github.com/spacycloud/spacycloud-go"

func main() {
    client := spacycloud.NewClient("en_core_web_sm", "4eC39HqLyjWDarjtT1zdp7dc")
    client.Entities("John Doe is a Go Developer at Google")
}
```

And a full example that uses your own custom model `7894`:

```go
package main

import "github.com/spacycloud/spacycloud-go"

func main() {
    client := spacycloud.NewClient("custom_model/7894", "4eC39HqLyjWDarjtT1zdp7dc")
    client.Entities("John Doe is a Go Developer at Google")
}
```

An `Entities` struct is returned.

## Usage

### Client Initialization

Pass the spaCy model you want to use and the spaCy cloud token to the client during initialization.

The spaCy model can either be a spaCy pretrained model like `en_core_web_sm`, `fr_core_news_lg`... but also one of your custom spaCy models using `custom_model/<model id>` (e.g. `custom_model/2568`).

Your token can be retrieved from your [spaCy Cloud dashboard](https://spacycloud.io/home/token).

```go
package main

import "github.com/spacycloud/spacycloud-go"

func main() {
    client := spacycloud.NewClient("<model>", "<token>")
}
```

### Entities Endpoint

Call the `Entities()` method and pass the text you want to perform named entity recognition (NER) on.

```go
client.Entities("<Your block of text>")
```

The above command returns an `Entities` struct.


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

Call the `LibVersions()` method to know the versions of the libraries used behind the hood with the model (for example the spaCy version used).

```go
client.LibVersions()
```

The above command returns a `LibVersion` struct.
