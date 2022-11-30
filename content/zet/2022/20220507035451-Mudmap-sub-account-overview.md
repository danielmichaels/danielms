+++
title = "Mudmap sub-account overview"
categories = ["zet"]
tags = ["zet"]
slug = "Mudmap-sub-account-overview"
date = "2022-05-07 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Mudmap sub-account overview

A rough idea of what this change should achieve.

- `organisation` is has a collection of users
- `users` has a relation to `permissions`
- `device` belong to `organisation`, not `users`
- Existing `users`'s must be attached to a `organisation` (probably by creating a `organisation` using the
  `user.Name` field as `group.Name` and adding them as a `users` to that group)
- `organisation` owners can update `organisation` info
- `permissions` should be a Many-to-Many. `users` can have many `permissions` and the same
  permission can belong to many `users`.
- Auth0 should also be aware of `users` and `organisation`
- Auth0 should also be aware of permissions
- Auth0 should place `permissions` within the token (if possible)
- Updates to the `permisions` table must also be reflected in Auth0
- Updates to the `organisation` table must be reflected in Auth0 

**Next Steps**

1. Test assumptions with Auth0's Authorization and Group tooling 
2. Create `permissions` and `organisation` table
3. Update `users` table to have relation with `organisation` 
4. Create change `devices` relation from `users` to `organisation`
5. Write sync methods so that the DB and Auth0 remain in sync 
6. Middleware for checking a user can access a device based on JWT token fields

This is a decent sized change with a few moving parts. I can possibly 
remove the Auth0 lock-in but I see no good way of doing this without
forcing a DB lookup for each request (check permission and group). If
I use Auth0 I can hopefully keep it sync'd with the DB and embed that
data in the `app_metadata` field of the token. A assumption I need to 
test fully before too much investment in code.

Related:

- [20220505024039 Mudmap sub-accounts](/20220505024039/)

**Updates**

I've added `app_metadata` to the JWT. This can now be plucked off on the 
backend.

Its possible to use M2M API key to make changes to the Group, 
Permissions and Roles of Users and organisation. However, the free-tier only
gives you 1000 calls per month (from what I can tell)

API Docs: https://auth0.com/docs/api/authorization-extension
Getting a token: https://auth0.com/docs/api/authorization-extension?shell#get-an-access-token

Some calls examples:

```sh
# get all groups
curlie https://mudmap.au12.webtask.io/<URL>/api/groups -H "authorization: Bearer <TOKEN>"
```
Returns 

```json
{
    "groups": [
        {
            "_id": "8931fb3a-8163-4c8a-8534-b5cc43381172",
            "name": "Test-Group",
            "description": "For Testing Only",
            "members": [
                "auth0|61de4f3e8e3c6000710b8c0d"
            ],
            "mappings": [

            ]
        },
        {
            "description": "test-description",
            "name": "new-test",
            "_id": "0fdf83fe-32db-463c-b362-4e29c0817781",
            "mappings": [

            ],
            "members": [

            ]
        }
    ],
    "total": 2
}
```


Tags:

    #mudmap #planning
