+++
title = "A viable openapi documentation platform"
categories = ["zet"]
tags = ["zet"]
slug = "A-viable-openapi-documentation-platform"
date = "2023-03-15 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# A viable openapi documentation platform

I've been toying with how best to deploy my `openapi` spec file so that
users can integrate with the backend. Initially, I thought about hosting 
it on the API server itself. But, it seemed quite complicated and sub optimal
for a go server to do this.

Instead I think docusarus (which I already use in other projects) has
first class support for this exact use case.

Demo: https://docusaurus-openapi.netlify.app/docs/intro


Tags:

    #openapi #documentation

