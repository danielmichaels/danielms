+++
title = "Kubernetes local development musings: work edition"
categories = ["zet"]
tags = ["zet"]
slug = "kubernetes-local-development-musings:-work-edition"
date = "2022-06-20 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Kubernetes local development musings: work edition

At work we use K8s for development, staging and production. We have to,
our external dependencies are inaccessible without it.

This makes local development painful. We need to *shift left* and move
more of the stack closer to our local systems.

Possible ideas:

- minikube as the local cluster
- k3s as a local cluster
- kind/k3d as a dockerised local cluster

TBC

