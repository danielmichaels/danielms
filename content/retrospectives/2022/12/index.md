+++
date = "2023-01-07"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-dec-2022"
title = "December 2022 Retrospective"
draft = false
+++

# Summary

Two releases for [Mudmap] and two weeks in [Perth]

## Mudmap

1. System update available widget
2. Installed packages table

After finishing up the above features I took an extended break from all things code. I still
took my laptop away but never really felt the need to open it up. My holiday back home was
far too relaxing!

### System Update Available

In pfSense, on the main dashboard, it tells if the version is up-to-date or if a 
newer version is available.

Mudmap now does this too. A small quality of life improvement but another step in replicating
the pfSense experience within Mudmap.

This was *simple* for customers who are on the API version 1.5.0 and above. However, it needed
to be backwards compatible for users that aren't. Quite a few are using the older version and I
did not want to prompt them to upgrade. This meant checking their currently installed version 
and providing a fallback for them. This fallback indicates that they should update and provides
a link with how to do so.

### Installed Packages

Another quality of life improvement. Users can now check which third party packages are installed
and their versions. It is read-only, as in they cannot upgrade from Mudmap, but it is a start.

The inclusion of the system update and installed packages means auditing devices got a little easier.

## Random Mobility Workout of the Day (MWOD)

This was something I knocked up in two evenings to service my own simple needs.

I love watching and following along with [Kelly Starret][kstarr]. When I was in the military
the MWOD gave me a lot of relief from pain. Its tough work doing cool shit! Since I've become
a desk warrior I sort of let it go but that has brought its own problems. I have new musculoskeletal 
pains.

Anyway, I built this to prompt me each day with a random workout to perform. It's a Go app using 
templates to pull a list of 160+ videos from YouTube and embed them in the site. Go plus Sqlite 
makes life pretty easy. [Codebase](https://github.com/danielmichaels/rmwod)

Check it out, [here][rmwod]

## Recommendations

Book: 100 Go Mistakes and How to Avoid Them ([link](https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them))

I write a decent amount of Go but there are various parts of the language which have tripped my 
up in the past. This book covers all the issues I've ever footgunned myself with, and explains
not just why but also how to fix it.

For any Go developer I think its worthwhile reading. So far I've learned a lot more about slices 
(under the hood), strings (particularly the importance of Runes) and solidified my knowledge in
control structures.

There might be a lot of ways of achieving things but something I really like about this book is
they are opinionated about why *x* thing is the best way to do something for the general.

For instance, when checking if a slice is empty the best way is to check its length. This will 
trap both `nil` and `empty` slices.

```go
// checks for nil and empty
// empty slices have a len of zero and nil slices are always empty
if len(foo) == 0 { 
	return false
}
```

I'm only about 40% of the way through, and I've referenced it a couple of times already. Another 
thing which is really neat; succinct little sections for each *Go Mistake*. This means you can easily
pick it up whenever you've got a spare moment which having to read a wall of text to understand 
the previous context.

## Beach Mode

The rest of this month was me soaking up rays at Perth's beautiful beaches.

![](beach-mode.jpeg 'chilling at the beach')

[mudmap]: https://mudmap.io/?utm_campaign=retro-nov-22&utm_source=danielms&utm_medium=blog
[perth]: https://en.wikipedia.org/wiki/Perth 
[kstarr]: https://en.wikipedia.org/wiki/Kelly_Starrett
[rmwod]: https://randommwod.com?utm_source=danielms
