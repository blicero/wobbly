// /home/krylon/go/src/github.com/blicero/wobbly/model/model.go
// -*- mode: go; coding: utf-8; -*-
// Created on 08. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-09 16:27:36 krylon>

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
