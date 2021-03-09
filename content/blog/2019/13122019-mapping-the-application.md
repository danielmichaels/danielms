+++
title = "Mapping an Application"
categories = ["web"]
tags = ["web", "security"]
slug = "mapping-applications-web-pentest"
date = "2019-12-23"
draft = "true"
+++

# Mapping the Application

This chapter is all about information gathering, fingerprinting and enumerating the application.

## Enumerating Content and Functionality

### Web Spidering

Spidering is the automated testing of an application that seeks to find all the possible endpoints or pages within the app. It does this by requesting a page, following all the links and recursively repeating that proces until no new content is found.
Smart spiders will also parse HTML forms and submit them back to the application using presets or random values. Some will parse client-side JavaScript to get to more URL's.

**Some spiders will actively target the `robots.txt` page and use any URL's contained within as its starting seed, thereby making the use of `robots.txt` somewhat counterproductive.**

## Basic Spidering Steps

1. configure the browser to use BurpSuite or similar as a local proxy
2. Browse the entire application normally, attempting to vist all URL's/links, submitting every form and proceed through all multi-step functions to completion. Try it with caching, cookies, and javascript on and off.
3. Review the site map your scanner has generated to see if any hidden content or functionality exists of which you did not see manually. establish how the spider enumerated each item.

### Discovering Hidden Content

It is common for applications to have hidden content within the site. Some examples may be testing or development portions that have not been removed, different categories of users (admin, anon, regular, etc). Users with different levels of authorisation may provide priveldge escalation on the application if an attacker can find access. There are many examples of content that can be hidden within the application, both surruptiously and accidentally.

### Brute-Forcing

Many scanners are capable of brute forcing an application by spidering across many similar URL's, or using a custom  script/ wordlist that aids in their effectiveness. This can greatly aid the spidering process and speed up the enumeration of the application. Burp Intruder is one such way of doing this.

## Mapping the Attack Surface

There are numerous ways to identify a applications structure, processes, naming conventions, underlying technology, etc. Once enumeration is complete we are going to look for one or more of the following expliotable methods listed below:

- Client-side validation - checks that may not be replicated on the server
- Database interaction - SQL injection
- File upload/ download - Path traversal, XSS
- Display of user-supplied data - XSS
- Dynamic redirects - redirection and header injection attacks
- Social networking sites - undername enumeration, XSS
- Login - Username enumeration, weak passwords (bruteforcable)
- Multistage login - implementation flaws
- Session state - Predictable tokens, insecure handling of tokens
- Access controls - Horizaontal and vertical priv esc
- User impersonation - priv esc
- Use of cleartext communications - Session hijacking, capture of creds and sesitive data
- Off-site links - Leakage of query string parameters in `Referer` header
- Interfaces to external systems - Shortcut the handling of sessions and/or access controls
- Error messages - Information leakage
- E-mail interaction - E-mail and/ or command injection
- Native code components - buffer overflows
- Use of third party application componenets - Known vuln's
- Identifiable web server software - Common configuration weaknesses, known software bugs
