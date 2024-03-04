+++
title = "MITM proxying for request inspection"
categories = ["zet"]
tags = ["zet"]
slug = "MITM-proxying-for-request-inspection"
date = "2024-03-04 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# MITM proxying for request inspection

I was listening to an old [Running in Production](https://runninginproduction.com)
podcast
[episode](https://runninginproduction.com/podcast/98-games-directory-lets-you-sync-your-games-and-achievements-in-1-place)
in which the developer reversed engineered API's using Fiddler.

I'd used Fiddler briefly and found it cumbersome, along with `mitmproxy`.

So on a whim I searched and found two competing but excellent alternatives:

- [Requestly](https://requestly.com)
- [HTTP-Toolkit](https://httptoolkit.com)

Both offer very comparable features for my use case. For simple use its 
splitting hairs and likely comes to down to the number of methods they
support and their ergonomics.

For instance, Requestly is better to look at in my opinion but HTTP-Toolkit
is no frills and just works.

Listening to the podcast was eye opening because I'd been working on a problem
in my day job which required request inspection.
Wireshark was too heavy handed and all the HTTP traffic was encrypted plus 
I can't filter easily by only trapping certain processes, e.g. my current terminal.

With a MITM proxy I was instantly able to see every request unencrypted. Within
seconds I could see that requests that we thought we going out weren't. 

If you need to track network requests unencrypted from a process you control be it 
a browser or terminal, you should be using a MITM proxy to intercept it. 

Tags:

  #TIL #mitm

