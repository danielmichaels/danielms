+++
title = "Forking and contributing to a Go project"
categories = ["zet"]
tags = ["zet"]
slug = "Forking-and-contributing-to-a-Go-project"
date = "2023-11-06 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Forking and contributing to a Go project

I am using <http://github.com/slack-go/slack> for a project and it does not support
*Sign in with Slack* OpenID connect.

After writing the connector manually I decided I'd try contributing upstream.

After forking and making the changes I thought would work I needed to test them against my
working application and didn't know how to use my fork instead of the actual package.

Thankfully, it's pretty easy.

```Go 
go mod edit -replace "github.com/slack-go/slack=github.com/danielmichaels/slack@openid-connect"
```

Doing this and a `go mod tidy` pulled in my branch and a minute later I'd proven my
changes worked.

Hopefully it gets merged sometime. [pr](https://github.com/slack-go/slack/pull/1242)

Edit: My contribution got merged! 

Tags:

    #oss #til #go
