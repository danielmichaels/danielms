+++
title = "git log to PR message"
categories = ["zet"]
tags = ["zet"]
slug = "git-log-to-PR-message"
date = "2024-10-20 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# git log to PR message

`git log <sha>..HEAD | mods "write a concise pull request for these changes"`

And it produced a really decent summary of all the changes. I think the key is
create meaningful commit messages.

I like to add a lot of context to my commit bodies and it pays off here.

Example

```markdown
## Summary

This pull request introduces multiple updates and refactors across the dashboard
templates, email account management, and Slack integration. These changes aim to
enhance performance, usability, and maintainability.

## Key Changes

1. Dashboard Enhancements:\
   • Simplified, restructured, and cleaned up dashboard templates, notably\
   accounts.html and connections.html ([Commit 5bd7454]).\
   • Added accounts listing with filtering, sorting, and pagination ([Commit\
   a2b906d]).\
   • Implemented detailed connection display and profile updates ([Commit\
   d7cdda8]).
2. Account and Connection Management:\
   • Introduced delete functionality with confirmation for accounts and\
   connections ([Commit 72818d1]).\
   • Improved URL routing and messaging for account actions in views.py\
   ([Commit d968d7e]).\
   ..truncated
```

Tags:

    #ai #git
