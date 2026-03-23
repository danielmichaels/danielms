+++
title = "I built a Tauri app for my home lab"
categories = ["zet"]
tags = ["zet"]
slug = "i-built-a-tauri-app-for-my-home-lab"
date = "2026-03-23 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# I built a Tauri app for my home lab

I've been leaning a lot more into AI coding lately.

It's made me rethink a couple things:

- I can use **any** language
- Reliability and error detection at compile time is more important than ever

Rust is a trade-off between compile-time and memory safety at the cost of developer experience and cognitive load. It takes a lot longer to develop a Rust app than a Go one and the results
for *most* app's (that I would built) aren't worth the "Rust cost" IMO. 

AI kind of changes the game here.

It iterates so fast that Rust on-boarding and learning can be "just-in-time" or "just-enough". Now you can "write" Rust and learn it as you go - it's like learning to code by 
reading someone else's code except its *your* code for *your* idea.

And, now you get the sweet compile time safety. It also has the benefit of keeping AI in check - it writes something and breaks an contract somewhere, well it'll detect that and the knock out
the fix super fast.

I still love Go, I still opt for it (for 95% of things) but think AI is changing the way we look at certain objectives.

---

So I wrote a Tauri app. It's a MacOS dock based app for monitoring all my Knative applications status. It connects to your kube context (so technically could monitor any Knative stack).

It can:

- ping it (to wake it up)
- view its ingress (if its serving a page/API etc will show it)
- view logs 

I don't need too much more than that. 

I found the Tauri experience to be top-notch compared to Wails (a Go Tauri like app builder).

Repo: <https://github.com/danielmichaels/knative-dash>

Tags

    #rust #ai #tauri
