+++
date = "2022-01-02"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-dec-2021"
title = "December 2021 Retrospective"
draft = false
+++

# Summary

Mudmap is now completely rebuilt in Go (with extra features) and ready for official (re)launch in January 2022.

## Highlights

- Mudmap version 2 exceeds version 1 in features
- New authentication backend makes life a lot simpler for me (and will allow MFA for users)
- After a big year, I've taken a couple of weeks off and feel great

## Goal Performance

A review of last months three goals. See [Novembers's Retrospective][old-retro]

[old-retro]: /retrospectives/2021/retrospective-nov-2021/

### Plug in the user authentication to Mudmap

- **Appraisal**: Completed
- **Rating**: A

Mudmap's current iteration uses third party packages and my own implementation to handle the login, logout, password 
reset flow for the application. It took me a long time and added complexity, especially to the frontend. This 
complexity meant that sometimes users could be booted from a session randomly, which isn't great user experience. It 
also took a lot of my development time to create and maintain, time better served elsewhere. Hindsight is also 
crystal clear, and I knew that if I had my time again, I would simply use an existing third party service instead. Enter [Auth0](https://auth0.com). 

Off the bat, it has been a breeze. Setting up the frontend to authenticate a user with Auth0 and then allow them 
to access the application was very simple. Using a React based frontend (Next.js) has its perks, namely the first 
class support *most* big orgs provide for it. Auth0 was no exception here providing excellent libraries for both 
Next.js and React. 

Where I did hit some issues was the secure communication between the (now authenticated) frontend and my backend 
server. When a user gains access (logs in) via your frontend, that does not give them immediate access to anything 
else, they need to then request an access token. That access token must then be sent as a header with each request 
to the server and verified. The process is well understood and *should* have been easy enough, but if you want to 
use Auth0's `auth0-react` module **and** `axios` well you're probably going to struggle a little. 

Why? The library expects you to request an access token from Auth0 via a react hook. Unfortunately, this 
effectively means I needed to either create an axios request (object) per component, or use fetch. 

In the end I removed `axios` completely and replaced it with fetch. On record, I much prefer `axios` - fetch has 
pitiful error handling.

At the same time, my backend implementation also hit a problem causing a two pronged debugging session whereby I had 
two errors but thought there was only one. Go is a typed language and when you say you're only sending a `string` 
but instead send a `[]string`, you're going to get an error. What tripped my up the most was the documentation of 
what should be getting send to my server, did not meet reality. After trawling the Auth0 forums and decoding the 
token I found the solution and implemented it. What I love here is that it was a fix *I* could easily implement into 
a third party module. In the past, I've found this process much more difficult using python.

### Email customers about upcoming changes to Mudmap

- **Appraisal**: I emailed a few but not all (I actually forgot)
- **Rating**: C

Embarrassingly, I wrote a big email to let customers know about the upcoming changes and forgot to send it before 
Christmas. It's going out this week. 

I did however write a [blog] post stating much the same things which are in the email. 

I'm giving myself a **C** because I *did* write it and I also wrote a blog post about it. Cutting myself some slack 
because I just drank a litre of sangria with my wife at the beach tonight too.

[blog]: https://www.mudmap.io/blog/version-2

## Recommendations

If you're into mountaineering or other sports which push the limits of human potential, you will absolutely love the 
movie [14 Peaks](https://www.imdb.com/title/tt14079374/). Two things I loved about it; the team is entirely nepalese,
and Nims is ex SB. Having served with some SB blokes over the years makes it feel closer to home. Also, I love to see 
former warfighters do big things with their lives once they've left the service. 

Write your successes and failures throughout the year. I've now written a retrospective for each month of this year, 
and reading back over them has been hugely rewarding for me. Each week, I write about what I did over at 
[What Got Done][wgd], too. Without doing these I really would not remember what I've been up to at all over the last 
year. Cataloging my year is the single best thing I have to beat out imposter syndrome, and also to keep me accountable.

[wgd]: https://whatgotdone.com/dansult

## Wrap up

It's been a big year with lots ups and downs, wasted efforts, and successes. 

**What can I do better?**

- Remember to send emails after I've drafted them up!
- Step away from the keyboard on those really nice days and just take the family out somewhere instead - the keyboard won't miss me

**What have I done well?**

- I've learnt a lot, and have been able to implement things now that only 12 months ago would have taken me months - keep growing and stay the course
- Exercising - I've done a lot more this month, especially since signing up for a 2k ocean swim in Feb

## Next month's goals

- Push Mudmap version 2 to production
- Add at least one more core feature to Mudmap

## Analytics

**Mudmap**

![](mm-nov-retro.png 'Mudmap plausible stats for December 2021')

**Check-Redirects**

![](cr-nov-retro.png 'Check-Redirects.com plausible stats for December 2021')

**Check-Redirects**

![](danielms-nov-retro.png 'danielms.site plausible stats for December 2021')

**Twitter**

![](twit-nov-retro.png '@dansult twitter stats for December 2021')

[mudmap]: https://mudmap.io/?utm_campaign=retro&utm_source=danielms&utm_medium=blog
