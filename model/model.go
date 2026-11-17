// /home/krylon/go/src/github.com/blicero/wobbly/model/model.go
// -*- mode: go; coding: utf-8; -*-
// Created on 08. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-11 05:32:04 krylon>

// Package model provides data types used throughout the application.
package model

import (
	"time"

	"github.com/blicero/wobbly/model/period"
)

// Job is a single command with a schedule.
type Job struct {
	Name    string
	Command string
	Period  period.Period
	Delay   int64
	LastRun time.Time
}

// Tab represents the contents of a configuration file, including list of
// Jobs and optionally a list of variables.
type Tab struct {
	Path      string
	Timestamp time.Time
	Variables map[string]string
	Jobs      []Job
}
