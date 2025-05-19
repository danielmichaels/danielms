+++
title = "Golang: title casing first character only"
categories = ["zet"]
tags = ["zet"]
slug = "golang:-title-casing-first-character-only"
date = "2024-02-17 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Golang: title casing first character only

Technically I think title case means all major words have their first letter
capitalised.

This isn't that; its capitalising the **first work only**.

I use this when rendering error fields to users. I like to keep all my 
errors consistent and lower cased. This means some extra work when presenting
errors to users, for instance form field errors. 

```golang
package main

import (
	"fmt"
	"strings"
)

func WithTitleCase(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func main() {
	sentance := "today was a good day"
	newSentance := WithTitleCase(sentance)
	fmt.Println(newSentance)
}
```

go play: https://go.dev/play/p/r2auIbxsiRH?v=

Tags:

  #go
