+++
title = "Flask, forms and http requests"
slug = "Flask-forms-and-http-requests"
date = "2018-04-23"
tags = ["python", "flask"]
categories = ["python", "webdev"]
+++

Flask, Forms and HTTP Requests
------------------------------

### Preface

This week I have been building a small webapp that leverage's an
external API to populate values in the user pages. And it pains me to
say that I spent a silly length of time debugging what was a rather easy
problem.

The webapp uses Flask, jinja2 templates and the
[fuelwatcher](https://github.com/danielmichaels/fuelwatcher) API. The
issue started when trying to institute a search bar that allows the user
to select a series of parameters.

In the end the following issues were present:

-   My lack of understanding in regards to HTTP requests
-   Not knowing how the Flask Request context functions
-   Improper formatting of my HTML forms

### Learn You Some HTTP For Great Good

Foolishly, I thought I understood HTTP requests. Sometimes you need to
use your theoretical knowledge practically to fully grasp a concept. It
turns out I was mixing `GET` and `POST` when
submitting data to the server in conjunction with using HTML form
attributes incorrectly.

------------------------------------------------------------------------

 | Get | Post
---|---|---
1.| Sends Parameters in URL | Sends Parameters in body
2.|Used for fetching documents | Used for updating data
3.|Has max length URL limitations | No max length (theoretically)
4.|Often Cached |  Server does not cache
5.|Idempotent | Not idempotent
6.|Should not change server data | Can change server data

------------------------------------------------------------------------

A key point for said webapp is item number 1; I was sending parameters
in the URL when they really needed to be sent via the body. Expanding on
this a little, in a `GET` request the parameters are
transmitted using URI schema friendly syntax (see wiki
[here](https://en.wikipedia.org/wiki/Uniform_Resource_Identifier)). It
is common to see query delimiters in the URL such as `&`,`=`, `+` and `#` and should
of been a clue in the debugging process. Transmitting a
`GET` request like this fetches data from another resource,
where in contrast my data was needed to execute a function within the
server application to pull information from a third party.

### Flask Requests

This is deep subject but is tightly coupled to the previous point. The
request object in flask gives access to the global *request* object.
Meaning it parses the incoming request data for you. This is important
because it is checking for a `method` attribute within that
request. Depending on which `method` is sent will affect
how it/ you should check for the request object.

**TL;DR:** If you incorrectly set the method to `GET` when
its actually a `POST` you will have problems.

To receive the request object and parse for the data you want the
following should be used:

```python 
# displaying both POST and GET for clarity

# POST request // Use 'form'
@app.route("/test" , methods=['GET', 'POST'])
def test():
    select = request.form.get('attribute to be parsed')
    return select # or render_template with select=select etc

# GET request // Use 'args'
@app.route('/data')
def data():
    # here we want to get the value of user (i.e. ?user=some-value)
    user = request.args.get('user')
    return str(user) # or how ever you want to use that data
```

Given these snippets you can see how mixing `GET` with
`POST` will lead to calling the wrong request method.

### HTML form tags for Dummies

HTML, basic right? Don't let hubris fool you, not learning HTML deeply
is foolish. To be fair, this is the first app I have created that has
not required the use of `flask-wtf`, which kindly generates
forms and their tags for you.

```html 
<!-- standard form tags -->
<form action="{{ url_for('app.function')}}" method="post">
```

Simple, right? Not if you:

-   Mix up the HTTP request
-   Use the wrong flask request method
-   Don't use an attribute which you can parse in using
    `request.form.get('attribute')`

Tightly coupled (I think there is a SOLID principle about this...) and
frustrating when you don't connect the dots.

The last piece of the puzzle was the simplest of all, I had not set an
attribute that the request object could get. In this example a
`<select>` tag was used and simply appending
`name=attribute_to_be_parsed` worked.

For clarity, first the html.

```html 
<!-- html -->
<select name='item1'>                   <-- attribute flask looks for
    {% for item in items %}
        <option>
            {{ item }}                  <-- what we get in the response
        </option>
    {% endfor}
</select>
```

And the flask part.

```python 
# python: flask
@app.route('/test')
def test():
    item_to_get = request.form.get('item1')     <-- flask request parsing the response body
    ... snip ...
```

... My face when it worked.

![Face of surprise!](/img/its_alive.jpg)

### Learn from your mistakes

Everyone makes mistakes, learn from it, be humble and don't do it again.
