+++
title = "Web Scraping Javascript with Python"
categories = ["python", "web"]
tags = ["python", "web"]
slug = "web-scraping-javascript-with-python-xhr"
date = "2020-03-13"
draft = "false"
+++

# Scraping Dynamic Pages with Python

![](/images/scraping.png 'web scraping diagram')

## Web scraping

Python is a great tool for web scraping tasks, it is efficient, easy to read and fast. Whenever looking to grab data from a site, the canonical packages are [BeautifulSoup](https://www.crummy.com/software/BeautifulSoup/bs4/doc/) and [Requests](https://github.com/psf/requests/). Unfortunately, when our target site is dynamically rendered, BeautifulSoup can't "see" those parts leading to a lot of head scratching. The often touted answer to this is [Selenium](https://selenium-python.readthedocs.io/) which spins up a browser thereby rendering the JavaScript making it possible to scrape data from it.

While selenium does work it kinda sucks, can be a pain to setup and introduces more complexity. On big projects where scrolling, pagination or link traversal is required you may be best served by using Selenium. Though I would recommend [Scrapy](https://scrapy.org/), which is excellent and is built for scraping - Selenium is not.

In some cases, there is another way.

## An alternative

Many modern web applications render data from third party API's or other backend services. When a site uses JavaScript to do this, then we can get a lot of data from it by using only Requests, [Postman](https://www.postman.com/) and [cURL](https://curl.haxx.se/).

The basic steps:

- find the API endpoint using the browser's development tools,
- copy the URL as a curl command,
- import that command into postman, checking it works,
- get postman to auto generate the request into Python code, and
- plug that into your script, and profit!

This will provide you repeatable python code which will always return a response with the data you require.

### Step 1: Find the URL

Let's use an example [website](https://lic-investing.online) which uses an [XMLHttpRequest](https://en.wikipedia.org/wiki/XMLHttpRequest) pull data from another server and populate a table with stock data.

After navigating to the site, open the development tools and click on the network pane. You may need to refresh the page to populate the network pane with its requests.

Searching through the requests, we find what we are looking for, in this case the requests are named after the stock symbols. There is no standard but you can filter by XHR in the network pane to make finding juicy targets easier.

For the uninitiated, when clicking on the XHR request, select the response pane and look at the JSON data it has returned.

![](/images/scrape-net-tab.png 'dev tools network tab with json xhr response data')

### Step 2: Copy as cURL

After finding the appropriate XHR endpoint, we now need to replicate the GET request.

In the network tab, as seen above, right click and hover over *Copy* which will show you an option to *Copy as cURL*. Selecting this will output something like the following:

```sh
curl 'https://lic-investing.online/api/stocks/bki' \
-H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:74.0) Gecko/20100101 Firefox/74.0' \
-H 'Accept: application/json, text/plain, */*' \
-H 'Accept-Language: en-US,en;q=0.5' --compressed \
-H 'DNT: 1' -H 'Connection: keep-alive' \
-H 'Referer: https://lic-investing.online/' \
-H 'Pragma: no-cache' \
-H 'Cache-Control: no-cache' \
-H 'TE: Trailers'
# formatted for easier reading
```

You can test that this works by running the command in your console on MacOS or Linux. On Windows? 🤷 soz, I don't know.

### Step 3: Import into Postman

After copying as cURL above, open up [Postman](https://www.postman.com/).

In the top left will be a orange box titled *New*, and to the right of it will be *Import*. Select *Import* and then click on the *Paste Raw Text* tab. Paste in the cURL command and hit *Import*. This process should resemble the image below.

![](/images/scrape-postman.png 'Postman import raw text example')

After importing the command, Postman will populate all the GET parameters needed to make a request in its main window. Clicking on the blue *Send* button will fire the request and once Postman receives the response you will see it in the main body of the screen.

Right now we have returned a JSON object from our target without needing to be on their website proving we can simulate a request from their frontend to the backend. This is how we will scrape the site in a repeatable and reliable way.

### Step 4: Get the python code

So far, we've created a request to retrieve data from the server. But now we need to turn this request into something we can replicate using Python.

Let's use Postman to automatically generate some Python code for us.

To get this auto generated code simply select the *Code* text block which is directly below the *Save* drop-down on the right hand side of Postman's main screen.

Scroll down to find Python code nicely generated for us. In this example I have chosen Requests, though it also offers `http.client` as well.

![](/images/scrape-python.png 'Postman code auto generator example')

### Step 5: Win

Now we have fully functional Python code which can make a request to the endpoint we are targeting and it will return a response as if we were the website. What we do now is pure business logic but the main thing is we did not need to reach for Selenium, and that is worth its weight in code.

## Closing points

Be wary of hitting the endpoint repeatedly, you may get black listed by the site. To prevent this, use a library such as [Requests-cache](https://github.com/reclosedev/requests-cache). Give them a star if you can!

This is also helpful when playing with API's which have request or rate limits imposed.

Not related to XHR but always look over the source of a webpage in a browser, and then contrast that with BeautifulSoup's response data. I have had success on React sites by grabbing their Props data that I couldn't even see rendered in the browser. To get at custom, hard to scrape data like this, Regex is your friend.

If its on the internet, its possible to scrape (within legal limits of course).
