+++
title = "golang map references"
categories = ["zet"]
tags = ["zet"]
slug = "golang-map-references"
date = "2023-06-27 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# golang map references

The CS naming for this type of structure escapes me. But for now, I'll call it "dot notation".
I use this all the time in typescript (and in python).

I wanted a way to have const's in Go but structured so that I could call them elsewhere
in the following manner; `Foo[Bar]`. This is much cleaner than `Foo["bar"]` as the compiler
prevents typo's such as `Foo["baz"]`.

```go
type MapKey string

const (
    Key1 MapKey = "foo"
    Key2 MapKey = "bar"
    Key3 MapKey = "baz"
)

var MyMap = map[MapKey]string{
    Key1: "foo",
    Key2: "bar",
    Key3: "baz",
}

fmt.Println(MyMap[Key1]) // "foo"
```

It's a little repetitious but works well for me.

Tags:

    #go
