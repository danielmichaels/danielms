+++
title = "Go vendor for debugging dependencies"
categories = ["zet"]
tags = ["zet"]
slug = "go-vendor-for-debugging-dependencies"
date = "2025-06-22 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Go vendor for debugging dependencies

TIL how to debug dependencies in Go easily. 

One of my deps have a client timeout of `2 * time.Second` which meant debugging the connections was nearly impossible.

I ran `go mod vendor` and I was then able to make changes to the vendored code. I increased the timeout to something reasonable and was able to find the issue.

I could use the `replace` directive in my `go.mod` and clone down the repo but that seemed like a lot of work. If vendoring didn't work it would of been my next step but `vendor` is perfect for simple changes like this.

Tags:

    #go
