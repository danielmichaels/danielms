+++
title = "Format JSON on The CLI Easily"
categories = ["zet"]
tags = ["zet"]
slug = "format-json-on-the-cli-easily"
date = "2022-11-01 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Format JSON on The CLI Easily

`python -m json.tool <file>`

This can then be chained to overwrite an existing file, effectively
formatting it.

`python -m json.tool <file> | tee /tmp/file.json && cat /tmp/file.json > <file>`

This gives the command multiple out's without breaking the originating file. 
Probably a far better way but this works.

`jq` would work but it isn't installed on remote Unix systems whereas
`python -m json.tool` always works.

Tags:

    #linux #json

