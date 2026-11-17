// /home/krylon/go/src/github.com/blicero/wobbly/parser/parser.go
// -*- mode: go; coding: utf-8; -*-
// Created on 10. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-11 05:39:31 krylon>

// Package parser implements the, well, parser for the configuration file.
package parser

import "github.com/blicero/wobbly/model"

// Parser is the common interface for Parsers of the configuration file.
// Since I might experiment with different implementations, this is the
// interface all of them provide to the application.
type Parser interface {
	Read(path string) (*model.Tab, error)
}
