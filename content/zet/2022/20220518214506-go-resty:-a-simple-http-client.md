+++
title = "Go-resty: A Simple HTTP Client"
categories = ["zet"]
tags = ["zet"]
slug = "go-resty:-a-simple-http-client"
date = "2022-05-18 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Go-resty: A Simple HTTP Client

When you need a decent HTTP client and don't want to craft your own, then
go-resty is a good choice. It is great when you need retries or an way
to send requests based on the response, i.e. token expiration and renewal.


How I implement a JWT re-authentication check before each request.

```go
import (
	"github.com/go-resty/resty/v2"
  "github.com/form3tech-oss/jwt-go"
)

func NewAuth0() Auth0 {
	c := resty.New()
	c.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		fmt.Println(c.Token)
		if !IsTokenValid(c.Token) {
			r, err := Authenticate()
			if err != nil {
				return err
			}
			request.Token = r.AccessToken
		}
		return nil
	})
	return Auth0{
		Groups: GroupApiStruct{
			client: RestClient{*c, *config.AppConfig(), c.Token},
		},
	}
}
func IsTokenValid(token string) bool {
	println(token)
	if token == "" {
		return false
	}
	claims := &jwt.StandardClaims{}
	p := &jwt.Parser{}
	_, _, err := p.ParseUnverified(token, claims)
	if err != nil {
		return false
	}
	err = claims.Valid()
	if err != nil {
		return false
	}
	return true
}

func Authenticate() (AuthResponse, error) {
	c := config.AppConfig()
	client := resty.New()

	var result AuthResponse
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(AuthPayload{
			GrantType:    "client_credentials",
			ClientId:     c.Secrets.Auth0ApiClientId,
			ClientSecret: c.Secrets.Auth0ApiClientSecret,
			Audience:     c.Secrets.Auth0ApiClientAudience,
		}).
		SetResult(&result).
		Post("https://mudmap.au.auth0.com/oauth/token")
	if err != nil {
		fmt.Println(err.Error())
		return AuthResponse{}, err
	}
	return result, nil
}

type AuthResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	Scope       string `json:"scope,omitempty"`
}

type AuthPayload struct {
	GrantType    string `json:"grant_type,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Audience     string `json:"audience,omitempty"`
}
```

This lets me check that the JWT is valid and if not, to request a new
one from the service. 

Coding this myself using just the standard library while possible is not
a pragmatic use of my time. I vendor my code so even if `resty` went away
tomorrow, it'll keep working long enough to replace it.

Tags:

    #go #rest
