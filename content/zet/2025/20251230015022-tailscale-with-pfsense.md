+++
title = "Tailscale with pfSense"
categories = ["zet"]
tags = ["zet"]
slug = "tailscale-with-pfsense"
date = "2025-12-30 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Tailscale with pfSense

I **finally** set up Tailscale on pfSense and what a massive improvement to my life it is!

Being able to connect to Tailscale and then get instant access to everything running on my LAN is massive.

Now I can do things like host apps in my public Coolify and have that container connected to TS giving it access to non-public assets running in my home network. It also doesn't consume my tailscale machine limits because I'm just exposing the subnets - that being said those services do not get a TS routable address. I don't mind though as the direct subnet connection is fine as these services are on static IPs.

I dislike the Cloudflare tunnels approach - this is much better for *this* use case. That being said, I do host Pangolin for other access patterns.

Tags:

    #tailscale

