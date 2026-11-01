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

Those are settled process memory figures from short local tests using the same files and operations, not a promise that every document will use that little memory. It's still experimental and not every operation is faster.

But this matters for my [Amp orbs](https://ampcode.com/what-are-orbs). The MCP server shares memory with the agent, builds, tests and whatever else I'm running. I want that memory going towards the actual work, not starving the orb just so I can send a PDF to my tablet.

No Python runtime needed either. Same idea as before: send a doc, scribble on it, get Amp to read it back. Just with a much smaller memory footprint.

Tags:

    #amp #remarkable #go #mcp
