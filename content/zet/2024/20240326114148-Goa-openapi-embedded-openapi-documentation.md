+++
title = "Goa openapi embedded openapi documentation"
categories = ["zet"]
tags = ["zet"]
slug = "Goa-openapi-embedded-openapi-documentation"
date = "2024-03-26 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Goa openapi embedded openapi documentation

A snippet for how I am storing an OpenAPI documentation page within a Goa
project docker image.

### `design/openapi.go`

```go
// design/openapi.go
package design

import . "goa.design/goa/v3/dsl"

var _ = Service("openapi", func() {
	Description("OpenAPI endpoints for debugging and demonstration")
	HTTP(func() {
		Path("/openapi")
	})
	Method("file", func() {
		Result(func() {
			Attribute("length", Int64, "Length is the downloaded content length in bytes.", func() {
				Example(4 * 1024 * 1024)
			})
			Attribute("encoding", String, func() {
				Example("application/json")
			})
			Required("length", "encoding")
		})

		Error("invalid_file_path", ErrorResult, "Could not locate file for download")
		Error("internal_error", ErrorResult, "Fault while processing download.")

		HTTP(func() {
			GET("/openapi3.json")
			SkipResponseBodyEncodeDecode()
			Response(func() {
				Header("length:Content-Length")
				Header("encoding:Content-Type")
			})
			Response("invalid_file_path", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
		})
	})
	Method("documentation", func() {

		Result(func() {
			Attribute("length", Int64, "Length is the downloaded content length in bytes.", func() {
				Example(4 * 1024 * 1024)
			})
			Attribute("encoding", String, func() {
				Example("application/json")
			})
			Required("length", "encoding")
		})

		Error("invalid_file_path", ErrorResult, "Could not locate file for download")
		Error("internal_error", ErrorResult, "Fault while processing download.")

		HTTP(func() {
			GET("/docs")
			SkipResponseBodyEncodeDecode()
			Response(func() {
				Header("length:Content-Length")
				Header("encoding:Content-Type")
			})
			Response("invalid_file_path", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
		})
	})
})
```

### `openapi.go`

The `example` generated from the above DSL 

```go 
package goagithub

import (
	"context"
	"embed"
	"github.com/danielmichaels/goa-github/gen/openapi"
	"io"
	"log"
)

//go:embed gen/http/openapi3.json
var openapijson embed.FS

//go:embed assets/static/docs.html
var assets embed.FS

// openapi service example implementation.
// The example methods log the requests and return zero values.
type openapisrvc struct {
	logger *log.Logger
}

// NewOpenapi returns the openapi service implementation.
func NewOpenapi(logger *log.Logger) openapi.Service {
	return &openapisrvc{logger}
}

func (o openapisrvc) File(ctx context.Context) (res *openapi.FileResult, body io.ReadCloser, err error) {
	f, err := openapijson.Open("gen/http/openapi3.json")
	if err != nil {
		return nil, nil, openapi.MakeInvalidFilePath(err)
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, nil, openapi.MakeInternalError(err)
	}
	return &openapi.FileResult{
		Length:   fi.Size(),
		Encoding: "application/json",
	}, f, nil
}

func (o openapisrvc) Documentation(ctx context.Context) (res *openapi.DocumentationResult, body io.ReadCloser, err error) {
	f, err := assets.Open("assets/static/docs.html")
	if err != nil {
		return nil, nil, openapi.MakeInvalidFilePath(err)
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, nil, openapi.MakeInternalError(err)
	}
	return &openapi.DocumentationResult{
		Length:   fi.Size(),
		Encoding: "text/html",
	}, f, nil
}
```

### `assets/static/docs.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Open API</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui@5.10.0/dist/swagger-ui.css" integrity="sha256-IgNmYaATqY6z6AuD6hvz9XN0OyeAc94gsTa+lK8ka1Y=" crossorigin="anonymous">
  <style>
      /* Fast dark mode https://github.com/swagger-api/swagger-ui/issues/5327 */
      @media (prefers-color-scheme: dark) {
          body {
              background: #1f1f1f;
          }
          .swagger-ui {
              filter: invert(88%) hue-rotate(180deg);
          }
          .swagger-ui .microlight {
              filter: invert(100%) hue-rotate(180deg);
          }
      }
  </style>
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://cdn.jsdelivr.net/npm/swagger-ui@5.10.0/dist/swagger-ui-bundle.js" integrity="sha256-i050FsZ0MSwm3mVMv7IhpfCdK90RKaXPS/EmiWxv8vc=" crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/swagger-ui@5.10.0/dist/swagger-ui-standalone-preset.js" integrity="sha256-IGoJVXW7MRyeZOsKceWVePAShfVpJhnYhDhOQp+Yi38=" crossorigin="anonymous"></script>
<script>
    // Initialize Swagger UI
    window.onload = function() {
        const ui = SwaggerUIBundle({
            url: "/openapi/openapi3.json",  // Provide the URL to your OpenAPI spec file
            dom_id: '#swagger-ui',
            deepLinking: true,
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIStandalonePreset
            ],
            layout: "BaseLayout"
        });
        window.ui = ui;
    };
</script>
</body>
</html>
```

Tags:

  #goa #dsl #go


