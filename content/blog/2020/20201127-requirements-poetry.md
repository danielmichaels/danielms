
+++
title = "From requirements.txt to poetry's pyproject.toml"
categories = ["python"]
tags = ["python", "poetry"]
slug = "requirements-txt-to-poetry-pyproject-toml"
date = "2020-10-27"
draft = "false"
+++

# Poetry and Requirements.txt

I am now a poetry convert, opting for it in any reasonably large projects. Personally, I have found the ability to pin pacakges and
produce a `lock` file invaluable for getting a complete picture of an applications dependencies. It also prevets needless docker build time 
spikes if a package is updated. 

Unfortunately, converting `requirements.txt` to `pyproject.toml` is not yet integrated natively by `poetry`.

## The problem

I needed to change a `requirements.txt` file of this structure:

```bash
-r base.txt
django-debug-toolbar==3.1.1  # https://github.com/jazzband/django-debug-toolbar
django-extensions==3.0.9  # https://github.com/django-extensions/django-extensions
django-coverage-plugin==1.8.0  # https://github.com/nedbat/django_coverage_plugin
pytest-django==4.1.0  # https://github.com/pytest-dev/pytest-django
... truncated ...
```

to a `pyproject.toml` file like this:

```bash
[tool.poetry.dependencies]
python = "^3.8"
gunicorn = "^20.0.4"
psycopg2 = "^2.8.6"
sentry-sdk = "^0.19.3"
django-storages = {extras = ["boto3"], version = "^1.10.1"}
django-anymail = {extras = ["sendgrid"], version = "^8.1"}
pytz = "^2020.4"
python-slugify = "^4.0.1"
... truncated ...
```

## The solution

All it took was one line:

`cat requirements.txt | grep -E '^[^# ]' | cut -d= -f1 | xargs -n 1 poetry add`

It will loop over each line and strip out any `#` then call `poetry add` followed by the name of each package. Test it by calling `cat requirements.txt | grep -E '^[^# ]' | cut -d= -f1` to see what it will output before trying to `poetry add` it.

 If it encounters something like `-r base.txt` which is not a package, poetry will throw an error to stdout but continue looping over the file. If you're really keen, you could strip lines starting with `-r` but why bother?

The caveats, this will install the latest version of each package because `xargs` will only return the package name and not any accompanying `==n.n.n`. Again, this could be extended if needed. 

To install development dependancies the line can be amended with a `-D` or `-dev`. 

## Poetry for the win

If you are considering switching to poetry but have many legacy `requirements.txt` files for various build states, its now almost trivial
to switch. If you're on docker, poetry or other tools such as `pipenv` or `pip-tools` are going to potentially save you a lot of lost time from 
your pip layer caches being busted when an unpinned subpackage gets updated. Github's dependabot can also interpret `pyproject.toml` files meaning you will get automated pull requests for package updates just like you would with a `requirements.txt` file. 

Get poetry [here](https://python-poetry.org/docs/#installation).
