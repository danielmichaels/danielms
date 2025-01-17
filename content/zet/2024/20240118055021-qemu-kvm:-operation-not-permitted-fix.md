+++
title = "qemu kvm: operation not permitted fix"
categories = ["zet"]
tags = ["zet"]
slug = "qemu-kvm:-operation-not-permitted-fix"
date = "2024-01-18 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# qemu kvm: operation not permitted fix

When starting a bridged VM and hit with:

```
Error starting domain: internal error: /usr/libexec/qemu-bridge-helper --use-vnet --br=virbr0 --fd=20: failed to communicate with bridge helper: Transport endpoint is not connected
stderr=failed to create tun device: Operation not permitted
```

It is likely that `qemu` has been updated and changed the bridge helper permissions.

Fix it with:

`sudo chmod u+s /usr/lib/qemu/qemu-bridge-helper`

Tags:

  #qemu #vm #kvm

