+++
title = "Mudmap org and new user registration"
categories = ["zet"]
tags = ["zet"]
slug = "Mudmap-org-and-new-user-registration"
date = "2022-06-16 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Mudmap org and new user registration

The rough registration process for adding a user to an `Organisation`.

1. Admin account enters email into Org settings tab of user they wish to invite
2. Email is sent to them with a short lived key
3. User opens email and either signs up or logs into Mudmap via Auth0
4. User navigates to Settings -> Organisation and clicks *Join Org*
5. User enters key from email
6. If key is expired, alert user and admin
7. Otherwise, add user to Org and remove them from existing

Scope of work (no particular order)

- Invite form
- Invite Email (sendgrid template & Go code)
- Registration flow documentation with video
- Token/Key table for tracking purposes, auth'ing Invites on backend 
- Token auth endpoint for checking validity
- Endpoint ^ add user to Org in DB, and in Auth0 Groups
- Email alerting Admins user has joined
- Email alerting Admins user attempted to join but token was out of date?
- Or, just allow user to get new token?
- Email welcoming User to Org

Tags:

    #mudmap #architecture

