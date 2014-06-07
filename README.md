# go-opml

 [![Build Status](https://travis-ci.org/gilliek/go-opml.png?branch=master)](https://travis-ci.org/gilliek/go-opml)

go-opml aims to be a Go package for parsing OPML files.

## Installation

```go get github.com/gilliek/go-opml/opml```

## Usage

Parse OPML from file:

```go
package main

import (
	"github.com/gilliek/go-opml/opml"
	"log"
)

func main() {
	doc, err := opml.NewOPMLFromFile("path/to/file.xml")
	if err != nil {
		log.Fatal(err)
	}

    //...
}
```


Parse OPML from URL:

```go
package main

import (
	"github.com/gilliek/go-opml/opml"
	"log"
)

func main() {
	doc, err := opml.NewOPMLFromURL("http://www.example.com/file.xml")
	if err != nil {
		log.Fatal(err)
	}

    //...
}
```

## License

BSD 3-clauses
