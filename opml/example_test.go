// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opml_test

import (
	"github.com/gilliek/go-opml/opml"
	"log"
)

func ExampleNewOPMLFromFile() {
	doc, err := opml.NewOPMLFromFile("path/to/file.xml")
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleNewOPMLParserFromURL() {
	doc, err := opml.NewOPMLFromURL("http://www.example.com/file.xml")
	if err != nil {
		log.Fatal(err)
	}
}
