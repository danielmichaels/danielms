+++ 
date = "2022-07-06"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-june-2022"
title = "June 2022 Retrospective"
draft = false 
+++

# Summary

Moved house and still haven't quite finished my one goal from last month.

## Goal Performance

A review of last months three goals. See [May's Retrospective][old-retro].

[old-retro]: /retrospectives/2022/retrospective-may-2022/

### Deploy the multi-account feature

- **Appraisal**: Some hiccups slowed development
- **Rating**: B

All of this month's Mudmap development was dedicated to multi-account feature. 
Unfortunately, that work is still ongoing but at least right now its down to fleshing out 
the user interface. So far, users can invite members to join their organisation, join other 
organisations, remove members from their own, update org details, and manage devices within 
their organisation based on simple permissions. 

At the beginning of the month I spent a lot of time building on top of the [Auth0 Authorization]
extension. This worked well in theory as it matched much of my vision for what I wanted to achieve. 
Unfortunately, it did not fit well into the application of that vision; simple actions require many
API calls. For instance, I could find no easy method to get the `_id` of a `group` without iterating
over every `group` and searching for its mutable `name`. To do most operations against the `group`
you need this `_id`, meaning lots of API calls and then loops to find the `_id`. This means 
latency but also requires a fair bit of error handling. In all, it *felt bad* and without 
storing that data in the database was slow.

Luckily, I went back to the much simpler and easier to use [Auth0 Management] API. This in hindsight
does everything I wanted the [Auth0 authorization] extension to do but with 1/10th the complexity. 
Building, integrating, and testing the authorization extension took me about a week of evenings but
replacing it with the management API took about three hours, and it removed a significant amount 
of code.

In hindsight, the toil and wasted effort building out the organisations using Auth0's 
authorization extension wasn't wasted. I think it actually helped to design a much simpler and 
easier to maintain product. The effort helped shape the *how* and *what* of its current form. Still,
it definitely slowed the deployment of the feature.

I also started on a side project to service my own needs - [storeman] - which cut into the feature 
development.

## Recommendations

- If you can afford it, get removalists to lift and shift your house. My hubris and tightwad-ism 
meant
I moved my entire house on my own minus a mate who helped lift the real heavy stuff. The entire 
process absolutely cooked my lower back, and I'm only now starting to recover. 
- Peaky Blinders, the latest season. I watch about 30-60 minutes of television a day so I'm picky 
about what to tune into on the idiot box. This season of Peaky Blinders was the best by far, 
really hoping they make another season.

## Wrap up

This is the second month when I've bombed on my targets. Still, I'm happy to be continually 
making progress on my objectives. It might be costing me some business in the short term, as 
*user-facing* features aren't being developed quick enough. There isn't much more time I can steal
from my days to get more done, but I think I could be smarter in my planning - hindsight, of course.


**What can I do better?**

- Plan more, code less - meaning, write more documentation of what I want to achieve before I try to achieve it.

**What have I done well?**

- Accomplished a lot in terms of functionality especially when I lost a number of days due to the move.

## Next month's goals

- Deploy the organisations feature
- Allow users to back up their device from within Mudmap

## Analytics

{{< plausible start=2022-06-01 end=2022-06-30 site_id=mudmap.io >}}


{{< plausible start=2022-06-01 end=2022-06-30 site_id=check-redirects.com >}}


{{< plausible start=2022-06-01 end=2022-06-30 site_id=danielms.site >}}


[wgd]: https://whatgotdone.com
[mudmap]: https://mudmap.io?ref=danielms.site
[Auth0 Authorization]: https://auth0.com/docs/api/authorization-extension
[auth0 management]: https://auth0.com/docs/api/management/v2
[storeman]: https://github.com/danielmichaels/storeman
