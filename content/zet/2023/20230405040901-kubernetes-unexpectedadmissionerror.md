+++
title = "kubernetes UnexpectedAdmissionError"
categories = ["zet"]
tags = ["zet"]
slug = "kubernetes-unexpectedadmissionerror"
date = "2023-04-05 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# kubernetes UnexpectedAdmissionError

Today we lost our storage from underneath us. This spawned hundreds of
pods which failed with `UnexpectedAdmissionError`.

To fix this, all we needed to do was restart our nodes after the storage
issue was fixed.

To delete all those pods I ran:

```shell
for pod in $(kubectl get pods -n <namespace> | grep -i UnexpectedAdmissionError | cut -d' ' -f 1);
do kubectl delete pod $pod -n <namespace>; done
```

Tags:

    #kubernetes


