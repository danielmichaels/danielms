+++
title = "Go template iterators"
categories = ["zet"]
tags = ["zet"]
slug = "go-template-iterators"
date = "2022-06-10 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Go template iterators

How to use Go templates `range` and `$`.

**range**

This loops over the length of a struct setting the dot value to element
being looped over. 

```go
type Example struct {
  Id string
  Name string
}

{{ range Example }}
ID: {{.Id}} & Name: {{.Name}}
{{end}}
```

If you need to access a value via the `.` notation and that value is outside
the current iteration, you need to use `$.`. 

```go
// example 
{{.Container.ID}}
// range over an Item struct to build a URL
{{range .Items}}
// the $ is required so that the template can inspect global state
// not just the current iteration. Item does not have a .Container.ID 
// field and will fail without $
<a href="/containers/{{$.Container.ID}}/items/{{.ID}}">url</a>
{{end}}
```

Tags:

    #go #templates
