+++
title = "Design First API Development with Goa"
categories = [""]
tags = [""]
slug = "Design-First-API-Development-with-Goa"
date = "2024-01-09"
draft = "false"
ShowToc = "true"
+++

[Goa.design][goa] is a golang tool for developing APIs using a design first approach. By leveraging Goa's approach it is
possible to generate server and client code automatically, documentation through OpenAPI (version 2 and 3 are supported) 
documents as well as gRPC code. 

This blog post introduces Goa, it's concepts and showcases some short examples. It does not walk through the
installation, or intend to supersede its tutorial/walkthrough which you should look over, 
[here](https://goa.design/learn/getting-started/)

Goa is broken into three parts; the design language (DSL), code generation and the Go package itself.

## What is a DSL?

> A DSL (Domain-Specific Language) is a programming language or a set of rules and syntax specifically designed to solve problems within a particular domain or industry. It provides a higher-level abstraction that allows programmers to express solutions in a more concise and declarative manner.                                                                     
          
This statement sums it up succinctly. In our case the *domain* or *industry* specifics relate to HTTP and RPC communication
via network calls. A lot of API work is boilerplate and can be hard to implement all the things considered *best practise*.
Goa strives to reduce the amount of programmer work required to build out well crafted APIs. 

Instead of writing out an OpenAPI document and then converting it to Go code such as [oapi-codegen](https://github.com/deepmap/oapi-codegen),
Goa uses its DSL to generate it (and the code).

Personally, I find the DSL quite powerful and easy to understand. What's more, it starts simple but provides rich ways
to extend it as your requirements dictate. The creator ([raphael][rr]) is also very active on github and golang's slack.

Here's a snippet of the DSL.

```golang
package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("check-redirects", func() {
	Title("HTTP Redirection Detection Service")
	Description("HTTP service detecting and reporting any and all redirects in a HTTP request")
	Server("server", func() {
		Host("localhost", func() { URI("http://localhost:9090") })
	})
})

var _ = Service("health", func() {
	Description("endpoints for determining service uptime and status")
	HTTP(func() {
		Path("/")
	})
	Method("healthz", func() {
		HTTP(func() {
			GET("/healthz")
			Response(StatusOK)
		})
		Result(Empty)
	})
	Method("version", func() {
		HTTP(func() {
			GET("/version")
			Response(StatusOK)
		})
		Result(AppVersion)
	})
})
var AppVersion = Type("version", func() {
	Description("Application version information")
	Attribute("version", String, "Application version", func() {
		Example("1.0")
		Example("6b51bebe0f965a5fffa8ff9db5aa702c76ec47f2")
	})
})
```

Without going too deep right now, this will create an API called *server*, a service called *health* and within the *health*
service build two endpoints; `healthz` and `version`. A custom type called `version` will also be created and used in the
`/version` endpoint.

The DSL provides an abstraction which lets you craft your API in a declarative way. One big benefit is its just Go code
meaning you can easily extend, simplify or create generic functions when working with it. 

## Code generation

Writing the DSL by itself does not generate code or documentation. The `goa` CLI does that. Specifically, `goa gen`.

After [installing](https://goa.design/learn/getting-started/) Goa, you can generate the client and server code. 

To generate the code run `goa gen`. Typically, Goa suggests a pattern of placing all DSL files into
a directory called `design`. You don't have to do this, but I think it makes sense for most use cases.

For the above code snippet, if it were at this path `design/design.go` you would run 
`goa gen github.com/danielmichaels/checkredirects/design`. In this example `github.com/danielmichaels/checkredirects` is
my module that I used during `go mod init`.

### Example generation

With this directory structure:

```shell
.
├── design
│   └── design.go
├── go.mod
└── go.sum
```

After I run `goa gen github.com/danielmichaels/checkredirects/design`, it will create a directory called `gen` with the following files:

```shell
├── design
│   └── design.go
├── gen # New
│   ├── health
│   │   ├── client.go
│   │   ├── endpoints.go
│   │   └── service.go
│   └── http
│       ├── cli
│       │   └── server
│       │       └── cli.go
│       ├── health
│       │   ├── client
│       │   │   ├── client.go
│       │   │   ├── cli.go
│       │   │   ├── encode_decode.go
│       │   │   ├── paths.go
│       │   │   └── types.go
│       │   └── server
│       │       ├── encode_decode.go
│       │       ├── paths.go
│       │       ├── server.go
│       │       └── types.go
│       ├── openapi3.json
│       ├── openapi3.yaml
│       ├── openapi.json
│       └── openapi.yaml
├── go.mod
└── go.sum
```

Now we have our Go code which we can import into our services. Note, this is all auto generated and **should not** be edited as
it'll be overwritten whenever we run `goa gen`.

### Creating the services

Now that the package code has been generated we can create the entrypoint, and service files automatically with another command;
`goa example`. Running this results in some new files:

```shell
.
├── cmd # New
│   ├── server
│   │   ├── http.go
│   │   └── main.go
│   └── server-cli
│       ├── http.go
│       └── main.go
├── design
│   └── design.go
├── gen # Truncated for brevity
├── go.mod
├── go.sum
└── health.go # New
```

Unlike `goa gen` this is a one-shot deal; if the generated files already exist it will not re-create them. This is because all your business logic
will be inside these files and it may override things you don't want overridden.

If we peek at `health.go` it will have stubbed out all the handlers, ready to be populated which your business logic.

```go
package checkredirects

import (
	"context"
	"log"

	health "github.com/danielmichaels/checkredirects/gen/health"
)

// health service example implementation.
// The example methods log the requests and return zero values.
type healthsrvc struct {
	logger *log.Logger
}

// NewHealth returns the health service implementation.
func NewHealth(logger *log.Logger) health.Service {
	return &healthsrvc{logger}
}

// Healthz implements healthz.
func (s *healthsrvc) Healthz(ctx context.Context) (err error) {
	s.logger.Print("health.healthz")
	return
}

// Version implements version.
func (s *healthsrvc) Version(ctx context.Context) (res *health.Version2, err error) {
	res = &health.Version2{}
	s.logger.Print("health.version")
	return
}
```

So after writing only 40 lines (`design/design.go`) we were able to auto generate a complete and working web server with 
two endpoints.

The responses as expected do not return anything but work as demonstrated below.

```shell
$ curlie :9090/healthz
HTTP/1.1 200 OK
Date: Tue, 09 Jan 2024 04:44:21 GMT
Content-Length: 0

$ curlie :9090/version
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 09 Jan 2024 04:44:28 GMT
Content-Length: 3

{
    
}
```

## Documentation

Another great feature of Goa is how well it documents the code and that it can generate valid OpenAPI documents.

This is the document it created:

```yaml
openapi: 3.0.3
info:
    title: HTTP Redirection Detection Service
    description: HTTP service detecting and reporting any and all redirects in a HTTP request
    version: 0.0.1
servers:
    - url: http://localhost:9090
paths:
    /healthz:
        get:
            tags:
                - health
            summary: healthz health
            operationId: health#healthz
            responses:
                "200":
                    description: OK response.
    /version:
        get:
            tags:
                - health
            summary: version health
            operationId: health#version
            responses:
                "200":
                    description: OK response.
                    content:
                        application/json:
                            schema:
                                $ref: '#/components/schemas/Version'
                            example:
                                version: 6b51bebe0f965a5fffa8ff9db5aa702c76ec47f2
components:
    schemas:
        Version:
            type: object
            properties:
                version:
                    type: string
                    description: Application version
                    example: 6b51bebe0f965a5fffa8ff9db5aa702c76ec47f2
            example:
                version: 6b51bebe0f965a5fffa8ff9db5aa702c76ec47f2
tags:
    - name: health
      description: endpoints for determining service uptime and status
```

Goa also provides powerful constructs enhance the documents. For example to define a field on a `schema` we can use
`Attribute` or  `Field` type. For this post we're only focusing on HTTP which uses `Attribute` whereas `Field` is for both
HTTP and gRPC.

Example of a Goa `Payload` and how we can add more context to it. 

```go
// Truncated
Payload(func() {
    Attribute("username", String, "Username", func () {
        Example("MyUsername")
		Pattern("^user_[a-zA-Z0-9]{12}$")
    })
    Required("password")
})
// Truncated
```

This will create an OpenAPI document which provides an example of `MyUsername`. The `Pattern` will also enforce the regex
and automatically handle payload validation without the need to write any logic.

## Conclusion

This post sought to introduce Goa in the most simple of terms. It hardly scratches the surface of its capabilities. I chose
the most simplistic example I could because it will lay the groundwork for some follow-up posts which show how to add more
realistic endpoints. In the next post I will create a service which accepts and returns JSON leveraging Goa's `Type`, `Error`
and `Payload` DSL primitives. This will be published with source code.

[goa]: https://goa.design
[rr]: https://github.com/raphael
