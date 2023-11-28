+++
title = "Looping over embedded structs in html template"
categories = ["zet"]
tags = ["zet"]
slug = "Looping-over-embedded-structs-in-html-template"
date = "2023-11-28 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Looping over embedded structs in html template

I am using the `slack-go` package and needed to pull out information from slacks Conversations
API.

`slack-go` uses the following structs to store this data:

```go
// Channel contains information about the channel
type Channel struct {
	GroupConversation
	IsChannel bool   `json:"is_channel"`
	IsGeneral bool   `json:"is_general"`
	IsMember  bool   `json:"is_member"`
	Locale    string `json:"locale"`
}
// GroupConversation is the foundation for Group and Channel
type GroupConversation struct {
	Conversation
	Name       string   `json:"name"`
	Creator    string   `json:"creator"`
	IsArchived bool     `json:"is_archived"`
	Members    []string `json:"members"`
	Topic      Topic    `json:"topic"`
	Purpose    Purpose  `json:"purpose"`
}
```

Getting the conversation is easy;

```go
s := slack.New(botToken)
	channels, _, _ := s.GetConversationsContext(ctx, &slack.GetConversationsParameters{
		TeamID: "id",
	})
```

And it returns `[]slack.Channel`

In a html template I wanted to get all the available channels and put them into a `<select>`.
It took a long time before I realised that `GroupConversation` is an embedded struct and thus
you cannot reference it by a struct value.

To get the `GroupConversation.Name` from a slice of `Channel` in `.tmpl` I used the following loop:

```go-template
<select class="select select-bordered w-full max-w-xs">
    <option disabled selected>Slack Channel</option>
      {{with .Channels}}
          {{range .}}
            <option>{{.Name}}</option>
          {{end}}
      {{end}}
  </select>
```

The `{{ range . }}` seems a little off putting to me; magical even. But, it works. I'm hoping
there is a better way but this works quite well for now.


Tags:

  #TIL #go

