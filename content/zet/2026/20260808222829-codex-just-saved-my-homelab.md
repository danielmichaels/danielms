+++
title = "codex just saved my homelab"
categories = ["zet"]
tags = ["zet"]
slug = "codex-just-saved-my-homelab"
date = "2026-08-08 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# codex just saved my homelab

At 0541 alerts started flying that several of my apps were going offline.

When I noticed later and started investigating I saw that the vm's on my proxmox were reporting `io-error` - the disk had either crashed or was full.

I'd done a lot of remediation work recently with the talos nodes (migrating from 1.9 to 1.11 all via codex with minimal downtime) and thought it was me the latter.

Instead of pulling the thread manually I just hit codex with luna to investigate. It found the culprit immediately; `lvm` was full. 

It gave me some ideas and I prompted it to investigate some old vm's which I didn't need anymore, and to delete them to reclaim space if possible.

Codex did it and got us out of the issue but also told me the `woodpecker` was 40gb which was odd. It dawned on me my CI agent must have no clean up scripts and the images and volumes are growing unbounded. Tasked codex to check and create the cron to clean up. It dutifully did so! 

I suggested we setup alerts using my custom in-cluster notifier (which as a human I would of knocked up quickly anyway) but it informed me that wouldn't of fired because this storage issue actually put the cluster into a bad state! 

Instead we setup a cloudflare worker to provide out of band telegram alerts for this sort of thing. Yes, proxmox supports this but I just paid to get CF email and it comes with workers so I want to use it for no other reason than... why not.

This is all stuff I would of fixed and identified myself in an hour or so (minus the cf worker as its written in Rust and I'm a rust newb). But, with codex it took less than 15m to go from down to back online. I host some friends stuff so getting back online is more than just a want it's a need - I don't want my friends to be affected by my mistakes in my homelab!

A codex success. I pretty much use `gpt-5.6-luna` exclusively these days and it excelled today as well. I haven't used `Sol` recently as `luna` and `terra` if I really need horsepower have been enough, and they are fast!

Tags:

    #homelab #outage #ai
