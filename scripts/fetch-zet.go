package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var (
	repo        = "zet"
	owner       = "danielmichaels"
	repoUrl     = fmt.Sprintf("https://github.com/%s/%s.git", owner, repo)
	cloneDir    = "/tmp/zet-clone"
	excludeDirs = []string{".git", "LICENSE", ".github"}
)

type ZetEntry struct {
	Title       string    `json:"content"`
	Path        string    `json:"path"`
	Sha         string    `json:"sha"`
	HtmlUrl     string    `json:"html_url"`
	DownloadUrl string    `json:"download_url"`
	Date        time.Time `json:"date"`
	Content     string    `json:"name"`
}

type ZetTemplate struct {
	Title  string
	Slug   string
	Date   time.Time
	Body   string
	IsoSec string
}

func main() {
	log.Println("Starting zet fetch process")
	start := time.Now()
	defer func() {
		finish := time.Since(start)
		log.Println("Execution Time:", finish)
	}()

	err := cloneRepository()
	if err != nil {
		log.Fatalln("Failed to clone repository:", err)
	}
	defer cleanupClone()

	entries, err := processZetEntries()
	if err != nil {
		log.Fatalln("Failed to process zet entries:", err)
	}

	if noNewZets(entries) {
		fmt.Println("No new zet's. Skipping...")
		return
	}

	err = writeToFile(entries)
	if err != nil {
		log.Fatalln("writeToFile failed", err)
	}

	createZetMarkdownFiles()
}

// cloneRepository clones the zet repository to the temporary directory
func cloneRepository() error {
	log.Printf("Cloning repository %s to %s", repoUrl, cloneDir)

	// remove if already exists
	_ = os.RemoveAll(cloneDir)

	cmd := exec.Command("git", "clone", repoUrl, cloneDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git clone failed: %s, output: %s", err, output)
	}

	return nil
}

// cleanupClone removes the temporary clone directory
func cleanupClone() {
	log.Printf("Cleaning up %s", cloneDir)
	_ = os.RemoveAll(cloneDir)
}

// processZetEntries processes all zet entries in the cloned repository
func processZetEntries() ([]*ZetEntry, error) {
	var entries []*ZetEntry

	// Get list of directories (isosec format)
	dirs, err := os.ReadDir(cloneDir)
	if err != nil {
		return nil, err
	}

	for _, dir := range dirs {
		if !dir.IsDir() || isExcluded(dir.Name()) {
			continue
		}

		dirPath := filepath.Join(cloneDir, dir.Name())
		readmePath := filepath.Join(dirPath, "README.md")

		if _, err := os.Stat(readmePath); os.IsNotExist(err) {
			log.Printf("No README.md in %s, skipping", dirPath)
			continue
		}

		title, err := readZetTitle(readmePath)
		if err != nil {
			log.Printf("Error reading title from %s: %v", readmePath, err)
			continue
		}

		sha, err := getGitSha(dirPath)
		if err != nil {
			log.Printf("Error getting SHA for %s: %v", dirPath, err)
			continue
		}

		date, err := time.Parse("20060102", dir.Name()[:8])
		if err != nil {
			log.Printf("Error parsing date from %s: %v", dir.Name(), err)
			continue
		}

		entry := &ZetEntry{
			Title:       title,
			Path:        dir.Name(),
			Sha:         sha,
			HtmlUrl:     fmt.Sprintf("https://github.com/%s/%s/tree/main/%s", owner, repo, dir.Name()),
			DownloadUrl: fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/main/%s/README.md", owner, repo, dir.Name()),
			Date:        date,
			Content:     title,
		}

		entries = append(entries, entry)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Date.After(entries[j].Date)
	})

	log.Printf("Found %d zet entries", len(entries))
	return entries, nil
}

// isExcluded checks if a directory should be excluded
func isExcluded(name string) bool {
	for _, exclude := range excludeDirs {
		if name == exclude {
			return true
		}
	}
	// Also exclude entries that don't start with numeric isosec format
	if len(name) < 8 {
		return true
	}
	_, err := strconv.Atoi(name[:8])
	return err != nil
}

// readZetTitle reads the first line of a README.md file and extracts the title
func readZetTitle(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		title := scanner.Text()
		title = strings.TrimPrefix(title, "# ")
		title = strings.TrimSpace(title)
		return title, nil
	}

	return "", errors.New("empty file or error reading file")
}

