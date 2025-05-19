+++
title = "Slack Bolt Python oauth install fix"
categories = ["zet"]
tags = ["zet"]
slug = "slack-bolt-python-oauth-install-fix"
date = "2024-09-26 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Slack Bolt Python oauth install fix

https://github.com/slackapi/bolt-python/issues/492

It was not apparent to me that the `slack-bolt` python app requires you to start
the oauth flow from <url>/slack/install

At that address it will load a page with a "Install" button (not your own) and
only clicking that to start the flow will work.

Otherwise it don't work! Super confusing and wasted a number of precious hours
on this.

Tags:

    #slack #python
