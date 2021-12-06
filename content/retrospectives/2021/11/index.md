+++
date = "2021-11-07"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-nov-2021"
title = "November 2021 Retrospective"
draft = false
+++

# Summary

Screw it, I'm rebuilding Mudmap's backend.

## Highlights

- 80% feature parity between Mudmap version 1 and now
- Decided to withdraw from MSFT interviews
- Started planning new user interface enhancements

## Goal Performance

A review of last months three goals. See [October's Retrospective][old-retro]

[old-retro]: /retrospectives/2021/retrospective-oct-2021/

### Mudmap feature proof-of-concept

- **Appraisal**: A PoC turned into a rewrite, and it's working out great
- **Rating**: A

Firstly, I want to acknowledge two things.

1. Everyone says don't rewrite, it is a time sink or that it's chasing waterfalls
1. There is nothing wrong with Django - it's great

That out of the way, let's continue. This started as a prototype to see if I could use Go for 
sending SSH traffic more efficiently. It pretty quickly devolved into a bunch 
of HTTP handlers with SSH connections inside them.

Somewhere within the [first half of the month][0], I started to feel a lot more productive and in 
control by re-implementing sections of Mudmap in Go. It was at this point that I realised it's 
time to commit and just do it. Despite this, I did have a number of self-conscious 
thoughts about undertaking this task. But, ultimately they all boiled down to me *caring about what 
others think* which is an indicator that you're doing, or not doing, something for the *wrong* 
reasons.

After a month of work, Mudmap V2 is about 80% feature parity with my in production V1 codebase. 
It also has more tests and better coverage than the current application and some components 
even have additional features. Something I'm really pleased with is how easy it is to add new 
features using Go - its type system and interfaces are great for this. 

This is a personal thing and not slight on Django, but I've found that Go has forced me to think 
more critically about how the application should be built. Django's beauty can also be its curse 
as there is a lot of *magic*. For a CRUD application (and especially one without DRF), it's 
wonderful. For me, some of that magic didn't work well for what I'm trying to achieve with 
Mudmap. For instance, my endpoints make calls to another API using SSH. This is not something 
Django expects and takes away the brevity of [Class Based Views][2]. I'm also not a huge fan of 
the serializers used by DRF (which may as well be a Django core module these days), whereas I am 
a big champion of Pydantic/FastAPI's type system. Go feels like a step above both options but 
it's a subject comparison. 

In hindsight, Go feels like the better option for a couple of other reasons which are mostly 
related to the developer experience. It *has* made me think more critically about what I am 
trying to accomplish at every step along the way. How? Firstly, in Go, a little copying 
is better than a little dependency. This has made me do a lot more research where in the 
past I may have just `pip install`'d something, and it has an added benefit of making me read a 
lot more source code, a powerful educator. I also love seeing how changing a functions' signature 
can ripple across the entire app, which has also made me plan further ahead. Lastly, Go is a 
concurrent language. Mudmap sends emails and executes things in background workers without Redis 
or Celery, and it can handle a boatload of requests. I'm still coming to grip with the fact 
that I don't need Gunicorn, its dev server is its prod server, and it's fast.

In all, I think this has been a good move and feel more confident about feature development and 
maintaining this codebase than I do with the current app. 

[0]: https://whatgotdone.com/dansult/2021-11-12
[2]: https://docs.djangoproject.com/en/3.2/topics/class-based-views/intro/

### Study at least 2 pomodoro’s each day

- **Appraisal**: This or rewrite Mudmap. This lost.
- **Rating**: C

This spawned from interviewing with Microsoft for a position within Azure. I put in a number of 
hours at the start of the month, working through the typical developer interview type questions. 
After about 10 days, I gave it away to spend more time on Mudmap but in the process did brush up 
on some of the basics.

I guess, this is a two-part reason as to why I stopped studying; I don't really want that job,
mostly because they want me to stay local and work in the office. That's not something my family,
or I want anymore. We're keen to head back to the coast where we lived for ten years before coming 
here. Also, that crap is utterly boring and saps my energy - if it means I never work for a Big 
Corp, so be it. 

## Recommendations

We started watching [The Expanse][3] on Prime and after three seasons am totally hooked. If you 
like space, it's actually really good.

I like to listen to music when working but get distracted by certain genre's. It means I have to 
pick music that is rather mellow. As much as I love 90's/2000's hip hop, it is too stimulating 
for me and makes it hard to concentrate. This month I found a really [chill playlist][4] which 
has replaced my [Interstellar][5] and [Tron][6] soundtracks.

[4]: https://open.spotify.com/playlist/35fMnNReBETCnQ0CH5CHug
[5]: https://open.spotify.com/album/3B61kSKTxlY36cYgzvf3cP
[6]: https://open.spotify.com/album/3AMXFnwHWXCvNr5NCCpLZI
[3]: https://en.wikipedia.org/wiki/The_Expanse_(TV_series)

## Wrap up

I feel like I am back this month after a sluggish August-September-October. It might be the end of 
lockdown, the 
increasingly nicer weather or just because I took some time to slow down but whatever it is I 
feel much more motivated. 

**What can I do better?**

- More exercise
- More engagement with customers

**What have I done well?**

- Focused on completing one task well before starting the next
- Prioritised what matters and brushed off meaningless tasks

## Next month's goals

- Plug in the user authentication to Mudmap
- Email customers about upcoming changes to Mudmap

## Analytics

**Mudmap**

![](mm-nov-retro.png 'Mudmap plausible stats for November 2021')

**Check-Redirects**

![](cr-nov-retro.png 'Check-Redirects.com plausible stats for November 2021')

**Check-Redirects**

![](danielms-nov-retro.png 'danielms.site plausible stats for November 2021')

**Twitter**

![](twit-nov-retro.png '@dansult twitter stats for November 2021')

[mudmap]: https://mudmap.io/?utm_campaign=retro&utm_source=danielms&utm_medium=blog
