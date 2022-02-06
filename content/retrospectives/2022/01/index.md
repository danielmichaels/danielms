+++ 
date = "2022-02-01"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-jan-2022"
title = "Jan 2022 Retrospective"
draft = false 
+++


# Summary

This month I deployed the new version of [Mudmap] - a complete re-write using Golang.

## Highlights

- Deployed version 2 of Mudmap 

## Goal Performance

A review of last months three goals. See [Decembers's Retrospective][old-retro]

[old-retro]: /retrospectives/2021/retrospective-dec-2021/

### Push Mudmap version 2 to production

- **Appraisal**: Completed
- **Rating**: A+

After checking my git history, I was surprised to see that my first commit for the new version of 
Mudmap was on the 5th of November. On January 10, I officially pushed v2 to production. 
Considering I was maintaining an existing product *and* rewriting it in another language - which 
is still relatively new to me - this feels like quite an achievement. 

#### Did it improve performance, and what was your reasoning?

Here is a [twitter thread](#tweets) with visuals. But the tl;dr is **yes**, memory consumption 
alone is remarkable. With my Django app, even a small load on the server could run the risk of 
pushing the application above the plan limits. Now, I am nowhere near it.

The difference between container image sizes are ridiculous as well. Somewhere around 15mb for 
the Go image and 1GB for Django. Build times feel days apart - from ~10 minutes using 
python to under a minute using Go. A pet peeve of mine is how slow python image building is, and 
I am glad to be rid of it. Another thing I'm happy with is the cost reduction in terms of the 
number of services needed. I save at least $14 per month by not needing redis or celery. I'm 
even considering removing Postgres for [Litestream] as I now only store basic user and device data.

There is a [notion document][r] with some of my reasons for switching from Django to Go. Of course, 
in almost any of those statements, a counter-point could be constructed. That is fair, but 
regardless, they are **my** reasons for such an undertaking.

### Add at least one more core feature to Mudmap

- **Appraisal**: Completed
- **Rating**: B

Mudmap now supports three helpful but not *core* features; shutdown, reboot and shell command 
execution (see them in [action][vsrs]). Admittedly, these are useful, but I did not deliver the 
bigger feature I've been working on; the ability to create, read, update and delete interfaces. 

I did update the documentation, landing page and recorded five YouTube [explainer videos]. 
It is the first time I have recorded and uploaded anything to YouTube, and it shows in my 
delivery. I opted to push what I recorded rather than waste time finessing and re-recording it 
again and again. Reaching for pragmatism not perfection. 

As I've transitioned (but still support version 1) to version 2, the documentation needed 
updating. This meant updating all the screenshots, links and explanations. I also added the 
videos to the docs as some people prefer videos over a wall of text.

Mudmap's [landing page][mudmap] also got a little better. Previously, I was using a placeholder 
image in the main section but replaced it with a screenshot of the dashboard. The [overview][vov]
video is now embedded into the page as well.

[explainer videos]: https://www.youtube.com/channel/UCtRlcQftzThqR5Q5iaOVYTA
[vsrs]: https://docs.mudmap.io/videos/demo-diagnostics?ref=retro-jan-2022
[vov]: https://docs.mudmap.io/videos/overview-video

## Recommendations

[My First Million][fc] episode on [Kevin Van Trump] and [FarmCon] - its incredible. This guy 
runs a **daily** agricultural investing/analysis newsletter which brings in ~30M USD a year. He's 
been writing it for 17 years and now runs [FarmCon]. I've been binge-watching all the videos on 
YouTube and cannot praise this enough - talk about finger on the pulse.

[OpenFaaS] for serverless functions. I've never been much interested in serverless, mostly 
because I don't like using AWS (etc) for small workloads/one-off tasks. And, when I've looked 
at it before it didn't really gel with me - another skill to learn. Then I found [OpenFaaS] - 
an open source serverless technology. Using OpenFaaS, I can write functions in almost any language with 
my two preferred options being python (Flask) and Go. It also supports NET, C#, Node, Java and 
more (I think). I am using it to render the [analytics](#analytics) charts below. It supports 
Kubernetes and a standalone binary called [faasd] for deployment to a small server - I run it on 
a $5 droplet.

## Wrap up

A good start to the year. I took the Christmas and New Year period off and also took a 
four-day long weekend at the end of this month. Only during the four-day break by the sea with 
just my little family did I *start* to feel rested. If I can, I will try to schedule more of 
these little breaks to ward off burn out. 


**What can I do better?**

- Not get as distracted by other things; [OpenFaaS] took some time away from Mumdap this month
- I need to read more books and not look at my phone before bed

**What have I done well?**

- Integrated Stripe and deployed Mudmap's new version early in the month
- Wrote several emails to customers, a blog post and newsletter announcing it

## Next month's goals

- Deliver the Interfaces feature for Mudmap
- Write at least one blog post for Mudmap

 
## Analytics

{{< plausible start=2022-01-01 end=2022-01-31 site_id=mudmap.io >}}


{{< plausible start=2022-01-01 end=2022-01-31 site_id=check-redirects.com >}}


{{< plausible start=2022-01-01 end=2022-01-31 site_id=danielms.site >}}

### Tweets

{{< twitter id=1480562848791105536 user=dansult >}}

[mudmap]: https://mudmap.io?ref=danielms.site
[pfmonitor]: http://pfmonitor.com?ref=danielms.site
[post]: https://www.reddit.com/r/PFSENSE/comments/9u1w4d/is_pfcentre_centralised_cloud_management_still_on/
[Newcastle]: https://www.visitnewcastle.com.au/?ref=danielms.site
[dvassallo]: https://twitter.com/dvassallo
[fc]: https://www.youtube.com/watch?v=ho22vkGFljg&ab_channel=MyFirstMillion
[Kevin van trump]: https://twitter.com/kevinvantrump
[farmcon]: https://www.farmcon.com/
[openfaas]: https://openfaas.com
[faasd]: https://docs.openfaas.com/deployment/faasd/
[r]: https://mudmapio.notion.site/Version-2-d78ca9bd813541738f7c71cfb9c95c9e
[litestream]: https://litestream.io
