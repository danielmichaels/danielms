+++
title = "Linux Environment Variables & PATH"
slug = "Linux Environment Variables"
date = "11 aug 2018"
+++

# Environment Variables

Simply put, environment variables are a collection of key value strings held by the current shell or process. Each process has access to an array of environment variables held in its user-space memory.

The environment variables are available to all applications, and allow these programs or scripts use them within the shell environment. They can allow custom modification to defaults that permeate throughout the user-space rather than be homed within each applications configuration.

Each process that is created via `fork()`, will inherit a copy of its parent's environment. This ensures the communication and transfer of variables from the parent is a once-only one-way transaction. This is important to prevent updated environment variables from one child polluting the global user-space should it go wrong.

# PATH 

Each environment will have access to several default or standard variables such as:

- PATH
- HOME
- LOGNAME
- SHELL
- EDITOR
- MAIL

By typing `echo $HOME` the users home location will be printed to the terminal. Environmental variables are **case sensitive**, and by convention the key is written in capitals. 

echoing `$PATH` on a unix platform will output a `:` separated list of file paths such as:

```shell
/usr/local/bin:/usr/local/sbin:/usr/bin
```

The `PATH` is set by system and user start-up scripts - the exact process can change between implementations. 
Its function is to search in the `PATH` locations for the name of any program being called for execution. 
In linux everything is file, and when a file is set with the execution bit, that file can be called with the `exec()` method. 
What the `PATH` does is search recursively through each location, from left to right as its printed, looking for that file if its not given with an absolute path. 

## Example

```shell
$ which echo
/bin/echo
$ ls -l /bin/echo
-rwxr-xr-x 1 root root 30k Dec 29 2017 /bin/echo
$ echo $PATH
/home/userName/bin:/usr/local/bin:/usr/bin:/bin  # /bin in PATH
```

Above we see that calling `which echo` returns the location of `echo` (in my distro it returns `echo: shell built-in command` but `/bin/echo` will call the file nonetheless).
Now when we call `echo` which is not an absolute path, the `PATH` variable is consulted. Moving left to right the system checks for an executable file named `echo`, which is finds inside `/bin`. Once found it executes and ceases the search. If we called `/bin/echo` the `PATH` would be searched.

## Example Two

In a real world case, I use Go's `/bin` location as an exemplar of why you may need to add to the `PATH` variable.

Dave Cheney has a pretty sweet Go knock off of an awesome package by the same name `httpstat`. The TL;DR is that it gives you statistic on a site by simply calling `httpstat https://reorx.com`. Astute readers will see that this requires `httpstat` to be on the `PATH` or it will generate an error.

Fixing this was simple as we see below.

# Creation & Deletion


Environment variables are key value pairs, case sensitive and by convention have the key written in caps.

## Creation

Adding to your current shell environments list of variables is simple.

```bash
# method 1
$ VAR=foo       # create key:value
$ export VAR    # export is called on VAR
$ echo $VAR     # echo VAR to terminal
foo

# method 2
$ export VAR=foo # combine export and variable assignment
$ echo $VAR
foo
```

Bourne shells accept the second method, and its probably the most common and convenient method. `export` is unix command to set the attribute for a variable. See `man export` for more information.

## Deletion

```bash
$ unset VAR
$ echo $VAR
# will echo empty string
```

Continuing with the `httpstat` example we will append the location of Go's package binaries to our `PATH`. 

```bash
$ export PATH=~/go/bin:$PATH # now added to PATH
```

Once the variable is exported, it will only live within the current shell. If we close it down, that updated `PATH` is now lost.

To make the variables persistent we add them to our shell's `.profile` or `.bashrc` if using bash. In my case, I use Zsh so its added to `.zshrc`. Calling `source .zshrc` will reload `.zshrc` and make the new variables available if not already. And, because each process inherits form the parent, removing it form the file will prevent it being added to the environment list if needed in the future.

# Security tidbits

Of note linux implements a security measure within the Superusers `PATH`. The current working directory is normally omitted from the `PATH` meaning files must be called using `./` as the prefix or as an absolute path. This prevents malicious users from placing an executable file with the same name as a system-wide executable such as `ls` being run as root.

# Conclusion

- Environment variables allows the system to access a list of values for use in applications
- Users can set and unset variables
- `PATH` is important when you need to access an executable without calling its absolute path