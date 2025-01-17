+++
title = "Go: simple background task"
categories = ["zet"]
tags = ["zet"]
slug = "Go:-simple-background-task"
date = "2023-04-26 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Go: simple background task

A really simple background worker snippet. Far from perfect!

```go
func background(fn func()) {
	// Launch a background goroutine.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// Recover any panic.
		defer func() {
			if err := recover(); err != nil {
        // zerolog
				logger.Error().Err(fmt.Errorf("%s", err)).Msg("background-task")
			}
		}()

		// Execute the arbitrary function that we passed as the parameter.
		wg.Done()
		fn()
	}()
	wg.Wait()
}
```

Tags:

    #go #snippet
