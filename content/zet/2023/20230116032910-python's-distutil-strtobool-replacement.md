+++
title = "Python's distutil strtobool replacement"
categories = ["zet"]
tags = ["zet"]
slug = "python's-distutil-strtobool-replacement"
date = "2023-01-16 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Python's distutil strtobool replacement

From 3.10 `distutil.util.strtobool` is deprecated with removal slated for
3.11.

This is a trivial function to replace its functionality and useful for kubernetes
helm values.

```python

def strtobool(value: str) -> bool:
  value = value.lower()
  if value in ("y", "yes", "on", "1", "true", "t"):
    return True
  return False
```

Tags:

    #til #python #helm
