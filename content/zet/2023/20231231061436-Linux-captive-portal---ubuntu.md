+++
title = "Linux captive portal - ubuntu"
categories = ["zet"]
tags = ["zet"]
slug = "Linux-captive-portal---ubuntu"
date = "2023-12-31 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Linux captive portal - ubuntu

Whenever I fly I *always* have captive portal issues with my Linux laptop.
The initial connection pops a captive portal but if I close the laptop and 
reopen it, it'll kill the connection.

So far this has managed to get a reconnection.

```bash
sudo dhclient
```

Then navigate to `https://captive.apple.com` or `neverssl.com` so that
it can re-open the captive portal.


Tags:

  #qantas #captive-portal

