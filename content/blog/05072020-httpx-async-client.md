
+++
title = "HTTPX is the new Requests with Async"
categories = ["http", "python"]
tags = ["python","programming"]
slug = "httpx-is-the-new-requests-with-async"
date = "2020-06-25"
draft = "false"
+++

# Long live requests welcome HTTPX


> HTTPX is a fully featured HTTP client for Python 3, which provides sync and async APIs, and support for both HTTP/1.1 and HTTP/2. 


Far and away, requests stands as *the* most notable third party package. It's almost considered part of what makes up python these days. 
Unfortunately, little forward progress has been made in recent times to incorporate python's async capabilities.

This is where HTTPX is stepping up to the plate. And it aims to be compatible with *most* of [requests API][0]. 

## Why use HTTPX over Requests?

In addtion to async, HTTPX also supports:

- HTTP/2
- Timeouts by default
- Full type annotation (yay, code editors)
- Direct requests to [WSGI][1] and [ASGI][2] applications
- Plus, all the standard features of requests - list [here][3]



## Simple example - sync

As HTTPX supports the core requests API, to do a simple get of a webpage we can use `httpx.get()`

```python
>>> import httpx

>>> url = 'http://httpbin.org/get'
>>> resp = httpx.get(url)
>>> resp.json()
# json output cleaned up for ease of reading.
{'args': {},
 'headers': {'Accept': '*/*',
             'Accept-Encoding': 'gzip, deflate',
             'Host': 'httpbin.org',
             'User-Agent': 'python-httpx/0.13.3',
             'X-Amzn-Trace-Id': 'Root=1-5f014d41-5cef343010793f4686d7c258'},
 'origin': '19.13.3.36',
 'url': 'http://httpbin.org/get'}
}
```

## Client is the new Session

When using requests, if you needed to do anything more than receieve simple data then the Session instance was necessary. 

HTTPX does not have a Session API instead it has `Client` and `AsyncClient` for obvious use cases.

From the HTTPX docs:

> If you do anything more than experimentation, one-off scripts, or prototypes, then you should use a Client instance.

For more reasons why you should use `Client` read the [docs][4]. But the core reasons are:

- Reduced latency across requests (no handshaking).
- Reduced CPU usage and round-trips.
- Reduced network congestion.
- Without `Client` you don't get HTTP/2 support

To use a the `Client` interface, its recommended to use a context manager.

```python
>>> with httpx.Client() as client:
    ...
```

To make a request with the `Client` its as simple as:

```python
>>> with httpx.Client() as client:
    resp = httpx.get('https://httpbin.org/get')

>>> resp
<Response [200 OK]>
```

One of the really cool features is `base_url` which allows for URL prepending.

```python
# base_url example
>>> with httpx.Client(base_url='http://httpbin.org') as client:
    resp = client.get('/headers')

>>> resp.request.headers
Headers({'host': 'httpbin.org', 'user-agent': 'python-httpx/0.13.3', 'accept': '*/*', 'accept-encoding': 'gzip, deflate', 'connection': 'keep-alive'})
```

## AsyncClient simple

To make an async call using the `Client` interface is as simple as adding the `async` keyword to the context manager and swapping `Client` with `AsyncClient`.

```python
# AsyncClient example
>>> async with httpx.AsyncClient() as client:
     resp = await client.get('http://httpbin.org/get')
>>> resp
<Response [200 OK]>
```

There is some API differences when using the `AsyncClient` that are worth being [aware][5] of.

## AsyncClient real example

I am not a python async expert so I did have some issues getting it to function as expected. 
When I needed to concurrently get several json endpoints worth of data a lot of experiementation was needed on my part. 

[Florimond Manca][6]s [post][7] set out the ground work for solving my problem; I needed a way to call a list of url's looping over the endpoint from a list id's. 

```python
# Florimond's example
>>> import httpx
>>> import asyncio
>>> # We're going to fetch tag pages concurrently...
>>> async def fetch(tag, client):
        return await client.get(f'https://dev.to/t/{tag}')

>>> async with httpx.AsyncClient() as client:
        responses = await asyncio.gather(
            fetch('hacktoberfest', client),
            fetch('python', client),
            fetch('opensource' client),
        )

>>> responses
[<Response [200 OK]>, <Response [200 OK]>, <Response [200 OK]>]
>>> urls = [r.url for r in responses]
>>> urls
[URL('https://dev.to/t/hacktoberfest'),
 URL('https://dev.to/t/python'),
 URL('https://dev.to/t/opensource')]
```

This worked perfectly in an Ipython terminal but I needed to integrate into a sync class (swapping out the slow synchronous api call's for a faster async version).

### My version

I needed three things:

1. A method to call each endpoint,
2. AsyncClient for executing the requests, and
3. Something to kick off these actions.

The following is a snippet which meets all three critieria. It takes Florimond's demonstration example and shows how it could be implemented into a real piece of callable code.

```python
# 1
async def fetch(item_id, client):
    """
    Return async response object for HackerNews item.
    :param item_id: json id for HackerNews item
    :param client: httpx.AsyncClient method
    :return: httpx.AsyncClient response object
    """
    return await client.get(
        f"https://hacker-news.firebaseio.com/v0/item/{item_id}.json"
    )
# 2
async def story_metadata():
    """
    Return a list of responses from the given list of HackerNews item id's.
    :return: list of HackerNews response objects
    """
    _list = [23709004, 23717964, 23723433, 23725506, 23713605]
    async with httpx.AsyncClient() as client:
        responses = await asyncio.gather(*[fetch(item_id, client) for item_id in _list])
        return responses
# 3
def run_async_task(coro):
    """
    A helper method to create async coroutines.
    :return: result set from the async function
    """
    results = asyncio.run(coro)
    return results
```

Adding async here led to a speed increase of 50% over a relatively small selection of items - it only grabs the top ten articles. The return time went from 3+ seconds to ~1.1-1.5, which is consistent with network latency rather than python synchronicity.

## Wrap up

Over the last month or so, whenever I have needed a requests like interface I have elected to use [HTTPX][8] instead.
In this time there has been only one thing missing; a caching adapter. Previously, I would lean heavily upon [requests-cache][9].

Apart from this, the experience has been seemless in fact I find HTTPX easier to use. 
It has a very active [community][10], great [leadership][11] and highly active development.
Being a relatively new package it is very possible to get involved and contribute as well.

### Gist

A full copy of the working code is below as well as [here][12]

{{< gist danielmichaels b2ffb53736a6a9157b12ea4a1388673b >}}

[0]: https://www.python-httpx.org/compatibility/
[1]: https://www.python-httpx.org/advanced/#calling-into-python-web-apps
[2]: https://www.python-httpx.org/async/#calling-into-python-web-apps
[3]: https://www.python-httpx.org/#features
[4]: https://www.python-httpx.org/advanced/#why-use-a-client
[5]: https://www.python-httpx.org/async/#api-differences
[6]: https://github.com/florimondmanca
[7]: https://dev.to/florimondmanca/httpx-for-hacktoberfest-help-build-the-future-of-python-http-16pj
[8]: https://www.python-httpx.org/
[9]: https://github.com/reclosedev/requests-cache
[10]: https://gitter.im/encode/community
[11]: https://github.com/encode
[12]: https://gist.github.com/danielmichaels/b2ffb53736a6a9157b12ea4a1388673b