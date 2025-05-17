+++
title = "Bad Request v Unprocessible Entity"
categories = ["zet"]
tags = ["zet"]
slug = "Bad-Request-v-Unprocessible-Entity"
date = "2023-06-01 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Bad Request v Unprocessible Entity

I think this Stack Overflow answer sums it up perfectly.

```shell
Case study: GitHub API

https://docs.github.com/en/rest/overview/resources-in-the-rest-api#client-errors

Maybe copying from well known APIs is a wise idea:

There are three possible types of client errors on API calls that receive request bodies:

Sending invalid JSON will result in a 400 Bad Request response:

HTTP/1.1 400 Bad Request
Content-Length: 35
{"message":"Problems parsing JSON"}
Sending the wrong type of JSON values will result in a 400 Bad Request response:

HTTP/1.1 400 Bad Request
Content-Length: 40

{"message":"Body should be a JSON object"}
Sending invalid fields will result in a 422 Unprocessable Entity response:

HTTP/1.1 422 Unprocessable Entity
Content-Length: 149

{
  "message": "Validation Failed",
  "errors": [
    {
      "resource": "Issue",
      "field": "title",
      "code": "missing_field"
    }
  ]
}
```

Tags:

    #http #error
