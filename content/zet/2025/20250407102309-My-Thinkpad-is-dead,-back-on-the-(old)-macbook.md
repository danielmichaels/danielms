+++
title = "My Thinkpad is dead, back on the (old) macbook"
categories = ["zet"]
tags = ["zet"]
slug = "My-Thinkpad-is-dead,-back-on-the-(old)-macbook"
date = "2025-04-07 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# My Thinkpad is dead, back on the (old) macbook

My trusty Thinkpad is finally so slow that I've calling it dead.

This is bringing great shame to me as I know am resorting to using my partners old Mac.

I think I've been a linux master race user since leaving mac around 2017-2018. Despite mac
being based on linux it still sucks to use. 

No docker (docker destop/orbstack/colima are all workarounds), inexplicable usage of the "Command" key in the "Alt" key position, weird security prompts requiring clicky clicky.

But, I do get it, people enjoy this because some things "just work". I mean they work because millions and millions are poured into its development.

Now I need to figure out how to manage Chezmoi with two different operating systems. Should I keep using mise, or just leverage brew. I feel like the answers to these questions are only a footgun away!

On a good note, rebuilding this machine showed me that my `zet` tool no longer builds as an underlying package has been deleted upstream. I'm guessing its cached on my other machines so I didn't notice - some learnings in there. I rewrote the tool using <https://github.com/alecthomas/kong> over the weekend to get it back online.

<https://github.com/danielmichaels/zet-cmd>

Tags:

    #mac 
