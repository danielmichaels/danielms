
+++
title = "Flask Jsonify, do I need it anymore?"
categories = ["python", "flask"]
tags = ["flask", "python"]
slug = "flask-jsonify-you-dont-need-it"
date = "2020-05-14"
draft = "false"
+++

# Just use a plain dict

For a long time, returning JSON in flask required using the `flask.jsonify` API.
However, since this [PR][0] Flask will by default call `jsonify` [under the hood][1] on any
dictionary it receives on a `make_response` call. 

## How does it work

When you need to return json in a response, you can simply use a plain python dictionary instead of the `jsonify` API.

```python
# plain dictionary
@app.route('/plaindict')                                                              
    def testplain():                                                                          
        test = 'testing'                                                                 
        name = 'roger'                                                                   
        age =  30                                                                        
        _list = [1,2,3]                                                                    
        return {                                                                         
                "development": test,                                                     
                "name": name,                                                            
                "age": age,                                                              
                "list": _list                                                              
                }       
```

The above code will return the following:
```shell
# response
HTTP/1.0 200 OK
Content-Length: 105
Content-Type: application/json
Date: Thu, 14 May 2020 21:20:06 GMT
Server: Werkzeug/1.0.1 Python/3.8.2

{
    "age": 30,
    "development": "testing",
    "list": [
        1,
        2,
        3
    ],
    "name": "roger"
}

```

And the same code but with `jsonify` (for comparison) will return the exact same thing.

```python
@app.route('/jsonify')
   def testjsonify():                                                                   
        test = 'testing'                                                                 
        name = 'roger'                                                                   
        age =  30                                                                        
        lst = [1,2,3]                                                                    
        return jsonify({                                                                 
                 "development": test,                                                     
                 "name": name,                                                            
                 "age": age,                                                              
                 "list": lst                                                              
                 })  
```

## How does it work

As stated above, Flask now does an `isinstance` check during the `make_response` API call. In the conditional, Flask
checks if the body is of type `dict` and if so calls `jsonify`. To see this in action refer to the [previously][1] mentioned line in the core Flask `app.py`. 

So, while we no longer *need* to explicitly call `jsonify` it is still very much being used by the application itself.


## Whats the catch?

There is one small gotcha, as seen from [pgjones][2] pull request.

> This doesn't support returning anything other than an associate array at the top level in the JSON response. I'm ok with this as in practice APIs are only extensible if the top level is an associate array.

An example of when you will still need to user `jsonify` is when the response is not of type `dict`. As an example, doing a list comprehension and returning that directly will be required to be wrapped in `jsonify`.

```python
@app.route("/users")
    def users_api():
        users = get_all_users()
        return jsonify([user.to_json() for user in users])
```

## Fin

While its not a *complete* replacement for `flask.jsonify`, it will probably reduce its explicit usage by a huge proportion.
Personally, I removed `jsonify` from most of my applications and now just drop in plain dictionaries.


[0]: https://github.com/pallets/flask/pull/3111
[1]: https://github.com/pallets/flask/blob/master/src/flask/app.py#L2017
[2]: https://pgjones.dev
