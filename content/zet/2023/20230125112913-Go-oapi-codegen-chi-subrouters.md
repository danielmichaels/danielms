+++
title = "Go oapi-codegen chi subrouters"
categories = ["zet"]
tags = ["zet"]
slug = "Go-oapi-codegen-chi-subrouters"
date = "2023-01-25 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Go oapi-codegen chi subrouters

I'm experimenting with open-api code generation. So far its pretty positive; write some yaml
and it generates a bunch of interfaces. Those interfaces are where you write the handler
business logic. It'll handle missing data in and out and error if so. 

As a `chi` user I've started making a simple todo app to test it all out. I was disappointed
to learn that `oapi-codegen` doesn't create the swagger UI for you but it will render the
yaml file for you.

Except it doesn't work out of the box for `chi`. In fact no other endpoint works - only
those specified the specification.

To get around this you must create a subrouter and use that to do the specification
validation else it'll always fail for other routes. It makes sense in the end. 

This is how I've done it in a simple way.

```golang
func main() {
	logger := httplog.NewLogger("mudmap", httplog.Options{
		JSON:     false,
		Concise:  true,
		LogLevel: "DEBUG",
	})
	logger = logger.With().Caller().Logger()
	swagger, err := spec.GetSwagger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil
	si := spec.NewTodoStore()

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(5))
	r.Use(httplog.RequestLogger(logger))
	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(swagger)
	})
  // Add n number of non-specification endpoints. But use sparingly.

	r.Mount("/", subRouter(swagger))

	spec.HandlerFromMux(si, r)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:9090"),
	}

	log.Println("starting server")
	log.Fatal(s.ListenAndServe())
}


// subRouter holds the oapi-codegen validators which are 
// mounted to the parent chi router.
func subRouter(s *openapi3.T) http.Handler {
	r := chi.NewRouter()
	r.Use(oapimiddleware.OapiRequestValidator(s))
	return r
}

```

Tags:

    #chi #go #openapi #codegen
