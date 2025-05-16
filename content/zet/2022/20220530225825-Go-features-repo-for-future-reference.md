+++
title = "Go features repo for future reference"
categories = ["zet"]
tags = ["zet"]
slug = "Go-features-repo-for-future-reference"
date = "2022-05-30 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Go features repo for future reference

I need to write a simple web app that exposes some of the more useful
Go features such as Mutexes, gorotinues, `select` and `NewRequestWithContext`. 

The rough outline:

- webserver
- each endpoint with a single *feature*

This way I, or anyone else can pull down the repo and run the server.
Using curl they can hit each endpoint to see how it works whilst reading
over the code. Each endpoint should have good logging to showcase what
is happening.

Tags:

    #golang #development

