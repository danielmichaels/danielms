+++
title = "Go's Functional Options pattern is great"
categories = ["zet"]
tags = ["zet"]
slug = "Go's-Functional-Options-pattern-is-great"
date = "2023-10-02 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Go's Functional Options pattern is great

Some consider this a spicy take but I think it is a fantastic and idiomatic way to build functions/methods which are backwards compatible.


Consider this code from `go-gitlab`. By using the `...ClientOptionFunc` it is possible to add new options without breaking compatibility.

```go
// NewClient returns a new GitLab API client. To use API methods which require
// authentication, provide a valid private or personal token.
// ref: https://github.com/xanzy/go-gitlab/blob/2692fa8f0c4e16c36af8ebdc571da3a0d4ce2d19/gitlab.go#L240-L250
func NewClient(token string, options ...ClientOptionFunc) (*Client, error) {
	client, err := newClient(options...)
	if err != nil {
		return nil, err
	}
	client.authType = PrivateToken
	client.token = token
	return client, nil
}
```

Here is how I've extended it with my own custom HTTP client in a project I am working on.

```go
func NewGitlab(token string, url string, insecure bool, timeout time.Duration, options ...gitlab.ClientOptionFunc) (*Gitlab, error) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure}}
	hc := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
	options = append(options, gitlab.WithBaseURL(url))
	options = append(options, gitlab.WithHTTPClient(hc))
	cl, err := gitlab.NewClient(
		token,
		options...,
	)
	if err != nil {
		return nil, err
	}
	return &Gitlab{Client: cl}, err
}
```

With this pattern in one caller I can pass in options which others shouldn't have. 

```go
// I do not want the default retries so I disable it.
glab, err := providers.NewGitlab(
    form.ClientToken,
    form.GitLabURL,
    insecure,
    providers.GitlabClientDefaultTimeout,
    gitlab.WithoutRetries(),
)
```

In every other instance I do not pass in the `gitlab.WithoutRetries` because the defaults are what I want.

A simple example but to me it makes life a lot easier and I've started adopting it more in my own work.

Tags:

    #go #code
