+++
title = "Taskfile is my new Makefile"
categories = ["zet"]
tags = ["zet"]
slug = "Taskfile-is-my-new-Makefile"
date = "2023-02-17 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Taskfile is my new Makefile

For every project one of the first things I do is create a Makefile.
Generally this works well except when I need to do some sort of validation, for 
example checking the existence of a directory. It's possible but I have to
google it every time.

Enter `Taskfile` a yaml based Makefile written in Go. I am very familiar 
with Gitlab pipelines so writing yaml like this comes naturally. To me,
this is so much easier to read and write. 

Example

```yaml
# https://taskfile.dev

version: '3'

vars:
  REGISTRY: docker.io/me
  IMAGE: myimage

tasks:
  default:
    desc: |
      List all available tasks
    cmds:
      - task --list
    silent: true

  env:
    desc: |
      Print all environment variables sorted alphabetically
    cmds:
      - env | sort
    silent: true

  dev:
    deps:
      - check
    desc: |
      Run the local development environment
    cmds:
      - air

  docker-build:
    desc: |
      Build the docker image
    cmds:
      - docker build . -f deploy/Dockerfile -t "{{.REGISTRY}}/{{.IMAGE}}"

  docker-push:
    deps:
      - docker-build
    desc: |
      Push the docker image
    cmds:
      - docker push "{{.REGISTRY}}/{{.IMAGE}}"

```

This has been incredibly useful in the development of a Django-channels
application where I needed a lot of setup. I was able to pass this over
to my team and the concept well received.

Tags:

    #taskfile #go #makefile #til
