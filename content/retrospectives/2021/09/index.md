+++
date = "2021-10-04"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-sept-2021"
title = "September 2021 Retrospective"
draft = false
+++

# Summary

Finally, released my firewall feature to Mudmap. It's a small step towards having a more fully 
featured application for pfSense users.

## Highlights

- Deployed a big feature to prod
- Took on a lot of learning for my day job - cutting into Mudmap's time

## Goal Performance

A review of last months three goals. See [August's Retrospective][old-retro]

[old-retro]: /retrospectives/2021/retrospective-aug-2021/

### Release firewall rules page

- **Appraisal**: I've done it, finally!
- **Rating**: A

After what feels like a lifetime I've pushed big feature to [Mudmap]. Even saying that makes me 
shudder, it's a bad thing - features shouldn't take so long to deploy. Regardless, I am happy 
its out and glad to move on to the next thing. I keep asking myself why it took so long. I think 
it comes down to two main things; time and complexity.

**Time**

It is just me working on [Mudmap], and I already work a full-time job. On one hand, it's amazing 
what you can accomplish over time as your small increments compound. On the other, if it was my 
full-time job those increments would compound much faster. Acknowledging this fact is important 
but it's only half the battle. I also spend more time than I should on things which *do not* add 
value to the project. Battling with tooling issues or lack of foresight for issues that would 
cost me precious time later, are especially damaging. As I won't be quiting full-time employment 
any time soon, time or the lack thereof will be a continued factor in the velocity of release.

**complexity**

I have definitely made things more complex than they need be in certain aspects. In hindsight, 
it is obvious but only now do I possess such clarity. Further, pfSense is a complex beast in of 
itself and building an API plus management dashboard for it isn't trivial. I cannot decrease the 
system complexity of pfSense but I can reduce *my* own systems. It starts with more time allotted 
to design and proof of concept work instead of ploughing head first into things I *think* can work. 

![](name-of-image.png)

### Write a Mudmap blog post

- **Appraisal**: I have it written just not published
- **Rating**: C

I can deploy what I have written to prod, but it feels half-baked or even nebulous in hindsight. 
I've decided to sit on it and maybe leave that post as a forever draft instead. As a result, 
this is unfinished.


## Recommendations

After years of pestering by Amazon, I finally trialled Prime. Before I mention the show I really 
enjoyed I want to say two things about Prime Video which suck.

1. Paying for movies and shows (in addition to the yearly subscription)
2. Search absolutely sucks

When I searched for a show that my daughter was begging me to watch, all I could remember is 
that it had "princess" in it. So I searched for "princess". No results. Tears ensued. A week 
later, my wife found it - "princess and the dragon". Now what crappy search engine cannot return 
"princess and the dragon" when searching "princess". 

Onto my recommendation. [Clarksons Farm](https://en.wikipedia.org/wiki/Clarkson%27s_Farm) is 
hilarious and well worth a watch. I'd also recommend watching all episodes from season four of 
[The Grand Tour](https://en.wikipedia.org/wiki/The_Grand_Tour) - I haven't laughed like that in 
a while.

## Wrap up

The first half of September was me working really hard to push out the firewall feature with the 
second half mostly devoted to things unrelated to it. My job - the one that pays my bills - is 
still important to me and must take priority when needed. We're moving to a new platform, have 
several big features to produce and need to upskill all at the same time. For me personally, as 
soon as I deliver my current scope of work, I could be seconded to another team to aid their 
work on the container platform. So this has required me to commit some of my in and out-of-hours 
work to Kubernetes education. In between that, time has also been spent increasing my 
proficiency in Go for potential job opportunities that (have) and may come up in the future.


**What can I do better?**

- Plan and prove that features or ideas are going work/fail before jumping in. Patience

**What have I done well?**

- Deployed without issue a long awaited feature

## Next month's goals

No goals, or rather I am taking the month to freestyle the next thing to do. Not setting goals 
is setting a goal to fail, or something like that, but I need a month to relax my mind. My wife 
has had some health issues, COVID garbage is constant and working from home whilst homeschooling 
ain't easy. But, that should all ease off towards the end of the month, so I'm taking until then 
to just experiment.


## Analytics

**Mudmap**

![](mm-sept-stats.png 'Mudmap plausible stats for September 2021')

**Check-Redirects**

![](cr-sept-stats.png 'Check-Redirects.com plausible stats for September 2021')

**Twitter**

I really haven't been playing the Twitter game much lately. Struggle thinking of things to say!
![](twit-sept-stats.png '@dansult twitter stats for September 2021')



[mudmap]: https://mudmap.io/?utm_campaign=retro&utm_source=danielms&utm_medium=blog
