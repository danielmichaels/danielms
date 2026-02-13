+++
title = "SCP nats-box creds k8s"
categories = ["zet"]
tags = ["zet"]
slug = "scp-nats-box-creds-k8s"
date = "2026-02-13 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# SCP nats-box creds k8s

Note to self, for k8s local testing a quick way to do some debugging with access to nats is via `nats-box`. In control plane you need creds so I just cp a creds file into the container.

```
kubectl cp ~/Downloads/my.creds nats-box-664bcb786c-6wt22:/tmp/my.creds -n nats
```

Then use it from inside the container.

If you have network path to the cluster/control plane server then you can just use `nats` CLI as is. E.g. `nats -s $NATS_URL --creds my.creds rtt`

Tags:

    #nats #debug #k8s

