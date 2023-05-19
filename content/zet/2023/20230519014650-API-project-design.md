+++
title = "API project design"
categories = ["zet"]
tags = ["zet"]
slug = "API-project-design"
date = "2023-05-19 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# API project design

A rough outline of how I'll implement an API for a project I'm developing
using `oapi-codegen`.

Endpoints:

```yaml
/healthz:
    - GET
/users:
    - GET
    - POST
    - PATCH
    - DELETE
/monitors:
    - GET
/monitors/{monitor-type}:
    - GET
    - POST
    - PATCH
    - DELETE
/monitors/{monitor-type}/{monitor:id}
    - GET
    - POST
    - PATCH
    - DELETE
```

The barebones are there but implementation needs some work.

Tags:

    #planning #openapi #go
