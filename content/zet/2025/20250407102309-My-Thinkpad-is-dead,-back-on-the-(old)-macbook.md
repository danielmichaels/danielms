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

Edit: adding apps that have similar ergonomics on Mac as they do linux:

## Terminals 

*Dropdown*:

`iterm2` is (so far) the only terminal I can configure to have dropdown like functionality similar to `guake` or `ddterm` on Linux. This really 
breaks my flow. I have f12 mapped to dropdown a terminal meaning I can keep my IDE/Browser whatever in view whilst I have terminal access. Iterm2 supports this 
BUT if you use mac's desktop's it'll yank you from a fullscreen in desktop 3 back to desktop 1 if thats where you "main" desktop is. Iterm won't dropdown into 
the current desktop - only the "main", typically desktop 1. I'm sure theres a mac way of explaining this but its a massively negative experience for my flow.

Edit: I think I've fixed this by doing [this](https://gitlab.com/gnachman/iterm2/-/issues/10695#note_1267078045), which is to set `Profiles -> <profile> -> Keys -> Configure Hotkey WindowFloating Window -> Floating Window` then restart iTerm. Works well so far..

*Terminal*:

`wezterm` and `Iterm2` so far are the only terminals that support the Zellij `option` key (`alt` rebinding). 

Terminals that do not support my workflow:

- `ghostty`: doesn't work well with Zellij, no alt key alternative
- `alacrity`: as above
- `kitty`: as above 

Tags:

    #mac #linux
