+++
title = "HTTP Security Basics"
date = "06 Jul 2019"
draft = false
slug = "http-security-basics"
tags = ['HTTP','security','web']
categories = ['HTTP','eli5']
+++

# HTTP Basics

![](/images/httpie.png 'http put image output')

HTTP is an application level message based model where the client sends a request and the server returns a response. HTTP uses a stateful protocol - TCP - but is _stateless_, where each request is connectionless avoiding the need for servers to hold an open connection. 

## HTTP Requests

The first line of every HTTP request has three (3) items, separated by spaces:

- The verb - `GET`,`PUT`,`DELETE`,`POST`,`TRACE`,`CONNECT`,`OPTIONS`, or `HEAD`
- The requested URL, typically the name of the resource plus an optional query string. Query strings are noted by the `?` eg. `/auth/448/yourdetails.ashx?uid=345`
- The HTTP version in use. Most common is 1.1 and will look like this `HTTP/1.1`

```shell
# Example request
 
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

## HTTP Responses

Like the request, each responses first line will return three items:

- The HTTP version in use,
- The status code - `200`, `302` and `404` being most common,
- The response "phrase" that further describes the status of the response - legacy and not required by browsers.

Responses can contain many more items after the first line, such as:

- `Server`: This may contain information about the server such as what engine it is running or modules installed. It may or may not be correct.
- `Set-Cookie`: Will set a cookie for future use, and it will be submitted via the `Cookie` header in future requests.
- `Pragma`: Can be used to instruct the browser not the store the response in its cache. Likewise the `Expires` header says when the is set to cache expire, and if expired it won't load from the cache.
- `Content-Type`: Almost all responses will contain a message body after the headers (separated by a single blank line). The `Content-Type` header indicates what is in the message body.
- `Content-Length`: The total size of the response in bytes.

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

The most common method. When clicking on a hyperlink, you are likely sending a `GET` request. `GET` requests can also be used to send information via a URL query string such as `?`.

```shell
http://example.com/over/there?name=ferret
```

URL's are displayed on screen and also logged in places such as browser history and web access logs. They are also transmitted via the `Referrer` header when clicking on an link that takes you an external resource. In such an example sensitive information in a query string would be logged in an external servers logs. Further, these strings (if unencrypted, or SSLstripped) could be sniffed by a malicious actor. Never use `GET` requests for transmission of sensitive data.

#### POST

`POST` is used to submit data to a resource. A `POST` can submit this information via the URL query string, or via the message body (preferred method, imo). If the `POST` is coupled with URL query strings, any bookmarking would remove the query from the URL, in an attempt to provide security. Using this method also ensures that only one request is sent with that information at a time. That is, if a user submits a `POST` request and then hits the 'back' button, they will likely see a prompt stating something along the lines of 'having to resubmit'. This ensures that the user can only submit the `POST` once per action, as the browser will not automatically reissue the request - unlike a `GET`.

#### HEAD

This functions in the same manner as a `GET` except it only returns the header information.

#### TRACE

Used for diagnostic purposes only. The server should return in the response body the exact contents that it received. Often used to diagnose issues and check if proxies are working correctly/ intercepting data unexpectedly. Commonly disallowed option returning `405 Method Not Allowed`

#### OPTIONS

Returns what HTTP methods are available for a particular resource.

#### PUT

Allows uploading of resources to the server. Used by API's to update existing records or data.  Can provide the attacker with an avenue to upload malicious scripts to the server for execution.

## Uniform Resource Locator (URL)

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

Within the `authority` a port number is delimited by `:`, this is optional and generally only included if not default for the protocol in use.

### REST

Representational State Transfer (ReST) is a style of architecture for distributed systems in which requests and responses contain representations of the current state of the systems resources. 
What is means is that RESTful applications will have their parameters within the path, rather than in the query string.


__Non-RESTful__


`https://example.com/search?manuf=samsung&model=galaxy`


__RESTful__

`https://example.com/search/samsung/galaxy`


