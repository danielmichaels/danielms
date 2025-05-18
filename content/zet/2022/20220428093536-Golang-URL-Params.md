+++
title = "Golang URL Params"
categories = ["zet"]
tags = ["zet"]
slug = "Golang-URL-Params"
date = "2022-04-28 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Golang URL Params

To create URL encoded parameters on a URL in Go.

```golang
// example from a Zulip webhook
params := url.Values{}
params.Add("to", "stream-name")
params.Add("topic", "topic")
params.Add("type", "stream")
params.Add("content", "here is the content")
encoded := params.Encode()

u := fmt.Sprintf("https://example.com?%s",encoded)

fmt.Println(u)
```

This will create the URL as var `u` with all the parameters 
correctly encoded. This could then be used in `http.NewRequest`
like so:

```golang
// skip err handling for brevity
req, _ := http.NewRequest("POST", u, nil)
```

Tags:

    #golang #web
