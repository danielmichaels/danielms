+++
title = "docker buildkit breaks 'docker build -o'"
categories = ["zet"]
tags = ["zet"]
slug = "docker-buildkit-breaks-'docker-build--o'"
date = "2022-09-13 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# docker buildkit breaks 'docker build -o'

I wasted a day only to discover that enabling `DOCKER_BUILDKIT=1` will
break the `docker build -o <bla.tar.gz>...`. Instead of creating a 
tar file, it instead creates a directory! The only reason I figured this
out was by running it manually and seeing that it was a directory called
`bla.tar.gz`. Debugging this when it was used in our runners was a nightmare.

I turned off that feature entirely out of spite, rather than fix it. Besides,
our pipelines are doing well without buildkit. 

Tags:

    #rant #docker #spite

