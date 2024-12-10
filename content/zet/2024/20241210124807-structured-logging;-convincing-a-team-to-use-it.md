+++
title = "structured logging; convincing a team to use it"
categories = ["zet"]
tags = ["zet"]
slug = "structured-logging;-convincing-a-team-to-use-it"
date = "2024-12-10 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# structured logging; convincing a team to use it

My $DAYJOB deals with millions of messages through the message bus, each
messsage kicking off many more IO bound operations.

We use Loki and Grafana to monitor the logs.

Our python code relies heavily on threading and its hard to know whats happening
in each thread.

And we don't use structured logging. Which ever since I started left a sour
taste in my mouth as I've seen it used very well in other environment which 10x
less load and events.

Today, after someone made a big change which included what I consider bad
changes to the plain loggers I decided to make the effort to convince them that
structured logging is worth the effort.

My core driver is the time wasted finding issues in our threaded code which we
could easily have detected/debugged with a trace ID spanning the events.
OpenTelemetry in this team would be a very very hard sell. But a trace ID? Thats
doable and has a lot of upside.

I know I will get a lot of pushback. Pretty much all the 12 factor type
improvements I've made have received little interest.

So this time I made a **huge** effort to push this. Incredibly detailed
reasoning, issues, example PoC's in the core libraries. Tomorrow, I'll start
making the merge requests to get the initial work done. I think if I can get our
core rabbit library cut over and then our two "hot path" projects 80% of the way
there I'll have the majority interested.

And, if not. Well, I'll just document how I did it and keep all those learnings
for future $DAYJOB if needed.

Tags:

    #logging #software #team
