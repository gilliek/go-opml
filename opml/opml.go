// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package opml provides all the required structures and functions for parsing
OPML files, as defined by the specification of the OPML format:

	[OPML 1.0] http://dev.opml.org/spec1.html
	[OPML 2.0] http://dev.opml.org/spec2.html

It is able to parse both, OPML 1.0 and OPML 2.0, files.
*/
package opml

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// OPMLNode is the root node of an OPML document. It only has a single required
// attribute: the version.
type OPMLNode struct {
	Version string   `xml:"version,attr"`
	Head    HeadNode `xml:"head"`
	Body    BodyNode `xml:"body"`
}

// HeadNode holds some meta information about the document.
type HeadNode struct {
	Title           string `xml:"title"`
	DateCreated     string `xml:"dateCreated,attr"`
	DateModified    string `xml:"dateModified,attr"`
	OwnerName       string `xml:"ownerName,attr"`
	OwnerEmail      string `xml:"ownerEmail,attr"`
	OwnerID         string `xml:"ownerId,attr"`
	Docs            string `xml:"docs,attr"`
	ExpansionState  string `xml:"expansionState,attr"`
	VertScrollState string `xml:"vertScrollState,attr"`
	WindowTop       string `xml:"windowTop,attr"`
	WindowBottom    string `xml:"windowBottom,attr"`
	WindowLeft      string `xml:"windowLeft,attr"`
	WindowRight     string `xml:"windowRight,attr"`
}

// BodyNode is the parent structure of all outlines.
type BodyNode struct {
	Outlines []OutlineNode `xml:"outline"`
}

// OutlineNode holds all information about an outline.
type OutlineNode struct {
	Outlines     []OutlineNode `xml:"outline"`
	Text         string        `xml:"text,attr"`
	Type         string        `xml:"type,attr"`
	IsComment    string        `xml:"isComment,attr"`
	IsBreakpoint string        `xml:"isBreakpoint,attr"`
	Created      string        `xml:"created,attr"`
	Category     string        `xml:"category,attr"`
	XMLURL       string        `xml:"xmlUrl,attr"`
	HTMLURL      string        `xml:"htmlUrl,attr"`
	Language     string        `xml:"language,attr"`
	Title        string        `xml:"title,attr"`
	Version      string        `xml:"version,attr"`
	Description  string        `xml:"description,attr"`
}

// OPML represents an OPML document.
type OPML struct {
	Root *OPMLNode
}

// NewOPML creates a new OPML structure from a slice of bytes.
func NewOPML(b []byte) (*OPML, error) {
	var root OPMLNode
	err := xml.Unmarshal(b, &root)
	if err != nil {
		return nil, err
	}

	return &OPML{Root: &root}, nil
}

// NewOPMLFromURL creates a new OPML structure from an URL.
func NewOPMLFromURL(url string) (*OPML, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return NewOPML(b)
}

// NewOPMLFromFile creates a new OPML structure from a file.
func NewOPMLFromFile(filePath string) (*OPML, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return NewOPML(b)
}

// Outlines returns a slice of the outlines.
func (doc *OPML) Outlines() []OutlineNode {
	return doc.Root.Body.Outlines
}
