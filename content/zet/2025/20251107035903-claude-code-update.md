+++
title = "Claude Code update"
categories = ["zet"]
tags = ["zet"]
slug = "claude-code-update"
date = "2025-11-07 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Claude Code update

> What cannot be understood will be rewritten - Bill Kennedy 

Okay I've come around on this tool a bit. 

Used correctly with intent and under firm guidance, it can be very powerful.

It's still a massive idiot often enough that its brilliance becomes offset to the point where its untrustworthy.

I am not convinced anyone blindly accepting this code won't be bitten in the arse badly in the future. It writes excellent legacy code.

That being said, I believe it is a massive help more often than not (when wielded correctly). My main use cases:

- Rubber ducking
- Architectural conversations
- Writing code when **I know exactly** what I want, and ensuring it does exactly that
- Questioning my assumptions/code (basically rubber ducky as well)

Where I **really** love it is writing throwaway code.

Like a lot of dev's there is literally not enough time to do all the fun/interesting things you can think up. Writing code takes time 
and I have a life outside of my chair. But, with Claude I can speed run some of these and get it to write pure legacy code because I literally 
don't give two hoots if it blows up later - it's an experiment, fun, exploration, single-shot script etc 

For that, Claude is amazing and worth the money.

As an example, I have a Brother thermal printer and hate that I have to plug it into a laptop/desktop to print stuff. So I decided I'd turn it into a 
always available service on my network. Plug it into a NUC/pi, write a flask server with a couple endpoints, use those to talk over `pyusb` and print
the labels my wife needs. Tell claude to create the HTML forms and then use the forms submit data to the flask server. Put both on tailscale. Dockerise so I can run the flask server on my LAN (connected to printer), and let other services interact with it. 

I could absolutely do all this without AI but it went from a weekends work to a couple hours (80% of that burnt just trying to get the printer to work because honestly f*ck printers bro).

So yeah, I'm warmed to Claude Code but its still just a tool to be used. My next big challenge is to learn `zig` without just relying on AI - if I can learn `zig` with the help of AI but not be so reliant on it that I don't actually learn it, then I'll chalk that up as a big win too. 

Tags:

    #claude #ai

