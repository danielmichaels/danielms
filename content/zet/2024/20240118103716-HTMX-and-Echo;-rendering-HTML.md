+++
title = "HTMX and Echo; rendering HTML"
categories = ["zet"]
tags = ["zet"]
slug = "HTMX-and-Echo;-rendering-HTML"
date = "2024-01-18 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# HTMX and Echo; rendering HTML

I am experimenting with replacing a Next.js frontend with HTMX (and Templ).
Typically, I use `chi` as my router but this time I've opted for `echo`.

For one reason, by default it uses context and all handlers must return an error.
This prevents any missed returns when error handling (which can be caught with
linters but its an extra step).

I'm used to `html/template` but getting it to work with `echo` was a different.

Following the [guide](https://echo.labstack.com/docs/templates) and adding my own
spin I came up with the following.

```go
type Template struct {
	templates *template.Template
}

func NewTemplate(patterns []string) *Template {
	for i := range patterns {
    // assets/view/ contains all my templates
		patterns[i] = "view/" + patterns[i]
	}

	ts, err := template.New("").Funcs(nil).ParseFS(assets.EmbeddedFiles, patterns...)
	if err != nil {
		log.Fatal().Err(err).Msg("template.New err")
	}

	return &Template{templates: ts}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
  e := echo.New()
	e.Use(middleware.Recover())
  // this is how Echo learns about our templates.
  // I can pass in multiple locations for templates but I only care about
  // the HTMX fragments for this renderer - everything else is handled by Templ
	e.Renderer = NewTemplate([]string{"fragments/*.tmpl"})
  // all routes are loaded here
  app.routes(e)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Server.Port),
		Handler:      e,
		IdleTimeout:  app.Config.Server.TimeoutIdle,
		ReadTimeout:  app.Config.Server.TimeoutRead,
		WriteTimeout: app.Config.Server.TimeoutWrite,
	}
	app.Logger.Info().Msgf("Started Webserver on '%d'", app.Config.Server.Port)
	e.Logger.Fatal(srv.ListenAndServe())
```

Now I can render HTML (with HTMX) using Echo.

Tags:

  #til #go #echo #templates
