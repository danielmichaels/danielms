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
worth checking that your command is going to do what you *think* it is
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

`--exclude=important_file.txt`: Can be used to omit files
or directories from being synced.

`--exclude=backups/ --include=backups/most_recent`: Inside
the exclusion we can explicitly include certain file, folders or
patterns that fall inside the broader exclude.

example:

    rsync -azvnP source_dir/ destination_dir/

### Syncing Across Remote Systems

Syncing can be either a "pull" or "push".

The "push":

    rsync -a ~/local_source_dir username@remote_host:/home/username/destination_dir

*In this example weare*\* copying the directory, not just its contents
so we omit the trailing slash.\* Here we are sending data from the
source to the destination.

The "pull":

    rsync -a username@remote_host:/home/username/destination_dir local_source_dir

This pull operation is syncing the remote directory with the local
system. We might use this to backup a small database file for instance.

### SSH

Syncing between systems if made much easier if key based authentication
is enabled. If not, the user will be prompted with a password.
