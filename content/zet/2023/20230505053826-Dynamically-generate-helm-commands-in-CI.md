+++
title = "Dynamically generate helm commands in CI"
categories = ["zet"]
tags = ["zet"]
slug = "Dynamically-generate-helm-commands-in-CI"
date = "2023-05-05 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Dynamically generate helm commands in CI

We run `helm` for our deployments (eventually switching to Argo thankfully)
and it works OK. What it doesn't do well in is re-usability in CI.

We have a project which all other projects inherit from for CI and it works
pretty well. It takes some organisation but once its done you can basically
set and forget your deployments. Except that helm sucks for that and 
we ended up with each project needing a bespoke `helm-deploy` job.

This week I sat down and designed a better solution using python templating.
Next week I think I'll re-do it in Go because its in the standard lib.

Here's how I'm doing it python.

```python
import os
import re

from jinja2 import Template

helm_set_prefix="HELM_SET_VALUE_"
helm_set_file_prefix="HELM_SET_FILE_"
helm_values_prefix="HELM_VALUES"

def prefix(re_prefix: str):
  return re.compile(r"{prefix}\w".format(prefix=re_prefix))

helm_set = {key.replace(helm_set_prefix,"").replace("_","."):val for key, val in os.environ.items() if prefix(helm_set_prefix).match(key)}
helm_set_file = {key.replace(helm_set_file_prefix,"").replace("_","."):val for key, val in os.environ.items() if prefix(helm_set_file_prefix).match(key)}
# this is just an array
helm_values = [v for k,v in os.environ.items() if k.startswith(helm_values_prefix)]

template = """
#!/bin/bash 

helm upgrade --install {{ release }} {{ chart_path }} \
  {%- if helm_set -%}
    {%- for k,v in helm_set.items() -%}
  --set {{ k }}={{ v }} \
    {% endfor %}
  {% endif %}
  {%- if helm_set_file -%}
    {%- for k,v in helm_set_file.items() -%}
  --set-file {{ k }}={{ v }} \
    {% endfor %}
  {% endif %}
  {%- if helm_values -%}
    {%- for v in helm_values -%}
  --values {{ v }} \
    {% endfor %}
  {% endif %}
  --atomic --timeout {{ timeout }}
"""

data = {
  "timeout": os.getenv("TIMEOUT", "300s"),
  "release": os.getenv("RELEASE"),
  "chart_path": os.getenv("CHART_PATH", "helm"),
  "helm_set": helm_set,
  "helm_set_file": helm_set_file,
  "helm_values": helm_values,
  }

output = Template(template)

# run me like so:
# python main.py > script.sh
print(output.render(data))

```

I'll bypass the how of getting this going in CI jobs but to pass values
its as simple as providing environment variables with prefixes.

```yaml
job:
  variables:
    RELEASE: MyRelease
    HELM_SET_db_name: foobar
    HELM_SET_FILE_cacert: ca.crt
    HELM_VALUES: helm/values.yml
```

Which would create a bash script which looks like

```bash
helm upgrade --install MyRelease helm \
  --set db.name='foobar' \
  --set-file cacert='ca.crt' \
  --values 'helm/values.yml' \
  --atomic --timeout 300s
```

This script can then be executed in another job as needed when output
as an artifact which makes its highly extensible.


Tags:

    #python #ci #helm
