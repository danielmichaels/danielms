+++
title = "k3s: tuning local development"
categories = ["zet"]
tags = ["zet"]
slug = "k3s:-tuning-local-development"
date = "2022-12-19 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# k3s: tuning local development

A couple of tunables for single node master/agent
k3s. 

*Fix high CPU on gnome*

Higher usage on Gnome/KDE can be caused by the large
number of overlay volumes.

`systemctl stop --user gvfs-udisks2-volume-monitor`

*Disable leader election*

For a single node development environment removing leader
election can reduce CPU usage.

```shell
cat << EOF | sudo tee /etc/rancher/k3s/config.yaml
kube-controller-manager-arg:
  - "leader-elect=false"
  - "node-monitor-period=60s"

Tags:

    #k3s #kubernetes #TIL
