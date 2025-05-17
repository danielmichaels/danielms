+++
title = "Go:embed reading individual files"
categories = ["zet"]
tags = ["zet"]
slug = "Go:embed-reading-individual-files"
date = "2022-12-18 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Go:embed reading individual files

I was having trouble accessing a particular file (json, in this case).
In the past, I've had no drama getting at templates and the like but was
clearly missing something here.

To access a single file, you can use the `ReadFile` method on the embedded
filesystem.

For instance,

```shell
# tree .
├── main.go
└── embeds
  └── stuff.json
```

I want to explicitly unmarshal `stuff.json` in my `main.go`

An abridged version of how to do that.

```go
package main

import "embed"

//go:embed files
var jsonFiles embed.FS

fun main() {

  js, _ := jsonFiles.ReadFile("embeds/stuff.json")
  
  var p map[string]any
  err := json.Unmarshal(js, &p)
  if err != nil { log.Fatalln("failed to read embedded file") }
  ```

  This opens the file for reading and gets all of its contents. A perfect
  solution for my needs.

  To read the entire directory just swap out `ReadFile` with 
  `ReadDir("embeds")`.

  Tags:

      #go #embed #TIL
