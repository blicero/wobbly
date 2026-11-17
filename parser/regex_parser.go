// /home/krylon/go/src/github.com/blicero/wobbly/parser/regex_parser.go
// -*- mode: go; coding: utf-8; -*-
// Created on 11. 11. 2026 by Benjamin Walkenhorst
// (c) 2026 Benjamin Walkenhorst
// Time-stamp: <2026-11-17 19:29:34 krylon>

package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/blicero/wobbly/common"
	"github.com/blicero/wobbly/logdomain"
	"github.com/blicero/wobbly/model"
	"github.com/blicero/wobbly/model/period"
)

var (
	patComment = regexp.MustCompile("#.*")
	patEmpty   = regexp.MustCompile(`^\s*$`)
	patVar     = regexp.MustCompile(`^(\w+)=(.*)`)
	patJob     = regexp.MustCompile(`^(\w+)\s+(\d+)\s+(\d+)\s+(\w+)\s+(.*)$`)
)

// RegexParser uses simple-minded regular expressions to parse
// the configuration file.
type RegexParser struct {
	log *log.Logger
}

// NewRegexParser creates and returns a RegexParser.
func NewRegexParser() (*RegexParser, error) {
	var (
		err error
		rp  = new(RegexParser)
	)

	if rp.log, err = common.GetLoggerStdout(logdomain.Parser); err != nil {
		return nil, err
	}

	return rp, nil
} // func NewRegexParser() (*RegexParser, error)

// Read reads a job table.
func (rp *RegexParser) Read(path string) (*model.Tab, error) {
	var (
		err    error
		fh     *os.File
		reader *bufio.Scanner
		tab    *model.Tab
	)

	if fh, err = os.Open(path); err != nil {
		rp.log.Printf("[ERROR] Cannot open %s: %s\n",
			path,
			err.Error())
		return nil, err
	}

	defer fh.Close() // nolint: errcheck
	reader = bufio.NewScanner(fh)
	tab = &model.Tab{
		Path:      path,
		Variables: make(map[string]string),
		Jobs:      make([]model.Job, 0, 4),
	}

	for reader.Scan() {
		var (
			line   = reader.Text()
			pieces []string
		)

		line = patComment.ReplaceAllString(line, "")

		if patEmpty.MatchString(line) {
			continue
		} else if pieces = patVar.FindStringSubmatch(line); len(pieces) >= 3 {
			tab.Variables[pieces[1]] = pieces[2]
		} else if pieces = patJob.FindStringSubmatch(line); len(pieces) >= 6 {
			var job = model.Job{
				Name:    pieces[4],
				Command: pieces[5],
			}

			switch strings.ToLower(pieces[1]) {
			case "hourly", "hour":
				job.Period = period.Period{
					Unit:  period.Hour,
					Count: 1,
				}
			case "daily", "day":
				job.Period = period.Period{
					Unit:  period.Day,
					Count: 1,
				}
			case "weekly", "week":
				job.Period = period.Period{
					Unit:  period.Week,
					Count: 1,
				}
			default:
				err = fmt.Errorf("unknown period definition %q",
					pieces[1])
				rp.log.Printf("[ERROR] %s\n",
					err.Error())
				return nil, err
			}

			if job.Period.Count, err = strconv.ParseInt(pieces[2], 10, 64); err != nil {
				rp.log.Printf("[ERROR] Cannot parse period %q: %s\n",
					pieces[2],
					err.Error())
				return nil, err
			} else if job.Delay, err = strconv.ParseInt(pieces[3], 10, 64); err != nil {
				rp.log.Printf("[ERROR] Cannot parse period %q: %s\n",
					pieces[3],
					err.Error())
				return nil, err
			}

			tab.Jobs = append(tab.Jobs, job)
		}
	}

	return tab, nil
} // func (rp *RegexParser) Read(path string) (*model.Tab, error)
