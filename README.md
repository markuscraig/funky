# Funky

Funky is both a CLI tool and API to simplify Golang-based Google Cloud functions.  Simplicity is divinity...

1. API
   * Simple callback handlers for all trigger types (HTTP, Storage, PubSub)
   * No requirement for Go 1.8 and plugins
   * No requirement to build using Docker on OSX (since plugins are Linux only)
1. CLI
   * The 'funky' cli tool simplifies Golang cloud function management
   * Builds, packages, deploys, and tests
   * Uses Google's own 'gcloud' SDK tool for all cloud operations

## Acknowledgements

This work is derived from Kelsey Hightower's (Google) excellent work and insights...

* https://github.com/kelseyhightower/google-cloud-functions-go

## Installation

Install the Google Cloud SDK...

* https://cloud.google.com/sdk/downloads

```bash
# update the 'gcloud' sdk
$ gcloud components update
```

Install Funky...

```bash
$ go get github.com/markuscraig/funky
$ go install github.com/markuscraig/funky
```

## Usage

```bash
$ funky help

TBD
```

## Examples

Create a new project...

```bash
# Create a project named 'simple'. This creates the 'funky.yaml' project file.
$ funky init simple [--bucket <NAME>]
```

Create an HTTP handler function...

```bash
$ funky create howdy --type http
```

Create a 'background' handler function with storage trigger...

```bash
$ funky create howdy --type bg --trigger storage --bucket myBucket
```

List all functions...

```bash
$ funky list
```

Deploy / Update all of the functions...

```bash
# deploys all functions (updates if existing)
$ funky deploy
```

Deploy / Update a set of functions...

```bash
# deploys the 'howdy' and 'ping' functions (updates if existing)
$ funky deploy hello ping
```

## Why?

Some things are better together...

1. Google Cloud Functions
2. Go
3. Declarative automation

Made with :green_heart: in Campbell, CA
