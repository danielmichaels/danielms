+++
title = "k3s nginx ingress over traefik"
categories = ["zet"]
tags = ["zet"]
slug = "k3s-nginx-ingress-over-traefik"
date = "2022-08-23 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# k3s nginx ingress over traefik

I have elected to use nginx-ingress over the k3s standard of Traefik.
This was mainly driven because OpenFaaS doesn't yet support it for 
custom domain names. I also do not really use Traefik features so I am
unsure what the big trade-offs are yet. 

Tags:

    #kubernetes #traefik #openfaas
