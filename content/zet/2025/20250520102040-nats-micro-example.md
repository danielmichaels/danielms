+++
title = "NATS Micro example"
categories = ["zet"]
tags = ["zet"]
slug = "nats-micro-example"
date = "2025-05-20 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# NATS Micro example

[NATS Micro][nm] is a super simple way to create microservices using NATS as the communication layer.

We're using them a lot at Synadia internally but I'd never used them before.

It only takes about 5 minutes to go from an empty file to running service that is handling load. It's really impressive.

Over the weekend I wrote a small (imperfect) [example application][demo] to learn enough to get started. It fetches the top posts 
from HackerNews over the last 24 hours, saves them to NATS KV, sorts by *points*, saves that days sorted posts to Object Store and 
exposes a service that prints the top *n* posts.

Once the service is running you can inspect most things with just three commands:

- `nats micro ls`
- `nats micro info HackerNews` - HackerNews being the service I'm running
- `nats micro stats HackerNews`

One thing I did come across was mixing of `publish` and `request` semantics. In one of my endpoints I am publishing items as they come in. This allows
the consumer to execute on each item instead of a large list - typical work queue behaviour. But, if I `nc.Publish` inside of a `req` and don't respond 
it will cause the service to report an error. This is more about how I structured the code, and it's easy to fix. I just created an two endpoints; one 
for publishing/subscribing and one that will respond. The endpoint which only exists to pub/sub could probably be excluded from the services framework
but I included it because it gives great discoverability and statistics out of the box.

Going forward I see myself using `nats micro`/`nats service` a lot. It's such a simple, yet powerful construct around NATS.

One thing I thought of was how this could be a OpenFaaS-like replacement, which also uses NATS, where you could bootstrap "functions" and do it 
all in `micro`. It would be easy to expose a single HTTP gateway like OpenFaaS's which proxies/transforms the data into NATS protocol. These things 
kind of already exist but not quite the same. 

References:

- https://www.youtube.com/watch?v=AiUazlrtgyU
- https://www.youtube.com/watch?v=byHGNUqIONw
- https://www.youtube.com/watch?v=s2seyKyQ_Zw

[nm]: https://github.com/nats-io/nats.go/blob/main/micro/README.md
[demo]: https://github.com/danielmichaels/nats-micro-hackernews

Tags:

    #nats #micro #services #openfaas
