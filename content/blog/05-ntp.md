+++
title = "TIL How Network Time Protocol Works"
slug = "TIL-How-Network-Time-Protocol-Works"
date = "2018-03-23"
tags = ["linux", "TIL"]
categories = ["TIL"]
+++

Network Time Protocol
---------------------

So today I watched a [talk](https://www.youtube.com/watch?v=MDmNvVG9AnQ)
on NTP and it was amazing. I will do my best to summarise the core parts
in this post.

------------------------------------------------------------------------

### What is it?

Network Time Protocol allows clocks between computers to be synchronised
over the internet or local area network. It was developed by [David L.
Mills]
(https://en.wikipedia.org/wiki/David_L._Mills) in
1985 and is currently in its fourth version. Accurate to a few
milliseconds of Coordinated Universal Time, NTP is a fundamental
instrument in modern networking.

### Why do we need it?

Its pretty important that we all agree on what time it is. It is even
more important that banks and other financial institutions agree on a
universal time stamp when conducting transactions, especially if they
are ordered chronologically. Things like SSL certificates could be
spoofed or bypassed if NTP did not ensure a universal time. There is
probably many other security related issues that NTP solves but I think
the point is clear, its a necessary protocol.

### How does it get the "right" time?

Most people have heard of [atomic clocks]
(https://en.wikipedia.org/wiki/Atomic_clock) or GPS
time. Basically, these methods are considered as accurate as humans can
calculate to 'real' time. It would be simplistic to say that NTP polls
these sources for the current time - but nonetheless that's the basic
premise.

Diving deeper, NTP actually uses things called a 'Stratum'. An atomic
clock is a Stratum 0, also known as a reference clock. The Stratum's go
from 0 all the way to 16. Stratum 1 is the closest an NTP server can get
to the reference clock, and is usually within a few milliseconds. These
are referred to as primary time servers. Each Stratum refers to the
preceding number, and polls that server for its time. Stratum 3 will
synchronise to a Stratum 2 server for example. They can and often do
peer with servers in their own Stratum as a sanity check and backup. Of
note Stratum 16 means 'unsynchronised'.

------------------------------------------------------------------------

### The Process

Roughly the process follows something like this:

-   ask for the time,
-   get the roundtrip times,
-   figure out if you trust the response,
-   make any adjustments,
-   repeat every 64 seconds, forever.

#### Roundtrip Times

After asking the time from a server, NTP needs to factor in how long it
took to get the response back because the time from when it sent the
response to receiving it will now be out of sync.

```shell 
# The four timestamps needed for calculating the time
t1 = timestamp of request packet
t2 = timestamp when server received packet
t3 = timestamp of servers reply transmission
t4 = clients response packet reception timestamp
```

Example of NTP packet

```shell 
13:11:58.155997 IP (tos 0x0, ttl 56, id 42684, offset 0, flags [none], proto UDP (17), length 76)
    cpe-110-141-196-84.vic.asp.telstra.net.ntp > client.local.lan.ntp: [udp sum ok] NTPv4, length 48
    Server, Leap indicator:  (0), Stratum 1 (primary reference), poll 10 (1024s), precision -23
    Root Delay: 0.000000, Root dispersion: 0.001953, Reference-ID: PPS^@
      Reference Timestamp:  3730684254.163282214 (2018/03/22 13:10:54) # t1
      Originator Timestamp: 3730684318.091658531 (2018/03/22 13:11:58) # t2
      Receive Timestamp:    3730684318.125586175 (2018/03/22 13:11:58) # t3
      Transmit Timestamp:   3730684318.125623665 (2018/03/22 13:11:58) # t4
        Originator - Receive Timestamp:  +0.033927643
        Originator - Transmit Timestamp: +0.033965133
```

To calculate the current time from the servers response NTP does the
following calculation:

    t4 - t1 = roundtrip time
    roundtrip time / 2 = one-way latency
    t3 + one-way latency = current time

The next part of this is whether or not the client trusts the NTP
server. This is done in a few ways. Firstly, by sending out several
queries to several servers rather than trusting that the response from
one server is correct. NTP then favours the lowest latency and discards
any outliers. Secondly, NTP uses some statistical analysis of its
responses over a period of minutes to determines who is accurate and who
isn't based off those statistics.

#### Make Adjustments

What NTP tries to never do is go backwards in time. Sometimes it has to
and we will get to that. But for the most part what it does it 'slew'
the clock. Simply it just slows or speeds up the clock to match the
correct time and does this in a gradual way with small increments.

At its max adjustment speed of 500ppm it would take 2000 seconds to make
an adjustment of just one second!

Given the slow slew rate, slewing is capped to 128 milliseconds.
Anything above that cannot be slewed and must be 'stepped' or jumped to
the correct time, be it forward or backward. This does not happen often
except in cases such as bringing a machine back online after maintenance
or during initial setup. Any machine that is over 1000 seconds out must
be manually configured within that threshold or it will not be able to
receive adjustments.

------------------------------------------------------------------------

### Conclusion

That's Network Time Protocol in a nutshell. I had never paid much heed
to NTP prior to [Joel Potischman's](https://twitter.com/jpotischj) talk
at !!Con. He gave a great talk and it only goes for ten minutes and uses
some good graphs and visualisations that are missing from this post. If
you want to see NTP in action on your computer you can use
`tcpdump -vv port 123` or check it out in wireshark. Whilst
writing this I found a bad response from one server that was +1023
seconds out and thus dropped as an outlier - so it does happen.
