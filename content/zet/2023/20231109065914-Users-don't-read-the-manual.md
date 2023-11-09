+++
title = "Users don't read the manual"
categories = ["zet"]
tags = ["zet"]
slug = "Users-don't-read-the-manual"
date = "2023-11-09 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Users don't read the manual

In my workplace I've written a couple of tools. Now its more than just me
committing code to the project but I still do 90% of the ongoing work.

It bootstraps a local kubernetes cluster with all our teams resources 
and applications. We're a dev team, not platform or infra. Most of them
don't know kubernetes which is why I created this tool.

It's a `pip install` and within 3 minutes you're entire dev env is stood up.

Yet, I get constant questions. Some of the answers are literally printed
to the screen once the `bootstrap` step completes.

It's hard not to get frustrated. On one hand, you can't expect users to 
understand everything. On the other, there's team members who do read 
the notices and take the time to investigate the tools `--help` before
throwing their hands up saying "it doesn't work".

All I can do is try to educate and keep working on it to make it better.

Still, its a real epidemic in the open source world. I see so many issues
on GitHub which are borderline aggressive towards project maintainers when
their problem was literally solved in the `--help` text.

Tags:

  #oss #people

