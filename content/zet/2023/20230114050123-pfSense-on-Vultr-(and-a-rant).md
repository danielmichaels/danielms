+++
title = "pfSense on Vultr (and a rant)"
categories = ["zet"]
tags = ["zet"]
slug = "pfSense-on-Vultr-(and-a-rant)"
date = "2023-01-14 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# pfSense on Vultr (and a rant)

I run test pfSense firewalls on Vultr. It is a pain in ass hand cranking
them each time.

I have a terraform module which can create the instance but not fully. Two issues
arise using terraform.

- Can only create an instance with an OS or an ISO
- pfSense ISO instances break (it cannot detect the disk) and OS installs don't have pfSense

So even though I can create the instance using the OS, then switch it to ISO via
two terraform `apply`'s (there is probably a better one-shot way). It still requires
even more manual steps.

Next problem is once its started you have to manually step through the TUI installer.
After its installed you need to reboot and *manually* remove the ISO in the Vultr 
admin dashboard (*settings>iso*).

Yet another step, after that you have to configure the interfaces via the TUI.

And finally, then go to https://<ipaddress>, login with *admin:pfsene* and run through
another installer.

Now you're ready to actually configure it the firewall itself. Such as SSH rules etc.

A massive pain the arse.

Also a pain the arse; poor terraform errors like:

```terraform
// this error is actually related to the network block about 30 lines
// below the resource at line 1
// this is up there with ansibles useless 'syntax error here or maybe somewhere else'
// errors
 Error: vm '101' not found
│ 
│   with proxmox_vm_qemu.pfsense_node[1],
│   on main.tf line 1, in resource "proxmox_vm_qemu" "pfsense_node":
│    1: resource "proxmox_vm_qemu" "pfsense_node" {
│ 
╵
```

Here is a really helpful guide: https://jarrodstech.net/project-pfsense-setup-on-vultr-with-private-lan/

All this BS makes me wish their was a great infrastructure as code firewall which is
open source. A lot of people will say they'd never want their firewall managed using
IaC but I say they're living in the past.

If it ain't in code, it ain't tracked - spreadsheets, device42 etc aren't the same.

If someone can log on and add or delete a rule and something can't detect that then
its, in my opinion, a liability.

If anyone reads this and wants to educate me, please (and I'm not joking)
email me and let's talk.

Tags:

    #pfsense #rant #terraform #iac
