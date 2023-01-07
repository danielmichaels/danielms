+++
title = "Kong is an amazing CLI for Go apps"
categories = ["zet"]
tags = ["zet"]
slug = "Kong-is-an-amazing-CLI-for-Go-apps"
date = "2023-01-07 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Kong is an amazing CLI for Go apps

Firstly, I want to acknowledge that I love Bonzai for command composition
and *most* CLI's. It's my goto and is well designed.

However, for complex applications in a team (who need get-opt like tooling)
Bonzai doesn't quite fit the bill.

Why not Cobra? It bloated as hell. The code generator creates in my opinion
code that quickly gets messy and is hard to test. 

I was using urfave/cli for quite a while and like it but started hitting some
issues when I structured it how I like it. Funny formatting issues, global
flags not working correctly, or overwriting subcommand flags.

Now I use alecthomas/kong. Here's an example from its git repo using a
struct based approach (it's kind of similar to mitchellh/cli in this way).

This is a Docker like CLI.

```go
// ref: https://github.com/alecthomas/kong/blob/master/_examples/docker
package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type Globals struct {
	Config    string      `help:"Location of client config files" default:"~/.docker" type:"path"`
	Debug     bool        `short:"D" help:"Enable debug mode"`
	Host      []string    `short:"H" help:"Daemon socket(s) to connect to"`
	LogLevel  string      `short:"l" help:"Set the logging level (debug|info|warn|error|fatal)" default:"info"`
	TLS       bool        `help:"Use TLS; implied by --tls-verify"`
	TLSCACert string      `name:"tls-ca-cert" help:"Trust certs signed only by this CA" default:"~/.docker/ca.pem" type:"path"`
	TLSCert   string      `help:"Path to TLS certificate file" default:"~/.docker/cert.pem" type:"path"`
	TLSKey    string      `help:"Path to TLS key file" default:"~/.docker/key.pem" type:"path"`
	TLSVerify bool        `help:"Use TLS and verify the remote"`
	Version   VersionFlag `name:"version" help:"Print version information and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type CLI struct {
	Globals

	Attach  AttachCmd  `cmd:"" help:"Attach local standard input, output, and error streams to a running container"`
	Build   BuildCmd   `cmd:"" help:"Build an image from a Dockerfile"`
  // omitted all the other options
}

type AttachCmd struct {
	DetachKeys string `help:"Override the key sequence for detaching a container"`
	NoStdin    bool   `help:"Do not attach STDIN"`
	SigProxy   bool   `help:"Proxy all received signals to the process" default:"true"`

	Container string `arg required help:"Container ID to attach to."`
}

func (a *AttachCmd) Run(globals *Globals) error {
	fmt.Printf("Config: %s\n", globals.Config)
	fmt.Printf("Attaching to: %v\n", a.Container)
	fmt.Printf("SigProxy: %v\n", a.SigProxy)
	return nil
}

type BuildCmd struct {
	Arg string `arg required`
}

func (cmd *BuildCmd) Run(globals *Globals) error {
	return nil
}
func main() {
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag("0.1.1"),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("docker"),
		kong.Description("A self-sufficient runtime for containers"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": "0.0.1",
		})
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
```

Example usage:

```shell
# go run . --help
Usage: docker <command>

A self-sufficient runtime for containers

Flags:
  -h, --help                              Show context-sensitive help.
      --config="~/.docker"                Location of client config files
  -D, --debug                             Enable debug mode
  -H, --host=HOST,...                     Daemon socket(s) to connect to
  -l, --log-level="info"                  Set the logging level (debug|info|warn|error|fatal)
      --tls                               Use TLS; implied by --tls-verify
      --tls-ca-cert="~/.docker/ca.pem"    Trust certs signed only by this CA
      --tls-cert="~/.docker/cert.pem"     Path to TLS certificate file
      --tls-key="~/.docker/key.pem"       Path to TLS key file
      --tls-verify                        Use TLS and verify the remote
      --version                           Print version information and quit

Commands:
  attach    Attach local standard input, output, and error streams to a running container
  build     Build an image from a Dockerfile

Run "docker <command> --help" for more information on a command.
```

Calling `attach`

```shell
# go run . attach --sig-proxy=false container123   
Config: /home/dan/.docker
Attaching to: container123
SigProxy: false
```

What I love about it is the ease of use when structing your application. 
I like to put things into `internal` and `Kong` makes it easy to place
my CLI command structs outside of the `main`/`run` command within
`cmd/cli/main.go`. These CLI structs are then placed into the above
examples `CLI` struct as fields and everything *just works*.

I swapped out urfave/cli for Kong in about 40 minutes - about 7 commands
in total.

Highly recommend it.

Tags:

    #cli #go #kong
