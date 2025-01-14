package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type Blog struct {
	Title   string
	Slug    string
	DateNow string
	Draft   bool
	ShowToc bool
}

func (b *Blog) createTemplate() ([]byte, error) {
	t := template.New("blog")

	parse, err := t.Parse(strings.TrimSpace(`
+++
title = "{{.Title}}"
categories = [""]
tags = [""]
slug = "{{.Slug}}"
date = "{{.DateNow}}"
draft = "{{.Draft}}"
ShowToc = "{{.ShowToc}}"
+++
	`))
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = parse.Execute(&buf, b)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (b *Blog) slugify() {
	b.Slug = strings.ReplaceAll(b.Slug, " ", "-")
}

func (b *Blog) hugoDateFormatter() {
	t := time.Now()
	b.DateNow = t.Format("2006-01-02")
}

func missingArgs(f string) {
	fmt.Printf("must provide %q argument", f)
	flag.Usage()
	os.Exit(1)
}

func (b *Blog) writePost(path string) {
	tpl, err := b.createTemplate()
	if err != nil {
		log.Fatalln(err)
	}
	f := fmt.Sprintf("%s/%s-%s.md", path, b.DateNow, b.Slug)
	mkdirp(path)
	err = os.WriteFile(f, tpl, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("wrote %q to file", f)
}

func mkdirp(p string) {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		// Create the directory
		err = os.Mkdir(p, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Directory created successfully!")
	}
}

func main() {
	title := flag.String("title", "", "Blog post title")
	slug := flag.String("slug", "", "Blog post url slug")
	draft := flag.Bool("draft", false, "Is post a draft")
	toc := flag.Bool("show-toc", true, "Show table of contents")

	flag.Parse()
	if *title == "" {
		missingArgs("title")
	}
	if *slug == "" {
		*slug = *title
	}
	blogPost := Blog{
		Title:   *title,
		Slug:    *slug,
		Draft:   *draft,
		ShowToc: *toc,
	}
	blogPost.slugify()
	blogPost.hugoDateFormatter()
	p, _ := os.Getwd()
	by := fmt.Sprintf("content/blog/%d", time.Now().Year())
	blogPath := filepath.Join(p, by)
	blogPost.writePost(blogPath)
}
