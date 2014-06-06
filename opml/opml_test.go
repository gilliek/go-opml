// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opml

import (
	"os"
	"testing"
)

func TestNewOPMLFromFile(t *testing.T) {
	testSuccess(t)
	testFailure(t)
}

func testSuccess(t *testing.T) {
	doc, err := NewOPMLFromFile(
		os.Getenv("GOPATH") + "/src/github.com/gilliek/go-opml/testdata/feeds.xml")
	if err != nil {
		t.Fatal(err)
	}

	version := doc.Root.Version
	if version != "1.0" {
		t.Errorf("Wrong OPML version: expected '1.0', found '%s'", version)
	}

	title := doc.Root.Head.Title
	if title != "Foobar" {
		t.Errorf("Wrong title version: expected 'Foobar', found '%s'", title)
	}

	outlines := doc.Outlines()
	if len(outlines) != 1 {
		t.Fatalf("Invalid number of outlines: expected 1, found %d", len(outlines))
	}

	if outlines[0].Text != "foo" {
		t.Errorf("Wrong outline text: expected 'foo', found '%s'", outlines[0].Text)
	}

	if outlines[0].Title != "bar" {
		t.Errorf("Wrong outline title: expected 'foo', found '%s'", outlines[0].Title)
	}

	if outlines[0].Type != "rss" {
		t.Errorf("Wrong outline type: expected 'rss', found '%s'", outlines[0].Type)
	}

	if outlines[0].XMLURL != "http://www.gilliek.ch/feeds" {
		t.Errorf("Wrong outline XML URL: expected 'http://www.gilliek.ch/feeds', found '%s'",
			outlines[0].XMLURL)
	}

	if outlines[0].HTMLURL != "http://www.gilliek.ch" {
		t.Errorf("Wrong outline HTML URL: expected 'http://www.gilliek.ch', found '%s'",
			outlines[0].HTMLURL)
	}

}

func testFailure(t *testing.T) {
	_, err := NewOPMLFromFile(
		os.Getenv("GOPATH") + "/src/github.com/gilliek/go-opml/testdata/does_not_exist.xml")
	if err == nil {
		t.Error("Expected failure!")
	}
}
