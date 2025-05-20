+++
title = "Makefile check args are passed in"
categories = ["zet"]
tags = ["zet"]
slug = "makefile-check-args-are-passed-in"
date = "2023-02-12 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Makefile check args are passed in

TIL how to get a Makefile to abort if arguments are not supplied.

```shell
## sops/decrypt
.PHONY: sops/decrypt
sops/decrypt:
	@test $(file) || (echo "file= not set" ; exit 1)
	@test $(regex) || (echo "regex= not set" ; exit 1)
	@echo "decrypting ${file} with regex (data|${regex})"
	@sops --decrypt --age $$(cat "${SOPS_AGE_KEY_FILE}" | grep -oP "public key: \K(.*)") --encrypted-regex "^(data|${regex})$$" --in-place "${file}"
```

This will error with `regex= not set` if not supplied.

A simple solution that works *good enough*

Tags:

    #TIL #Makefile
