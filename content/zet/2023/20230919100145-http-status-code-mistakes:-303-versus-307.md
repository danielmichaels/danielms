+++
title = "http status code mistakes: 303 versus 307"
categories = ["zet"]
tags = ["zet"]
slug = "http-status-code-mistakes:-303-versus-307"
date = "2023-09-19 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# http status code mistakes: 303 versus 307

I just wasted a lot of time trying to debug an issue with my template driven web app.

When redirected after making a database change to another page, it kept deleting one of the fields and setting
it back to the `NOT NULL DEFAULT 0` as defined in my table schema.

I was questioning my understanding of SQL and the [Post/Redirect/Get](https://en.wikipedia.org/wiki/Post/Redirect/Get) 
pattern.

Well it turns out I have fat fingered a tab complete and instead of `http.StatusSeeOther` I had inadvertently used
`http.StatusTemporaryRedirect`.

Quote from [rfc7231](https://datatracker.ietf.org/doc/html/rfc7231#section-6.4.7):

```shell
The 307 (Temporary Redirect) status code indicates that the target
resource resides temporarily under a different URI and the user agent
MUST NOT change the request method if it performs an automatic
redirection to that URI.  Since the redirection can change over time,
the client ought to continue using the original effective request URI
for future requests.
```

The upper cased *MUST NOT* precending *change the request method* is a strong signal about what was happening to my requests.

TIL: 303 and 307 are very different and I should pay more attention to the redirect status code when redirecting.

Tags:

    #http #redirects #til
