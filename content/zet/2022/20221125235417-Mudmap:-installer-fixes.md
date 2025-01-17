+++
title = "Mudmap: installer fixes"
categories = ["zet"]
tags = ["zet"]
slug = "Mudmap:-installer-fixes"
date = "2022-11-25 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Mudmap: installer fixes

This week I've been working on fixing Mudmap's installer. 

I've done the math and nearly 50% of installs have failed inexplicitly 
in the last 3 months. For users which installation worked without issue
they're happy and never complain. Obviously, the inverse has a lot of 
complaints. I am no closer to diagnosing exactly what causes this. A 
consistent (>60%) of failures trigger at the user account install step.

Given this I am exploring that step in detail but also changing the way
the installer is handled on the front end. The error users receive just
confuses them and that needs to go. 

I'm also focusing on revert the device back to its pre-mudmap state with
better consistency. This is important for users, and for future installs.
To do this without storing users root password, I've elected to use an 
in memory data store - `go-memdb` - which is super lightweight and easy to
understand.

Tags:

    #mudmap #go

