+++
title = "Go test -race requires CGO_ENABLED=1"
categories = ["zet"]
tags = ["zet"]
slug = "go-test--race-requires-cgo_enabled=1"
date = "2022-05-01 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Go test -race requires CGO_ENABLED=1

To run `go test -race` `CGO_ENABLED=1` must be set. I do not use `CGO`
as I do not want any linking.

The workaround is simple, export it for that command.

```golang 
CGO_ENABLED=1 go test -race ./...
```

Tags:

    #golang #cgo #testing

