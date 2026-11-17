// /home/krylon/go/src/github.com/blicero/wobbly/engine/engine.go
// -*- mode: go; coding: utf-8; -*-
// Created on 17. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-17 21:59:28 krylon>

// Package engine implements the processing and execution of Jobs.
package engine

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/blicero/krylib"
	"github.com/blicero/wobbly/common"
	"github.com/blicero/wobbly/logdomain"
	"github.com/blicero/wobbly/model"
	"github.com/blicero/wobbly/parser"
)

const timestampFormat = "20060102150405"

// ErrMissingTab indicates that no Job Tab has been read.
var ErrMissingTab = errors.New("Missing Job Tab")

// Engine reads a tab file and processes the Jobs defined therein.
type Engine struct {
	log    *log.Logger
	parser parser.Parser
	tab    *model.Tab
}

// New creates a new Engine
func New() (*Engine, error) {
	var (
		err error
		e   = new(Engine)
	)

	if e.log, err = common.GetLogger(logdomain.Engine); err != nil {
		return nil, err
	} else if e.parser, err = parser.NewRegexParser(); err != nil {
		e.log.Printf("[ERROR] Failed to create Parser: %s\n",
			err.Error())
		return nil, err
	}

	return e, nil
} // func New() (*Engine, error)

// ReadTab reads a Job Tab.
func (e *Engine) ReadTab(path string) error {
	var (
		err error
		tab *model.Tab
	)

	if tab, err = e.parser.Read(path); err != nil {
		e.log.Printf("[ERROR] Cannot parse %s: %s\n",
			err.Error())
		return err
	}

	e.tab = tab
	return nil
} // func (e *Engine) ReadTab(path string) error

// ProcessTab processes a Job Tab.
func (e *Engine) ProcessTab() error {
	var (
		err error
	)

	if e.tab == nil {
		return ErrMissingTab
	}

	// for _, job := range e.tab.Jobs {

	// }

	return krylib.ErrNotImplemented
} // func (e *Engine) ProcessTab() error

func (e *Engine) ProcessJob(idx int) error {
	var (
		err              error
		lastRun, nextRun time.Time
		job              model.Job
	)

	if lastRun, err = e.checkJobLastRun(idx); err != nil {
		return err
	}

	job = e.tab.Jobs[idx]
	nextRun = lastRun.Add(job.Period.Duration())

	if nextRun.After(time.Now()) {
		e.log.Printf("[DEBUG] Job %s is not due until %s\n",
			job.Name,
			nextRun.Format(common.TimestampFormat))
		return nil
	}

	return krylib.ErrNotImplemented
} // func (e *Engine) ProcessJob(idx int) error

func (e *Engine) checkJobLastRun(idx int) (time.Time, error) {
	var (
		err                error
		fh                 *os.File
		buf                bytes.Buffer
		job                *model.Job
		spoolfile, tstring string
		t                  time.Time
	)

	job = &e.tab.Jobs[idx]
	spoolfile = filepath.Join(common.SpoolDir, job.Name)

	if fh, err = os.Open(spoolfile); err != nil {
		if os.IsNotExist(err); err != nil {
			return t, nil
		}

		e.log.Printf("[ERROR] Cannot open timestamp file for Job %s: %s\n",
			job.Name,
			err.Error())
		return t, err
	}

	defer fh.Close() // nolint: errcheck

	if _, err = io.Copy(&buf, fh); err != nil {
		e.log.Printf("[ERROR] Cannot read from timestamp file for Job %s: %s\n",
			job.Name,
			err.Error())
		return t, err
	}

	tstring = krylib.Chomp(buf.String())

	if t, err = time.Parse(timestampFormat, tstring); err != nil {
		e.log.Printf("[ERROR] Cannot parse time stamp for Job %s (%q): %s\n",
			job.Name,
			tstring,
			err.Error())
		return t, err
	}

	return t, nil
} // func (e *Engine) checkJobLastRun(idx int) (time.Time, error)
