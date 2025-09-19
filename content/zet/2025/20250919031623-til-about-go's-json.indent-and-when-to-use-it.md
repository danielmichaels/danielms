+++
title = "TIL about Go's json.Indent and when to use it"
categories = ["zet"]
tags = ["zet"]
slug = "til-about-go's-json.indent-and-when-to-use-it"
date = "2025-09-19 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# TIL about Go's json.Indent and when to use it

I needed to `json.Unmarshal` then `json.MarshalIndent` a large amount of JSON. In this was a set of very large numbers.

I learnt that this causes precision issues because unmarshalling to an `interface{}` will use `float64` for numbers. 

This caused a lot of head scratching, and all I can say is thank god for the debugger! Those who "don't use debuggers", well, I tip my hat to you because `print` debugging would not have go me through this pickle.


Example JSON

```json
"account_details": [
  {
    "reserved_memory": 18446744073709551615,
    "reserved_storage": 18446744073709551615
  }
]
```

Was raising:

```go
decode 0001.json: json: cannot unmarshal number 18446744073709552000 into Go struct field Foo.account_details.Bar.reserved_memory of type uint64
```

To achieve my aims and keep the numbers exactly the same, `json.Indent` can be used.

My takeaway:

- `json.Indent` is for pure text manipulation
- `json.MarshalIndent` is for parsing

I made a quick <https://go.dev/play/p/qFKb23mUYuz> link showing it in action.

Tags:

    #go #json
