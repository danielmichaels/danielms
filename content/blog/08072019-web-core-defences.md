
+++
title = "Web Application Defence Mechanisms"
categories = ["web", "security"]
tags = ["security", "web"]
slug = "web-application-defence-mechanisms"
date = "2019-08-08"
draft = "true"
+++

# Core defences

_The fundamental security problem of all web applications: **user input is untrustworthy**_

Defensive mechanisms web applications utilise:

- Handling user access to application data and functions to prevent unauthorised access,
- Handling user input to prevent malformed input causing undesirable behaviour,
- Handling attacks or ensuring that the application works as intended when being directly targeted&mdash;including taking steps to frustrate the attacker, and
- Handling the application itself so everything can be monitored, logged and configured correctly.

These mechanisms are the main defence against the fundamental security problem but also account for a vast majority of an applications attack surface.

> If you know the enemy and know yourself, you need not fear the result of a hundred battles - Sun Tzu

In the same vein, in order to circumvent these defensive measure, you must understand them thoroughly.

## User Access

In applications that have users, there will be several privileges or roles within the user space. For instance, a social-networking site might have the base user, a group administrator and overall application administrator accounts. Likewise, regular account owners should be able to view their own data but not that of others. Mitigations, or security mechanisms in place within this spectrum are:

- Authentication,
- Session Management,
- Access Control.

### Authentication

The simplest to understand; a method to check the validity of a user, such as a username and password. Authentication should scale appropriately with the needs of the applications, for instance, MFA. Many authentication mechanisms also have password resets, account recovery, and self-registrations. All of which add to the attack surface.

### Session Management

On busy web applications there can be several other users sending requests simultaneously to the server requesting identical things. For instance, hundreds of people may be viewing the same page on Amazon, adding items to their basket all whilst browsing without being logged in. The application must keep track of theses requests and be able to distinguish each unique visitor.

This is done via tokens, or sessions that track each user and are stored on the webserver. The server then forwards the unique token to the client, which then returns that token in each subsequent HTTP request. This allows both sides to session track. Cookies are one such method for session management.

Common attacks against session management is stealing tokens, guessing tokens or exploiting defects in how the token is used to authenticate the owner. Stealing a session token could allow an attacker to use an authenticated users token effectively allowing them to masquerade as that user.

### Access Control

Simply put, the applications ability to assess whether a user has the required bona-fides to gain access to the system resources they are requesting. This often requires some fine-grained logic and thus can be exploited if developers make too many assumptions about how a "user" will access the resources. 


## User Input

All user input is untrusted. Therefore, applications must handle all input with strict enforcement of this rule. Sanitisation, restriction of inputs and reducing assumptions about what a "user" would submit are necessary precautions. But, protecting against all forms of malicious input is a difficult task and many applications are vulnerable to this.


### Varieties of Input

Limiting what a user can input is common practice. For example, a user sign up page may use an email account that requires the use of an @ symbol and must only contain alphanumeric letters with a minimum length of four characters. Other examples may be the requirement to submit a textual paragraph that allows all alphabetical characters plus special characters but will not allow HTML tags. Many other possibilities exist such as markup languages and how they are delimited so as not to allow arbitrary code to be run.

## Approaches to Input Handling

### "Reject Known Bad"

Blacklisting known bad things. The weakest form of defence, as spoofing or encoding arbitrary code blocks could slip past this form of defence. 

For example, if `SELECT` is blocked, `SeLeCt` might not be or if `alert('xss')` is blocked maybe `prompt('xsss')` isn't. In addition, some tokensing parsers which detect these blacklisted words may stop searching when the detect a NULL Byte. `%00<script>alert(1)</script>` may get through.

### "Accept Known Good"

This method is very effective against stopping code injection as it is based off only allowing accepted elements to be input. The flaw in this is where accepting a known good can conflict with the requirements of the application itself. For instance, a name field in a form will likely necessitate the use of `-`, as some people legitimately have hyphenated names. But, `-` are commonly used to attack databases on the backend, so it may not be used in all cases where it could detract from the user experience. 

### Sanitisation

A common defence mechanism that relies on the fact that users will try malicious code, so before processing data, the application will sanitise the input in various ways. Methods such as removing dangerous characters or escaping them before processing. 

## Boundary Validation

Simplistically, this is the process of validating, or checking user input on the peripheries first prior to the input reaching the backend. Think taking dirty input on the web side and cleaning it before sending to the server side, and therefore trusting all server side code to be "clean". 

This has some serious faults. Often attackers can sidestep or exploit such a scheme by chaining exploits in a way that no validation method could feasibly defend against without serious detriment to the overall functionality of the application itself. 

The best solution to defending against this, is boundary validation where for each layer of processing, a validation check is conducted appropriate to that layer; **defence in depth**.

## Handling Attackers

All applications should be of the mindset that they will be targeted and attacked by dedicated and skilled attackers. With this mindset, the developers must have a methodology for handling and monitoring such incidents. These measures should include:

- Error Handling,
- Logging,
- Alerts,
- Reacting to attackers.

### Handling Errors

Robust testing of the application prior to launch should be conducted, and if done properly identify many bugs and errors in the design and codebase. But, not all issues can be counted for in testing and as such there must be a way to capture and respond to errors in the application. 

Debug and system messages should never be presented to the outer layers - they greatly assist the attackers in identifying weaknesses in the application. All errors should be gracefully raised, and catered for so as to not give away too much information to the attacker.

### Audit Logs

All activities taken on the system should be logged for analysis later. Items that definitely require logging are:

- All authentication actions; failed, successful, and forgotten or changed passwords,
- Transactions such as payments,
- Activities blocked by access control mechanisms,
- Any attack strings such as malicious input.

It is important that these logs are protected against unauthorised read and writes. This can sometimes be achieved by the use of autonomous systems that only accept input data from the servers. 

### Alerting Administrators

Audit logs are retroactive; they allow a detailed look into the past for forensic analysis but provide little real-time information. Alert mechanisms that should be brought up to the administrators attention are things like:

- Anomalies, such as a large number of requests originating from a single IP,
- Business anomalies like unusual funds transfers or out of hours actions on a machine,
- Requests containing attack strings,
- Requests which data is hidden with the intent to obfuscate the true intent of the data, or request.

### Reacting to Attacks

Alerts signifying an attack are useless without a response. Reactive approaches should be unique to the application in question but many are generic. Defences designed to slow or block simple attacks will not defeat a prudent and patient attacker, but should stop a lot of automated broad brush attacks.
