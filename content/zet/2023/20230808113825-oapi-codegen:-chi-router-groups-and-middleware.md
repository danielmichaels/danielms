+++
title = "oapi-codegen: chi router groups and middleware"
categories = ["zet"]
tags = ["zet"]
slug = "oapi-codegen:-chi-router-groups-and-middleware"
date = "2023-08-08 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# oapi-codegen: chi router groups and middleware

Today I learned a valuable lession in assumptions.

I assumed that putting the `oapivalidator` with the `router.Group` where the endoints *live* was
how it worked. Not realising that **zero** validation was being done. Not until I started implemented
JWT authentication and noticing that no errors were propagating.

Here is the code that works.

```go
router := chi.NewRouter()
// middlewares
router.Use(oapimiddleware.OapiRequestValidator(openapiSpec))
// endpoints
router.Group(func(api chi.Router){
    server.HandlerFromMuxWithBaseURL(s, router, "/api/v1")
})
```

Tags:

    #go #til #aar
