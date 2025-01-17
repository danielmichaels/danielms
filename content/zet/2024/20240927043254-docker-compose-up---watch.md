+++
title = "docker compose up --watch"
categories = ["zet"]
tags = ["zet"]
slug = "docker-compose-up---watch"
date = "2024-09-27 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# docker compose up --watch

I used to volume mount to get "hot-reload" behaviour. Now I've switched to
`docker compose watch`.

Its much simpler to understand. Here's an example of syncing code and restart a
celery worker container:

```yaml
services:
  worker:
    build:
      context: .
      network: host
    network_mode: host
    command: celery
    env_file:
      - .env
    depends_on:
      db:
        condition: service_healthy
      rabbit:
        condition: service_healthy
    volumes:
      - worker_data:/usr/src/app
    develop:
      watch:
        - action: sync+restart
          path: ./
          target: /app
          ignore:
            - .venv
volumes:
  worker_data:
```

A simplified `compose.yml` file but the core part I want to showcase is the
`develop` key.

Using `develop.watch` lets you define the path to "watch", target to sync the
changes to and the action to take.

For this container I have to restart it but its very robust and simple to
follow. You can also just `sync` which would work well for some applications.

It's possible to achieve similar things with other tools (`watchfiles` is a good
one) and methods but I prefer to lean on a `docker` builtin these days.

Tags:

    #docker
