---
title: "NATS ADS-B"
date: 2025-07-03T00:00:00Z
draft: false
---

What you are seeing here is live [ADS-B] data from a sensor located at my house. It is connected to a NATS server
as a [leaf node][ln] which then connects to [Synadia] [Cloud][ngs]. 

--- 

## Usage

Clicking on *Watch Aircraft* will start a live tail of aircraft seen in the last 30 minutes.
There may be times when no aircraft are flying due to noise, typically 0000-0500 AEST.

There are two function that can be made to look up aircraft by their Callsign or [mode s][ms] number.
These make use of NATS [services/micro][sm] framework.

{{< nats-websocket >}}

[sm]: https://docs.nats.io/using-nats/nex/getting-started/building-service
[ads-b]: https://en.wikipedia.org/wiki/Automatic_Dependent_Surveillance%E2%80%93Broadcast
[ms]: https://skybrary.aero/articles/mode-s
[synadia]: https://synadia.com
[ngs]: https://www.synadia.com/cloud
[ln]: https://docs.nats.io/running-a-nats-service/nats_docker/ngs-leafnodes-docker