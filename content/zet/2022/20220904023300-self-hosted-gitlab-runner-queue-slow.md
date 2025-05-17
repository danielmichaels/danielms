+++
title = "self-hosted gitlab-runner queue slow"
categories = ["zet"]
tags = ["zet"]
slug = "self-hosted-gitlab-runner-queue-slow"
date = "2022-09-04 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# self-hosted gitlab-runner queue slow

Gitlab self-hosted runners seem to take a long time to pick up jobs
from the queue. In many cases its was over 3 minutes before the pipeline
would even start and then have lengthy pauses between stages.

Changing two lines in the `/etc/gitlab-runner/config.toml` file seems to
have fixed this issue.

```shell
concurrent = 5
check_interval = 5
```

[link](https://gitlab.com/gitlab-org/gitlab-runner/-/issues/4567)

Also, as my runner is a docker-runner inside a container I could not 
edit this file. The Dockerfile is locked down so I had to `docker cp` the
file out of the container, edit it, and then `docker cp` it back in.


```shell
docker cp <Container-ID>:/etc/gitlab-runner/config.toml .
vim config.toml
docker cp config.toml <Container-ID>:/etc/gitlab-runner/config.toml
```

Tags:

    #cicd #gitlab #runners #docker

