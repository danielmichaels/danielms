+++
title = "microservice learning task"
categories = ["zet"]
tags = ["zet"]
slug = "microservice-learning-task"
date = "2022-08-26 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# microservice learning task

I am building out my Proxmox server and local kubernetes cluster. I don't
have a huge workload to put on either, nor a sophisticated arrangement. So,
I figured this is a good opportunity to replicate the type of work I do
in my day job. Unfortunately, despite working on a distributed microservice
I don't get to work on all parts so I have some knowledge gaps - this should
remediate that.

Here's what I'm thinking:

**Weather God**:

> a simple weather service

It should:

- weather now for location x
- weather historical for location x
- snow falls AU
- barometric
- weather chart

This is a very simple application but could contain many small microservices,
and/or functions.

Spitballing;

- temperature
- barometric
- chart generation
- database
- event bus
- user interface

All deployed on to my local cluster

Tags:

    #project #kubernetes #microservices #proxmox
