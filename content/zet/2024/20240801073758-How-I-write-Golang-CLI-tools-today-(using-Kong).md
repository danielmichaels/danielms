+++
title = "How I write Golang CLI tools today (using Kong)"
categories = ["zet"]
tags = ["zet"]
slug = "How-I-write-Golang-CLI-tools-today-(using-Kong)"
date = "2024-08-01 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# How I write Golang CLI tools today (using Kong)

I've ditched `cobra` for `alecthomas/kong` for good. Smaller, easier to grok,
great interface design.

This is a short snippet on how I layout a basic CLI (which I do for every
project, nearly). CLI's power most of my app's, including web servers.

Here is the most simple layout of a useless but instructive example CLI.

```shell
.
├── cmd
│   └── app
│       └── main.go
├── go.mod
├── go.sum
└── internal
    └── cmd
        ├── cmd.go
        └── echo.go
```

Each `*.go` file in detail.

```go
// cmd/app/main.go

package main

import (
        "fmt"
        "me/my-cli/internal/cmd"
        "os"

        "github.com/alecthomas/kong"
)

const appName = "my-cli"

var version string

type VersionFlag string

func (v VersionFlag) Decode(_ *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                       { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
        fmt.Println(vars["version"])
        app.Exit(0)
        return nil
}

type CLI struct {
        cmd.Globals

        Echo cmd.EchoCommand `cmd:"" help:"Example of an Echo command"`
        Version VersionFlag `       help:"Print version information and quit" short:"v" name:"version"`
}

func run() error {
        if version == "" {
                version = "development"
        }
        cli := CLI{
                Version: VersionFlag(version),
        }
        // Display help if no args are provided instead of an error message
        if len(os.Args) < 2 {
                os.Args = append(os.Args, "--help")
        }

        ctx := kong.Parse(&cli,
                kong.Name(appName),
                kong.Description("My new CLI"),
                kong.UsageOnError(),
                kong.ConfigureHelp(kong.HelpOptions{
                        Compact: true,
                }),
                kong.DefaultEnvars(appName),
                kong.Vars{
                        "version": string(cli.Version),
                })
        err := ctx.Run(&cli.Globals)
        ctx.FatalIfErrorf(err)
        return nil
}

func main() {
        if err := run(); err != nil {
                os.Exit(1)
        }
}
```

```go
// internal/cmd/cmd.go

package cmd

type Globals struct {
        Format string `help:"Output format" default:"console" enum:"console,json"`
}
```

```go
// internal/cmd/cmd.go

package cmd

import (
        "encoding/json"
        "fmt"
)

type EchoCommand struct {
        Text string `arg:"" help:"text to echo"`
}

type Formatter interface {
        Output(text string) string
}

type OutputFunc func(text string) string

func (o OutputFunc) Output(text string) string {
        return o(text)
}

func echoMessage(text string, format Formatter) {
        fmt.Println(format.Output(text))
}

func (e *EchoCommand) Run(g *Globals) error {
        ConsoleFormatted := OutputFunc(func(text string) string {
                return text
        })
        JSONFormatted := OutputFunc(func(text string) string {
                jsonData, _ := json.Marshal(map[string]string{"message": text})
                return string(jsonData)
        })

        if g.Format == "json" {
                echoMessage(e.Text, JSONFormatted)
        }
        if g.Format == "console" {
                echoMessage(e.Text, ConsoleFormatted)
        }

        return nil
}
```

Once these files exist, from the root level, run `go run cmd/app/main.go` and it
will output:

```shell
Usage: my-cli <command> [flags]

My new CLI

Flags:
  -h, --help                Show context-sensitive help.
      --format="console"    Output format ($MY-CLI_FORMAT)
  -v, --version             Print version information and quit ($MY-CLI_VERSION)

Commands:
  echo    Example of an Echo command

Run "my-cli <command> --help" for more information on a command.
```

The `echo` command can be called with:

```shell
go run cmd/app/main.go echo "this is going to be written to the console"
# this is going to print to the console
```

and to print JSON:

```shell
go run cmd/app/main.go echo --format=json "this is going to print out some json"
# {"message":"this is going to written to the console"}
```

Tags:

    #go #cli #kong
