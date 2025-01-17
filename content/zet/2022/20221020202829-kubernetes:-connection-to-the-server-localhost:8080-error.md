+++
title = "kubernetes: connection to the server localhost:8080 error"
categories = ["zet"]
tags = ["zet"]
slug = "kubernetes:-connection-to-the-server-localhost:8080-error"
date = "2022-10-20 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# kubernetes: connection to the server localhost:8080 error

This annoying error is the result of a bad `$KUBECONFIG` file/ reference.
To fix this you will need to get a new copy from the server. 

For k3s this will mean scp'ing `/etc/rancher/k3s/k3s.yaml` to your client
and then changing the server IP from `127.0.0.1` to the actual IP of the
remote.

`oc` helpfully does all this for you but `k3s` does not. 

Tags:

    #kubernetes #troubleshooting


