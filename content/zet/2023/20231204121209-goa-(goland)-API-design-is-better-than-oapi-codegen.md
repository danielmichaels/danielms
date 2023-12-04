+++
title = "goa (goland) API design is better than oapi-codegen"
categories = ["zet"]
tags = ["zet"]
slug = "goa-(goland)-API-design-is-better-than-oapi-codegen"
date = "2023-12-04 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# goa (goland) API design is better than oapi-codegen

I'm a weekend in switching from `deepmap/oapi-codegen` to <goa.design> and couldn't
be happier.

So far it ticks the boxes the `oapi-codegen` provided but is much simpler to implement.

It uses a DSL instead of writing the `openapi.yaml` file. This DSL generates `openapi`
version 2 and 3. It also generates `gRPC` protobuf files. I find the DSL quite idiomatic
to use and beautiful to look at - unlike yaml.

Using `goa` so far has been much smoother and more expressive. Not to mention it 
forces better architecture and abstractions from the get go. I typically use a 
"god" struct for my application. Instead `goa` uses a service orientated approach
except it generates 80% of it for you. It's the opinionated yet open framework Go
lacks.

So far really happy. Early days but will be recommending it to everyone.

Tags:

  #go #goa #api

