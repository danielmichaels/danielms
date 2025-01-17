+++
title = "Migrating from CapRover to Coolify"
categories = ["zet"]
tags = ["zet"]
slug = "Migrating-from-CapRover-to-Coolify"
date = "2025-01-03 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Migrating from CapRover to Coolify

I think I've been running a CapRover instance since 2019 or 2020. So long that
my instance is still running Ubuntu 18.04. And, thats why I'm migrating.

Recently, I blew away all my machines and installed Ubuntu 24.04. In doing this
I accidentally deleted my old CapRover SSH keys. The machine is SSH access only
so I don't have anyway to manage it, or upgrade CapRover.

I decided to spin up a new instance on Hetzner and try the backup/restore
process that CapRover recommends. The process is quite clunky IMO and ultimately
didn't work. I have a number of services running on the CapRover and knowing I
can't migrate them over gave me the impetus to move completely over to Coolify.

Side note, one thing that annoyed me about CapRover is when services went down
due to node issue, such as running out of RAM (underprovisioning), all the
services would not automatically come back up (I miss kuberenetes!). This meant
I had to go into each app and manually restart them - theres no API or CLI for
this. Coolify has an API and I'm hopeful I could write scripts to do this.

So far I'm really enjoying Coolify. I like the approach to deploying services
from GitHub. CapRover used webhooks but each service needed to manually sign in
and then copy/paste the webhook address into the repo. Coolify instead creates a
GitHub App so you can add whichever repo's you want and it automatically has
access to them. Much better DX.

I use basic auth a lot on my CapRover services (e.g. NTFY and
ChangeDetection.io). Coolify doesn't have a button/automation for this. You have
to manually add a label to the service file (compose file) which Traefik then
enforces. Its not too hard but its also not as easy as a button and form with
username/ password fields for the auth like CapRover does. Also I'm not sure how
it handles websockets which, again, CapRover sets up with a single click. I
suspect this might another Traefik config I'll have to set.

So far, so good. Glad I made the switch.

Tags:

    #caprover #coolify #paas
