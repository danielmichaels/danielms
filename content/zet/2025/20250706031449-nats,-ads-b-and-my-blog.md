+++
title = "NATS, ADS-B and my blog"
categories = ["zet"]
tags = ["zet"]
slug = "nats,-ads-b-and-my-blog"
date = "2025-07-06 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# NATS, ADS-B and my blog

With the help of Claude Code, I (mostly) "vibed" a new page on my static Hugo blog. 

The page connects to a server over websockets and that server is connected to Synadia Cloud over NATS. 
At my house I have a Raspberry Pi 2 with an [RTL-SDR][rtl] configured to listen for [ADS-B][adsb] beacons 1090Mhz.

The program running on the Pi is called [dump1090] (created by the Antirez who wrote Redis btw). It outputs raw AVR data onto port 30003 which is what I collect.

So I wrote Go monolith with two entrypoints; agent and server. The agent is installed/ran on a sensor (rpi2) and its configured to listen on TCP 30003 for the
raw data. It also connects to a NATS server and then publishes the data to NATS. The agent uses core NATS so its fire and forget - no persistence. I set it up to reconnect
using a backoff mechanism incase my local network or its wifi dongle goes down. It's connected to my IOT VLAN which is firewalled from the rest of the network but because it can make outbound requests, such as a NATS server, we get around all the complexities of firewalls, network address translation etc. Go NATS!

The server is also connected to NATS - though it can be configured to run as an embedded NATS server. It's subscribing to the subjects that contain the 1090 raw data. It then does some parsing and munging to turn this into useful info. ADS-B packets are transmit frequently and have a structure but not every packet needs to fill all the "fields" - you can get packets with a mix match of these fields and we have to construct them in order to build out useful data structures. Once each "packet" is rolled up into a valid go struct we save that into a [NATS JetStream][js] [Key Value][kv] store. The aircrafts [ICAO] is the key and it has a TTL of 30 minutes.

A websocket is made available on the server which serves as a proxy for its NATS connection. Messages can be sent up the socket and as long as they match the whitelist they'll be sent to NATS and the reply returned down the socket to the caller. I found this to be the safest way to do it in a static site which cannot hide secrets.

I've also created a couple of NATS [micro] endpoints which allow users to lookup aircraft via their callsign or ICAO number. Another great feature of NATS.

Access to this is being exposed (and here's where Claude came in clutch) on my Hugo blog at <https://danielms.site/nats-adsb>. It uses a Hugo shortcode which then contains
a lot of plain javascript to connect to the servers websocket and render the responses on the page. 

On the site you can "watch" all aircraft in real-time (as long as their are planes fly near my house). And lookup their callsign/ICAO number in the table, or do a search of any planes data via a form field. I'm leveraging a free ADSB API in the server to do these requests over NATS micro - the API sometimes has issues but, hey, its free.

Tags:

    #nats #websockets #hugo

[adsb]: https://en.wikipedia.org/wiki/Automatic_Dependent_Surveillance%E2%80%93Broadcast
[rtl]: https://www.rtl-sdr.com/
[dump1090]: https://github.com/antirez/dump1090
[js]: https://docs.nats.io/nats-concepts/jetstream
[kv]:https://docs.nats.io/nats-concepts/jetstream/key-value-store
[icao]: https://en.wikipedia.org/wiki/International_Civil_Aviation_Organization#:~:text=ICAO%20is%20also%20responsible%20for%20issuing%20two,type%20designators%20B741%2C%20B742%20and%20B743%20respectively.
[micro]: https://docs.nats.io/using-nats/nex/getting-started/building-service
