// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opml

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type OPMLNode struct {
	Version string   `xml:"version"`
	Head    HeadNode `xml:"head"`
	Body    BodyNode `xml:"body"`
}

type HeadNode struct {
	Title           string `xml:"title,attr"`
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

type BodyNode struct {
	Outlines []OutlineNode `xml:"outline"`
}

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

type OPML struct {
	Root *OPMLNode
}

func NewOPML(b []byte) (*OPML, error) {
	var root OPMLNode
	err := xml.Unmarshal(b, &root)
	if err != nil {
		return nil, err
	}

	return &OPML{Root: &root}, nil
}

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

func NewOPMLFromFile(filePath string) (*OPML, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return NewOPML(b)
}

func (doc *OPML) Outlines() []OutlineNode {
	return doc.Root.Body.Outlines
}
