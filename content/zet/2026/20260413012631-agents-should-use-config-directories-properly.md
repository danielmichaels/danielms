+++
title = "Agents should use config directories properly"
categories = ["zet"]
tags = ["zet"]
slug = "agents-should-use-config-directories-properly"
date = "2026-04-13 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Agents should use config directories properly

Why do all the agents/LLM's use `~/.$foo` instead of `~/.config/$foo`.

Its lazy programming, teaches others that it's okay to do this (it's not) and basically ensures for the rest of time we have our home directories filled up with shit.

Do a `ls -la` on your `$HOME` and look at home much junk there is. `$XDG_CONFIG_HOME` exists for a reason. 

I feel these companies have a duty of care because everyone will follow them blindly.

Thing is we can plainly see none of them give a shit. Look at Anthropic stealing IP. Scam Altman (say no more).

Talk about grinding my gears.

At least OpenCode gets right - literally one of the few shining lights in this shitshow.

Culprit list:

- ollama
- openai (.codex)
- anthropic (.cluade)
- pi 
- amp
- crush
- google (.gemini)

Tags:

    #rant #ai

