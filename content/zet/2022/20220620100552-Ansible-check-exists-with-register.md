+++
title = "Ansible check exists with register"
categories = ["zet"]
tags = ["zet"]
slug = "Ansible-check-exists-with-register"
date = "2022-06-20 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Ansible check exists with register

In building my Ubuntu disaster recovery ansible script, I discovered the
`register` condition.

To prevent needless tasks from running, for instance, `shell` commands 
which would be redundant. It is possible to check a condition exists
and only execute the task based on the response.

```yaml
- name: Download GH cli install script
  stat:
    path: gh
  register: gh_exists

- name: Install gh cli
  become: true
  when: gh_exists.stat.exists == False
  ansible.builtin.get_url:
    url: https://raw.githubusercontent.com/danielmichaels/dot/master/installers/install-gh
    dest: /tmp/gh.sh
  tags:
    - install
    - productivity

- name: Install GH Cli
  become: true
  when: gh_exists.stat.exists == False
  ansible.builtin.shell:
    cmd: cat /tmp/gh.sh | sh -s -- -y
  tags:
    - install
    - productivity
```

The above code block installs the `gh` CLI tool. It executes a script
from my github `dot` repo only if the `gh` binary **does not** exist in 
the `PATH`.

I also achieved the same thing for installing Go. Using this method I
didn't need to install a `ansible-galaxy` third party role. 

```yaml
# vars:
#    go_version: "1.18.3"
- name: Check go exists and version
  stat:
    path: /usr/local/go/bin/go
  register: go_exists

- name: Remove existing /usr/local/go
  when: go_exists.stat.exists == False
  become: true
  ansible.builtin.file:
    path: /usr/local/go
    state: absent

- name: Install Golang
  when: go_exists.stat.exists == False
  become: true
  ansible.builtin.unarchive:
    src: "https://go.dev/dl/go{{ go_version }}.linux-amd64.tar.gz"
    dest: /usr/local
    remote_src: yes
```

It might not be elegant but it works well for me.

Tags:

    #ansible #productivity
