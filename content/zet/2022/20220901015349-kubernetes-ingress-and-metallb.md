+++
title = "kubernetes ingress and metallb"
categories = ["zet"]
tags = ["zet"]
slug = "kubernetes-ingress-and-metallb"
date = "2022-09-01 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# kubernetes ingress and metallb

For my homelab, I'm teaching myself how to setup a bare metal cluster using
k3s, nginx-ingress and metallb. It's show soo many holes in my knowledge 
and understanding of k8s. When you work on established systems such as 
OpenShift at your day job, you take for granted all the things it does for you.

## I misunderstand Ingress vs Services

At work, we use Ingress and NodePorts to allow traffic to our applications
and I thought I understood it. But, now that I'm throwing Metallb into the
mix, I clearly don't.

My understanding at the minute

1. Create a metallb static IP for my nginx-ingress
2. Define a Service for my app
3. Use the container port on the Ingress and set a host
4. Create a DNS host override in pfSense
  1. Use the metallb static IP of the nginx-ingress service for all the hostnames
  2. 'argocd.home.lab` and `tekton.home.lab` map to `192.168.20.199` and nginx will resolve it

So far that works. It took me a long time to figure that out and this is just for
the LAN. Exposing applications later will be another thing I likely misunderstand.

Tags:

    #kubernetes #homelab #til

