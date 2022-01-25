+++
date = "2021-08-04"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-july-2021"
title = "July 2021 Retrospective"
draft = false
+++

# Summary

This month really got away from me. Mudmap's development has not slowed, but it has not been 
hitting the goals I've set this month. Instead, I've adapted to a changing landscape and user needs.

## Highlights

- Mudmap gets a free tier
- Static IP addresses are now standard for Mudmap's servers
- A shift in focus and support for pfSense versions 

## Goal Performance

A review of last months three goals. See [June's Retrospective][old-retro]

[old-retro]: https://danielms.site/retrospectives/2021/retrospective-june-2021/

### Cold email at least 15 people

- **Appraisal**: Probably set the bar too low here; only a few replies.
- **Rating**: A

I did this, and got probably the market expected percentage of responses. All the responses 
mentioned a need for locking down their SSH port with some static IP addresses. 

It was from this feedback that I chose to prioritise the deployment of a proxy in front of 
Mudmap. Setting the proxy allowed for a pair of static IP's to be placed in front of Mudmap.
The proxy is completely transparent to users even with SSH tunnelling. By doing this, users can 
now set source addresses for SSH - a huge security increase.

Some might wonder why my servers don't have a static IP already? The hosting platform, 
[Render.com](https://render.com), does not provide static addresses for any of its containers. 
A third party service, [QuotaGuard](https://quotaguard.com) integrates well with Render and 
hooking it into my application took about 2 hours. I am pretty happy with the service and whilst 
they don't offer a free tier, I feel the starter plan is quite generous. 


### Add firewall rules to Mudmap

- **Appraisal**: C
- **Rating**: I have built the read-only pages but paused development due to important issues that arose.

So far, Mudmap's feature set is limited to a only a portion of pfSense's. The vision is grand 
and will eventually cover the majority of its features. Firewall rules, including the ability 
to read, create, update and delete them is a high priority. Unfortunately, other issues came up 
that forced this into the paused, or blocked state. I could deploy it as a read-only copy but 
have chosen to instead deploy it once its finished properly.

A work in progress look.

![](https://mudmapio.s3.us-west-2.amazonaws.com/public-images/marketing/mm-firewall-rules-wip.png 'Mudmap firewall rules work-in-progress image')

### Record some videos for onboarding new clients

- **Appraisal**: Did not start this.
- **Rating**: F

I did not even get started on this. 

I did create a landing page for new users inside the application itself. Initially, when a user 
register an account and logged in they were presented an empty table. That's not very welcoming 
and makes the assumption that users *know* how to use Mudmap. Now they get presented with a few 
helpful links to get started - this will eventually be upgraded again, likely with a welcome video.

The welcome page. Unfortunately, I didn't get screenshot of the empty tables as a comparison.

![](https://mudmapio.s3.us-west-2.amazonaws.com/public-images/marketing/mm-no-devices.png 'Mudmap firewall rules work-in-progress image')

## Free Tier

Mudmap is new, lacks social proof and is chasing a market full of security minded folks. 
Initially, to get started using the platform I was asking for a subscription. But, this was, as 
you would expect, not working too well. Before someone could even evaluate the performance, 
reliability and security of Mudmap they were asked for a payment method. I wanted to reduce 
this friction and increase user uptake.

So I decided to offer a *free tier* for all users.

Every user can now add their first **two** devices for free. No credit card needed and no time 
limit. It also lets small businesses or hobbyists who have a couple of devices test it out.

In hindsight this should have been released from the start. I'm hoping this will bring in some 
more feedback of the system too.

## Social Media

This month I spent some time upping Mudmap's social media presence. You can now find it at:

- [LinkedIn](https://linkedin.com/company/mudmap)
- [Twitter](https://twitter.com/mudmapio)
- [Buttondown](https://buttondown.email/mudmapio)

It's always that fine line between development of the product and marketing. The developer in me 
always thinks it is not good enough to spruik yet. This, I feel, is a problem all developers have.

## A problem and shift in Mudmap's offering

Late this month, I started to see a uptick of users and a correlating spike of errors. The 
investigation into this lead me to a resolution I always felt was probable but hopefully unlikely.

**tl;dr** Mudmap is ceasing support for pfSense+

At the start of this year, Netgate elected to split pfSense into a closed and open-source model.
The newer and closed source [pfSense+] is for the time being largely the same as pfSense 
Community Edition (CE). Unfortunately for Mudmap (yet, fortunate for everyone else) they've 
made changes that have effected the product.

Initially, I attempted to reverse engineer a workaround to the most recent breaking change 
between pfSense+ and Mudmap. But after careful consideration, I decided this course of action 
is not in the best interested of customers. 

Why? Providing a reliable and safe platform from which customers can manage multiple 
pfSense devices remotely is Mudmap's core mission statement. This cannot be achieved without a 
stable platform to build from. Providing an interface that could potentially cause disruption 
or worse for customers is a risk I *will* not take.

Naturally, I am disappointed, and it will reduce Mudmap's market appeal - possibly quite 
significantly - but I'm not losing hope.

So where to from here?

Mudmap will continue to support pfSense CE and develop functionality to meet the needs of that 
user base. I'll also be keeping a close eye on the developments of pfSense+. I think it is 
important to disclose how excited I am for it and that I fully support Netgate's decision to 
close source the project. I expect pfSense+ to be a cut above the rest when it re-launches 
after its rewrite in [Golang]. This should also deploy with an API very similar to [TNSR]'s. As 
both a user and developer, this makes me really happy. It should also allow Mudmap to push back 
into the market of supporting pfSense+ clients as I will (hopefully) be able to hook into the 
new API.

I have written a page in the [documentation] explaining the change in support, updated the 
[homepage][mudmap] and made a toggle in the application itself ensuring that users are aware of 
this. I will be writing a blog post, newsletter and LinkedIn post in response as well.

## Recommendations

**Interesting topics**

Whilst I am not a pentester or bug bounty hunter, I thoroughly enjoy watching videos on the 
subjects. Having binge watched most of [ippsec]'s videos, I recently started on [John Hammond]'s 
stuff. If you're a developer, there is a lot of information to glean from how these guys go 
about exploiting *your* servers.

Worth a mention, [ippsec] has a handy [search page][ippsec-search] where you can filter topics 
and find which video to watch based on it. For instance, if you searched `jwt`, it will show 
you a selection of boxes he's pwned with JWT's in there somewhere.

**React development cheat code**

If you aren't using `swr` from [Vercel], maybe you should be! `react-query` is apparently just 
as good - and I believe it, it's from [Tanner Linsley]. I can only talk about `swr` and say that 
its made my Next.js app blazing *fast* - it literally feels like a cheat code for client side 
web. It integrates perfectly with `axios` too. As a heavy user of `axios` interceptors for JWT 
refreshing, I initially saw all the docs using `fetch` and thought maybe interceptors wouldn't 
work. Thankfully it worked flawlessly.

Podcast of the month:

- Everything on the [How to Take Over The World Podcast].  My favourite is the Putin series and Caesar.

## Wrap up

This month was a bit of test. It started quite well but sadly ended on a sour note. Nonetheless, 
I'm looking forward to August and producing more content for users of Mudmap to enjoy. 

**What can I do better?**

- Market Mudmap's potential rather than internally focus on what it's not providing right now

**What have I done well?**

- Adapted to changing circumstances and took action when needed.

## Next months goals

- Actually release the firewall rules pages
- Publish a Mudmap blog post

## Analytics


{{< plausible start=2021-07-01 end=2021-07-31 site_id=mudmap.io >}}


{{< plausible start=2021-07-01 end=2021-07-31 site_id=check-redirects.com >}}


{{< plausible start=2021-07-01 end=2021-07-31 site_id=danielms.site >}}


**Twitter**

![](twitter-stats.png 'Mudmap plausible stats for July 2021')

[mudmap]: https://mudmap.io/?utm_campaign=retro&utm_source=danielms&utm_medium=blog
[documentation]: https://docs.mudmap.io/pfsense-community-and-plus
[ippsec]: https://www.youtube.com/channel/UCa6eh7gCkpPo5XXUDfygQQA
[ippsec-search]: https://ippsec.rocks/?#
[john hammond]: https://www.youtube.com/channel/UCVeW9qkBjo3zosnqUbG7CFw
[How to Take Over The World Podcast]: https://httotw.com/
[pfsense+]: https://www.netgate.com/blog/announcing-pfsense-plus
[golang]: https://www.netgate.com/blog/pfsense-plus-pfsense-ce-dev-insights-direction
[tnsr]: https://docs.netgate.com/tnsr/en/latest/api/
[tanner linsley]: https://react-query.tanstack.com/
[vercel]: https://swr.vercel.app/