// getGitSha gets the git SHA of the last commit that modified the specified path
func getGitSha(path string) (string, error) {
	relPath, err := filepath.Rel(cloneDir, path)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("git", "-C", cloneDir, "log", "-n", "1", "--pretty=format:%H", "--", relPath)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	sha := strings.TrimSpace(string(out))
	if sha == "" {
		// fallback to the HEAD commit
		cmd = exec.Command("git", "-C", cloneDir, "rev-parse", "HEAD")
		out, err = cmd.Output()
		if err != nil {
			return "", err
		}
		sha = strings.TrimSpace(string(out))
	}

	return sha, nil
}

// noNewZets checks if the current zet.json file contains the same entries as the ones we just found
func noNewZets(entries []*ZetEntry) bool {
	zetFilePath := "./assets/zet.json"
	_, err := os.Stat(zetFilePath)
	if err != nil {
		log.Println("zet.json does not exist", err)
		return false
	}

	zetFile, err := os.ReadFile(zetFilePath)
	if err != nil {
		log.Println("error checking zet.json contents", err)
		return false
	}

	var existingZets []*ZetEntry
	err = json.Unmarshal(zetFile, &existingZets)
	if err != nil {
		log.Println("error unmarshalling zet.json", err)
		return false
	}

	if len(existingZets) != len(entries) {
		log.Printf("zet count changed: %d existing vs %d new", len(existingZets), len(entries))

		// Create maps to track entries by path
		existingMap := make(map[string]*ZetEntry)
		for _, z := range existingZets {
			existingMap[z.Path] = z
		}

		newMap := make(map[string]*ZetEntry)
		for _, z := range entries {
			newMap[z.Path] = z
		}

		// Find added zets
		for path, entry := range newMap {
			if _, exists := existingMap[path]; !exists {
				log.Printf("New zet added: %s - %s", path, entry.Title)
			}
		}

		// Find removed zets
		for path, entry := range existingMap {
			if _, exists := newMap[path]; !exists {
				log.Printf("Zet removed: %s - %s", path, entry.Title)
			}
		}

		return false
	}

	// Check for changes in SHA or content
	changesFound := false
	for i, v := range entries {
		if i >= len(existingZets) {
			log.Printf("New zet added at index %d: %s", i, v.Title)
			changesFound = true
			continue
		}

		if v.Sha != existingZets[i].Sha {
			log.Printf("SHA changed for %s: %s → %s", v.Path, existingZets[i].Sha[:8], v.Sha[:8])
			changesFound = true
		}

		if v.Title != existingZets[i].Title {
			log.Printf("Title changed for %s: %q → %q", v.Path, existingZets[i].Title, v.Title)
			changesFound = true
		}
	}

	return !changesFound
}

func writeToFile(entries []*ZetEntry) error {
	// Make sure assets directory exists
	err := os.MkdirAll("./assets", 0755)
	if err != nil {
		return err
	}

	j, err := json.MarshalIndent(entries, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile("./assets/zet.json", j, 0644)
	if err != nil {
		return err
	}
	log.Printf("wrote %d entries to zet.json", len(entries))
	return nil
}

func createZetMarkdownFiles() {
	p, _ := os.Getwd()
	zetPath := filepath.Join(p, "content/zet")

	zetFile, err := os.ReadFile("./assets/zet.json")
	if err != nil {
		log.Fatalln("failed to parse zet.json")
	}

	var zets []ZetEntry
	err = json.Unmarshal(zetFile, &zets)
	if err != nil {
		log.Fatalln("failed to unmarshal zet.json")
	}

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
		readmePath := filepath.Join(cloneDir, z.Path, "README.md")
		zbody, err := os.ReadFile(readmePath)
		if err != nil {
			log.Printf("failed to read content from %s: %v", readmePath, err)
			// fallback, attempt to fetch from github
			zbody, err = getZetContents(z.DownloadUrl)
			if err != nil {
				log.Fatalf("failed to retrieve zet body: %v", err)
			}
		}

		zt := ZetTemplate{
			Title:  z.Title,
			Slug:   slugify(z.Title),
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
func getZetYears(zets []ZetEntry) []int {
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

func getZetContents(url string) ([]byte, error) {
	// Use Go's HTTP client instead of external curl command
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
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

func slugify(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "-", -1)
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
mermaid = "true"
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
