+++
title = "reMarkable MCP in Go"
categories = ["zet"]
tags = ["zet"]
slug = "remarkable-mcp-in-go"
date = "2026-11-01 19:45:18 +1000"
draft = "false"
ShowToc = "true"
+++

# reMarkable MCP in Go

A follow up to [Amp + reMarkable]({{< ref "/blog/2026/2026-10-25-remarkable-pdf-preview.md" >}}): I rewrote the reMarkable MCP server in Go.

It's on GitHub: [go-remarkable-mcp](https://github.com/danielmichaels/go-remarkable-mcp).

The big win for me is memory usage. The repo has the full comparison but the short version is:

- Small PDF read/render/export: Python **193 MiB**, Go **32 MiB**.
- EPUB + nine Markdown renders/uploads: Python **245 MiB**, Go **71 MiB**.

This matters for my [Amp orbs](https://ampcode.com/what-are-orbs). As memory matters for all things *cloud*. We pay in memory and CPU. I converted this from python to go so that the limited memory by Orbs have is dedicated to work not idle processes. 

This also removes any runtime requirements. No `uvx` just a small binary ready and willing whenver I call it.

Tags:

    #amp #remarkable #go #mcp
