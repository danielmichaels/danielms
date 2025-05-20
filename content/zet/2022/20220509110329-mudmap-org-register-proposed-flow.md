+++
title = "Mudmap Org Register Proposed Flow"
categories = ["zet"]
tags = ["zet"]
slug = "mudmap-org-register-proposed-flow"
date = "2022-05-09 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Mudmap Org Register Proposed Flow

The biggest issue I've come to experience with the new multi-account
design is re-arranging how Users fit.

I've decided `organisation` is a better name than `account` or `business`.

## Flow

1. User signs ups (entered in DB)
2. Org is created, User assigned to Org as Admin (Org name is user.email)
3. Users are still tracked in DB using their `user_id` from Auth0 
4. Org's will now get the Stripe `sub_id` (**Not** the User) 
5. User's are added to Org's by invitation 
6. User's must accept the invite, triggering the move. 
7. Devices belong to Org's
8. All members of an Org can *see* Devices 
9. Existing devices will be moved to each User's (new) Org 
10. Org delete will CASCADE 

Everything is tracked in the database but "synced" with Auth0 using the
Authorization Extension. This allows Mudmap to review a User's access
to an Org and Device (permissions are omitted for now) by inspecting
the JWT. If the JWT has some limitation that I've not accounted for then
each request will require a DB lookup. For latency reasons I'd prefer 
not to do that. If, that is the case it's probably a prime reason to 
transition to Litestream (assuming replication is working well).

This is a ten thousand foot view and still needs a few things ironed out.

Tags:

    #mudmap #research #planning
