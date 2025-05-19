+++
title = "Goa websocket with subprotocols (auth)"
categories = ["zet"]
tags = ["zet"]
slug = "goa-websocket-with-subprotocols-(auth)"
date = "2024-09-07 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Goa websocket with subprotocols (auth)

How I figured out to pass an API key to my websocket endpoint using `goa` and
`React`. It wasn't straight forward!

Firstly, you have to setup an `websocket.Upgrader{}` and assign `CheckOrigin`
and `Subprotocols` values.

```go
upgrader := &websocket.Upgrader{}
upgrader.CheckOrigin = func(r *http.Request) bool { return true } // setup authorised origins; this is a demo

// You have to tell it that Sec-Websocket-Protocol and your custom auth header are valid subprotocols
// or it won't work! and you will get an error like this:
// Error during WebSocket handshake: Sent non-empty 'Sec-WebSocket-Protocol' header but no response was received
upgrader.Subprotocols = []string{"Sec-Websocket-Protocol", "x-api-key"}
```

Then you have to add some headers to your request. But before we do that, in Goa
we need a custom middleware to read the subprotocol. Using
`Authorization: Bearer asdadad`, or in my case `x-api-key: 123` like a
traditional HTTP request doesn't work.

```
# authorization (subprotocol) header on a websocket 101 Switching
Sec-WebSocket-Protocol: x-api-key, key_000000000000
```

Because its `x-api-key, key_000000000000`, you have to split the string and do
some manipulation. Things you don't have to do in normal HTTP requests.

Here's how I do it using a Goa middleware

```go
// set custom context keys
type xApiKeyWSType string
const apiKeyCtxKey xApiKeyWSType = "x-api-key"

func WebsocketConnection() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := r
            // if not a websocket connection, continue
			if upgrade := r.Header.Get("Connection"); upgrade != "Upgrade" {
				// Not a websocket connection
				h.ServeHTTP(w, req)
			}
            // Get the header with the key
			if auth := r.Header.Get("Sec-WebSocket-Protocol"); auth != "" {
                // the header will return a string: x-api-key, key_000000000000
                // we need to split, then trim white space
				header := strings.Split(auth, ",")
				header[1] = strings.TrimSpace(header[1])
				ctx := context.WithValue(r.Context(), apiKeyCtxKey, header[1])
				req = r.WithContext(ctx)
			}
			h.ServeHTTP(w, req)
		})
	}
}
```

Then in your security scheme, you'll need to get that context key we just set.

```go
func (a *ApiKey) Validate(
	ctx context.Context,
	key string,
	scheme *security.APIKeyScheme,
	db *repository.Queries,
) (context.Context, error) {
	if key == "" {
		key = GetXApiKeyWS(ctx)
	}
    // truncated but do something with the key.
    // a typical HTTP request will have the key available already but
    // we need to grab it via the GetXApiKeyWS function
	return ctx, nil
}
```

It took me a few hours of reading issues, code, websocket specs to figure out
why `wscat` and `insomnia` could connect but not React. Inspecting the headers
was a big clue - no response headers were set from the server. But, if I
connected to a globally available unauthenticated server such as
`wss://echo.websocket.events` they were.

It's all working now and apart from this I'm really happy with how easy it is to
get going with Goa and websockets. From a `design.go`/DSL prespective its
simple. I just replaced `Result` with `StreamingResult` in my Services'
`Method`. Had to plumb in the `upgrader` but it was simple. The hard part was
making sure access was still secure behind my API keys (and soon to be JWTs).

Tags:

    #goa #websockets #react #go
