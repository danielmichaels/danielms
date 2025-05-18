+++
title = "Go's peterbourgon ff CLI package snippet"
categories = ["zet"]
tags = ["zet"]
slug = "Go's-peterbourgon-ff-CLI-package-snippet"
date = "2024-07-11 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Go's peterbourgon ff CLI package snippet

This snippet will create a simple CLI using
<https://github.com/peterbourgon/ff>.

It will read CLI arguments in the following order:

- flag
- environment variable
- config file

This isn't perfect (its a POC) but is a good starting point.

```go
package main

import (
	"context"
	"fmt"
	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"
	"github.com/peterbourgon/ff/v4/ffyaml"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"text/template"
)

const appName = "pvenotify"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	ret := run(ctx)
	os.Exit(ret)
}

type config struct {
	username    string
	password    string
	host        string
	noTLSVerify bool
	verbose     bool
	config      string
	help        bool
}

func run(ctx context.Context) int {
	var cfg config

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("could not determine user config dir: %v", err)
	}
	pveConfigDir := fmt.Sprintf("%s/%s", userConfigDir, "pvconfig")
	err = CreateDirectoryIfNotExist(pveConfigDir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "could not create pveconfig dir: %v\n", err)
	}
	fd := FileData{cfg}
	tfile, err := os.ReadFile("./config.yaml")
	tmpl := template.Must(template.New("config").Parse(string(tfile)))
	err = generateDefaultConfigFile(pveConfigDir, tmpl, fd)

	fs := ff.NewFlagSet(appName)
	fs.StringVar(&cfg.username, 'u', "username", "un", "username for authentication")
	fs.StringVar(&cfg.password, 'p', "password", "pw", "password for authentication")
	fs.StringVar(&cfg.host, 'a', "api", "http://localhost:8006/api2/json", "PVE api host address")
	fs.BoolVarDefault(&cfg.verbose, 'v', "verbose", false, "verbose logging")
	fs.BoolVarDefault(&cfg.noTLSVerify, 't', "no-tls-verify", false, "do not verify TLS connections")
	fs.StringVar(&cfg.config, 'c', "config", fmt.Sprintf("%s/config.yaml", pveConfigDir), "location of config file")
	fs.BoolVarDefault(&cfg.help, 'h', "help", false, "show help information")

	root := &ff.Command{
		Name:        appName,
		Flags:       fs,
		Usage:       fmt.Sprintf("%s [OPTIONS]", appName),
		Subcommands: nil,
		Exec: func(_ context.Context, args []string) error {
			if cfg.help {
				_, _ = fmt.Fprintln(os.Stderr, ffhelp.Flags(fs))
				return nil
			}
			fmt.Printf("Config: %+v\n", cfg)
			return nil
		},
	}

	if err := root.ParseAndRun(
		ctx,
		os.Args[1:],
		ff.WithEnvVarPrefix("PVE"),
		ff.WithEnvVars(),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigAllowMissingFile(),
		ff.WithConfigFileParser(ffyaml.Parser{}.Parse),
	); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	return 0
}
func doesNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func CreateDirectoryIfNotExist(dirPath string) error {
	if doesNotExist(dirPath) {
		if err := os.Mkdir(dirPath, 0755); err != nil {
			return err
		}
	}
	return nil
}

type FileData struct {
	config
}

func generateDefaultConfigFile(dirPath string, tmpl *template.Template, data FileData) error {
	defaultFileName := "config.yaml"
	filePath := filepath.Join(dirPath, defaultFileName)
	if doesNotExist(filePath) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		if err := tmpl.Execute(file, data); err != nil {
			return err
		}
	}
	return nil
}
```

And the config file would look like this:

```yaml
no-tls-verify: true
username: username_via_config
password: password_via_config
host: http://localhost_via_config:8006/api2/json
verbose: true
```

Tags:

    #go #cli
