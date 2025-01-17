+++
title = "hstr: bash history fuzzy finder"
categories = ["zet"]
tags = ["zet"]
slug = "hstr:-bash-history-fuzzy-finder"
date = "2023-01-14 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# hstr: bash history fuzzy finder

I usually opt for `history | grep -i <term>` or `history | ag <term>`
depending on which system I'm on.

For local development, I've started using `hstr`.

ref: https://github.com/dvorka/hstr

Fedora: `dnf install hstr`

Ubuntu 22.04 `apt instll hstr`

Or, just use `nix-env -i hstr`

Personally, I've opted for `nix` because I'm still on 20.04 and I am 
tired of installing `ppa`'s only to have their keys expire and break
all updates on the system.

Tags:

    #nix #tools

