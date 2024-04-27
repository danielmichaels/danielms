+++
title = "golang ssh client snippet"
categories = ["zet"]
tags = ["zet"]
slug = "golang-ssh-client-snippet"
date = "2024-04-27 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# golang ssh client snippet

Here's an SSH snippet I've used successfully across projects.

```go
package ssh

import (
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

type Auth []ssh.AuthMethod

type Config struct {
	Username    string
	HostAddress string
	Auth        goph.Auth
	Port        int32
}

// Client is an SSH client capable of sending and receiving through a tunnel.
func Client(sc *Config) (*goph.Client, error) {
	client, err := NewClientWithPort(sc)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewClientWithPort returns a client containing an SSHConfig struct and allows
// for a custom port to be set on the goph.Client.
func NewClientWithPort(sc *Config) (*goph.Client, error) {
	cf := goph.Config{
		Auth:     sc.Auth,
		User:     sc.Username,
		Addr:     sc.HostAddress,
		Port:     uint(sc.Port),
		Timeout:  goph.DefaultTimeout,
		Callback: ssh.InsecureIgnoreHostKey(),
	}
	cl, err := goph.NewConn(&cf)
	// todo(dan) need to handle dead servers otherwise 60s + timeout before EHOSTUNREACH(113)
	if err != nil {
		return nil, err
	}
	return cl, nil
}

type ReturnedError struct {
	Status string
	Reason string
	Detail string
	Code   int32
}
```

Tags:

    #go #ssh

