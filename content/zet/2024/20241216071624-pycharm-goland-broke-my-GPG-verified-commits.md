+++
title = "pycharm-goland broke my GPG verified commits"
categories = ["zet"]
tags = ["zet"]
slug = "pycharm-goland-broke-my-GPG-verified-commits"
date = "2024-12-16 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# pycharm-goland broke my GPG verified commits

1. Open `~/.gnupg/gpg-agent.conf`
2. Comment the line that uses `pinentry-program <path>/pinentry-ide.sh`
3. Reload gpg agent with `gpgconf --reload gpg-agent`

Hopefully they fix this in the next release (they usually are very responsive to
these issues) because each time you open the IDE it's going to prompt you to
install this script again.

ref:
https://youtrack.jetbrains.com/issue/IJPL-173525/Git-GPG-signing-fails-with-errors-like-Bad-CA-certificate-or-failed-to-write-commit-object

Tags:

    #ide #gpg #bug
