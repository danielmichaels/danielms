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

I check my git history and was surprised to see that my first commit for the new version of 
Mudmap was on the 5th of November. On January 10, I officially pushed v2 to production. 
Considering I was maintaining an existing product *and* rewriting it in another language - which 
is still relatively new to me - this feels like quite an achievement. 

#### Did it improve performance, revenue, or was this just a waste of time?

Here is a [twitter thread](#tweets) with visuals. But the tl;dr is **yes**, memory consumption 
alone is remarkable. With my Django app, even a small load on the server could run the risk of 
pushing the application above the plan limits. Now, I am nowhere near it.

The difference between container image sizes are ridiculous as well. Somewhere around 15mb for 
the Go image and 1GB for Django. Build times are days apart in time - from ~10 minutes using 
python to under a minute using Go. A pet peeve of mine is how slow python image building is, and 
I am glad to be rid of it. 

### Add at least one more core feature to Mudmap

- **Appraisal**: Completed
- **Rating**: B

Mudmap now supports three helpful but not *core* features; shutdown, reboot and shell command 
execution (see them in [action][vsrs]). Admittedly, these are useful, but I did not deliver the 
bigger feature I've been working on; interfaces. 

I did update the documentation pages, landing page and recorded five YouTube [explainer videos]. 
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
been writing it for 17 years and now runs [FarmCon]. 

## Wrap up

A good start to the year. I took the Christmas and New Year period off and also took a 
four-day-long weekend at the end of this month. Only during the four-day break by the sea with 
just my little family did I *start* to feel rested. If I can, I will try to schedule more of 
these little breaks to ward off burn out. 


**What can I do better?**


**What have I done well?**


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
