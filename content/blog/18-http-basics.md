+++
title = "HTTP basics"
date = "2019 apr 06"
draft = false
slug = "http-security-basics"
tags = ['HTTP','security','web']
categories = ['HTTP','eli5']
+++

# WHH - Chapter 3 - Web App Technologies

this will be primer of many of the underlying technologies that power the web and associated technologies.

## The HTTP Protocol

**BLUF**: Message based model where the client sends a request and the server retruns a response. HTTP uses a stateful protocl - TCP - but is stateless, where each request is autonomous and may use a different TCP connection.

### HTTP Requests

The first line of every HTTP request has three (3) items, separated by spaces:

- The verb (GET/PUT/DELETE/POST/TRACE/CONNECT/OPTIONS/HEAD)
- The Requested URL typically the name of the resource plus an optional query string. Query strings are noted by the `?` eg. `/auth/448/yourdetails.ashx?uid=345`
- The HTTP version in use. Most commonn is 1.1 and will look like this `HTTP/1.1`

an example request:
```shell
GET /Protocols/rfc2616/rfc2616-sec5.html HTTP/1.1
Host: www.w3.org
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br
Referer: https://duckduckgo.com/
DNT: 1
Connection: keep-alive
Upgrade-Insecure-Requests: 1
```

### HTTP Responses

Like requests, all responses should return three items:

- The HTTP version in use
- The status code - 200, 302 and 404 being most common
- The response "phrase" that further describes the status of the response - legacy and not required by browsers

Responses can contain many more items after the first line, such as:

- `Server`: this may contain information about the server such as what engine it is running or modules installed. It may or may not be correct
- `Set-Cookie`: will set a cookie for future use, and it will be submitted via the `Cookie` header in future requests
- `Pragma`: can be used to instruct the browser not the store the response in its cache. Likewise the `Expires` header says when the is set to cache expire, and if expired it won't load from the cache.
- `Content-Type`: Almost all responses will contain a message body after the headers (separated by a single blank line). The `Content-Type` header indicates what is in the message body.

```shell
# Example response

HTTP/1.1 200 OK
Date: Mon, 18 Jun 2018 04:29:32 GMT
Server: Apache/2.2.22 (Debian)
Content-Location: rfc7230.html
Vary: negotiate,Accept-Encoding
TCN: choice
Last-Modified: Sun, 17 Jun 2018 07:19:43 GMT
ETag: "3c9fe1-417e4-56ed145dec1c0;56ee301a89430"
Accept-Ranges: bytes
Cache-Control: max-age=604800
Expires: Mon, 25 Jun 2018 04:29:32 GMT
Content-Encoding: gzip
Strict-Transport-Security: max-age=3600
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block
X-Content-Type-Options: nosniff
X-Clacks-Overhead: GNU Terry Pratchett
Keep-Alive: timeout=5, max=100
Connection: Keep-Alive
Transfer-Encoding: chunked
Content-Type: text/html; charset=UTF-8
```

### HTTP Methods

#### GET

The most common method. Most often when clicking on a hyperlink, you are sending a `GET` request. `GET` requests can also be used to send information via a URL query string such as `?`.

```shell
http://example.com/over/there?name=ferret
```

URL's are displayed on screen and also logged in places such as browser history and web access logs. They are also transmitted via the `Referrer` header when clicking on an link that takes you an external resource. In such an example sensitive information in a query string would be logged in an external servers logs. Further, these strings (if unencrypted, or SSLstripped) could be sniffed by a malicous actor. Never use `GET` requests for transmission of sensitive data.

#### POST

`POST` is used to submit data to a resource. A `POST` can submit this information via the URL query string, or via the message body (preferred method, imo). If the `POST` is coupled with URL query strings, any bookmarking would remove the query from the URL, in an attempt to provide security. Using this method also ensures that only one request is sent with that information at a time. That is, if a user submits a `POST` request and then hits the 'back' button, they will likely see a prompt stating something along the lines of 'having to resubmit'. This ensures that the user can only submit the `POST` once per action, as the browser will not automatically reissue the request - unlike a `GET`.

#### HEAD

This functions in the same manner as a `GET` except it only returns the header information.

#### TRACE

Used for diagnostic purposes only. The server should return in the response body the exact contents that it recieved. Often used to dianose issues and check if proxies are working correctly/ intercepting data unexpectedly.

#### OPTIONS

Returns what HTTP methods are available for a particular resource.

#### PUT

Allows uploading of resources to the server. Often cited or used when a resource exists, and is being updated. Can provide the attacker with an avenue to upload malicious scripts to the server for execution.

