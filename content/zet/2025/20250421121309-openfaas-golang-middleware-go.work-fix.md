+++
title = "OpenFaaS golang-middleware go.work fix"
categories = ["zet"]
tags = ["zet"]
slug = "openfaas-golang-middleware-go.work-fix"
date = "2025-04-21 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# OpenFaaS golang-middleware go.work fix

This is an annoying problem where golang-middleware has a go.work referencing go
1.23 but I'm using 1.24 and it want versioning like 1.23.0

```shell
| 2025/04/21 12:11:59 stdout: go: module function listed in go.work file requires go >= 1.23.0, but go.work lists go 1.23; to update it:
| 2025/04/21 12:11:59 stdout:  go work use
| 2025/04/21 12:11:59 stdout: failed to build, error: exit status 1
```

To fix this I had to go into the `template/golang-middleware/go.work` file and
change `1.23` to `1.23.0` or it won't build locally.

This is janky but it works for now.

Note: I use docker compose for local development with openfaas but even using
`faas-cli local-run` this issue occurs. In fact, my fix didn't fix this issue
using `local-run`.

Tags:

    #openfaas
