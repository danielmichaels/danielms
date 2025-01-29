+++
title = "Taskfile dynamic command variables"
categories = ["zet"]
tags = ["zet"]
slug = "Taskfile-dynamic-command-variables"
date = "2025-01-29 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Taskfile dynamic command variables

[Taskfile](https://taskfile.dev) is amazing and I use it as the task runner for
all my projects.

Why?

- no .env pollution between operations; define envs per task
- concatenate many commands together, for instance, regenerate tailwind when
  files change and restart server
- simplify my life by reducing the amount of typing

There more reasons but those are the highlights.

One thing that I only learned a couple months ago is that you can supply dynamic
variables to tasks.

As example, I have a bunch of python files that need various environment
variables - each file needing a different one to the next. Thats really annoying
to configure manually. Here's a simplistic version using `Taskfile`.

Here's how I'd do that in a `Taskfile`

```yaml
version: "3"

vars:
  PG_PORT: 5672

tasks:
  benchmark:*:
  env:
    PG_HOST: foo.bar
    TRACEMALLOC: 1
  vars:
    ARG_1: "{{index .MATCH 0}}"
  cmds:
    - python -m {{.ARG_1}}

  production:*:
  env:
    PG_HOST: prod.foo.bar
    TRACEMALLOC: 0
  vars:
    ARG_1: "{{index .MATCH 0}}"
  cmds:
    - python -m {{.ARG_1}}
```

Now I can run my benchmark or production code for any python module by passing
the name of the module like so:

```shell
task benchmark:pg_connection
# or
task production:pg_connection
```

Thats a really silly example, heres one I actually use.

```yaml
version: "3"

tasks:
  nats:cname:*:
  desc: Publish a message to NATS to trigger a CNAME lookup
  vars:
    ARG_1: "{{index .MATCH 0}}"
  cmds:
    - nats pub scanners.dangling_cname.{{.ARG_1}}
```

Now I can publish messages to NATS easily in development to the queue I want
with any domain name I want to target.

```shell
task nats:cname:google.com
task nats:cname:tesla
```

This has been a massive quality of life improvement. Its possible to use more
than one ARG here too but for simplicity I've left it at just one.

Tags:

    #taskfile
