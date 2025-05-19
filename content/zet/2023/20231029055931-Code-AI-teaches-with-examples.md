+++
title = "Code AI teaches with examples"
categories = ["zet"]
tags = ["zet"]
slug = "code-ai-teaches-with-examples"
date = "2023-10-29 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Code AI teaches with examples

I use AI probably at least once a day. I use it a lot to help with `kubectl --output`, `jq` and
shell scripts.

But, where it really shines, at least for me, is teaching me through examples.

Here's a snippet which I really like. I was thinking about a better way to initialise 
a bunch of NATS subscriptions. AI gave me a good solution that was different to what I
was thinking. And much better; I was considering a `map` of funcs...

My ever growing list of subscriptions.

```go
func (n *Nats) InitSubscribers() error {
	err := n.inboundMailReceived()
	if err != nil {
		return err
	}
	err = n.processInboundMail()
	if err != nil {
		return err
	}
	err = n.eventUserAssignedToEmail()
	if err != nil {
		return err
	}
	err = n.actionPostMessage()
	if err != nil {
		return err
	}
	err = n.eventChannelMessage()
	if err != nil {
		return err
	}
	return nil
}
```

What I asked JetBrains AI: `Refactor this code for me`. Here's its output:

```go
func (n *Nats) InitSubscribers() error {
	functions := []func() error{
		n.inboundMailReceived,
		n.processInboundMail,
		n.eventUserAssignedToEmail,
		n.actionPostMessage,
		n.eventChannelMessage,
		// add more function calls here if necessary
	}

	for _, fn := range functions {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}
```

I love this because I had a problem. Thought through it myself and then asked AI. 
It showed me a way that I wasn't initially considering, and I prefer it to my idea.

Now I have a good mental model for approaching this problem in the future. It's
like my own little peer reviewer or rubber ducky. 

Tags:

  #til #jetbrains #ai