See [here](https://en.wikipedia.org/wiki/Representational_state_transfer) for more information.

### HTTP Headers

Some interesting and important headers:

**General Headers**

- `Connection`: Lets the other end know whether to shutdown communication or keep it open for future messages. Defaults to `keep-alive` in HTTP/1.1 but _must not_ be used in HTTP/2 as header metadata is handled differently due to HPACK.
- `Content-Encoding`: Specifies what encoding is used. `gzip` being very common as it compresses the data for faster transmission.
- `Content-Length`: Provides the message body length in bytes.
- `Content-Type`: What type of contents are in the message body such as `text/html`, or `application/json`. Very important during `POST` requests in determining the type of data that will be sent in the body, and how the server should expect to parse it. One type is `multipart/form-data` which will set a `boundary=` string, allowing the server to accept each field as a new _part_ and separate them according to the boundary defined in the `Content-Type`. Another method is the `application/x-www-form-urlencode` type which will place the data from the form into the body of the message, rather than in the query string.
- `Transfer-Encoding`: If any encoding has been done on the message body, typically seen when chunking has been performed.

**Request Headers**

- `Accept`: Tells the server what kind of content the client is willing to accept, such as image types etc.
- `Accept-Encoding`: What encoding types the server will accept.
- `Authorization`: Submission of credentials to the server for one the built-in HTTP authentication types.
- `Cookie`: Submits cookies that the server has previously issued.
- `Host`: Hostname that appeared in the requested URL.
- `If-Modified-Since`: When the browser last received the requested resource. If no changes have happened since then the server may issue a 304 status code telling it to render from its cached copy. Compared against `Last-Modified` and determines the cached variant is older or newer than the current resource.
- `If-None-Match`: A sort of hash that the server issues to the browser which it checks to see if is valid prior to telling the browser to render from cache. This specifies an _entity tag_ and it is the _ETag_ that is validated. It cares only if the information is identical or not and takes precedence or `If-Modified-Since` as the `ETag` is considered stronger validator. See [RFC 7232] for more information.
- `Referer`: From which URL this request has originated.
- `User-Agent`: Information about the browser or application making this request.

**Response Headers**

- `Access-Control-Allow-Origin`: Whether this resource can be retrieved via a cross-domain request.
- `Cache-Control`: Passes caching directives to the browser.
- `ETag`: The entity tag submitted to server. A special header that specifies the version of a resource. Used to verify if the server needs to send a full response, or if the client can render a cached copy. Each time a resource changes, and new `ETag` value must be generated.
- `Expires`: How long the contents of the message body are valid for. Cached copies may be used until this time.
- `Location`: Used in redirect responses (3XX) to specify the target of the redirect.
- `Pragma`: Passes caching directives to the browser.
- `Server`: Information about the server's software.
- `Set-Cookie`: Issues cookies to the browser which it will send back to the server in subsequent requests.
- `WWW-Authenticate`: Provides details on types of authentication that server supports after 401 status code.
- `X-Frame-Options`: Whether responses can be loaded within browsers frame and how to do it.

### Cookies

![](/images/cookies.png 'cookie output request headers')

Many web applications rely on cookies as allow servers to send data to the client, and then have the client return that data in its transmissions within the server. Since HTTP is stateless, cookies allow functionality to persist over time. Login data, shopping carts and site settings all rely on cookies for a better user experience.

```shell
# Example of setting a cookie

Cookie: tracking=dw34tad23590ddfsawWcsTs
```

Cookies are usually made up of a key value pair but may consist of any string that does not contain a space. A server can send more than one `Set-Cookie` per response by using a semi-colon to separate the cookies.

`Set-Cookie` can also include optional attributes such as:

- `expires`: The date at which the cookie will become invalid. This means that cookies are stored by the browser on disk and may be reused repeatedly until that date. If no expires is set, it will only remain valid for the current browser session.
- `domain`: Which domain the cookie is valid for.
- `secure`: Whether that cookie must be sent via HTTPS or not.
- `HttpOnly`: if this is it set, it prevents javascript from accessing it.

### Status Codes

Five (5) groupings:

- `1xx` - Informational.
- `2xx` - Request successful.
- `3xx` - Redirection.
- `4xx` - Request error of some type.
- `5xx`- Server error of some type.

The most common status codes are:

- `100 continue`: Sometimes sent when a server receives a request containing a body. It means that the headers were received and the client should continue sending the message. Once all is received a further response will be sent.
- `200 OK`: Successfully received.
- `201 Created`: The request was successful and a new resource was created. Typically the result of `POST` and `PUT` requests.
- `301 Moved Permanently`: Redirection to the new URL of the requested resource. A `Location` header will specify the new resource location and that it should use that in future.
- `302 Found`: Redirection to a temporary URL, again specified in the `Location` header.
- `304 Not Modified`: Instruction to browser saying it should use its cached copy of the requested resource. `If-Modified-Since` and `If-None-Match` headers are sent to determine if client has most recent copy.
- `400 Bad Request`: Invalid HTTP request was received.
- `401 Unauthorized`: Invalid credentials or none were supplied. A reply with `WWW-Authenticate` headers will be sent stating what authentication types are supported.
- `403 Forbidden`: No one is allowed access.
- `404 Not Found`: Resource does not exist.
- `405 Method Not Allowed`: Method used is not supported by the specified URL. Both `HEAD` and `GET` must never be disabled, and should not return this code.
- `413 Request Entity Too Large`: Some endpoints or browsers will limit the string length - usually seen in buffer overflow attacks.
- `414 Request URI Too Long`: Same as above.
- `500 Internal Server Error`: The server hit some error processing your request. A good hint that you are trying to exploit something it is not configured to handle properly. Always try and identify the nature of this error.
- `503 Service Unavailable`: An indication that the server is responding but the application is not.

### State and Sessions

Although HTTP is stateless there is a requirement for the client and server to process long running requests. An example of this is a sites Shopping Cart. A user can browse and add items to their cart whilst making several requests over a period of time. To make this possible the server maintains a User Session, which allows it to track the users actions across the site.

In some cases, state information is stored on the client side. Instead of the server storing the session, now the data is sent to the client in each server response and the client will send it back in each request. It is important that the server does not trust the client, as it may alter the session information for nefarious reasons. One mitigation is using a hash of the state, which the server validates before accepting the clients session data.

Given the stateless nature of HTTP, applications need a method in which to re-identify clients. Typically, this is achieved through the use of tokens - a unique identifier for each user such as cookie.

### Encoding Schemes

HTTP has several encoding schemes to ensure the safe delivery of data. The following is some of the most common.

#### URL Encoding

URL's can only contain characters from the US-ASCII set; 0x20 through to 0x7e. Several characters within this set are restricted as they have special meaning in the scheme or HTTP.
All URL encoded characters are prefixed by `%` followed by their hexadecimal representation of the ASCII character.

Examples:

- `%3d` &mdash; =
- `%25` &mdash; %
- `%0a` &mdash; New Line (\n)
- `%00` &mdash; Null Byte

**note**: `+` represents the 'space' in a URL.

#### Unicode Encoding

Unicode is the system that allows us to use many different character sets. For English speakers, `utf-8` will be the most common, but many more exists.
Encoding 16-bit unicode in a URL requires the prefix `%u` followed by the hexadecimal.

Example unicode

- `%u2215` &mdash; /
- `u00e9` &mdash; é

UTF-8 is a variable length encoding that uses one or more bytes for each character. To send this using URL encoding, each byte is delimited by a `%`.

For instance:

- `%c2%a9` &mdash; ©
- `%e2%89%a0` &mdash; ≠

Unicode is an important part of attacking web applications as it can defeat input validation schemes.

#### HTML Encoding

Safe incorporation of HTML within web applications is important. Several HTML characters have special meaning and must be encoded correctly.
When attacking an application, HTML encoding will be most evident when probing for XSS vulnerabilities. If an application returns user input unmodified, it is likely vulnerable.

#### Base64

Base64 allows any binary data to be represented as using only ASCII characters. It permits the sending of ASCII strings over the wire for safe reassembly on the other side in the original format. It is also used heavily in basic HTTP user authentication.

Base64 splits the bytes into 6bit streams, allowing for 64 possible permutations - each chunk of 6 bits allows for 64 possible characters. It allows only the following characters:

```ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789/+```

If the final block results in fewer than three chunks (16 bits or 2 bytes) then it will be marked by `=` or `==` which says that their may be one or two trailing blocks of zero's.

Base64 is prevalent across the web and is often used to transmit binary data within cookies and other parameters. It is often used to obfuscate data in transit through security by obscurity. Always decode any intercepted Base64 data, it could be a goldmine. Base64 can often be identified quickly by the `=` trail, or if it is JSON it will start with `ey` which represents `{`.

## Fin

HTTP is an important protocol not just for security and network engineers but for developers too. Having a cursory understanding aides every one who uses the web. Next time you look at the network tab in your browsers developer tools be sure to look at the headers, you might notice something worth exploring. 

[RFC 7232]: https://www.rfc-editor.org/rfc/rfc7232.txt
