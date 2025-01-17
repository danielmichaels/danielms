package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
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

var (
	gh         = "https://api.github.com"
	repo       = "zet"
	owner      = "danielmichaels"
	contentUrl = fmt.Sprintf("%s/repos/%s/%s/contents", gh, owner, repo)
	readmeUrl  = fmt.Sprintf("%s/repos/%s/%s/readme", gh, owner, repo)
)

type Content struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Sha         string      `json:"sha"`
	Size        int         `json:"size"`
	Url         string      `json:"url"`
	HtmlUrl     string      `json:"html_url"`
	GitUrl      string      `json:"git_url"`
	DownloadUrl interface{} `json:"download_url"`
	Type        string      `json:"type"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		Html string `json:"html"`
	} `json:"_links"`
}
type Readme struct {
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

func main() {
	log.Println("retrieving zets from github")
	start := time.Now()
	defer func() {
		finish := time.Since(start)
		log.Println("Execution Time:", finish)
	}()
	content, err := fetchContents()
	if err != nil {
		log.Fatalln("fetchContents failed", err)
	}
	paths, err := parseIsosec(content)
	if err != nil {
		log.Fatalln("parseIsosec failed", err)
	}
	rData, err := fetchReadmeData(paths)
	if err != nil {
		log.Fatalln("fetchReadme failed", err)
	}

	readme, err := parseTitle(rData)
	if err != nil {
		log.Fatalln("parseTitle failed", err)
	}

	if noNewZets(readme) {
		fmt.Println("No new zet's. Skipping...")
		return
	}
	err = writeJSONToFile(readme)
	if err != nil {
		log.Fatalln("writeJSONToFile failed", err)
	}
	// only create new markdown entries if a new zet has been created
	createZetMarkdownFiles()
}

// fetchContents retrieves each file or folder from the repository.
func fetchContents() ([]*Content, error) {
	cl := http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(http.MethodGet, contentUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GH_ZET_PAT")))
	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	var content []*Content

	err = decodeJSONBody(res, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// parseIsosec grabs the folder name which is an 'isosec' format. Each isosec
// is appended to an array for use in other functions.
func parseIsosec(c []*Content) ([]string, error) {
	var s []string
	for _, v := range c {
		s = append(s, v.Path)
	}
	return s, nil
}

// fetchReadmeData retrieves the repositories readme subfolders by looping
// over each directory and pulling out the data. This uses the GitHub API /readme
// endpoints. This is limited to 1000 objects. If this is exceeded or rate limiting
// is applied to this task, the git tree endpoint must be used.
func fetchReadmeData(i []string) ([]*Readme, error) {
	cl := http.Client{Timeout: time.Second * 2}

	var r []*Readme

	for _, path := range i[:] {
		var readme *Readme

		url := fmt.Sprintf("%s/%s", readmeUrl, path)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Accept", "application/vnd.github.v3+json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GH_ZET_PAT")))
		res, err := cl.Do(req)
		if err != nil {
			return nil, err
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		err = json.NewDecoder(res.Body).Decode(&readme)
		if err != nil {
			return nil, err
		}
		r = append(r, readme)
		log.Printf("fetched: %q", readme.Path)
	}

	return r, nil
}

// parseTitle loops over an array of Readme objects to munge the data into a
// format we desire for showing users.
func parseTitle(r []*Readme) ([]*Readme, error) {
	for k := range r {
		c, err := base64.StdEncoding.DecodeString(r[k].Content)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to decode for %q", r[k].Path))
		}
		scanner := bufio.NewScanner(strings.NewReader(string(c)))
		var line int
		for scanner.Scan() {
			if line >= 1 {
				break
			}
			// Remove the prefix # from the first line
			content := strings.Replace(scanner.Text(), "# ", "", -1)
			content = strings.TrimSpace(content)
			r[k].Content = content
			// Remove the /README.md from the path struct field
			r[k].Path = strings.Replace(r[k].Path, "/README.md", "", -1)

			r[k].Date = parseDate(r[k])
			line++
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func parseDate(r *Readme) time.Time {
	d := r.Path[:8]
	date, _ := time.Parse("20060102", d)
	r.Date = date
	return r.Date
}

// noNewZets is a checking function to ensure that the zet.json file is only regenerated
// if the current file is not equal to the number of zets found in the zet repo on GitHub.
// This prevents needless commits via GitHub actions.
func noNewZets(r []*Readme) bool {
	z, err := os.Stat("./assets/zet.json")
	if err != nil {
		log.Println("zet.json does not exist", err)
		return false
	}

	zetFile, err := os.ReadFile(z.Name())
	if err != nil {
		log.Println("error checking zet.json contents", err)
		return false
	}

	var existingZets []*Readme
	err = json.Unmarshal(zetFile, &existingZets)
	if err != nil {
		log.Println("error unmarshalling zet.json", err)
		return false
	}

	for k, v := range existingZets {
		if v.Sha != r[k].Sha {
			log.Printf("updated zet %q (old: %s, new: %s) found\n", v.Content, v.Sha, r[k].Sha)
			return false
		}
	}

	if len(existingZets) == len(r) {
		return true
	}

	return false
}

func writeJSONToFile(s []*Readme) error {
	j, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile("./assets/zet.json", j, 0644)
	if err != nil {
		return err
	}
	log.Printf("wrote %d entries to zet.json", len(s))
	return nil
}

// readJSON is helper for trapping errors and return values for JSON related
// handlers
func decodeJSONBody(r *http.Response, dst interface{}) error {
	// Set a max body length. Without this it will accept unlimited size requests
	maxBytes := 1_048_576 // 1MB

	// Init a Decoder and call DisallowUnknownFields() on it before decoding.
	// This means that JSON from the client will be rejected if it contains keys
	// which do not match the target destination struct. If not implemented,
	// the decoder will silently drop unknown fields - this will raise an error instead.
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// decode the request body into the target struct/destination
	err := dec.Decode(dst)
	if err != nil {
		// start triaging the various JSON related errors
		var syntaxError *json.SyntaxError
		var unmarshallTypeError *json.UnmarshalTypeError
		var invalidUnmarshallError *json.InvalidUnmarshalError

		switch {
		// Use the errors.As() function to check whether the error has the
		// *json.SyntaxError. If it does, then return a user-readable error
		// message including the location of the problem
		case errors.As(err, &syntaxError):
			return fmt.Errorf(
				"body contains badly-formed JSON (at character %d)",
				syntaxError.Offset,
			)

		// Decode() can also return an io.ErrUnexpectedEOF for JSON syntax errors. This is
		// checked for with errors.Is() and returns a generic error message to the client.
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		// Wrong JSON types will return an error when they do not match the target destination
		// struct.
		case errors.As(err, &unmarshallTypeError):
			if unmarshallTypeError.Field != "" {
				return fmt.Errorf(
					"body contains incorrect JSON type for field %q",
					unmarshallTypeError.Field,
				)
			}
			return fmt.Errorf(
				"body contains incorrect JSON type (at character %d)",
				unmarshallTypeError.Offset,
			)

		// An EOF error will be returned by Decode() if the request body is empty. Use errors.Is()
		// to check for this and return a human-readable error message
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// If JSON contains a field which cannot be mapped to the target destination
		// then Decode will return an error message in the format "json: unknown field "<name>""
		// We check for this, extract the field name and interpolate it into an error
		// which is returned to the client
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		// If the request body exceeds maxBytes the decode will fail with a
		// "http: request body too large".
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)

		// A json.InvalidUnmarshallError will be returned if we pass a non-nil pointer
		// to Decode(). We catch and panic, rather than return an error.
		case errors.As(err, &invalidUnmarshallError):
			panic(err)

		// All else fails, return an error as-is
		default:
			return err
		}
	}

	// Call Decode() again, using a pointer to anonymous empty struct as the
	// destination. If the body only has one JSON value then an io.EOF error
	// will be returned. If there is anything else, extra data has been sent
	// and we craft a custom error message back to the client
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}
	return nil
}

func createZetMarkdownFiles() {
	p, _ := os.Getwd()
	zetPath := filepath.Join(p, "content/zet")

	zetFile, err := os.ReadFile("./assets/zet.json")
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
	cl := http.Client{Timeout: 10 * time.Second}
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
