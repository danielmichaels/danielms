+++
title = "reproducible random strings in python"
categories = ["zet"]
tags = ["zet"]
slug = "reproducible-random-strings-in-python"
date = "2023-02-16 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# reproducible random strings in python

```python
import hashlib
import uuid

seed = 'Type your seed_string here' #Read comment below

m = hashlib.md5()
m.update(seed.encode('utf-8'))
new_uuid = uuid.UUID(m.hexdigest())
```

I used this to create kube resources programatically whilst keeping
their names consistent for easier deletion and management.

Tags:

    #python #kubernetes

