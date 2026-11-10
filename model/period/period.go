// /home/krylon/go/src/github.com/blicero/wobbly/model/period/period.go
// -*- mode: go; coding: utf-8; -*-
// Created on 09. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-09 16:46:57 krylon>

package period

import "time"

//go:generate stringer -type=Unit

// Unit represents a time interval.
type Unit uint8

// Hour, Day and Week represent an hour (3600 seconds), day (86400 seconds),
// and a week (604800 seconds).
const (
	Hour Unit = iota
	Day
	Week
)

// Period is an amount of time.
type Period struct {
	Unit  Unit
	Count int64
}

// Duration returns the Periods ... well, Duration.
func (p *Period) Duration() time.Duration {
	var res time.Duration

	switch p.Unit {
	case Hour:
		res = time.Hour
	case Day:
		res = time.Hour * 24
	case Week:
		res = time.Hour * 24 * 7
	}

	res *= time.Duration(p.Count)

	return res
} // func (p *Period) Duration() time.Duration
