+++
title = "Mudmap Org Permissions"
categories = ["zet"]
tags = ["zet"]
slug = "Mudmap-Org-Permissions"
date = "2022-05-05 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Mudmap Org Permissions

Mudmap must be able to provide user accounts under an account root. This
account root will be called an `organisation`. All users will belong to
an `organisation`.

RBAC policies are applied per-user according to their level of access.

As a general guide, a `root` account is created from which all accounts 
for that company or group are then underneath. This account is the 
`organisation` owner for lack of better wording. It might be prudent to
allow multiple `root` account holders.

This means devices are attached to an account or at least an account user. Users
associated to that user can see any device in the hierarchy. At the moment, only
a single user account can have and see its own devices. This will need to be radically
altered to accommodate these changes. 

## RBAC Policies 

| Account Type | Payments | Privileges | Read | Write |
|---|---|---|---|---|
| Root | y | y | y | y |
| Privileged | n | y | y | y |
| Manager | n | n | y | y |
| User | n | n | y | n |
| Custom | ? | ? | ? | ? |

* managers cannot elevate privileges (only privileged account or root) 

## Questions

- How to create this type of account setup with Auth0? Do I need to hack something together
- How to grant RBAC - JWT or DB look up on each request? JWT best case, DB worse but either way RBAC should be stored in a DB for reference in admin panels 
- How to associate users to the root account programmatically with backward compat 

Tags:

    #mudmap #research
