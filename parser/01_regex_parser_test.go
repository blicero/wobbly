// /home/krylon/go/src/github.com/blicero/wobbly/parser/01_regex_parser_test.go
// -*- mode: go; coding: utf-8; -*-
// Created on 17. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-17 19:47:51 krylon>

package parser

import (
	"testing"

	"github.com/blicero/wobbly/model"
)

var rp Parser

func TestCreateRegexpParser(t *testing.T) {
	var (
		err error
	)

	if rp, err = NewRegexParser(); err != nil {
		rp = nil
		t.Fatalf("Cannot create RegexParser: %s",
			err.Error())
	}

} // func TestCreateRegexpParser(t *testing.T)

func TestRegexParserReadTable(t *testing.T) {
	const (
		filename = "testdata/test001.tab"
		expPath  = "/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin"
	)

	var (
		err  error
		tab  *model.Tab
		sh   string
		path string
		ok   bool
	)

	if rp == nil {
		t.SkipNow()
	} else if tab, err = rp.Read(filename); err != nil {
		t.Fatalf("Failed to read %s: %s",
			filename,
			err.Error())
	} else if tab == nil {
		t.Fatalf("Parser returned tab == nil")
	} else if sh, ok = tab.Variables["SHELL"]; !ok || sh != "/bin/sh" {
		t.Fatalf("Parser misread SHELL: %q",
			tab.Variables["SHELL"])
	} else if path, ok = tab.Variables["PATH"]; !ok || path != expPath {
		t.Fatalf("Parser misread PATH: %s",
			tab.Variables["PATH"])
	}
} // func TestRegexParserReadTable(t *testing.T)
