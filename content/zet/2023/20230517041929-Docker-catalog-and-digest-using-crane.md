+++
title = "Docker catalog and digest using crane"
categories = ["zet"]
tags = ["zet"]
slug = "Docker-catalog-and-digest-using-crane"
date = "2023-05-17 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Docker catalog and digest using crane

I've written a number of bash scripts over my time to check container registry
catalogs and then parse them for certain tags.

TIL that google's crane supports this out of the box.

Install crane with `arkade`:

```shell
arkade get crane
```

Here's the output of `ls` and `digest`:

```shell
# list all tags
crane ls danielmichaels/sitesearch
915194c0cd6350048a4b6c3a24eb555bae7e0a11
9c7d76991e79a8dd0fa1f081e70b5c2daf6838e1
a12e6f4326f0a3397df42cade3e4382dd1915f83
ab658d594ab00b36eb7a66d9493654d2d0f7d971
d7a22105650736aaa3f3fbf8a81a063824be6d0a
ed22c15b0057c1a5ad93a836e1793b01c7f2e195
latest
okteto
staging
```

```
# get a tag digest
crane digest danielmichaels/sitesearch:latest
sha256:2fdca04e110e75df31c22a33fcd6db9edda6fb36cc86c44bece01d0725346e72
```

Tags:

    #til #crane #docker #registry

