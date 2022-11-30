+++
title = "inlets troubleshooting"
categories = ["zet"]
tags = ["zet"]
slug = "inlets-troubleshooting"
date = "2022-11-14 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# inlets troubleshooting

[inlets.dev](https://inlets.dev) troubleshooting.

So far, I've had some issues getting to work out of the box. I think this
mostly comes down to the documentation drifting out of sync with the project.

```shell
inletsctl create \
--access-token-file ~/.doctl-token \
--provider digitalocean \
--region sgp1 \
--letsencrypt-domain inlets.cupscanteen.com \
--letsencrypt-domain t.cupscanteen.com \
--letsencrypt-domain slack.cupscanteen.com \
--letsencrypt-email webmaster@cupscanteen.com
```

Sometimes this doesn't work. So to get it working, I jump on the host. For
DO the password is emailed when you create the host (which the above step
does for you).

```shell
vim /etc/default/inlets-pro
# add any extra domains with --letsencrypt-domain <sub>.<domain>.<tld>
# after the DOMAINS=
```

The documentation states its `--letsencrypt-domain=foo.com,bar.com` but
this does not seem to work.

`systemctl restart inlets-pro` and see if the process has started correctly. It
can take time to get the cert from Let's Encrypt.

If this isn't working, try changing `/etc/default/inlets-pro` `ISSUER`
from `prod` to `staging` and restart the service. This should connect
and if so, you know its a Let's Encrypt issue. Potentially the IP has been
temporarily rate limited, or the DNS entry does not exist so it cannot do
a HTTP01 interrogation

Tags:

    #inlets #tunnelling #troubleshooting
