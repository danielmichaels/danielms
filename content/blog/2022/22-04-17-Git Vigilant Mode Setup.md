+++
title = "Github Vigilant Mode Setup"
categories = ["git", "software"]
tags = ["til"]
slug = "Github Vigilant Mode Setup"
date = "2022-04-17"
draft = "false"
ShowToc = "true"
+++

## Setting up Github's vigilant mode 

1. create GPG key

- `gpg --full-generate-key`
- Pick `RSA`
- Enter a key of at least `4096`
- Default expiration
- Enter `git config --global user.name` value
- Enter `git config --global user.email` value
- Add a comment describing it as `Github` or similar

2. Add the key to Github 
   
- Retrieve the key ID and copy it into the clipboard.
- Run the follwing and grab the key ID on the `ssb` line.
    ```shell
    ❯ gpg --list-secret-keys --keyid-format=long dan@danielms.site
    sec   rsa4096/XXXXXXXXXXXXXXXX 2022-04-15 [SC]
          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
    uid                 [ultimate] Daniel Michaels <dan@danielms.site>
    ssb   rsa4096/XXXXXXXXXXXXXX8B 2022-04-15 [E] <-- copy after rsa4096/
    ```

Github expects the Private key file so next retrieve it using `gpg --armor --export XXXXXXXXXXXXXX8B`. Copy the output to the clipboard. 

Paste it into the GPG keys page inside Github using their
[guide](https://docs.github.com/en/authentication/managing-commit-signature-verification/adding-a-new-gpg-key-to-your-github-account)

1. Ensure your commits (and tags) are signed

The easiest way is to setup every commit as signed lest you forget to sign a commit. 

- `git config --global commit.gpgsign true`
- `git config --global user.signingkey ABCDEF01` (where ABCDEF01 is the fingerprint of the key to use)
- `git config --global alias.logs "log --show-signature"` (now available as $ git logs)

## Copying GPG key between devices

If you have more than one device and do not want to have several keys but instead use the same
key across all devices, there are a few steps.

`gpg --list-secret-keys user@example.com`

`gpg --list-secret-keys user@example.com > private.key`

Copy the key to the device using `scp` or similar and then install it into the `gpg` keychain. 

`gpg --import private.key` is touted to work however this did not work for me on Ubuntu
20.04.Instead, I found `gpg --batch --import private.key` worked as expected.

## Deleting tags

If you're like me and push tags without signing them first (fixed by following the above guide)
how do you delete them?

```shell

# delete local tag '12345'
git tag -d 12345
# delete remote tag 
git push --delete origin tagName
```

This is a quick down and dirty on setting up GPG keys, git and GitHub. 
