+++
title = "Defending your apps"
categories = ["web"]
tags = ["web", "security"]
slug = "defence-mechanisms-for-web-apps"
date = "2019-10-13"
draft = "false"
+++

# Web Application Defence Mechanisms

> one rule; do not trust user input.

![](/images/defence-mechanisms-post.jpg "binary globe")

## Core defences

The methods by which we can defend against attackers.

- Handling user access
- Handling user input
- Handling attacks or ensuring that the application works as intended when being directly targeted, and
- Handling the application itself so everything can be monitored, and configured correctly

### User Access

In most applications that have users, there will be several privileges or roles within the user space. For instance, a social-networking site might have the base user, a group administrator and overall application administrator accounts. Likewise, the owner of an account should be able to view their own data but not that of others. Mitigation's, or security mechanisms in place within this spectrum are:

- Authentication,
- Session Management,
- Access Control.

### Authentication

The simplest to understand; a method to check the validity of a user - username and password. Authentication can scale appropriately with the applications need, such as MFA. Many authentication mechanisms also have password resets, account recovery, and self-registrations. All of which add to the attack surface.

### Session Management

On busy web applications there can be several other users sending requests simultaneously to the server requesting identical things. For instance, hundreds of people may be viewing the same page on amazon and adding an item to their basket whilst browsing without being logged in. The application must keep track of theses requests and be able to distinguish each unique visitor.

This is done via tokens, or sessions that track each user and are stored on the web server. The server then forwards the unique token to the client, which then returns that token in each subsequent HTTP request. This allows both sides to session track. Cookies are one such method for session management.

Common attacks against session management is stealing tokens, guessing tokens or exploiting defects in how the token is used to authenticate the owner. Stealing a session token could allow an attacker to use an authenticated users token effectively allowing them to masquerade as that user.

### Access Control

Simply put, the applications ability to assess whether a user has the required Bona-Fides to gain access to the system resources they are requesting. This often requires some fine-grained logic and thus can be exploited if developers make too many assumptions about how a "user" will access the resources.


## User Input

All user input is untrustworthy. Therefore, applications must handle all input with strict enforcement of this rule. Sanitation and restriction of inputs and reducing assumptions about what a "user" would submit are necessary precautions. But, protecting against all forms of malicious input is a difficult task and many applications are vulnerable to this.


### Varieties of Input

Limiting what a user can input is common practice. For example, a user sign up page may use an email account that requires the use of an @ symbol and must only contain alphanumeric letters with a minimum length of four characters. Other examples may be the requirement to submit a textual paragraph that allows all alphabetical characters plus special characters but will not allow HTML tags. Many other possibilities exist such as markup languages and how they are delimited so as not to allow arbitrary code to be run.

## Approaches to Input Handling

### Reject Known Bad

Blacklisting known bad things. The weakest form of defence, as spoofing or encoding arbitrary code blocks could slip past this measure.

For example, if `SELECT` is blocked, `SeLeCt` might not be! or if `alert('xss')` is blocked maybe `prompt('xsss')` isn't. In addition, some tokensing parsers which detect these blacklisted words may stop searching when the detect a NULL Byte. `%00<script>alert(1)</script>` may get through.

### Accept Known Good

This method is very effective against stopping code injection as it is based off only allowing accepted elements to be input. The flaw in this is where accepting a known good can conflict with the requirements of the application itself. For instance, a name field in a form will likely necessitate the use of `-`, as some people legitimately have hyphenated names. But, `-` are commonly used to attack databases on the backend, so it may not be used in all cases where it could detract from the user experience.

### Sanitisation

A common defence mechanism that relies on the fact that users will try malicious code, so before processing the application will sanitise the input in various ways. Methods such as removing dangerous characters or escaping them before processing. This is something all applications should do, and their are many libraries in every language to support this.

## Boundary Validation

Simplistically, this is the process of validating, or checking user input on the periphrases first prior to the input reaching the backend. Think taking dirty input on the web side and cleaning it before sending to the server side, and therefore trusting all server side code to be "clean".

This has some serious faults. Often attackers can exploit such a scheme by chaining exploits in a way that no validation method could feasibly defend against without serious detriment to the overall functionality of the application itself.

The best solution to defending against this, is boundary validation where for each layer of processing, a validation check is conducted appropriate to that layer; defence in depth.

## Handling Attackers

All applications should be of the mindset that they will be targeted and attacked by dedicated and skilled attackers. With this mindset, the developers must have a methodology for handling and monitoring such incidents. These measures should include:

- Error Handling,
- Logging,
- Alerts,
- Reacting to attackers.

### Handling Errors

Robust testing of the application prior to launch should be conducted, and if done properly identify many bugs and errors in the design and code base. But, not all issues can be counted for in testing and as such there must be a way to capture and respond to errors in the application.

Debug and system messages should never be presented to the outer layers - they greatly assist the attackers in identifying weaknesses in the application. All errors should be gracefully raised, and catered for so as to not give away too much information to the attacker.

### Audit Logs

All activities taken on the system should be logged for analysis later. Items that definitely require logging are:

- All authentication actions; failed, successful, and forgotten or changed passwords,
- Transactions such as payments,
- Activities blocked by access control mechanisms,
- Any attack strings such as malicious input.

It is important that these logs are protected against unauthorised read and writes. This can sometimes be achieved by the use of autonomous systems that only accept input data from the servers.

### Alerting Administrators

Audit logs are retroactive; they allow a detailed look into the past for forensic analysis but provide little real time information. Alert mechanisms that should be brought up to the administrators attention are things like:

- Anomalies, such as a large number of requests originating from a single IP,
- business anomalies like unusual funds transfers or out of hours actions on a machine,
- requests containing attack strings,
- requests which data is hidden with the intent to obfuscate the true intent of the data, or request.

### Reacting to Attacks

Alerts signifying an attack is nothing without a response. Reactive approaches will should be unique to the application in question but many are generic. Things such as slowing or putting up things to block simple attacks will not defeat a prudent and patient attacker, but should stop a lot of automated broad brush attacks.


## Summary

The above methods for protecting web applications are nothing new, in fact, much of this content is taken from old resources such as the [web application hackers handbook][1]. That is because the advice still holds true even in today's "modern" web app's. Attacks and methods may change slightly, or evolve to overcome new frameworks securities mechanisms but the vast majority of attacks simply exploit what could be defended against by implementing whats contained in this post.

Things don't change that much.

[1]: https://portswigger.net/web-security/web-application-hackers-handbook
