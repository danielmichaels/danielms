+++
title = "GH Actions versus GitlabCI"
categories = ["zet"]
tags = ["zet"]
slug = "GH-Actions-versus-GitlabCI"
date = "2024-03-24 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# GH Actions versus GitlabCI

I am convinced that people persist with GH actions because its simply
available on Github. 

It's objectively worse in everyway than GitlabCI. I've used Drone and Woodpecker
which are also ten times the product that GH Actions is.

Can't use custom container images! What a PITA.

I make heavy use of "god" images for CI. These are images that have everything you 
need for all your CI activities. Think terraform, packer, taskfile, python, golang and so on.

It creates a huge image but it also means I only have to pull the image and then everything is
available.

In Github, I have to run a step and then install that stuff (go and python, are available
though). This means my pipelines are spent installing things like `golines`, `betteralign` and
`go-task` - for every single invocation.

In Gitlab I pull the image once, its cached on the runner and boom, it fires up the pipeline
and have access to all those things immediately.

Gitlab is better, as are any other providers who **don't** work the same as Github.

Tags:

  #rant #ci #github

