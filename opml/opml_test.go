// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opml

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNewOPMLFromURL(t *testing.T) {
	testNewOPMLFromURLSuccess(t)
	testNewOPMLFromURLFailure(t)
}

func TestNewOPMLFromFile(t *testing.T) {
	testNewOPMLFromFileSuccess(t)
	testNewOPMLFromFileFailure(t)
}

func testNewOPMLFromURLSuccess(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile(
			os.Getenv("GOPATH") + "/src/github.com/gilliek/go-opml/testdata/feeds.xml")
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(w, string(b))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	doc, err := NewOPMLFromURL(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	testDoc(t, doc)
}

func testNewOPMLFromURLFailure(t *testing.T) {
	_, err := NewOPMLFromURL("1.2.3.4")
	if err == nil {
		t.Error("Expected failure!")
	}
}

func testNewOPMLFromFileSuccess(t *testing.T) {
	doc, err := NewOPMLFromFile(
		os.Getenv("GOPATH") + "/src/github.com/gilliek/go-opml/testdata/feeds.xml")
	if err != nil {
		t.Fatal(err)
	}

	testDoc(t, doc)
}

func testNewOPMLFromFileFailure(t *testing.T) {
	_, err := NewOPMLFromFile(
		os.Getenv("GOPATH") + "/src/github.com/gilliek/go-opml/testdata/does_not_exist.xml")
	if err == nil {
		t.Error("Expected failure!")
	}
}

func testDoc(t *testing.T, doc *OPML) {
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
