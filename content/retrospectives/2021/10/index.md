+++
date = "2021-10-05"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-oct-2021"
title = "October 2021 Retrospective"
draft = false
+++

# Summary

A month of learning and exploring, free from self-imposed pressure.

## Highlights

- Focused on learning Go
- Deployed my first Go app
- Interviewed with Microsoft

## Goal Performance

A snippet from last month's goal setting in the [September retro][old-retro]

[old-retro]: /retrospectives/2021/retrospective-sept-2021/

>No goals, or rather I am taking the month to freestyle the next thing to do. Not setting goals 
is setting a goal to fail, or something like that, but I need a month to relax my mind. My wife 
has had some health issues, COVID garbage is constant and working from home whilst homeschooling 
ain’t easy. But, that should all ease off towards the end of the month, so I’m taking until then 
to just experiment.    

Instead of reviewing and assessing my goals, I'll cover off on what I have done.

### Focused on learning Go

After several on and off attempts at learning Go over the years, I've actually committed to it 
over the last month or so. When it comes to learning new things - regardless of topic - I like 
to use the [four stages of competence model][comp] as my frame of reference.

[comp]: https://en.wikipedia.org/wiki/Four_stages_of_competence

| Stage | Description |
|---|---|
| Unconscious incompetence | The individual does not understand or know how to do something and does not necessarily recognize the deficit. They may deny the usefulness of the skill. The individual must recognize their own incompetence, and the value of the new skill, before moving on to the next stage. The length of time an individual spends in this stage depends on the strength of the stimulus to learn. |
| Conscious incompetence | Though the individual does not understand or know how to do something, they recognize the deficit, as well as the value of a new skill in addressing the deficit. The making of mistakes can be integral to the learning process at this stage. |
| Conscious competence | The individual understands or knows how to do something. However, demonstrating the skill or knowledge requires concentration. It may be broken down into steps, and there is heavy conscious involvement in executing the new skill. |
| Unconscious competence | The individual has had so much practice with a skill that it has become "second nature" and can be performed easily. As a result, the skill can be performed while executing another task. The individual may be able to teach it to others, depending upon how and when it was learned. |

*sourced from [wikipedia][comp] as it much more eloquent than anything I could write.*

Initially, I set about the fumbling and foolish unconscious incompetent phase. This is when I 
just start and know that anything I write will be bad, of poor quality and embarrassing in the 
future. In the past, I've spent way too long researching the *best* way and otherwise convincing 
myself that I could prevent this by just reading more. You can't, just like you can't theorise 
yourself into an NBA contract - you have to do the work.

When I trained new guys on their way into the special forces, you'd see all kinds of wacky stuff.
The old way of teaching was to punish them for silly mistakes. Thankfully, times are changing 
and we (eventually) came to realise that learning is a curve. Something akin to 60% of time 
spent in a learning continuum is spent in this first phase, with the next stages accelerating 
off the back of a solid foundation. This tangent is a poor attempt to say; embrace the *wacky* 
and uncomfortable start because beginners *are allowed* to make mistakes.

Eventually, I started to hit some rough spots where experience taught me that I was missing some 
fundamental Go building blocks. Instead of searching around, I just pulled the trigger on a book 
by [Alex Edwards] called [Let's Go Further][lgf]. It has been a wonderful resource which really 
helped to teach me how to set out my projects and make them extensible.

I think I am halfway between consciously incompetent and competent now. So I know when I'm 
making the right decision about 50 percent of the time now.

[alex edwards]: https://twitter.com/ajmedwards?lang=en
[lgf]: https://lets-go-further.alexedwards.net/

### Deployed my first Go app 

Off the back of Alex's [book][lgf], I was able to refactor my URL shortener application that I 
had been writing in Go. I decided to not only refactor it but also deploy it to my Caprover 
instance. This way I'd get the experience of deploying a dockerized Go app with a next frontend 
(hosted by Vercel). In addition to that, I tried out [Litestream] as the SQL backend and was 
able to replicate the database using S3 as the storage layer. 

I use Github for personal stuff, such as this but professionally use Gitlab (and Gitea in 
the past). Github is moving with some serious speed these days but my familiarity and 
productivity is firmly with Gitlab. This did trip me up during the deployment where I 
retrospectively added a CI pipeline. I stumbled upon some issues with secrets and contributors 
which I ultimately solved by refactoring *my* code. Nonetheless, as I understand it, secret 
management in public repo's can be a security issue which is why I opted to refactor rather than 
sort out GH actions. 

The process of building and deploying my first Go app was a breeze compared to *any* python 
project. I'm still coming to terms with Go having a webserver that's production ready after 
having spent a lot of time setting up `uvicorn` and `gunicorn` to serve python apps.

Check it out at [https://tars.run](https://tars.run).

[litestream]: https://litestream.io/

### Interviewed with Microsoft 

During the month I was approached by a Microsoft recruiter about a couple of potential jobs they 
*think* I might be suitable for. Flattered, I took the call and ended up having a couple of 
first round interviews. Quickly I discovered that my technical interview skills and knowledge 
are not up to scratch. Despite being recommended to continue along the pipeline, I elected to 
postpone any further interviews. In the meantime, I've decided to slowly brush up and in some 
spots actually learn new skills.

But, I'm an *Indie Hacker* now right? Yes, and no. I work full-time for someone else and still 
need to provide for my family. Until I can swing from one branch to the next without my kids 
skipping dinner, I'll need to be ready for new job opportunities. Getting slightly humbled by 
some easy questions really put my future employability in the spotlight. 

I think this was a great wake-up call too. I've been ignoring the fact that knowing how solve 
the types of problems used during these interviews will at worst make me learn some more 
fundamentals and at best get me a job. I've nothing to lose by studying a little each day.

## Recommendations

I am seriously late to the party but I've started watching some [Twitch](https://twitch.tv) 
streamers. Here are a couple I enjoy listening to:

- [rwxrob](https://www.twitch.tv/rwxrob) for his Linux, Go and k8s rants
- [Jordan Lewis](https://www.twitch.tv/large__data__bank) for his cockroach labs Go programming
- [Anthony Sottile](https://www.twitch.tv/anthonywritescode) because he writes python open source and sounds like a real person when streaming

## Wrap up

This month was a great unwind and refresh from working on Mudmap. It gave me the intellectual 
and emotional space to take stock of a few things. 

## Next month's goals

- Mudmap feature proof-of-concept
- Study at least 2 pomo's each day


## Analytics

**Mudmap**

![](mm-oct.png 'Mudmap plausible stats for September 2021')

**Check-Redirects**

![](cr-oct.png 'Check-Redirects.com plausible stats for September 2021')

**Twitter**

I really haven't been playing the Twitter game much lately. Struggle thinking of things to say!
![](twit-oct.png '@dansult twitter stats for September 2021')



[mudmap]: https://mudmap.io/?utm_campaign=retro&utm_source=danielms&utm_medium=blog
