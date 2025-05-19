+++
title = "Golang: find uniques"
categories = ["zet"]
tags = ["zet"]
slug = "golang:-find-uniques"
date = "2022-11-29 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Golang: find uniques

Coming from python where getting a Set from a list is as simple as calling
`s = set(array)` the go version is quite verbose.

Here is how I am finding unique values from an array of integers.

```go
func dedupeInts(ints []int) []int {
  // create a map of int:bool to track found int's
  all := make(map[int]bool)
  // create an array of int's which we'll add any unique entries to
  var list []int
  
  for _, item := range ints {
    // if value (int) is not in the map marked as true then add it the
    // the map and assign true
    // this is how we check for unique entries
    if _, value := all[item]; !value {
      all[item] = true
      list = append(list, item)
    }
  }
  return list
}
```

Tags:

    #golang
