+++
title = "Flush linux DNS cache"
categories = ["zet"]
tags = ["zet"]
slug = "Flush-linux-DNS-cache"
date = "2023-09-13 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Flush linux DNS cache

TIL how to flush my local DNS cache in Ubuntu. 

```
# check if cache is populated
resolvectl statistics
# flush cache
resolvectl flush-caches
# confirm, cache should say 0
resolvectl statistics
```

I was doing some local testing with cloud providers, I switched from one VPS provider to another but kept the same subdomain.
This was causing curl and Go.http errors. Flusing the cache fixed it.

Tags:

    #go #dns #curl
