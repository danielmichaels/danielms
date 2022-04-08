+++ 
date = "2022-04-06"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-march-2022"
title = "March 2022 Retrospective"
draft = false 
+++


# Summary

After taking some time off to have fun and experiment, I got back into the swing of Mudmap 
development.

## Highlights

- Released several updates to Mudmap
- Got Covid (not really a highlight but noteworthy)

## Goal Performance

A review of last months three goals. See [Feb's Retrospective][old-retro]

[old-retro]: /retrospectives/2022/retrospective-feb-2022/

I did not set any goals for this month. My partner has been quite unwell, and I decided to take 
some pressure off and be more available for her. 

## Month overview

I discovered the project [Alpine.js][alpine] this month and was really intrigued by how 
user-friendly it is. I am pretty comfortable using React these days but the added complexity of 
running a separate frontend can get a little annoying, especially for simple app's. Setting up 
Alpine is a [cinch] and can easily be integrated into any web application by either dropping in 
a link to the CDN, or by creating a bundle for it and linking it directly.

Where [Alpine] shines is when coupled with traditional web applications such as Django, 
or Go templates. To get more familiar with it, I refactored an existing app of mine - [tars.run].
It was originally built with a Next.js frontend but was overkill for what it does. Now, instead 
of needing two applications, the entire thing including database runs inside a single binary. It 
is refreshingly easy building app's this way.

## Mudmap updates

I have received a number of requests from customers asking Mudmap to support devices which do 
not have static IP addresses. I finally added this feature and it was so simple that I regret 
not doing it sooner. 

Mudmap now shows each device's DHCP lease table too. This feature is not fully finished as 
you cannot update any of the leases, only view them. I chose to release it in stages so 
that customers can start using it straight away, rather than having to wait for all aspects of 
that feature to be completed first. This is probably a theme I will continue with, and it allows 
for a faster capture of feedback too.

The application user interface also received a bit of love. Most of the tables now support 
pagination by default, have smaller margins and where appropriate have tool-tips to reduce 
ambiguity. Many of the table buttons are now smaller, or represented by a more familiar icon - the 
play button for example. These small fixes do a lot to make the user interface feel more polished.


[tars.run]: https://tars.run?ref=wgd
[alpine]: https://alpinejs.dev
[cinch]: /blog/alpine.js-and-tailwind-html-setup/

## Recommendations

Re-watch Seinfeld. It's been such a brilliant way to decompress before bed and even after 30 years 
it is still hilariously relevant. I am about half-way through the entire show after nearly a 
month of watching an episode or four each evening. So good.

I have completely moved from Arch (actually, Manjaro) to Ubuntu and couldn't be happier. 
Honestly, I'm not sure if I'm missing anything from the move - any package that I had on Arch I 
probably have on Ubuntu. And, I actually like [snap] packages - I said it! I've had a couple of 
little issues, but they've all been super easy fixes. 

**fighting words**: Ubuntu is the closest thing we have to linux desktop, and I'm 
now a *recommender* of the distro.

[snap]: https://snapcraft.io

## Wrap up

I expected this month to unproductive but even after getting Covid, it still ended up being quite
good. I was able to use my time to experiment with things and just have some fun which gave me 
the break I needed to get back into Mudmap development. It was also good just coming home and 
relaxing some nights, without the guilt I've put on myself to keep working on things like Mudmap.
This has been especially important as my day job's workload has increased. In all, March was a 
far better month than February.

**What can I do better?**

- Write in my journal every day, even if only a couple of words
- Read more often before bed

**What have I done well?**

- I Wrote two short blog posts this month
- Didn't take myself too seriously!

## Next month's goals

- Push two Mudmap features to production, no matter how small
- Walk with my daughter to the [telegraph station] whilst we're in [Alice Springs][alice] in April

[telegraph station]: https://en.wikipedia.org/wiki/Alice_Springs_Telegraph_Station
[alice]: https://en.wikipedia.org/wiki/Alice_Springs

## Analytics

{{< plausible start=2022-03-01 end=2022-03-31 site_id=mudmap.io >}}


{{< plausible start=2022-03-01 end=2022-03-31 site_id=check-redirects.com >}}


{{< plausible start=2022-03-01 end=2022-03-31 site_id=danielms.site >}}


[wgd]: https://whatgotdone.com
[mudmap]: https://mudmap.io?ref=danielms.site
