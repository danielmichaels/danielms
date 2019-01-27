+++
title = "Uploading to PyPI"
slug = "Uploading-to-PyPI"
date = "2018-04-07"
tags = ["python"]
categories = ["TIL", "python"]
+++

**UPDATE 27/01/2019: This area of python and the PyPI is under going rapid
development and as such the following may have parts which are no longer in date.
In time I will update this post to reflect these changes.**

Python Package Index
--------------------

Python has a wonderful community and package ecosystem. It currently has
over 130,000 packages for download and a large variety to choose from.
To download a python package via the `pip` command, the
package must be uploaded to the Python Package Index, or
[PyPI](https://pypi.org/). Going forward it may be referred to as the
"warehouse" as PyPI is going through an upgrade of its infrastructure
and for the better.

### Have application, now what?

PyPI offers two servers for the uploading of python packages; testing
and production.

Sending the package to the test server is a great idea as it allows you
to download your tarball onto any system for testing. This allows you to
do a few novel things like sharing it with friends or co-workers,
spinning up virtual machine's with different operating systems or
installing it into separate virtual environments with different versions
of python.

This article does assume you are using a version control system, and in
particular GitHub but this is not a requirement for PyPI.

### Productionise your code

Before looking at how to upload your modules, first it must be made
ready for release into the wild.

### Directory Structure

```bashj
root-dir/           # The directory which all your files live.
    setup.py        # covered below (Required)
    setup.cfg       # if using markdown rather than ReStructuredText
    LICENSE.txt     # should be required!
    README.md       # Also should be required!
    tests/          # tests are a good idea
        test.py
    your-package/
        __init__.py
        awesome.py
        wicked.py
```

If you look at any great package such as
[Requests](https://github.com/requests/requests) or
[Glances](https://github.com/nicolargo/glances) you will see a similar
(although much more intricate) structure. The key files we **need** are
`setup.py` and `setup.cfg` if using Markdown.

### Setup.py

``` python
# setup.py

from codecs import open
# ensure consistent encoding
from setuptools import setup
# always prefer over distutils
from os import path

VERSION = '0.1.0'
URL = 'https://github.com/username/package'
DOWNLOAD_URL = (URL + '/tarball/' + VERSION)

here = path.abspath(path.dirname(__file__))

with open(path.join(here, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
    name='yourpackage',
    packages=['yourpackage'],
    version=VERSION,
    description='blurb that users first see to decide if interesting',
    long_description=long_description,
    long_description_content_type='text/markdown',
    author='Optional',
    author_email='Optional',
    url=URL,
    download_url=DOWNLOAD_URL,
    classifiers=[
        #   3 - Alpha
        #   4 - Beta
        #   5 - Production/Stable
        'Development Status :: 3 - Alpha',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6', ]
)
```

The information contained within `setup.py` and in
particular its `setup()` function is what creates the
package's metadata for parsing by PyPI once uploaded.

This is a stripped down version of my own `setup.py`. Many
tutorials which are older will use `distutils` but PyPI's
[example](https://github.com/pypa/sampleproject) structure explicitly
states to favour the newer `setuptools`.

By using a context manager and the `codecs.open` method we
can read the README.md file for use in the
`long_description` parameter within `setup()`.
This is a requirement for the PyPI server as it by default only parses
ReStructuredText. If using README.rst, this can be ignored. Further, the
`long_description_content_type='text/markdown'` must be
included or it will not format the content correctly. This is a very
recent addition to PyPI - March 2018.

Also, the `version` is what sets the tarball filename, so
creating an easy to adjust global variable makes it a lot easier to
amend this file when updating your package. As you can see there is
three locations that need to be updated for it to function correctly.
(Thanks to [Dan Bader] (https://dbader.org) for the
idea.)

Please refer to the example page for more information, particularly if
your package is more complex than just a few modules.

### Setup.cfg

This is just required for Markdown parsing. As you may see it might just
be easier to utilise the default supported .rst files. Something I may
do in the future.

```bash
[metadata]
description-file = README.md
```

### Python setup.py sdist

Running this command will invoke the `setup.py` and create
a folder called `dist/` inside your root directory. This is
where your application's tarball will now live. It is also a good time
to create or update your git tags for your repository.

FYI, once you create a local tag it must be pushed to the remote.

```bash
git tag X.Y.Z -m "Add a message such as; First!"
git push X.Y.Z      # preferred option
git push --tags     # less preferred as it pushs ALL tags to the remote server
```

### Upload: y u no easy

In theory uploading to PyPI is just that simple. Unfortunately it isn't
that easy and a lot of the helpful blogs and references out there aren't
current with the new standards. As always the official user guide
([here](https://packaging.python.org/tutorials/distributing-packages/#uploading-your-project-to-pypi))
is the holy grail but isn't the easiest reading when completely unsure!

### 1. Register

To upload anything to PyPI you must first register with it. And if you
want to make use of the testing server, you must register with it
separately. Although they use the same software, each server uses a
separate database and this is why two different sign up's are required.
They can be found here for [Live](https://pypi.org/account/register/)
and [Test](https://test.pypi.org/account/register/).

### 2. Create .pypirc

This file allows your development machine to talk to the PyPI servers.
It should look something like this.

```bash
# ~/.pypirc

[distutils]
index-servers = 
  pypi
  pypitest

[pypi]
repository=https://upload.pypi.org/legacy/
username=username

[pypitest]
repository=https://test.pypi.org/legacy/
username=username
```

This is current to today's date but the repository url may change as the
PyPI warehouse continues its evolution. The file **must** be located in
the home directory. Both username and password can be set in this file,
or in environment variables.

### 3. Install Twine

What is it and why use it? Twine is a package written by the PyPI
maintainers that uses SSL by default when sending information to their
endpoint. Python versions before 2.7.9 and 3.2 do not use this by
default and spill user credentials over the air. Also twine separates
the creation of the package tarball and uploading into two logical
commands; setuptools does not - they are done in the same invocation.

Tarball? To send your package to the server it first must be compressed
into a single file. So basically, you zip your files to send and when
`pip install xxxx` is called your tarball is downloaded and
unzipped at the end user. This is an important point because any changes
you make after creating your tarball are not included in the package so
you will need to rezip it to include them.

Personally, I install twine on the system interpreter and update it
along with setup tools frequently.

### 4. Upload!

```bash
twine upload -r pypitest dist/package-version-i-select-explicitly.0.1.0.tar.gz
>>> Uploading distributions to https://test.pypi.org/legacy/
>>> Enter your password:
>>> Uploading package-version-i-select-explicitly.0.1.0.tar.gz
```

Using the `-r` flag allows you to set which server to send
the file too. This file name is setup in the `.pypirc` file
and if you have setup a username it will not prompt you for it.
Likewise, it will not prompt for a password should you choose to enter
that, and if you do consider `chmod 600` on that file for
security reasons.

In many examples you may see something like this:

```bash
twine upload -r pypitest dist/*
```

This will upload all of your tarball's located in the
`dist/` directory. I personally choose which distribution
to upload. Either, or. Once done, goto the test PyPI and check to see
that it looks as expected, or as previously stated download the test
file and check its functioning as expected.

This can be done like so:

```bash
pip install --index-url https://test.pypi.org/simple/ yourPackageName
```

Once happy send that baby to the production PyPI by repeating the
commands but this time specifying the `pypi` server like
so:

```bash
twine upload -r pypi dist/package-version-i-select-explicitly.0.1.0.tar.gz
```

You now have a production package in the wild. Check out your code at
[Libraries.io](https://libraries.io/).

### 5. Help PyPI

None of this would be possible with the tireless work of the Python
Software Foundation and the handful of volunteers that make PyPI a
reality. Join the PSF and maybe consider a donation, or convince your
employer to contribute if they rely on python software! Without PSF and
PyPI we wouldn't have python as we know it today. Please visit and sign
up here: [PSF](https://psfmember.org/)
