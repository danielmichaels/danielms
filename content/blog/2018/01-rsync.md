+++
date = "2018-01-24"
categories = ["administration"]
tags = ["rsync"]
slug = "rsync-cheatsheet"
title = "Rsync cheatsheet"
+++

This is a short primer on the most simple of rsync's capabilities.

Rsync is a fast and extraordinarily versatile file copying tool. It uses
a delta transfer algorithm, which reduces the amount of data sent over
the network by sending only the differences between the source and
destination files. It can contact remote systems via SSH or through a
rsync daemon over TCP.

### Basic Syntax

`-r`: The 'r' flag denotes that the operation is recursive.
It will copy across all files and folders inside the source directory.
However, if the source destination does not have a trailing slash it
will copy the folder, rather than the files inside the folder.

`-a` Archive mode is superior to recursive mode as it will
sustain symbolic links, special and device files, modification times,
group, owner and permissions.

`--dry-run` or `-n` As a sanity check it is
worth checking that your command is going to do what you _think_ it is
going to do. The dry run will not execute the command. It can (read:
should) be coupled with the next command.

`-v`: Verbose will print out the what actions were
undertaken by the command. If `-n` is not coupled with a
verbose flag it will print nothing to the screen.

`--delete`: Will remove extra items in the destination
folder that do not exist in the source directory. **Caution:** This can
lead to complete deletion of the destination folder if incorrectly
implemented.

`-z`: When transferring across the network rsync provides a
compression option to save on bandwidth.

`-P`: Outputs a progress bar to the terminal.

`-H`: Preserves hard links - very useful for back ups and snap shots as files that don't change are often hard linked to the original file.

`--stats`: will output a summary of what was sent. Use `--stats -h` to get a human readable output as seen below.

```sh
# rsync with --stats -h
sending incremental file list

Number of files: 2,265 (reg: 2,264, dir: 1)
Number of created files: 0
Number of deleted files: 0
Number of regular files transferred: 0
Total file size: 100.20M bytes
Total transferred file size: 0 bytes
Literal data: 0 bytes
Matched data: 0 bytes
File list size: 65.50K
File list generation time: 0.001 seconds
File list transfer time: 0.000 seconds
Total bytes sent: 108.52K
Total bytes received: 17

sent 108.52K bytes  received 17 bytes  19.73K bytes/sec
total size is 100.20M  speedup is 923.17
```

`--exclude=important_file.txt`: Can be used to omit files
or directories from being synced.

`--exclude=backups/ --include=backups/most_recent`: Inside
the exclusion we can explicitly include certain file, folders or
patterns that fall inside the broader exclude.


example:

    rsync -azvnP source_dir/ destination_dir/

### Syncing Across Remote Systems

Syncing can be either a "pull" or "push".

_In this example we are_ copying the directory, not just its contents and sending it from the local system to a remote system.

To copy the entire folder and not just the files within, we must **omit** the trailing slash.

The "push":

    rsync -a ~/local_source_dir username@remote_host:/home/username/destination_dir

Here we are retrieving data from a remote server and pulling it down to our local system. Again, we are taking the _entire_ directory so there's no trailing slash.

The "pull":

    rsync -a username@remote_host:/home/username/directory_we_want where_we_want_it_locally

### SSH

Syncing between systems is made much easier if key based authentication
is enabled. If not, the user will be prompted with a password.

To use `rsync` over a `ssh` connection we need to specify the protocol `rsync` needs to use.

`rsync -aznvPH -e "ssh -i ~/ec2_keyfile.pem" user@remote:/home/folder /tmp/local_system/`

The `-e` flag above, tells `rsync` to use a command, or in this case, another protocol to tunnel over.

We wrap the `-e` flag command in quotations to encapsulate our private key - this is particularly useful when accessing AWS. Any `.ssh/config` settings are also respected by `rsync` making commonly accessed systems far easier to use.
