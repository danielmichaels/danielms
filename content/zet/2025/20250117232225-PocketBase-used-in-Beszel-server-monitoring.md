+++
title = "PocketBase used in Beszel server monitoring"
categories = ["zet"]
tags = ["zet"]
slug = "pocketbase-used-in-beszel-server-monitoring"
date = "2025-01-17 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# PocketBase used in Beszel server monitoring

TIL: [beszel] uses PocketBase as its server backend!

Lately, I've been semi obsessing over PB due to its rich featureset. For a Go
dev, it's a game changer in terms of speed and ease of use.

My side projects will never amount to anything needing more than a single
instance. Instead of burning a lot of my limited (2 kids) free time with
boilerplate before even getting to the core logic of whatever I'm trying to
build, I usually run out of steam.

PocketBase goes a long way to fixing that. It gives so much out-of-the-box. IMO
its the Django of Go. Of course, unlike Laravel, Django, Ruby on Rails, it only
supports SQLite but.. do you really need PG or MySql? Probably not. Just get a
bigger box and vertically scale that bad boy!

So today when I found out [beszel] uses PB, I was kinda over the moon. This
isn't a _toy_ its a legit useful project. Even TailScale
[dev rel](https://tailscale.com/blog/video-beszel) have made a video on it.

If nothing else its made me even more adamant that PB is a underutilized tool in
the Go dev's toolkit.

I think the big reason for this is its kinda marketed as a JS dev's tool.
Effectively just a FireBase alternative - and it can be. But, its written in Go
and can be extended as little or as much as you like.

Want a message queue hanging off it - Done! Want a custom API - Done! The sky is
the limit. Only limitation is your imagination.

[beszel]: https://github.com/henrygd/beszel

Tags:

#til #go #pocketbase
