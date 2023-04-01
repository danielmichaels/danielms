+++
title = "Qantas Wifi and Docker"
categories = ["zet"]
tags = ["zet"]
slug = "Qantas-Wifi-and-Docker"
date = "2023-04-01 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Qantas Wifi and Docker

I'm currently in the air and spent ages trying to figure out why I could
connect to the in-flight wifi but not trigger the captive portal.

It was Docker.

The in-flight wifi shares an address space with docker's defaults. 

`172.18.0.0/23`

It did not occur to me that the `br` network sharing this was from Docker.
Instead I spent a bit of time trying to manually remove it, and more time
still `curl`'ing and pinging the gateway to reverse engineer the portal
URL.

In the end a `docker system prune -af` did the trick and immediately
fired the captive portal.

Tags:

    #captive-portal #qantas #networking #docker

