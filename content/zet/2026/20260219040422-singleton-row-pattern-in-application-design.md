+++
title = "Singleton row pattern in application design"
categories = ["zet"]
tags = ["zet"]
slug = "singleton-row-pattern-in-application-design"
date = "2026-02-19 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Singleton row pattern in application design

Anti-pattern? It depends.

I'm using it in a situation where only one instance of a key can ever exist. New key? We update the existing. It is a global state mechanic. The alternative is an Entity-Attribute-Value table - also a "it depends" anti-pattern.

In this case I'm creating an identifier with the same value, always. All `INSERT` and `DELETE` operations target the same known `id`. Has a nice benefit of making upsert behaviours deterministic.

Whilst I have this feeling of it being "wasteful" to use a single table for a single row - I think it works (we'll see what my workmates say!).

Some alternatives; `is_active` `BOOLEAN`. This is nice because we have a historical log. We just insert new keys with `is_active=1` and set all the others to `is_active=0`. Or, `ORDER BY created DESC LIMIT 1`. But, it adds some complexity like deactivating old rows when activating the new one, queries are a little more complex (I guess not *too* complex - `WHERE is_active == 1`). I think the downside here is AFAIK the database doesn't provide a guarantee only a single row is active.

> It depends

Is the right call here but for my current use case it works well enough. Should I ever need more historical analysis then moving to the `is_active` model is a no-brainer. 

Also, I think you could have the singleton and a separate history table. I have done this in the past as a cheap timeseries like purview of the changes - it that example it was tracking DNS records over time; current versus historical.

Tags:

    #database #design
