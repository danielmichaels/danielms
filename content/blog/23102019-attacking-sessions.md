+++
title = "Attacks against session management"
categories = ["web"]
tags = ["web", "security"]
slug = "session-management-attacks"
date = "2019-10-23"
draft = "true"
+++

# Attacking Session Management

Session management is fundamental to web applications, it provides a method to uniquely identify users across a number of requests and stores data about the interactions between the user and client. When a user logins in, the session is what gives the application persistent assurance of the users identity beyond the initial request sent with login credentials.

Because of this, session management is prime target for attacks against the application. Accessing another users session token will bypass any authentication effectively allowing unfettered access to that users account.

## The Need for State

HTTP is stateless, based upon a request-response model where each pair of messages is its own independent transaction. The protocol contains no mechanism for linking requests, nor distinguishing requests between other users.
Obviously, today we need sessions to ensure consistency and security - we no longer only have just static sites.
Almost all websites have sessions, even simple sites that do not require logging in are likely to have sessions.

Sessions work like so; the server issues a token, and on each request the client sends that token in its headers thereby allowing the server to track that session, and link its travels across the site.
Whilst cookies are the de facto standard, it is possible to see bespoke session management mechanisms. Despite this all session management schema's can be grouped into two categories:

1. Weakness in the generation of session tokens,
2. Weka ness in the handling of session tokens throughout their life cycle.

### Hack Steps

1. Applications can combine several different items of data collectively as a token, including cookies, URL params, and hidden form fields. Do not assume a particular parameter is the session token without proving it!
2. Sometimes, items that appear to be the application's session token may not be. For instance the standard session cookie may be generated but not actually used by the application as the session token.
3. Observe which new items are sent to the browser after authentication; new session tokens are often created after a successful login.
4. To identify which items are actually being employed as session tokens, navigate to a page that definitely uses the token (such as, "my details" pages). Make several requests for it, systematically removing each item that you suspect being the token. If removing that item causes the session-dependant page not to be returned, this *MAY* confirm that it is a session token. Burp Repeater is a good tool for this.

## Alternative to Sessions

Not all applications use sessions. Some alternatives are:

- **HTTP Authentication**: This form of session management acts much like a form field. It resends the credentials for each request. It is not as common as the better alternatives.
- **Session-less state mechanisms**: Session-less state management is a client based alternative where all session data is sent via a cookie, or hidden form field. It must be secure within context to prevent an attacker sniffing it and resubmitting it somewhere else in the application to circumvent the security. As such, it is often blobbed, encrypted and signed to prevent this. Expiration timers can also be implemented.

### Hack Steps

1. If HTTP authentication is being used, it may indicate that no session management mechanism is in use.
2. If session less state is in use the following are general indicators of it in use:
  - Token-like data items issued to the client are long (>100 bytes)
  - New token like items in response to every request
  - The data is encrypted or signed
  - Attempts to submit the same item in more than one request are denied
3. Should the evidence suggest that session management tokens are not in use then chose another avenue of attack as this will likely return nothing.


## Weakness in Token Generation

Tokens can be generated in an unsafe manner that allows the attacker to identify the values of tokens that have been issued to other users.

**NOTE: There are numerous locations where an applications security depends on the unpredictability of the tokens it generates, for example:**
- Password Recovery tokens sent to users registered email address
- Tokens within hidden form fields to prevent CSRF
- Tokens used to give one-time access to protected resources
- Persistent tokens used in "remember me"
- Tokens in shopping like application where a user does not have to login to have a "cart" that remembers their selected items

## Meaningful Tokens

Some tokens are created through the transformation of some identifying information about the user (username/ email etc). This information is the encoded or obfuscated in some way and combined with other data.

Example of a seemingly random string of characters that makes up a token

`dXNlcj10aGluZ3MxO2FwcD1hZG1pbjtlbWFpbD1hZG1pbkBhZG1pbi5jb20=`

On inspection, that data is possibly Base64, and when decoded

`user=things1;app=admin;email=admin@admin.com`

it reveals quite a bit of information for the attacker. In such an example, if an attacker can formulate possible tokens of users, they can construct a list of possible tokens from information they have already uncovered within the application.

Tokens which have meaningful data often have structure. Some common pieces of data used within tokens:
- account username
- numeric ID used to distinguish between accounts
- users first and last names
- users e-mail
- users role or group
- date time stamp
- incrementing or predictable number
- clients IP address

Each component or the entire token may be encoded in different ways. This can be deliberate, or just to conform to HTTP encoding/ transport standards. Common types include; XOR, Base64 and ASCII Hex.

**NOTE: not all parts of the token are used by the application. For instance, within the token the application might only use the "user" and "date" fields. All other fields could just be padding, particularly in the use of binary blobs. Narrowing down the token to its required bits reduces the entropy and complexity.**

### Hack Steps

1. Obtain a single token from the application, and modify it systematically in order to determine whether all or just parts are validated by the server. Try altering one byte (or bit) at a time to see what parts are noise. When confirmed omit the noise from future requests. Burp has the "char frobber" to assist in this
2. Login as several different users a different times recording tokens received from the server. If self-registration is available and you can choose your username, create a series of accounts with similar usernames but with minor variations between them; A, AA, AAA, AAAA, AAAB, AAAC etc.
3. Analyse the tokens and look for similarities, or that showcase user-controllable data is used during token generation
4. Check the token for any detectable encoding or obfuscation techniques. Where a username contains sequentially characters look for corresponding character sequences in the encoding - this could be XOR. Look for evidence of Hex. And, anything that ends in '=' or '==' which may be Base64.
5. If you believe you have enough information to guess the token, use Burp to fuzz the requests and look for requests where the page loaded correctly indicating a valid response.


## Predictable Tokens
