+++
title = "Using Github Actions to publish my zettelkasten notes"
categories = [""]
tags = [""]
slug = "github-actions-auto-publish-zettelkasten-notes"
date = "2025-01-14"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

I publish my zettelkasten notes to this website using GitHub Actions.

## What is a zettelkasten?

> A Zettelkasten is a personal tool for thinking and writing that creates an interconnected web of thought

From the zettelkasten website, a true *zet* should have the following properties:

1. Atomic - Each note contains one main idea
2. Unique identifiers - Every note has a unique ID for referencing
3. Linking - Notes are connected to other related notes via links
4. Emergence - Knowledge and insights emerge from the network of connected notes

It's meant to ensure that over time, the notes become more interconnected and the network of knowledge grows.

My zettelkasten isn't completely true to this model but its good enough for my use case.

## Why do I use a zettelkasten?

When I learn something, figure out a new way or approach, ill try put it in my zet. Keeping them short and to the point
makes this more likely to happen. I found with a blog I finessed it too much which built up a resistance.

My zets can be as polished or sloppy as I choose. I am the only person I write for.

Unlike obsidian or other tools my approach is just text. Linking beyond a simple search, is on me to do at the time
of search. This works better than an automated tool because it's the act of re-reading notes even if unrelated that
builds your mental model of your own knowledge base (spaced repetition)

## My zet-cmd tool

To write, read and edit my notes I have a custom CLI written in Go,
called [zet-cmd](https://github.com/danielmichaels/zet-cmd).

It uses GitHub as the backend to store the zets and your `$EDITOR` to write them. It's a simple tool that does one thing
and does it well enough for me.

Here's the output of `zet`:

```go
NAME
zet - zettelkasten commander

SYNOPSIS
zet COMMAND

COMMANDS
help          - display help similar to man page format
conf          - manage conf in /home/danielmichaels/.cache/zet/config.yaml
var - cache variables in /home/danielmichaels/.cache/zet/vars
new|c|create  - Create a new zet
l|latest|last - Get the most recent zet isosec and print it screen
e|edit        - edit a zet
g|get         - Retrieve a zet for editing
q|query       - create a searchable URL with a query string
f|find        - Find a zet title by search term
check         - check environment variables and configuration
t|tags        - Find zet(s) by tag'
git           - run git commands over the zet repo
v|view        - view command for zet entries.

DESCRIPTION
The zet command is Zettelkasten Bonzai branch used to create small slips of knowledge.Those slips are then uploaded to Github for public search-ability and ease of use.

CONTACT
Site:   danielms.site
Source: git@github.com:danielmichaels/zet-cmd.git
Issues: github.com/danielmichaels/zet-cmd/issues

LEGAL
zet (v0.5.0) Copyright 2022-2024 Daniel Michaels
License Apache-2.0
```

The tool is written in Go and uses [Bonzai](https://github.com/rwxrob/bonzai) for the CLI. Though, it's using an older
version and will need updating to the latest version - which has several breaking changes. Could equally be rewritten in
[kong](https://github.com/alecthomas/kong) another favourite of mine.

## How I use it

When I want to create a new zet I run `zet new`. This will create a new entry in my zet repo using a timestamp as a
folder name, and inside that folder a new README.md. The `zet new "my new zet"` will generate the README.md with
`my new zet` as the `#` (h1) heading and drop into the editor, in my case, `(n)vim`.

Saving and exiting will then prompt you to commit the changes to the repo, this will then push the changes to GitHub.

Editing a zet is as simple as running `zet edit`. This accepts `last` for the most recent zet, or a search term.

For example, `zet edit kong` returns:

```bash
# zet edit kong
0) 20230107005542 Kong is an amazing CLI for Go apps
1) 20240801073758 How I write Golang CLI tools today (using Kong)
```

Entering one of the numbers will open the zet in your editor.

Viewing a zet is as simple as running `zet view`. This accepts `last` for the most recent zet, or a search term. Same
functionality as `zet edit`. `zet view all` will list all zets in the repo.

Zets are rendered in the terminal using [glow](https://github.com/charmbracelet/glow) to make them visually appealing.

## Adding zet's to a hugo site

Sometimes I am not at my computer but want to reference something I've written about before. All my zets are public,
so I figured why not add them to my site. The added bonus is others can also view them, and it takes the *pressure* off
of me to write a blog post.

To do that I use a GitHub action to fetch the zets from my repo and add them to my site. The action runs every hour and
fetches the latest zets only triggering a rebuild if a new zet is added. This is done via the commit functionality of
[stefanzweifel/git-auto-commit-action](https://github.com/stefanzweifel/git-auto-commit-action) which just commits directly
to `master`. It's my blog, so I don't care if it commits straight to master.


All the tools and actions are available on [GitHub](https://github.com/danielmichaels/danielms) in the `scripts` and `.github`
directories. `./scripts/fetch-zet.go` is the script this action uses to fetch the zets.

## How it works

A picture paints a thousand words. This graph shows the *broad* flow of events. Refer to `./scripts/fetch-zet.go` for
the full flow.

{{< mermaid >}}
graph TD
subgraph "Github Actions hourly schedule"
    A[Start fetch-zet.go] --> E[Fetch Contents from GitHub API]
    E --> F{Check if New Zets}
    F -->|Yes| G[Write JSON to assets/zet.json]
    G --> H[Check if zet already exists, create if not]
    F -->|No| L[Exit - No Updates Needed]
    H --> M[Commit changes to repo]
end
    M --> N[Netlify rebuild on commit to Master]
{{< /mermaid >}}
