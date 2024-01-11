package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	slugRx  = `slug\s*=\s*"([^"]+)"`
	titleRx = `title\s*=\s*"([^"]+)"`
)

var (
	uri        = "https://plausible.io/api/v1/stats/breakdown?site_id=danielms.site&period=12mo&property=event:page&filters=event:page==/blog/**/&limit=5"
	fileName   = "top-posts.json"
	tmpFileNae = "/tmp/plausible-stats.json"
)

type PlausibleStats struct {
	Results []struct {
		Page     string `json:"page"`
		Visitors int    `json:"visitors"`
	} `json:"results"`
}

type Post struct {
	Title string `json:"title"`
	// Slug represents the 'slug' field from Hugo's metadata
	Slug string `json:"slug"`
	// SlugURL is the internet address from the root, e.g. /blog/rsync-cheatsheet
	// This the URL Plausible tracks in the analytics
	SlugURL string `json:"slug_url"`
	// Path represents path on disk. This is used to create the reference in custom Hugo shortcodes.
	Path string `json:"path"`
}

type Posts []Post

// fetchStats retrieves each file or folder from the repository.
func fetchStats(d *PlausibleStats) error {
	cl := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("PLAUSIBLE_API_KEY")))
	res, err := cl.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	err = json.NewDecoder(res.Body).Decode(&d)
	if err != nil {
		return err
	}
	return nil
}

func writeJSONToFile(d any, fn string) error {
	if d, ok := d.(*PlausibleStats); ok {
		j, err := json.MarshalIndent(d.Results, "", "\t")
		if err != nil {
			return err
		}
		err = os.WriteFile(fn, j, 0644)
		if err != nil {
			return err
		}
		log.Printf("wrote %d entries to %q", len(d.Results), fn)
	}
	if p, ok := d.(*Posts); ok {
		j, err := json.MarshalIndent(p, "", "\t")
		if err != nil {
			return err
		}
		err = os.WriteFile(fn, j, 0644)
		if err != nil {
			return err
		}
		log.Printf("wrote %d entries to %q", len(*p), fn)
	}
	return nil
}

func readHugoMetadata(path string) (Post, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return Post{}, err
	}

	startMarker := "+++"
	endMarker := "+++"

	// Find the start and end positions of the markers
	start := strings.Index(string(contents), startMarker)
	end := strings.LastIndex(string(contents), endMarker)
	if start == -1 || end == -1 || start >= end {
		return Post{}, fmt.Errorf("failed to find markers in file: %s", path)
	}
	contentsBetweenMarkers := string(contents)[start+len(startMarker) : end]
	slug := getFieldFromHugoMetadata(slugRx, contentsBetweenMarkers)
	title := getFieldFromHugoMetadata(titleRx, contentsBetweenMarkers)
	if slug != "" {
		post := Post{Slug: slug, Path: path, Title: title}
		post.toTitle()
		post.stripSlugPrefix()
		post.stripContent()
		post.generateSlugURL()
		return post, nil
	}
	return Post{}, nil
}
func (p *Post) toTitle() {
	p.Title = strings.ReplaceAll(p.Title, "title =", "")
	p.Title = strings.TrimSpace(p.Title)
	p.Title = strings.Trim(p.Title, `"`)
	p.Title = strings.ReplaceAll(p.Title, "-", " ")
}

func (p *Post) stripSlugPrefix() {
	p.Slug = strings.ReplaceAll(p.Slug, "slug =", "")
	p.Slug = strings.TrimSpace(p.Slug)
}
func (p *Post) generateSlugURL() {
	p.Slug = strings.Trim(p.Slug, `"`)
	p.SlugURL = fmt.Sprintf("/blog/%s", p.Slug)
	p.SlugURL = strings.ReplaceAll(p.SlugURL, " ", "-")
	p.SlugURL = strings.ToLower(p.SlugURL)
}
func (p *Post) stripContent() {
	p.Path = strings.ReplaceAll(p.Path, "content/", "")
}

func getFieldFromHugoMetadata(rx, entry string) string {
	re, err := regexp.Compile(rx)
	if err != nil {
		return ""
	}
	matches := re.FindStringSubmatch(entry)
	if len(matches) != 0 {
		return matches[0]
	}
	return ""
}
func main() {
	var data PlausibleStats
	if err := fetchStats(&data); err != nil {
		log.Fatalln("fetchStats err:", err)
	}
	if err := writeJSONToFile(&data, tmpFileNae); err != nil {
		log.Fatalln("writeJSONToFile err:", err)
	}
	var posts Posts
	directory := "content/blog"
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".md" {
			post, err := readHugoMetadata(path)
			if err != nil {
				return err
			}
			if post.Slug != "" {
				posts = append(posts, post)
			}
		}
		return nil
	})
	var matchedPosts Posts
	for i := range data.Results {
		result, _ := strings.CutSuffix(data.Results[i].Page, "/")
		for _, p := range posts {
			if strings.ToLower(p.SlugURL) == result {
				matchedPosts = append(matchedPosts, p)
			}
		}
	}
	err = writeJSONToFile(&matchedPosts, fileName)
	if err != nil {
		log.Fatalln("writeJSONToFile::posts", err)
	}
}
