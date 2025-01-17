+++
title = "Docker layers and ARG"
categories = ["zet"]
tags = ["zet"]
slug = "Docker-layers-and-ARG"
date = "2023-05-19 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Docker layers and ARG

I spent a lot of today debugging a PICNIC* error related to Docker's 
multi-stage builds.

> PICNIC: Problem In Chair, Not In Computer

Here is the example

```Dockerfile
# heavily truncated
FROM python:latest as base
ARG CERTIFICATES=certificates

FROM python:latest as builder

COPY $CERTIFICATES/* /usr/local/share/ca-certificates/
```

This example (in a more full examle at least) does not fail. It'll
copy. Except its going to copy `/*` into `/usr/local/share/ca-certificates/`
because `$CERTIFICATES` does not exist in the `builder` stage.

I spent a lot of time trying to figure out why my `certificates/` directory
was getting the certs **and** the source code during CI resulting in an image
full of junk in `/usr/../ca-certificates`.

Pretty obvious in hindsight.

The fixed example Dockerfile:

```Dockerfile
FROM python:latest as base

FROM python:latest as builder
ARG CERTIFICATES=certificates # moved from base

COPY $CERTIFICATES/* /usr/local/share/ca-certificates/
```

Also worth noting, when installing custom certs in a multistage 
you'll need to copy the certs into the final stage and re-run
`update-ca-certificates`.

```Dockerfile
# example of multistage with certs
FROM python as base
# do stuff
FROM python as builder

ARG CERTIFICATES=certificates
COPY $CERTIFICATES/* /usr/local/share/ca-certificates/
RUN update-ca-certificates
# do stuff that needs custom cert

FROM base

COPY --from=builder /usr/local/share/ca-certificates/ /usr/local/share/ca-certificates/
# copy whatever else
RUN update-ca-certificates
```

Tags:

    #til #docker #cert
