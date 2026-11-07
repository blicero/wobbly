// /home/krylon/go/src/github.com/blicero/wobbly/main.go
// -*- mode: go; coding: utf-8; -*-
// Created on 07. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-07 14:31:09 krylon>

package main

import (
	"fmt"

	"github.com/blicero/wobbly/common"
)

func main() {
	fmt.Printf("%s %s\n",
		common.AppName,
		common.Version)

	fmt.Printf("Nothing to see here, move along...\n")
}
