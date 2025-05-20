+++
title = "python Protocols are Go interfaces"
categories = ["zet"]
tags = ["zet"]
slug = "python-protocols-are-go-interfaces"
date = "2022-07-22 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# python Protocols are Go interfaces

This week I successfully used python's [Protocol](https://peps.python.org/pep-0544/)
to fake a dependency on Hashicorp's [Vault](https://www.vaultproject.io).

Background:
  - we need Vault access to save some data
  - init and sidecar must be running for the pod to be created
  - in development we don't need access to the real vault
  - in development access to the real vault crosses network boundaries via indirection

Instead, we have decided to use a *verified fake* and build an interface which 
we can use to swap out the real and fake Vault when required. 

After a suggestion from a colleague about Protocol, and some light reading,
this process was super simple. It works almost exactly the same as a Go interface
whereby you define the interface and other classes implement it.

Example:

```python

class SecretStore(Protocol):
  def save(self):
    pass

class VaultStore:
  def save(self):
    print("vault store save")

class FailStore:
  def no_save(self):
    print("not implemented")

class Saver:
  def do_save(self, obj: SecretStore):
    return obj.save()

def main():
  saver = Saver()
  saver.do_save(VaultStore())
  saver.do_save(FailStore()) # <-- AttributeError: BrokenStore object has no attribute 'save'

if __name__ == "__main__":
  main()
```

Tags:

    #python #interfaces #programming
