
+++
title = "Dropbear and AWS"
categories = ["AWS", "ssh"]
tags = ["ssh"]
slug = "dropbear-aws-pem-keys"
date = "2019-09-10"
draft = "false"
+++

# BusyBox, SSH and EC2

![terminal](/images/bash-terminal.png 'terminal graphic')

Accessing an EC2 instance from BusyBox's Dropbear SSH client isn't easy. Firstly, `.pem` files are not compatible with `dropbear`, nor can you convert them to dropbear's key format with the built-in `dropbearconvert`. Secondly, depending on your version of `openssh` it may not be immediately apparent that your private keys are incompatible with the conversion application either. Thankfully, workarounds are possible. 

## Dropbear SSH

Dropbear is a lightweight client and server application mostly seen on embedded devices. It is designed to replace OpenSSH in low memory footprint systems as it can be compiled down to [110kb][1]. 
It is compatible with `.ssh/authorized_keys`, however it does have limitations.

## Those limitations

Amazon Web Services such as Elastic Cloud Compute hand out private keys for passwordless login. Unfortunately for dropbear users, the format is `.pem` which is incompatible. And, `dropbearconvert` - the program provided for turning `openssh` keys into `dropbear` keys - does not accept `.pem` files.

Dropbear's conversion tool will also not accept private keys from `openssh`, as seen below.

```sh
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABlwAAAAdzc2gtcn
NhAAAAAwEAAQAAAYEAtNY9pkKAhGYD/qDIThmRd7Y8kMo7QvbSQj+tv5huz4hTqdgALmxP
....
FBuu7Rt2qQdJPBAAAADGRhbmllbEBBbm9hdAECAwQFBg==
-----END OPENSSH PRIVATE KEY-----
```

Instead, it must be an `rsa` key like so.

```sh
-----BEGIN RSA PRIVATE KEY-----
VY2Q002wQjJzfA783q0wPwPgdQVNBj8timSYHTmZLlZ54pPtBLhMvZ4tJ/AeXxSm
MIIEpQIBAAKCAQEA4N5r5Z+/rl2lmNdxsmcqyhfZ49m1g/5mIMSdPbTXgKcn2T3o
...
pEJt+8fBAoGBAM2KBHEA5RFnv812nGJG6f2scaMxufbQh5vtc0tf7DDAPqmHlnqr
-----END RSA PRIVATE KEY-----
```

## The workaround

I found the most reliable means to gain access to EC2 from `dropbear` on `busybox` was to create the keys elsewhere and then move them to the device manually.

### The steps

#### 1. Create a public and private key on either the AWS EC2 instance, or your local machine with:
- `ssh-keygen -m PEM -t rsa -b 4096 -f <new_privkey>`


#### 2. Copy the _private_ key to BusyBox:

- `scp <priv_key_filename> <user>@<busybox>:/<user>/.ssh/`

#### 3. Copy the private key to the cloud - _only if doing these steps on your local system_:

- `scp -i <aws-ec2-private-key>.pem <new_privkey> ec2-user@<ec2-dns-address>:/home/<ubuntu>/.ssh/`

#### 4. On the system you are running these steps, inside `.ssh/authorized_hosts` append the public key:

__from local machine__:

- `cat <new_pubkey>.pub | ssh -i <aws-ec2-private_key> | "cat >> /home/ubuntu/.ssh/authorized_keys"`

__inside ec2__:

- `cat <new_pubkey>.pub >> .ssh/authorized_keys`

#### 5. This will convert our rsa key into a dropbear compatible key 

- Login to busybox
- `dropbearconvert openssh dropbear <new_privkey> <new_privkey_dropbear_compat>`

#### 6. Login in to AWS

- `ssh -i <new_privkey_dropbear_compat> <ec2-user>@<ec2-dns-address>`
- profit

## Get to work

We can now access our AWS server from our `busybox` device. From here we can do whatever we need to do such as establish a reverse tunnel for administration 😉

[1]: https://lists.ucc.gu.uwa.edu.au/pipermail/dropbear/2004q3/000022.html
