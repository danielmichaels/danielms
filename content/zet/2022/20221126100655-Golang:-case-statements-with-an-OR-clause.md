+++
title = "Golang: case statements with an OR clause"
categories = ["zet"]
tags = ["zet"]
slug = "Golang:-case-statements-with-an-OR-clause"
date = "2022-11-26 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Golang: case statements with an OR clause

TIL how to use Go's switch and case statement with `||` (logical OR)
like functionality:

```go
switch foo {
  case "thing", "things":
    println("captured both")
  default:
    println(foo)
}
```

Reference: 

- https://go.dev/ref/spec#Switch_statements

Tags:

    #go

