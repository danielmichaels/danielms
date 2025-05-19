+++
title = "go parsing seconds into time.Duration"
categories = ["zet"]
tags = ["zet"]
slug = "go-parsing-seconds-into-time.duration"
date = "2022-10-16 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# go parsing seconds into time.Duration

If you have time in seconds and need to turn that into a 
duration such as `2h30m` use `time.ParseDuration`

```go
# https://go.dev/play/p/6C_eVWyTeAD
var seconds = 86900

h, _ := time.ParseDuration(fmt.Sprintf("%ds", seconds))
fmt.Println(h)
// 24h8m20
fmt.Println(h.Hours())
// 24.138888888888889
fmt.Println(h.String()
// 24h8m20s
```

Now you can do operations on this as a `time` object or cast it
to an float using `h.Hours()` and do math operations.

Tags:

    #go #time
