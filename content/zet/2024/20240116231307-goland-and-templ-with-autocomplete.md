+++
title = "GoLand and templ with autocomplete"
categories = ["zet"]
tags = ["zet"]
slug = "goland-and-templ-with-autocomplete"
date = "2024-01-16 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# GoLand and templ with autocomplete

To use `templ` with autocomplete via its
[Plugin](https://github.com/templ-go/templ-jetbrains) you must do the following:

- Go to settings>templ and enter the path to the `templ` binary (e.g. `$HOME/.local/bin/templ`)
- Close all GoLand instances
- Launch GoLand via the terminal inside the project directory (e.g. `goland .`)

It works but still needs some refining at the project level.

The alternative is to use `neovim` for the `templ` files or VSCode. As
a GoLand subscriber these are sub par alternatives but may be required in the
short term.

Tags:

  #go #templ
