+++
title = "Go-colly meta tags"
categories = ["zet"]
tags = ["zet"]
slug = "Go-colly-meta-tags"
date = "2023-01-31 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Go-colly meta tags

TIL how to scrape meta tags, specifically `og:title` and so on using `go-colly`.

```golang
c.OnHTML("html", func(e *colly.HTMLElement) {
  // note the use of backticks and double quotes - it has to be exactly like this
  title := e.ChildAttr(`meta[property="og:title"]`, "content")
  fmt.Println(title)
})
```

Tags:

    #scraping #go #colly

