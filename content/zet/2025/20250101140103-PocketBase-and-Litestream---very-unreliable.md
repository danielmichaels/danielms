+++
title = "PocketBase and Litestream - very unreliable"
categories = ["zet"]
tags = ["zet"]
slug = "PocketBase-and-Litestream---very-unreliable"
date = "2025-01-01 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# PocketBase and Litestream - very unreliable

After most of a day I've given up on a pure PocketBase Litestream setup.

It's really flaky and I experienced several instances where the restored database was subseqently
overritten by PocketBase. I couldn't ascertain why and don't have the time to understand the problem.

Frankly, I'm disappointed by it. I think PocketBase will be my new MVP/PoC stack with Go. Its an amazing
tool and offers so much out of the box.

I created <pocketshare.infra.ptco.rocks> with it in about a week. Its a file sharing app - which I felt
pushed all the PocketBase buttons. It was an excellent thing to build for learning PB.

For now, I'm resorting to volume mounting the DB but running it in a container. I'll rely on 
automated backups and manual intervention in case of emergency rather than automation via Litestream.

This will work great but I'm still down about not figuring out Litestream!

ref:
    - https://pocketbase.io
    - https://github.com/pocketbase/pocketbase/discussions/3080

Tags:

    #go #pocketbase

