+++
title = "get private ssl cert"
categories = ["zet"]
tags = ["zet"]
slug = "get-private-ssl-cert"
date = "2022-08-29 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# get private ssl cert

This is how I get private SSL certs which is really useful when using
private gitlab and argocd

```shell
openssl s_client -showcerts \
-servername gitlab.homelab.com \
-connect gitlab.homelab.com:443 \
</dev/null 2>/dev/null | sed -n -e \
'/BEGIN\ CERTIFICATE/,/END\ CERTIFICATE/ p' \
> /tmp/sslkey.pem
```

Tags:

      #openssl #ssl #argocd
