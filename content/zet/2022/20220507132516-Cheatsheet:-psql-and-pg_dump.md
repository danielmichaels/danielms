+++
title = "Cheatsheet: psql and pg_dump"
categories = ["zet"]
tags = ["zet"]
slug = "Cheatsheet:-psql-and-pg_dump"
date = "2022-05-07 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Cheatsheet: psql and pg_dump

In my day job I use MariaDB so I often forget the syntax for Postgres and
I'm trying not to leverage graphical DB explorers as often. When jumping
on to containers from a k8s orchestration platform, you don't often have
access to anything but the CLI.

## Commands

**Login**

In the container: `psql -U <username> -d <database>`
On local host to db: `psql -h localhost -p 5432 -U dbuser -d db`

Not specifying the database will cause an error, unlike MariaDB which
just wants a username and password.

**Show Tables**

`\dt`

**Desc Table**

`\d <table-name>`

**Dump DB into File**: 

`pg_dump -U <user> -d <database> >> sqldump.sql` then run `docker cp`
if its inside a container.

**Restore DB from file**:

`psql -f sqlfile.sql -U dbuser -p 5432 -h localhost -d db`

**Install `psql` without Postgres

`sudo apt install postgresql-client`

## Still stuck? 

Instead of reaching for google, run `curl cht.sh/psql` it is fantastic.

Tags:

    #database #psql #postgres
