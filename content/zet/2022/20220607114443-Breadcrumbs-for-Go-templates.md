+++
title = "Breadcrumbs for Go templates"
categories = ["zet"]
tags = ["zet"]
slug = "Breadcrumbs-for-Go-templates"
date = "2022-06-07 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Breadcrumbs for Go templates

How I am dealing with breadcrumbs for Go templates.

1. Create a `BreadCrumb` struct
2. Attach the `BreadCrumb` to the template, in my case I use a `TemplateData` struct in every
template.
3. Pass the `BreadCrumb` to the view and then reference it within the template.

```go 
type BreadCrumb struct {
  Name string
  Href string
}
type TemplateData struct {
  BreadCrumbs []BreadCrumb
  // ... trunc
}

func (s *Server) handleHome() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    // trunc
    crumbs := []BreadCrumb{
      {Name: "Page One", Href: "/one"},
      {Name: "Page Two", Href: "/two"},
    }
    s.render(w, https.StatusOK, "home.tmpl", TemplateData{BreadCrumbs: crumbs})
```

Then reference that in templates with 

```html
{{ if .BreadCrumbs }}
    {{range .BreadCrumbs }}
      <li>
        <div class="flex items-center">
          <svg class="flex-shrink-0 h-5 w-5 text-gray-400"
                  xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"
                  fill="currentColor" aria-hidden="true">
            <path fill-rule="evenodd"
                    d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                    clip-rule="evenodd"/>
          </svg>
          <a href="{{.Href}}" class="ml-4 text-sm font-medium text-gray-500 hover:text-gray-700">{{.Name}}</a>
        </div>
      </li>
    {{end}}
{{end}}
```

Tags:

    #go #templates
