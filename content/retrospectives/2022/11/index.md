+++
date = "2022-12-02"
categories = ["retrospectives"]
tags = ["retrospectives"]
slug = "retrospective-nov-2022"
title = "November 2022 Retrospective"
draft = false
+++

# Summary 

Two releases for [Mudmap]! :tada:

## Mudmap

1. Switched data model to [sqlc](https://sqlc.dev)
2. Added multi-stage installer

### Switching to [sqlc]

This isn't something I needed to do so why do it? Two reasons; type safety for the data model
and reduce friction when adding or updating models.

I experimented with gRPC for another project a couple of months ago and was awestruck by how much
the gRPC compiler does for you. It autogenerates thousands of lines of type safe code which you can
then call without needing to write anything yourself. As an average at best developer I loved 
this. Updating the yaml definition rebuilds the interface with a single command and all the callers
get updated. Of course, you might have to change your business logic depending on what you've 
updated, but you don't need to fiddle with any of the service layer. It's a huge time saver.

`sqlc` gives you this same experience but with sql. If tomorrow my `Device` model needs to include a
`Serial` field I would have to update every sql statement, all the `sql.Scan` method 
calls and then worry about the implementation within handler and service layers. Instead, by 
using `sqlc` I just update my sql queries and model definitions and the entire data model will be 
auto generated for me with those changes. Now all I need to do is focus on business logic 
changes (which the go type system will catch).

It was a bit of learning curve at first especially with `uuid` fields and `null`'s but that
speaks more to my flawed database schema design choices. *Rant alert:* for any product/project 
I produce in the future I will never use another `uuid`. A Stripe-like [id][ss] (e.g. `pi_123abc456`) 
or [Terraform][rp]'s `random_pet` (e.g. `epic_sawfish_123`) would be more than sufficient.

### Multi-Stage Modals

Perhaps my biggest cross the bear in Mudmap is the flaky installer. It really gets me down, and I've
still yet to find a good way to fix this issue - its churned one of my biggest customers too.

This feature does not **fix** the problem, instead it's part of my pathway to making the installation
process more pleasant. 

How it works (starts at *Password Confirmation*): 

![](mm-modal.svg 'A crude picture of how the process works')

*A crude drawing*

This change introduced a multi-stage modal instead of a single step. Using a multi-stage allowed for
better error handling and branching when prompting users for input. 

Additionally, the Device's root password is now being stored in an in-memory database. I've made an
ideological stand against storing this password in Mudmap's actual database. I feel that storing the
password in a memory datastore is a good practice and does not break the promise of not storing it
in **the** database. [go-memdb](https://github.com/hashicorp/go-memdb) was an easy choice for this.
The detractor is that in the future I may need to switch to an external key value store such as Redis
if I spin up multiple services.

By storing a Device root password in memory it can now be retrieved during the installation process 
across multiple stages. Without doing this, each stage would require the user to re-enter the 
password. This is not only inconvenient but also introduces needless complexities like what happens
if they re-enter it wrong - it is another error handler and conditional branch that needs to be 
dealt with.

Adding this also opens up future enhancements such as a proper Deletion event. When a Device is 
deleted or the installation fails, Mudmap must be removed from the firewall. This requires the 
removal of the Mudmap service account before removing the API. Once the Mudmap service account is
removed every action taken must be done as the root user. This requires root access which with this
new addition does not require user intervention.

A bad practice which I am trying to fix (and this fixes most of it) is not deleting Mudmap artifacts
when an installation fails. It is possible that these artifacts on subsequent re-installation 
attempts may even cause issues. But, the biggest problem is Mudmap isn't being a good citizen and 
reverting a user's device back to the state it was when they first tried the platform out.

I do not expect this will fix failed installations. What I might try next is to let users 
install it manually. This would require a few steps and may be too much friction, but it gives 
the user power to do something when an issue is encountered. Another thing I am contemplating is 
storing some metadata about each device in an attempt to correlations between successful and 
unsuccessful install attempts. I am unsure about this though.

todo: system logging etc features

## Zettelkastens

I use a [custom implementation][zet-cmd] of a [zettelkasten](https://en.wikipedia.org/wiki/Zettelkasten)
which I use to store snippets of information. The storage backend is GitHub and each Zet is
just a single markdown file stored in a timestamped directory within my [zet]
repo. Sometimes I want to review something I've
written, but I'm not at my computer and searching for it on GitHub is inefficient.
So I wrote a Go tool which would embed it into this [site](/zet).

At first, I wrote it to just link directly to the Zet's markdown file on GitHub.
I found this wasn't a great experience as it would take a few seconds to transition
and the layout shift was jarring. I did learn how to render a json file using
Hugo's custom [shortcodes].

Instead of this I decided host the markdown files directly on this site. I enjoyed
writing this because it leverages one of Go's truly great strengths; `text/template`.
If you've come from python you've probably heard of or used [Jinja2](https://jinja.palletsprojects.com/)
as the templating engine. It's a sweet module but its another dependency and having
an inbuilt templating engine within Go is one of its underappreciated features.

How I leverage Go's `text/template` to auto generate Hugo's frontmatter can be found [here][tp].
The code is pragmatic and works but don't look to it as an exemplar of *good* but instead
*practical*. [file1], [file2] and [GitHub actions].

## Recommendations

- [FIFA Uncovered](https://www.imdb.com/title/tt22872838/)

Shock horror FIFA is corrupt as hell. A well crafted overview of how FIFA quickly
became a money *under the table* organisation and how it continues to do so today.
When people can wield absolute power (without any oversight) they *will* abuse
it for their own gain.

## What's next

I'm testing out a new work/fun balance; 3-4 days on mudmap 3-4 days on whatever makes 
me happy per week. Going too hard too soon on Mudmap has been the thing that burns me out.
Also, once I release a feature or update I take a couple of days off to build whatever as a
reward. Doing this has actually lead to days where I've just jumped on [Excalidraw] and 
explored a Mudmap feature to build in the future. I'll play with this a see where it takes
me.

## Goals

I want to release another Mudmap update before Christmas.

## Analytics 

{{< plausible start=2022-11-01 end=2022-11-30 site_id=mudmap.io >}}

{{< plausible start=2022-11-01 end=2022-11-30 site_id=danielms.site >}}

[mudmap]: https://mudmap.io/?utm_campaign=retro-nov-22&utm_source=danielms&utm_medium=blog
[zet-cmd]: https://github.com/danielmichaels/zet-cmd/
[zet]: https://github.com/danielmichaels/zet/
[tp]: https://github.com/danielmichaels/danielms/blob/9b0d9d473196e47dc3d629b3552f77525d870839/create-zet-as-md.go#L166-L191
[shortcodes]: https://github.com/danielmichaels/danielms/blob/9b0d9d473196e47dc3d629b3552f77525d870839/layouts/shortcodes/render-zet.html
[file1]: https://github.com/danielmichaels/danielms/blob/master/fetch-zets.go
[file2]: https://github.com/danielmichaels/danielms/blob/master/create-zet-as-md.go
[github actions]: https://github.com/danielmichaels/danielms/blob/master/.github/workflows/zet-creator.yaml
[rp]: https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/pet
[ss]: https://gist.github.com/fnky/76f533366f75cf75802c8052b577e2a5
[sqlc]: https://github.com/kyleconroy/sqlc
[excalidraw]: https://excalidraw.com
