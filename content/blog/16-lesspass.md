+++
date = "15 apr 2019"
categories = ["password", "eli5"]
tags = ["eli5"]
slug = "lesspass-eli5"
title = "LessPass: A Primer"
+++

# LessPass

## The managerless password manager

LessPass, unlike other password managers does not store any information in a database or local cache. It simply takes in three pieces of information and then hashes that to return a password. If all three pieces of information haven't changed LessPass will always return the same password.

## Why use it over other choices

Straight from the creator himself:

- How do I synchronize this file on all my devices?
- How do I access a password on my parents’ computer without installing my password manager?
- How do I access a password on my phone, without any installed app?

For some of these I have my own workarounds but they are hacks and not very robust. For instance, I upload my password database to a cloud service and then pull it down on to my phone. It's a pain, not user friendly and exposes that database to a third-party vendor.

LessPass is also Free as in free beer, open source and secure. It has a command line implementation, a web application, android app (sorry iOS) and browser extensions. 

I would also add another reason to use it, its easy and what is easy gets used. Good luck getting your dad to download an applicaton that requires you to go to github, clone the repo and then run an install script or bat file. Even worse, good luck explaining to dad why 1Password is worth the 5 bucks a month or whatever it is. The smallest barriers to entry can be enough to be prohibitive when the user isn't computer savvy.

## How does it work

LessPass returns a unique password that it derives from three (or four, more on that later) pieces of information; `site`, `login` and a `master password`. 
If they are same it will always return the same output, a great example of idempotence or a *pure function*.


![alt text][lesspass]

An image from the creator's Medium [post] giving a great visual representation of what is happening under the hood.

![alt text][lesspass-gif]

The 'how' of LessPass and it's ease of use.

It also allows the user to set password rules such as using capitalised letters, special characters and numbers. So if certain websites require that the user *not* use special characters, LessPass can be configured for this. It defaults to 16 characters but can be set to a maximum of 32. 

You are also able to set whats called a `counter` meaning if you do not want to change your `master password` but are required to update your password you can issue a counter. It is an effective way to keep the three main pieces of information the same yet result in a new hash. Granted, this requires another level of cognitive load if you need to track several of these counters across sites.

**There must be a better way**

## LessPass Profiles

The application also offers something called a `profile` which allows you to store such information. This requires a sign up to LessPass which is free, or you can host your own [docker] image for extra security. Each `site` gets its own profile containing everything **except** the `master password` and the generated password. 

Each profile is stored as a `json` object like so:

```shell
{
    "login": "danielmichaels",
    "site": "www.github.com",
    "lowercase": true,
    "uppercase": false,
    "symbols": true,
    "numbers": true,
    "counter": 1,
    "length": 20
}
```

## Command Line Interface

I really like the CLI for LessPass - it suits my workflow by allowing me to set profiles on my machine via `alias`. I refer to this as a 'hardcoded' approach - it helps me none when I'm not on my box but I can just as easily hit lesspass.com for it in such a case so it's hardly a showstopper. As example this is how I grab my github password.

```shell
alias ghpw="lesspass -c github.com danielmichaels"
```

This drops me straight into a prompt for my `master password` and then copies it to my clipboard. I could go further and set an environment variable but I have my dotfiles on github and would rather not accidently expose such information. 

## I Don't Need A Password Manager

Said someone just waiting to get owned by a credential dump. We humans just simply cannot rememember suitably difficult passwords that are hard for a computer to guess and easy for us to remember.

Password manager public enemy number one is; if my master gets cracked they own everything. Yep thats true. If my house key gets cloned they've got unrestricted access to my house too. The mitigation lies in defeating the most likely avenue of exploitation and that is password attacks from credential dumps hoping to strike gold from someone that reuses their password on several sites. The bad guys are looking for the 'path of least resistance', just don't be that path! Using a password manager and some common sense will go a long way to protecting yourself online. It's not a panacea but its better than the alternative.

## Meta

You can find out more and get LessPass from [here]. The python package can be downloaded with `pip install lesspass` and both Chrome and Firefox extensions can be found in their respective stores. If you use LessPass be sure to give them a [star here] on Github.

[lesspass]: /images/lesspass-hash.png "LessPass image showing how hashing works."
[lesspass-gif]: /images/lesspass.gif "LessPass fullmotion gif of the app in action."
[post]: https://blog.lesspass.com/lesspass-how-it-works-dde742dd18a4
[docker]: https://github.com/lesspass/lesspass
[lesspass.com]: https://lesspass.com
[here]: https://lesspass.com
[star here]: https://github.com/lesspass/lesspass




