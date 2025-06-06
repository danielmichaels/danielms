+++
title = "Custom zerolog logger"
categories = ["zet"]
tags = ["zet"]
slug = "custom-zerolog-logger"
date = "2023-07-30 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Custom zerolog logger

Zerolog is my favourite logger for Go.

Chi comes with `httplog` which creates a very simple yet powerful 
`http` logger. For non-http apps which need logging I've copy pasted
my own logger below:

```go
package main


import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

func NewLogger(serviceName string, opts ...Options) zerolog.Logger {
	if len(opts) > 0 {
		Configure(opts[0])
	} else {
		Configure(DefaultOptions)
	}
	logger := log.With().Str("service", strings.ToLower(serviceName))
	if !DefaultOptions.Concise && len(DefaultOptions.Tags) > 0 {
		logger = logger.Fields(map[string]interface{}{
			"tags": DefaultOptions.Tags,
		})
	}
	return logger.Logger()
}

var DefaultOptions = Options{
	LogLevel:        "info",
	LevelFieldName:  "level",
	JSON:            false,
	Concise:         false,
	Tags:            nil,
	SkipHeaders:     nil,
	TimeFieldFormat: time.RFC3339Nano,
	TimeFieldName:   "timestamp",
}

// Configure will set new global/default options for the httplog and behaviour
// of underlying zerolog pkg and its global logger.
func Configure(opts Options) {
	if opts.LogLevel == "" {
		opts.LogLevel = "info"
	}

	if opts.LevelFieldName == "" {
		opts.LevelFieldName = "level"
	}

	if opts.TimeFieldFormat == "" {
		opts.TimeFieldFormat = time.RFC3339Nano
	}

	if opts.TimeFieldName == "" {
		opts.TimeFieldName = "timestamp"
	}

	// Pre-downcase all SkipHeaders
	for i, header := range opts.SkipHeaders {
		opts.SkipHeaders[i] = strings.ToLower(header)
	}

	DefaultOptions = opts

	// Config the zerolog global logger
	logLevel, err := zerolog.ParseLevel(strings.ToLower(opts.LogLevel))
	if err != nil {
		fmt.Printf("httplog: error! %v\n", err)
		os.Exit(1)
	}
	zerolog.SetGlobalLevel(logLevel)

	zerolog.LevelFieldName = strings.ToLower(opts.LevelFieldName)
	zerolog.TimestampFieldName = strings.ToLower(opts.TimeFieldName)
	zerolog.TimeFieldFormat = opts.TimeFieldFormat

	if !opts.JSON {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: opts.TimeFieldFormat})
	}
}

type Options struct {
	// LogLevel defines the minimum level of severity that app should log.
	//
	// Must be one of: ["trace", "debug", "info", "warn", "error", "critical"]
	LogLevel string

	// LevelFieldName sets the field name for the log level or severity.
	// Some providers parse and search for different field names.
	LevelFieldName string

	// JSON enables structured logging output in json. Make sure to enable this
	// in production mode so log aggregators can receive data in parsable format.
	//
	// In local development mode, its appropriate to set this value to false to
	// receive pretty output and stacktraces to stdout.
	JSON bool

	// Concise mode includes fewer log details during the request flow. For example
	// excluding details like request content length, user-agent and other details.
	// This is useful if during development your console is too noisy.
	Concise bool

	// Tags are additional fields included at the root level of all logs.
	// These can be useful for example the commit hash of a build, or an environment
	// name like prod/stg/dev
	Tags map[string]string

	// SkipHeaders are additional headers which are redacted from the logs
	SkipHeaders []string

	// TimeFieldFormat defines the time format of the Time field, defaulting to "time.RFC3339Nano" see options at:
	// https://pkg.go.dev/time#pkg-constants
	TimeFieldFormat string

	// TimeFieldName sets the field name for the time field.
	// Some providers parse and search for different field names.
	TimeFieldName string
}

```

This needs more tweaking but serves as a good starting point.

Tags:

    #go #logging

