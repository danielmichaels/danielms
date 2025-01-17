+++
title = "Initial Zettelkasten"
categories = ["zet"]
tags = ["zet"]
slug = "Initial-Zettelkasten"
date = "2022-04-26 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Initial Zettelkasten

This marks the very first Zettelkasten entry in this repo. To do this I've created my own
`Zet` tool called [zet-cmd] written in Golang. It uses [Bonzai] by [rwxrob] and is heavily
inspired by his own [zet] repo.

I needed a way to capture small, digestible nuggets of wisdom and other thoughts or learnings
in a manner that let's me do it without needing a full blog post. It also needs to be
searchable publicly and using GitHub makes that pretty simple.

So far [zet-cmd] can:

- Create new Zet's
- Edit existing ones
- Search across all Zet's
- Generate a url with a search term for the repo on GitHub
- Retrieve the latest zet for easier editing

In the future more commands will be added which make life easier. More helper commands to make
piping data into Vim will be added in the future.

Some example usages:

Creating a GitHub.com search query with `zet q "search term"` would return:
https://github.com/danielmichaels/zet/search?q=search+term

Tags:

    #zettelkasten #golang #bonzai #rwxrob

[rwxrob]: https://github.com/rwxrob
[zet]: https://github.com/rwxrob/zet
[Bonzai]: https://github.com/rwxrob/bonzai
[zet-cmd]: https://github.com/danielmichaels/zet-cmd
