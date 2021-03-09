
+++
title = "Getting Fedora 32 and Docker to Play Nice"
categories = ["linux", "docker"]
tags = ["fedora", "docker"]
slug = "getting-fedora-32-docker-play-nice-a-how-to"
date = "2020-08-29"
draft = "false"
+++

# Fedora 32 and Docker Are Incompatible

![Fedora and Docker logos](/images/fedora-docker.png "fedora and docker logo's combined")

Fedora is a great distribution for development, its very customizable, and supports a great number of packages and applications out of the box. If you're going to use a linux environment thats not Debian based, it ought to be fedora.

Sadly nothing is ever perfect, meaning Fedora too, has its downsides. Currently, at least to me, is its inability to run Docker out of the box.

## Its actually dockers fault

To be fair its not Fedora's fault but Dockers. Fedora no longer supports [control group][0] version 1 and has moved onto the much better version 2. Unfortunately, Docker only supports version 1 and they aren't backwards compatible. 

This effectively makes using Fedora with docker untenable without modification. 

## Podman to the rescue

Sure, you can use podman but the world still uses docker and more specifically docker-compose.

Your work environment might even negate both by just using kubernetes or openshift. But in the open source and hobby project world that is overkill.

So, let's instead support both.

## Why not have both?

While the current version of docker (19.03 as of writing) doesn't support cgroupv2, it does support it on its [moby][7] branches.

Unstable? Well, its either that or no docker for you! It appears docker will merge support and make a release sometime this year - 2020. But until then you need to use the binaries supplied from compiling the docker branches.

Thankfully, some [legend][4] has already done it. Go to the [github][5] page, download the latest release and follow the installation instructions - which is just `dnf install`. Once its installed, you will now have docker with cgroupv2 support.


Your new docker setup will still need to be configured as per the normal [docker instructions][6]

Profit!

## Docker and Podman in harmony

Now you can safely run both docker and podman without downgrading anything and still get all the benefits of fedora and cgroupv2.

I personally think Podman is great, but so is docker. For my outside-of-work humble purposes, I only need docker and am thankful there is a workaround. Because, unpopular opinion, macosx sucks and so does ubuntu.

Docker has a rich ecosystem, massive mind share and wonderful support in both vscode and pycharm. Something most other container platforms lack. Until podman has a better integration I think its uptake in the developer community at large will be slow.

## Recap

Fedora 31+ cant run docker unless you install the precompiled binaries from the Moby branch.

After installing these, you'll have docker and podman working seamlessly.

No excuse to not use fedora as a docker acolyte.

Also if you want more information as to *why* we should be switching to cgroup version 2 see [here][1] (this dude is legit, probably *the* authority on linux), [here][2] and [here][3].

[0]: https://wiki.archlinux.org/index.php/cgroups
[1]: https://www.youtube.com/watch?v=yZpNsDe4Qzg
[2]: https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/8/html/system_design_guide/understanding-control-groups_setting-limits-for-applications
[3]: https://0xax.gitbooks.io/linux-insides/content/Cgroups/linux-cgroups-1.html
[4]: https://twitter.com/_AkihiroSuda_?s=20
[5]: https://github.com/AkihiroSuda/moby-snapshot
[6]: https://docs.docker.com/engine/install/linux-postinstall/
[7]: https://github.com/moby/moby