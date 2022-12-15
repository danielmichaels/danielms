+++
title = "Python rich progress bar and httpx"
categories = ["zet"]
tags = ["zet"]
slug = "Python-rich-progress-bar-and-httpx"
date = "2022-12-15 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Python rich progress bar and httpx

How to create a progress bar using `rich` and `httpx`:

```python
path = "/tmp/file"
url = httpx.URL("https://releases.ubuntu.com/20.04/ubuntu-20.04.3-desktop-amd64.iso")

with httpx.stream("GET", url=url, verify=False, follow_redirects=True) as r:
  with Progress() as progress:
    size = int(r.headers["Content-Length"]
    with open(path, "wb") as dest:
      dl = progress.add_task("[red]Downloading Ubuntu[/red]", total=size)
      for data in r.iter_raw():
        dest.write(data)
        progress.update(dl, advance=len(data), description="[yellow]Downloading...[/yellow]")
      progress.update(dl, description="[green]Download Complete[/green]")
```

Tags:

    #python #httpx

