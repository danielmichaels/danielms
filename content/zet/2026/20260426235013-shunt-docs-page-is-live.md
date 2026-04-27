+++
title = "Shunt docs page is live"
categories = ["zet"]
tags = ["zet"]
slug = "shunt-docs-page-is-live"
date = "2026-04-26 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Shunt docs page is live

I created a `mkdocs` page for my NATS based rule router; [shunt].

I use this extensively in my homelab. For instance,

- Making all my Zigbee2MQTT events available in VictoriaMetrics
- Notifications automations by making HTTP requests into my Knative handlers
- Telling me when my plants need watering
- Various other 

Some examples use cases:

- Telegram hooks for Jellysearr so I know that I need to approve kids requests (and when they are done)
- Synadia Control Plane alert webhooks into Knative which trigger telegram
- Motion sensors lights for my office

An example shunt rule:

```yaml
- name: z2m-metrics
  trigger:
    nats:
      subject: "zigbee2mqtt.*"
  action:
    http:
      url: "http://fn-z2m-metrics.knative-functions.svc.cluster.local/ingest"
      method: POST
      headers:
        X-Z2M-Device: "{@subject.1}"
      passthrough: true
      retry:
        maxAttempts: 3
        initialDelay: "2s"
        maxDelay: "10s"
```

This will listen for every message on the `zigbee2mqtt.*` subject and make a POST with the `*` subject token as the header. Internally, my knative function will receive the request stripping the header for our metrics endpoint.

For the uninitiated, knative functions are HTTP based and thats why we POST into the handler rather than listen on another subject and have `shunt` publish to it. Knative eventing works basically the same way but `shunt` means we can avoid the overhead of `eventing` completely and rely solely on `serving`.

This is a noisy endpoint and shunt works pretty well here. 

Tags:

    #shunt #nats #knative
