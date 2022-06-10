+++ 
date = "2022-06-04"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-may-2022"
title = "May 2022 Retrospective"
draft = false 
+++

# Summary

## Highlight


## Goal Performance

A review of last months three goals. See [April's Retrospective][old-retro].

[old-retro]: /retrospectives/2022/retrospective-april-2022/

### Deploy pfSense+ capable Mudmap 

- **Appraisal**: Did not roll this out. 
- **Rating**: D

All of my this month was dedicated to multi-account feature. This is still
a great feature to roll out by I decided early on in the month that it would come after.

There are some compatibility issues remaining with the API's support for
pfSense+, namely with ARM devices. This has only become apparent in the
last few days, and is not something I can test either. So, in hindsight
waiting this one has saved Mudmap from some potentially unhappy customers.
Hoping a fix is released for this in due course.  When this does release it will have to come with that caveat.

### Get started on multi-account support (a big customer request)

- **Appraisal**: Still ongoing but should be finished this month.
- **Rating**: C

It is not finished yet, but at least it is work I am excited to do and release for my customers.
I've given this a high priority as numerous people have requested it. They
have staff who they want to manage devices but its currently not possible without sharing a single account. The sub-accounts must also have some permission settings, like payment details for instance.

So, I've settled on the following structure;

```
Organisation
      |
      |
  ---------
  |       |
Users   Devices
```

Each user has a `permission` applied which will grant them access to
certain views and API operations. Permissions haven't been started yet.

I've had to migrate Users to Organisations along with their subscriptions and Devices.

Considering the changes I'm probably going to do an outage window over a weekend, rather than roll it out during a weekday and potentially catch someone in a weird state.
Not ready for that yet but these are thing considerations I've been thinking about along with writing the code.

I am still working on the user interface, which as always, is the most tedious and slow part for me. 
The other hard part, is managing how users add new members to their Organisation.
To be a member they have to sign up via Auth0 - to get a valid id - and then
they can be added to the Org. So far, I've settled on admin users being able 
to invite members by email. Then they sign up and enter a code to be assigned
to the Org. It feels a little contrived and suboptimal but works for now.

I feel like I accomplished a lot towards achieving this goal.

## Recommendations

If you want an excellent security tool that scans and identifies everything on your network, 
check out [Rumble](https://rumble.run). It is a company that was started by [H.D Moore](https://en.wikipedia.org/wiki/H._D._Moore)
who created the [Metasploit](https://en.wikipedia.org/wiki/Metasploit_Project) project. I ran it
to see what is on my network but also to see how it identifies my many pfSense VM's. It found 
all of my assets with very detailed info within 3 minutes across two /24 networks. Definitely worth
a look, and you can run it from a raspberry Pi.

## Wrap up

A great month from a productivity standpoint. I've managed to work out a decent balance for 
work and out-of-work work, in addition to my life. Much of last year I was burning the candle 
from both ends; late night and early hours. Honestly, I got *more* done, but it wasn't as good as 
the work I'm doing now. I've settled into a sleep in until 6:30-7 schedule, and I feel much more 
rested with more clarity of thought. In the past, I was up at 5, and it took me 30 minutes just to 
wake up! 

We also got out for an overnight camping trip before it starts to get cold. Camped right next to 
a river watching the fire is a luxury I don't take for granted. What I take for granted is how 
quickly time (i.e. our life) flies by between such adventures. Keeping grounded and getting away 
from the house has been another good move over the last couple of months - if I'm at home, I 
can't help but be *plugged* in.

**What can I do better?**

- I need to get in front of issues sooner, and have those hard conversations with people more often. At times I'm too passive or indifferent.

**What have I done well?**

- Took several days to design and proof-of-concept the multi-account architecture and it made the development so much faster.

## Next month's goals

- Deploy the multi-account feature

## Analytics

{{< plausible start=2022-05-01 end=2022-05-31 site_id=mudmap.io >}}


{{< plausible start=2022-05-01 end=2022-05-31 site_id=check-redirects.com >}}


{{< plausible start=2022-05-01 end=2022-05-31 site_id=danielms.site >}}


[wgd]: https://whatgotdone.com
[mudmap]: https://mudmap.io?ref=danielms.site