### Uniform Resource Locator (URL)

URL's are unique identifiers for web resources.

```shell
# Example URI specification from RFC 3986

         foo://example.com:8042/over/there?name=ferret#nose
         \_/   \______________/\_________/ \_________/ \__/
          |           |            |            |        |
       scheme     authority       path        query   fragment
          |   _____________________|__
         / \ /                        \
         urn:example:animal:ferret:nose
```

Within the `authority` a port number is demlimited by `:`, this is optional and generally only included if not default for the protocol in use.

### REST

Representational State Transfer (ReST) is a style of archetecture for distributed systems in which requests and reponses contain representations of the current state of the systems resources. See [here](https://en.wikipedia.org/wiki/Representational_state_transfer) for more information.

### HTTP Headers

Some interesting and important headers:

**General Headers**

- `Connection`: lets the other end know whether to shutdown communication or keep it open for future messages
- `Content-Encoding`: specifies what encoding is used. `gzip` being very common as it compresses the data for faster transmission
- `Content-Length`: provides the message body length in bytes
- `Content-Type`: what type of contents are in the message body such as `text/html`, or `application/json`
- `Transfer-Encoding`: if any encoding has been done on the message body, typically seen when chunking has been performed

**Request Headers**

- `Accept`: tells the server what kind of content the client is willing to accept, such as image types etc
- `Accept-Encoding`: what encoding types the server will accept
- `Authorization`: submission of credentials to the server for one the built-in HTTP authentication types
- `Cookie`: submits cookies that the server has previously issued
- `Host`: hostname that appeared in the requested URL
- `If-Modified-Since`: when the browser last recieved the requested resource. If not changes since then, the server may issue a 304 status code telling it to render from its cached copy
- `If-None-Match`: a sort of hash that the server issued to the browser which it will check to see if is valid prior to telling the browser to render from cache
- `Referer`: from which URL this request has orginated
- `User-Agent`: information about the browser or application making this request

**Response Headers**

- `Access-Control-Allow-Origin`: whether this resource can be retireved via a cross-domain Ajax request
- `Cache-Control`: passes caching directives to the browser
- `ETag`: the entity tag submitted to server
- `Expires`: how long the contents of the message body are valid for. Cached copies may be used until this time
- `Location`: used in redirect responses (3XX) to specify the target of the redirect
- `Pragma`: passes cahcing directives to the browser
- `Server`: information about the server's software
- `Set-Cookie`: issues cookies to the browser which it will send back to the server in subsequent requests
- `WWW-Authenticate`: provides details on types of authentication that server supports after 401 status code
- `X-Frame-Options`: whether and how responses can be loaded within browsers frame

### Cookies

Many web applications rely on cookies as they can allows servers to send data to the client, and then have the client return that data in its transmissions within the server. This sort of bridges the stateless gap that exists between functionality and HTTP. And, rightly so they can be a huge vector for an attacker.

```shell
# Example of setting a cookie

Cookie: tracking=dw34tad23590ddfsawWcsTs
```

Cookies are usually made up of a key value pair but may consist of any string that does not contain a space. A server can send more than one `Set-Cookie` per response by using a semi-colon to separate the cookies.

`Set-Cookie` can also include optional attributes such as:

- `expires`: a date until which the cookie becomes invalid. This means that cookies are stored by the browser on disk and may be reused ad nuseum until that date. If no expires is set, it will only remian valid for the current browser session.
- `domain`: which domain the cookie is valid for
- `secure`: whether that cookie must be sent via HTTPS or not
- `HttpOnly`: if this is it set, it prevents javascript from accessing it

### Status Codes

Five (5) groupings:

- 1xx - informational
- 2xx - request successful
- 3xx - redirection
- 4xx - request error of some type
- 5xx - server error of some type

The main ones of importance are:

- `100 continue`: sometimes sent when a server recieves a request containing a body. It means that the headers were recieved and the client should continue sending the message. Once all is recieved a further response will be sent
- `200 OK`: succesfully recieved
- `201 Created`: usually seen after a `PUT` stating that the request was a success
- `301 Moved Permanently`: redirection to the new URL of the requested resource. A `Location` header will specify the new resource location and that it should use that in future
- `302 Found`: redirection to a temporary URL, again specified in the `Location` header.
- `304 Not Modified`: instruction to browser saying it should use its cached copy of the requested resource. `If-Modified-Since` and `If-None-Match` headers are sent to determine if client has most recenet copy
- `400 Bad Request`: invalid HTTP request was recieved
- `401 Unauthorized`: invalid credentials or none were supplied. A reply with `WWW-Authenticate` headers will be sent stating what authentication types are supported
- `403 Forbidden`: no one is allowed access
- `404 Not Found`: resource does not exist
- `405 Method Not Allowed`: method used is not supported by the specified URL
- `413 Request Entity Too Large`: some endpoints or browsers will limit the string length - usually seen in buffer overflow attacks
- `414 Request URI Too Long`: same as above
- `500 Internal Server Error`: the server hit some error processing your request. A good hint that you are trying to exploit something it is not configured to handle properly. Always try and identify the nature of this error.
- `503 Service Unavailable`: an indication that the server is responding but the application is not

### State and Sessions

Although HTTP is stateless there is a requirement for the client and server to process long running requests. An example of this is a sites Shopping Cart. A user can browser and add items to their cart whilst making several requests over a period of time. To make this possible the server maintains a User Session, which allows it to track the users actions across the site.
In some cases, state infromation is stored on the client side. Instead of the server storing the session, now the data is sent to the client in each server response and the client will send it back in each request. It is important that the server does not trust the client, as it may alter the session information for nefarious reasons. One mitigation is using a hash of the state, which the server validates before accepting the clients session data.
Given the stateless nature of HTTP applications need a method in which to reidentify clients. Typically, this is achieved through the use of tokens - a unique identifier for each user such as cookie.

### Encoding Schemes

HTTP has several encoding schemes to ensure the safe delivery of data. The following is some of the most common:

#### URL Encoding

URL's can only contain characters from the US-ASCII set; 0x20 through to 0x7e. Several characters within this set are restricted as they have special meaning in the scheme or HTTP.
All URL encoded characters are prefixed by `%` followed by their hexadecimal representation of the ASCII character.

Examples:

- `%3d` - =
- `%25` - %
- `%0a` - New Line (\n)
- `%00`- Null Byte

**note**: `+` represents the 'space' in a URL.

#### Unicode Encoding

Unicode is the system that allows us to use many differnet character sets. For english speakers, `utf-8` will be the most common, but many more exists.
Encoding 16-bit unicode in a URL requires the prefix `%u` followed by the hexadecimal.

example unicode

- `%u2215` - /
- `u00e9` - é

UTF-8 is a variable length encoding that uses one or more bytes for each character. To send this using URL encoding, each byte is delimited by a `%`.

for instance:

- `%c2%a9` - ©
- `%e2%89%a0` - ≠

Unicode is an important part of attacking web applications as it can defeat input validation schemes.

#### HTML Encoding

Safe incorporation of HTML within web applications is important. Several HTML characters have special meaning and must be encodeded correctly.
When attacking an application, HTML encoding will be most evident when probing for XSS vulnerabilities. If an application returns user input unmodified, it is likely vulnerable.

#### Base64

Its origins are in MIME or email; it allows the sending of ASCII strings over the wire for safe reassembly on the otherside in the original format. It is also used heavily in basic HTTP user authentication.
Base64 splits the bytes into 6bit streams, allowing for 64 possible permutations - each chunk of 6 bits allows for 64 possible characters. It allows only the following characters:

`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789/+`

If the final block results in fewer than three chunks (16 bits or 2 bytes) then it will be marked by `=` or `==` which says that their may be one or two trailling blocks of zero's.

Base64 is prevailent across the web and is often used to transmit binary data within cookies and other parameters. It is often used to obsfucate data in transit - security by obsecurity at its best. Always decode any intercepted Base64 data, it could be a goldmine. Base64 can often be identified quickly by the `=` trail, or if it is JSON it will start with `ey` whiich represents `{'`.

## Questions CHAPTER THREE

1. What is `OPTIONS` method used for?
- It will return all the HTTP methods the server is capable of.
2. What are the `If-Modified-Since` and `If-None-Match` headers used for? Why are we interested in these?
- These are used for caching purposes, they are sent on the request. If-Modified-Since gives the time it was last checked/ site was accessed and cached. If-None-Match is a hash of the cache that the server may have issued.
- deleting these headers from the request we ensure that each time we recieve a fresh copy of the page and not an older cached copy.
3. What is the significance of the `secure` flag when a server sets a cookie?
- It will only send the cookie via HTTPS
4. what is the difference between 301 and 302 status codes?
- 301: moved permenently
- 302: moved temporarily
5. How does the browser interoperate with a web proxy when SSL is being used?
- The proxy serves as a TCP low-level connection/ forwarding agent. It cannot inspect the HTTPS traffic but will create the connection at the socket level for transport via the proxy whilst maintianing security.
