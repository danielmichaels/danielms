+++
title = "Home server notes"
categories = ["zet"]
tags = ["zet"]
slug = "Home-server-notes"
date = "2024-03-29 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Home server notes

Collecting issues that I hit along the way to bootstrapping my new proxmox homelab.

**Paperless-ngx**

For a reverse proxy to work you need to override the `PAPERLESS_URL` variable. To update this,
open `/opt/paperless/paperless.conf` in the LXC container.

**Caddy**

Caddy doesn't come with a convenient LXC container. Instead I created a ubuntu 22.04
VM and installed Caddy as a systemd unit. This makes saving the Caddyfile
hard so I recommend saving the output once *sorted* and then configuring
it to be API driven thereafter.

To make a `server` **not** request a HTTPS certificate automatically then
use `http://` in the server name block.

## Proxmox Networking

Make sure IPv6 is set to Static and empty. If set to DHCP and your router doesn't provide it
boot times increase dramatically.

Set static IPv4, especially for Cloudflared LXC or `/etc/resolv.conf` will get the gateway
instead of DNS server. The server will fail to connect to Cloudflare and `/etc/resolv.conf`
will always be overwritten by DHCP.

Also make sure all services/VMs have a static IP so when the server reboots it doesn't give 
new IPs rendering the proxy useless. Do this in the GUI and not in the service as it'll be
overwritten by cloud-init.

Tags

  #proxmox #homelab #lxc
