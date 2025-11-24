+++
title = "Huma API demo repo"
categories = ["zet"]
tags = ["zet"]
slug = "huma-api-demo-repo"
date = "2025-11-24 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Huma API demo repo

I created and pushed a Huma demo application to GitHub.

I've been using Huma in my personal projects for a while now and so far am very happy with it. It "just works".

Unless you are writing a trivial API I think you **must** have an OpenAPI specification available to your users and your future self.

Huma takes a code first approach - similar to FastAPI. I really like its validation model and how nicely it outputs errors.

When leveraged with `restish` it is a great pairing. 

Repo is at <https://github.com/danielmichaels/huma-demo> and show cases unauthenticated routes alongside authenticated. 

API key and Cookie authentication is demonstrated via the API, API docs and a HTML index page.

Claude did a great job slopping out the HTML/JS to demo the cookie auth. 

Tags:

    #huma #go #api

