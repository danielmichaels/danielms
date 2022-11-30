package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func createZetMarkdownFiles() {
	p, _ := os.Getwd()
	zetPath := filepath.Join(p, "content/zet")

	zetFile, err := os.ReadFile("zet.json")
	if err != nil {
		log.Fatalln("failed to parse zet.json")
	}

	var zets []ZetJson
	err = json.Unmarshal(zetFile, &zets)

	// Get years and create directory paths for them if not exist.
	yrs := getZetYears(zets)
	for _, v := range yrs {
		v := strconv.Itoa(v)
		exist, err := exists(filepath.Join(zetPath, v))
		if err != nil {
			log.Fatalln(err)
		}
		if !exist {
			_ = os.MkdirAll(filepath.Join(zetPath, v), os.ModePerm)
			log.Printf("created %q directory", filepath.Join(zetPath, v))
		}
	}

	for _, z := range zets[:] {
		zbody, err := getZetContents(z)
		if err != nil {
			log.Fatalln("failed to retrieve zet body from github", err)
		}
		zt := ZetTemplate{
			Title:  z.Content,
			Slug:   slugify(z.Content),
			Date:   z.Date,
			Body:   string(zbody),
			IsoSec: z.Path,
		}

		_, err = createTemplate(&zt)
		if err != nil {
			log.Fatalln("failed to create template", err)
		}

		zetYear := strconv.Itoa(z.Date.Year())
		writeZetMarkdown(zt, filepath.Join(zetPath, zetYear))

	}
}
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// getZetYears extracts only unique years from an array of all zets
func getZetYears(zets []ZetJson) []int {
	var yrs []int
	for _, y := range zets {
		yrs = append(yrs, y.Date.Year())
	}

	return removeDupeYears(yrs)
}

// removeDupeYears whittles down an array of years as int's to a list of unqiue
// years.
func removeDupeYears(arr []int) []int {
	keys := make(map[int]bool)
	var l []int
	for _, item := range arr {
		if _, value := keys[item]; !value {
			keys[item] = true
			l = append(l, item)
		}
	}
	return l
}
func getZetContents(z ZetJson) ([]byte, error) {
	cl := http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(http.MethodGet, z.DownloadUrl, nil)
	if err != nil {
		return nil, err
	}
	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func writeZetMarkdown(zt ZetTemplate, path string) {
	tpl, err := createTemplate(&zt)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s-%s.md", path, zt.IsoSec, zt.Slug), tpl, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("wrote %q markdown file into zet directory", zt.Title)
}

type ZetJson struct {
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Sha         string    `json:"sha"`
	Size        int       `json:"size"`
	Url         string    `json:"url"`
	HtmlUrl     string    `json:"html_url"`
	GitUrl      string    `json:"git_url"`
	DownloadUrl string    `json:"download_url"`
	Type        string    `json:"type"`
	Content     string    `json:"content"`
	Encoding    string    `json:"encoding"`
	Date        time.Time `json:"date"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		Html string `json:"html"`
	} `json:"_links"`
}

type ZetTemplate struct {
	Title  string
	Slug   string
	Date   time.Time
	Body   string
	IsoSec string
}

func slugify(s string) string {
	return strings.Replace(s, " ", "-", -1)
}
func createTemplate(z *ZetTemplate) ([]byte, error) {
	t := template.New("zet")

	parse, err := t.Parse(strings.TrimSpace(`
+++
title = "{{.Title}}"
categories = ["zet"]
tags = ["zet"]
slug = "{{.Slug}}"
date = "{{.Date}}"
draft = "false"
ShowToc = "true"
+++

{{.Body}}
	`))
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = parse.Execute(&buf, z)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
