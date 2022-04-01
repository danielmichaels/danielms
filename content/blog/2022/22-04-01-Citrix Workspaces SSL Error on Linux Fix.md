+++
title = "Citrix Workspaces SSL Error on Linux Fix"
categories = [""]
tags = ["til"]
slug = "Citrix Workspaces SSL Error on Linux Fix"
date = "2022-04-01"
draft = "false"
ShowToc = "true"
+++

TIL: Citrix on linux does not work out of the box.

I don't know why it's like this and honestly don't really care what the excuse it - it feels like a
perfect example of how much care Citrix puts into its products.

This quick post goes over how to get it working.

## Overview

You need to download Citrix's certificate and put it in the ICACClient's keystore - which is
ridiculous.

## Download the Certificate

1. Head to [Citrix](https://www.citrix.com/)'s website.
2. In your browser's navigation bar there is likely a *lock* symbol. Clicking on this should give
   some info about the certificate. Click around here until you end up at a page with the cert's
   information.
   ![](/images/cert-details.png 'view certificate image')
3. In the certificate's *about* page, scroll down until you hit a PEM download button and click it.
   ![](/images/pem-dl.png 'PEM file download')

## Place certificate in Citrix keystore

1. Change the PEM file to a `.crt` using `cp cert.pem cert.crt`.
2. `sudo` up using `sudo su`
3. Change the permission to 644 with `chmod 644 cert.crt`
4. Make owner root with `chown root:root cert.crt`
5. Move this file into the keystore by running `cp cert.crt /opt/Citrix/ICAClient/keystore/cacerts/`
6. Run `/opt/Citrix/ICAClient/util/ctx_rehash` to make Citrix recognise the new cert

Doing these steps should resolve the SSL error and allow you to use Citrix as expected.

**NOTE:** These steps assume Ubuntu 20.04 and `icaclient_22.3.0.24_amd64.deb`

Reference: https://support.citrix.com/article/CTX231524
