+++
title = "nextdns rewrites"
categories = ["zet"]
tags = ["zet"]
slug = "nextdns-rewrites"
date = "2024-03-30 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# nextdns rewrites

I was humbled today; I thought nextdns rewrites meant you could point
all queries to a DNS server.

For instance, *.foo.bar -> 192.168.20.20 (private LAN DNS server) and that
would then pass on the request to the DNS server.

It doesn't work like this. Instead you have to map each host to the LAN host
address/IP (or external).

There's no API for this either so for my 15 internal domains I have to 
enter them in the GUI.

I have several profiles and its a massive PITA. Now I have one profile for
all nextdns enabled devices so that I only have to change one place.

A little disappointed about this not being exposed via an API (and thus scriptable).

Tags:

  #TIL #dns #nextdns

