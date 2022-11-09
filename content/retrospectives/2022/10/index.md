+++
date = "2022-11-05"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-oct-2022"
title = "October 2022 Retrospective"
draft = false
+++

# Summary 

A month of working on only things I find *fun*.

## Contract Hour Tracking

I am a full-time employee on a hourly contract basis. It's a funny position 
because I'm effectively paid casual rates but full-time. Unlike most other 
contracts the possibility of my contract being cancelled is nearly zero due 
to the extreme staff shortages in my workplace.

With that, I am given a set number of hours I can work within the contract 
period. I cannot go over but any hours remaining at contract end is simply 
money I'm robbing myself of. Unfortunately, I'm pretty lazy about tracking 
it all and when I did, I realised that I've under paid myself over this 
contract period. A project idea was spawned.

Enter a rather simple contract timesheet calculator; [timesheet-calculator][calc].
It's pretty simplistic in the output it provides but so far has been pretty 
helpful. Every day it sends an email with some stats to keep an eye on such 
as my mean daily required hours to reach contract zero. The most useful 
emails are my weekly and monthly ones. The monthly email is the most useful 
as it spits out the daily log in tabular format making filling out my 
monthly timesheet much easier. 

A trivial project but for those with my level of laziness automation is a 
solution worth the time investment.

## Nomad

At work, we're pretty tied to Kubernetes and are actively looking at 
replacing our OpenShift cluster. The reasons are varied, but it mostly comes 
down this fact; it's too much for us. We're a small team of developers who 
also maintain our own Kubernetes cluster which all of our workloads run on. 
Without a dedicated Ops team its starting to drain our resources.

Our initial investigations into alternatives mostly centred around other 
Kubernetes platforms. One of the key metrics is it must be on-prem without 
the need for an active internet connection - we live and die by the proxy. 

At first, we dismissed Nomad, mostly because of the time we'd spent on 
Kubernetes and the significant amount of work that's gone into creating its 
supporting assets; manifests and pipelines as the exemplars.

Sunk cost fallacy aside, Nomad, on the surface looks like it suits our 
needs more adeptly. I wrote a [couple of zets][z] about Nomad; some basics 
and what I think is interesting about it. I spent a bit of time playing 
around with the `nomad agent -dev` mode but couldn't commit too much more 
time to getting a home lab up and running. Though I think that is something 
I need to do in November, or over the Christmas holiday period.


## Mudmap

I have not done any work on Mudmap this month.

Honestly, I am still unsure what to do with it. People are still interested, 
but I am unsure if it's still right for me.

## Recommendations

- [Andor](https://en.wikipedia.org/wiki/Andor_(TV_series))

This is the first Disney spin-off that I have enjoyed. I wrote an entire rant
about it [here](https://github.com/danielmichaels/zet/blob/main/20221022005935/README.md).
It feels like the first *real* storyline to emerge from the Star Wars 
universe that isn't just a rehash of existing lore, or absolute garbage 
(Mandalorian and Obi-Wan).

## What's next

I'll keep on making things that are useful to me whilst I figure out the 
bigger picture of what I want to do.

## Goals

I think my main thing will be to continue this year seeking enjoyment in my 
downtime. 

## Analytics 

{{< plausible start=2022-10-01 end=2022-10-31 site_id=mudmap.io >}}

{{< plausible start=2022-10-01 end=2022-10-31 site_id=danielms.site >}}

[mudmap]: https://mudmap.io/?utm_campaign=retro&utm_source=danielms&utm_medium=blog
[z]: https://github.com/danielmichaels/zet/search?q=nomad
[calc]: https://github.com/danielmichaels/faas-templates/tree/main/timesheet-calculator
