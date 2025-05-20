+++
title = "Helm upgrade dynamically with Go"
categories = ["zet"]
tags = ["zet"]
slug = "helm-upgrade-dynamically-with-go"
date = "2023-05-07 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Helm upgrade dynamically with Go

I wrote a python file to create `helm upgrade` commands dynamically 
based on environment variables.

Not happy with that I also wrote it in Go. I much prefer it because its
easier to distribute and has zero dependencies.

```go
// Auto generate a helm deployment for use across multiple projects in CI.
// To use this the following is required:
//
//   - Environment variables set with the values needed. Read carefully for how to use each
//     type of variable.
//
//   - GitlabCI, this is primarily designed to be used in GitlabCI using its includes and
//     artifact generation functionality
//
//     See comments for how to use local environment variables to set values for helm.
//
// To use this in GitlabCI you only need to set the 'variables' block with your required values.
// Typically, this job will generate an artifact with the script as its output which a deployment
// job will have a needs relationship with and will then execute the script. YMMV
package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

// Environment variables, to use HELM_SET_FILE_PREFIX and HELM_SET_PREFIX
// the following considerations must be taken into a account.
// helm's '--set' and '--set-file' are used to override yaml fields inside a values.yml file.
// Using the following yaml as a example we'll illustrate how to use these envvars.
//
// example values.yml
// db
//
//	name: myDatabase
//	password: ""
//	caCert: {}
//
// Shell environment variables do not allow for "." notation. To get around this limitation
// it is expected that any "_" after the prefix will be translated into a "."
// For instance, HELM_SET_db_password=myPass would become '--set db.password=myPass'
// HELM_SET_FILE_db_caCert=/tmp/path/0/ca.crt becomes '--set-file db.caCert=/tmp/path/0/ca.crt'
//
// The outlier here is HELM_VALUES which is always provided as a string such as HELM_VALUES=helm/values.yml,helm/values2.yml
// which generates '--values helm/values.yml,helm/values2.yml'
const (
	HELM_SET_PREFIX      = "HELM_SET_VALUE_"
	HELM_SET_FILE_PREFIX = "HELM_SET_VALUE_"
	HELM_VALUES          = "HELM_VALUES"
)

type Values struct {
	Key   string
	Value string
}

func extractFileAndSetFile(re_prefix string) []Values {
	var values []Values
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, re_prefix) {
			trimmedValues := strings.SplitN(env, re_prefix, 2)
			splitOnEquals := strings.SplitN(trimmedValues[1], "=", 2)
			dotValue := strings.ReplaceAll(splitOnEquals[0], "_", ".")
			value := Values{
				Key:   dotValue,
				Value: splitOnEquals[1],
			}
			values = append(values, value)
		}
	}
	return values
}

func extractValuesFiles(re_prefix string) []string {
	var values []string
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, re_prefix) {
			kv := strings.SplitN(env, "=", 2)
			v := kv[1]
			values = append(values, v)
		}
	}
	return values
}

type TemplateData struct {
	Release     string
	ChartPath   string
	SetValue    []Values
	SetFile     []Values
	ValuesFiles []string
	// Allow for --dry-run to be passed in the script; usefule for debugging
	// or when wanting to use helm to template but kubectl to execute/apply
	DryRun any
}

func main() {
	setValues := extractFileAndSetFile(HELM_SET_PREFIX)
	setFiles := extractFileAndSetFile(HELM_SET_FILE_PREFIX)
	valuesFiles := extractValuesFiles(HELM_VALUES)
	release := os.Getenv("HELM_RELEASE")
	chartPath := os.Getenv("HELM_CHART_PATH")

	data := TemplateData{
		Release:     release,
		ChartPath:   chartPath,
		SetValue:    setValues,
		SetFile:     setFiles,
		ValuesFiles: valuesFiles,
		DryRun:      os.Getenv("DRYRUN"),
	}

	temp := template.Must(template.New("helm").Parse(helmTemplate))

	err := temp.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

}

var helmTemplate = `#!/bin/bash

helm upgrade --install {{ .Release }} {{ .ChartPath }} \
	{{ if .SetValue -}}
		{{- range .SetValue -}}
	--set '{{ .Key }}={{ .Value }}' \
		{{- end -}}
	{{- end }}
	{{ if .SetFile }}
		{{- range .SetFile -}}
	--set-file '{{ .Key }}={{ .Value }}' \
		{{- end -}}
	{{- end }}
	{{ if .ValuesFiles }}
		{{- range .ValuesFiles -}}
	--values '{{ . }}' \
		{{- end -}}
	{{- end }}
	--atomic --timeout 300s {{ if .DryRun }}--dry-run{{ end }}
`
```

To build this as a distributable binary is as simple as `go build main.go`.
You can name the binary with `go build -o <name> main.go`. The run the
binary and pipe its output to a script for execution when required. e.g.
`go run main.go > deploy-helm.sh`.

It is a simple solution and would be keen to here how others solve it.

Tags:

    #helm #go #templates
