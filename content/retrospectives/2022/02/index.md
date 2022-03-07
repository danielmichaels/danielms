+++ 
date = "2022-03-06"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-feb-2022"
title = "Feb 2022 Retrospective"
draft = false 
+++


# Summary

A number of external influences impeded my ability to do get much work done this month. As 
a result very little was achieved on Mudmap.

## Highlights

- Small fixes for Mudmap installer
- Built a Cobra CLI for interacting with an Australia crypto exchange

## Goal Performance

A review of last months three goals. See [January's Retrospective][old-retro]

[old-retro]: /retrospectives/2022/retrospective-jan-2022/

###  Deliver the Interfaces feature for Mudmap

- **Appraisal**: Failed
- **Rating**: F

I didn't even get close to completing this.

A bit of background. A number of things have been happening this month in my personal and work 
life - which obviously share resources with my development time. My entire family got quite sick 
though I suffered the worst, and the kicker, it was plain old flu. I actually had someone tell 
me I am wrong here and that it *must* have been COVID but just a false negative. Give it up 
people, the cold and flu still exist! 

Additionally, I've been moved to a senior position in another team at work. The team is at the 
point of complete atrophy with everyone jumping ship or taking jobs elsewhere, and unfortunately,
I am the *best* candidate to help out. I cannot begin to get started on how badly everything has 
fallen apart over the last 6 months - I feel like I'm working at Parts Unlimited from the 
[Phoenix Project](https://www.amazon.com.au/Phoenix-Project-DevOps-Helping-Business/dp/0988262592). 
We even have our own Brent!

Needless to say, my mind has been quite occupied over the last few weeks meaning time for Mudmap 
has been low. 

### Write at least one blog post for Mudmap

- **Appraisal**: Failed
- **Rating**: F

For the same reasons above I have had no time to do this task either.

## Recommendations

I am about two years late to the party but [Ted Lasso] has to be the most wholesome and 
feel-good show I've watched in a long time. Some of it really hits home for me, in the past I 
have been in a high performance team with high stakes. I'd deal with the team dynamics of highly 
talented, motivated and sometimes abrasive individuals, and it is tough sometimes. So this 
series made me happy, sad, reflective and proud all at the same time. The soundbite that I'll 
take away from it is *be curious, not judgemental*. It is something I need to remind myself daily.

After 6 weeks of training I competed in my first open water swim. It was a tough two kilometers 
in some decent swell; 3m at 10 second intervals. But, it feels like I've accomplished something
worth being proud of and most importantly it aligns with my true self. The points in my life that I 
look back upon with the most fondness where also the most painful and required a lot of 
commitment and dedication.

> “In some ways suffering ceases to be suffering at the moment it finds a meaning, such as the 
> meaning of a sacrifice.” - Viktor Frankl

I recommend doing something hard and painful every now and again. Get out of your chair, don't 
take the umbrella and feel rain on your skin or challenge yourself by going for a [Park Run]. 
Just do something where all you can think about for a few minutes is survival - most problems 
aren't that bad when you have such a contrast.

[ted lasso]: https://en.wikipedia.org/wiki/Ted_Lasso
[park run]: https://www.parkrun.com

## Wrap up

A short retro after a month of very little in the way of accomplishments. I did create a small 
[Go CLI][cli] using Cobra for a popular Australian crypto exchange (still a WIP). This was 
mostly to enable my friend's analysis of trade conditions but also to get more acquainted with 
writing Go CLI's. This is an area where I believe Go really shines, especially with its ease of 
producing cross-platform binaries. Whilst I may have accomplished little on Mudmap I did create 
something useful for a friend - a silver lining for the month.

In March, I hope to be a bit more productive on Mudmap, or if I am still unable to produce much 
assess my ability to keep going with the project. This is that is weighing on my 
mind, especially in light of Netgate's recent announcements. Effectively, they are looking to 
transition pfSense CE from development to maintenance mode. At least that is the feeling amongst the 
community, and whilst I do understand their desire to move in this direction, I do see that as a 
potential threat to Mudmap's long term viability.

**What can I do better?**

- Write my weekly [What Got Done][wgd] on Friday instead of delaying until the weekend or later

**What have I done well?**

- I've done a lot to help out a friend with their own project, mostly on the deployment side
- Allowed myself to take time off when needed

## Next month's goals

This month I'm setting no goal other than to help my sick partner recover. As much as it feels 
cheap from a product standpoint family trumps penny stocks.
 
## Analytics

{{< plausible start=2022-02-01 end=2022-02-28 site_id=mudmap.io >}}


{{< plausible start=2022-02-01 end=2022-02-28 site_id=check-redirects.com >}}


{{< plausible start=2022-02-01 end=2022-02-28 site_id=danielms.site >}}


[cli]: https://github.com/cupscanteen/swyftx-cli
[wgd]: https://whatgotdone.com
[mudmap]: https://mudmap.io?ref=danielms.site
