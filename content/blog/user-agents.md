+++
date = "15 may 2019"
categories = ["programming", "eli5"]
tags = ["eli5"]
slug = "user-agents-101"
title = "User Agents 101"
draft = false
+++

# User agents

## What

Every browser sends with its request, a string that says who it is. That string will contain information on its application type, operating system, and software version currently in use.

Heres mine: `Mozilla/5.0 (X11; Linux x86_64; rv:66.0) Gecko/20100101 Firefox/66.0` but it could just as easily be `Mozilla/5.0 (PlayStation 4 4.71) AppleWebKit/601.2 (KHTML, like Gecko)`.

If you've ever been surfing the net on your mobile and selected "Request Desktop Site" then you have changed your phone browsers user agent. And, if you have done that you might have noticed a significant change in that sites resolution. Its all to do with that string of data, and what the website does with that information.

Not only will each browser have its own user agent but so will any program used to query a web site. Google uses a series of "spiders" which crawl the internet and index it giving us Google search. Googlebot, cURL, wget, Postman and your latest smart light bulb all have user agents.

You can find yours by typing "user agent" into [DuckDuckGo]. Or, you could open your browsers *Development Tools* with `Ctrl + Shift I` and then select the *Network* tab. Once inside look for a subsection named *Headers*. Your user agent will be located in the request headers section.

## Why do I need to know this

Knowing is half the battle. My bank once wouldn't let me login unless I used Chrome or Edge which is ridiculous. I choose my browser not the web application, so I just changed the user agent, and wouldn't you know, *...I got in*.

Maybe you are an app developer, or for some reason need to verify something by assessing the user agent. At least now you know that it isn't a fool-proof way to confirm a device - you will need more than one method to ascertain the true origin. Don't rely on user agents alone as you can see below.

As a developer, or user if you repeatedly request a website it is possible to be served a `HTTP/1.1 429 Too Many Requests`. It's probable that your user agent in conjunction with your IP address has triggered that.

Web scraping is a common programming task, and allows developers to access information from sites that do not offer an application programming interface (API).
Some pages offer API's but a programmer may choose not to use it. Usually this is to circumvent having to pay for it, or maybe the API does not offer a certain, wanted piece of information. More often then not, scraping is in violation of the web applications terms of service. It's best to utilise an API whenever possible.

So how do people get around this?

## It's just a string

Thankfully, HTTP is stateless (yes, even HTTP/2) and thus we can influence each request and its headers.

We can demonstrate this with a simple script which will choose a user agent at random from a list and then send it with its request. Every language has the ability to get information from the internet and will likely already have a way to do the same thing. This example will be in python using the `requests` library.

**Step 1. Set the user agents to be selected randomly.**

```python
import random
import requests

def user_agent():
    user_agents = ['Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 6.1)',
                    'Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko',
                    'Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)',
                    'Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko',
                    'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36',
                    'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36',
                    'Mozilla/5.0 (Windows NT 5.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36',
                    'Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36']
    agent = random.choice(user_agents)

    return agent
```

**Step 2. Call the URL with our randomly selected user agents.**

```python
def check_ua():
    url = 'https://www.whatsmyua.info/api/v1/ua?='
    resp = requests.get(url, headers={'User-Agent': user_agent()})
    ua = resp.json()[0]['ua']['rawUa']
    return ua
```

We override the default `headers` with our `user_agent` function and are returned a response. In this case we've hit an endpoint that will return our user agent as a JSON object.

The source for this script can be found [here]

So changing a user agent is that easy. Firefox has extensions, Chrome would, Opera has it built-in. PC, laptop or mobile handset its all the same after all its just HTTP.

[DuckDuckGo]: https://duckduckgo.com/?q=user+agent
[here]: https://github.com/danielmichaels/utils/blob/master/user_agent_mixer.py
