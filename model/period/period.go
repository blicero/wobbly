// /home/krylon/go/src/github.com/blicero/wobbly/model/period/period.go
// -*- mode: go; coding: utf-8; -*-
// Created on 09. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-09 16:02:26 krylon>

package period

//go:generate stringer -type=Unit

// Unit represents a time interval.
type Unit uint8

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
