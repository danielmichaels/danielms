+++
title = "home server notes"
categories = ["zet"]
tags = ["zet"]
slug = "home-server-notes"
date = "2024-03-29 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# home server notes

Collecting issues that I hit along the way to bootstrapping my new proxmox homelab.

**Paperless-ngx**

For a reverse proxy to work you need to override the `PAPERLESS_URL` variable. To update this,
open `/opt/paperless/paperless.conf` in the LXC container.

**Caddy**

Caddy doesn't come with a convient LXC container. Instead I created a ubuntu 22.04
VM and installed Caddy as a systemd unit. This makes saving the Caddyfile
hard so I recommend saving the output once *sorted* and then configuring
it to be API driven thereafter.

Tags

  #proxmox #homelab #lxc
