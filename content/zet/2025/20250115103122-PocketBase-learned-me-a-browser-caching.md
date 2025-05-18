+++
title = "PocketBase learned me a browser caching"
categories = ["zet"]
tags = ["zet"]
slug = "PocketBase-learned-me-a-browser-caching"
date = "2025-01-15 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# PocketBase learned me a browser caching

Im writing a file sharing tool to explore pocketbases abilities. The files are
stored in PB and retrieved by creating their unique url. This is easy using the
JS SDK and works well.

I needed to extend this im go by using PBs hooks. Particularly the
OnRequestFileDownload hook. Everytime a file is downloaded we can create custom
logic before returning the file object

In this case, im creating a entry in the analytics table to track each download.
I also have options like download limits, expiry and password protection. So I
need to use the hook.

This is where the problem began. I could trigger the download by not the hook.

It took me a number of hours to figure it out. Browsers cache things. After the
first download that url is cached and therefore it won't fetch the contents from
the server again. Which is perplexing to me, but it wouldn't even show up in my
dev tools tab. This means no hook called, no analytics. Not helpful

The workaround is to cache bust by adding a querystrinh to the url. I used Unix
time which means every second the cache is busted and contents pulled freshly
from server. Perfect.

Tags:

    #til #browsers #pocketbase
