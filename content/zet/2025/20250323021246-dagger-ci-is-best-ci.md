+++
title = "Dagger CI is best CI"
categories = ["zet"]
tags = ["zet"]
slug = "dagger-ci-is-best-ci"
date = "2025-03-23 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Dagger CI is best CI

I've switched to using [dagger](dagger.io) for CI and couldn't be happier.

Without pushing code to the origin I can run my entire suite of CI tools
locally. This means I know everything will run when I push and CI will pass
should I want to know.

More importantly, I can debug/fix broken CI without needless commits. This saves
me a lot of time and if I was in a team, money wasted on compute/CI run minutes.

Whats more, you can "checkout" someones branch using dagger without git
pulling/git switching. Dagger isn't just limited to running CI - it can run
entire apps including supporting services. This means you can run the entire
project in the PR and check it works as it should with all the expected
services. Obviously, this is only as good as the effort that goes into writing
the Dagger files.

Doing it this way makes running complicated applications easier, for instance,
if you needed NATS, Postgres and Redis to run E2E/Integration tests. You can now
do that locally or in your CI, e.g. GH Actions or Gitlab Runners.

All that your CI runner does now is start your Dagger process and begin
everything inside Dagger. It's a far better user experience.

I need to sit down and write more about this amazing technology.

Tags:

    #dagger #ci
