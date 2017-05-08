# Funcy (\'fəŋ-kē\)

Funcy (\'fəŋ-kē\) is both a CLI tool and set of polyglot API's to simplify management of Google Cloud functions in various languages.  Simplicity is divinity...

## Acknowledgements

This work is derived from Kelsey Hightower's (Google) excellent work and insights...

* https://github.com/kelseyhightower/google-cloud-functions-go

Some concepts are heavily influenced by the excellent Apex project...

* https://apex.run

## Supported Languages

1. Node.js (Javascript)
2. Golang
   * No requirement for Go 1.8 and plugins
   * No requirement to build using Docker on OSX (since plugins are Linux only)
3. _tbd_

## API

* Simple callback handlers for all trigger types (HTTP, Storage, PubSub)

## CLI

* The 'funcy' cli tool simplifies cloud function management
* Builds, packages, deploys, and tests
* Uses the Google Cloud API's for all cloud operations...
  1. https://cloud.google.com/functions/docs/reference/rpc/
  2. https://cloud.google.com/storage/docs/json_api/

## Installation

Install Funcy...

```bash
$ go get github.com/markuscraig/funcy
$ go install github.com/markuscraig/funcy
```

Install the Google Cloud SDK...

* https://cloud.google.com/sdk/downloads

```bash
# update the 'gcloud' sdk
$ gcloud components update
```

## Usage

```bash
$ funcy help

TBD
```

## Examples

Create a new project...

```bash
# Create a project named 'simple'. This creates the 'funcy.yaml' project file.
$ funcy init simple [--bucket <NAME>]
```

Create an HTTP handler function...

```bash
$ funcy create howdy --type http
```

Create a 'background' handler function with storage trigger...

```bash
$ funcy create howdy --type bg --trigger storage --bucket myBucket
```

List all functions...

```bash
$ funcy list
```

Deploy / Update all of the functions...

```bash
# deploys all functions (updates if existing)
$ funcy deploy
```

Deploy / Update a set of functions...

```bash
# deploys the 'howdy' and 'ping' functions (updates if existing)
$ funcy deploy hello ping
```

## Why?

Some things are better together...

1. Google Cloud Functions
2. Go
3. Declarative automation

Made with :green_heart: in Campbell, CA
