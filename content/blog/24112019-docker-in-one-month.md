+++
title = "learn docker in one month"
categories = ["docker", "development"]
tags = ["docker"]
slug = "learn-docker-in-one-month"
date = "2019-11-24"
draft = "false"
+++

# Learning Docker in one month

![](/images/docker.png "docker image")

Yep, [one month][1].

## Docker basics

### What

docker is a product that uses operating system virtualisation to created packages and applications which live within a "container".
These containers bundle together applications and dependancies without the need for installation on the host operating system. Containerisation is meant to empower developers by providing a "written once and run anywhere" methodology.

### Why

Some of the core reasons to use docker.

- The ability to have standalone applications and packages with dependencies that are isolated for the base operating system and other applications is a boon for developers and users,
- Linux hosts can have several docker containers running without fear of contaminated package versions or dependancies causing errors,
- Sharing consistent environments between developers is possible as each docker container image is idempotent, meaning the result of building a docker image will be the same whether built once, or one thousand times,
- Version pinning is possible making compatibility issues on new release cycles less problematic, and testing of packages easier to implement,
- Continious Integration is a perfectly paired with docker allowing each new release, or push to the repository to be built automatically with a new docker image, and tested to ensure against regression.

### Virtual Machines aren't docker

Docker and virtual machines differ in their implementation and uses cases. They are complimentary rather than competitors and service different needs through each use-case.

![](/images/dockervm.png "docker versus virtual machines")

Containers leverage the base operating systems kernel whereas virtual machines initialise each machine with its own hardware virtualised operating system. Although docker utilises the host's operating system, its application are seperated from the kernel ensuring isolation via the docker engine. 

Regardless, each technology should be utilised as a method to achieve an outcome rather than be a dogmatic position for any one developer to live and die by. In fact, both methods can be used in conjunction, it is possible to run docker contianers within virtual machines - this being something I do often in vSphere. 

## Docker processes

An important distinction with docker is its process orientation. Containers are designed to execute one process and then quit - the length of the running process arbitary but once its is complete the container will effectively exit. 

this is an important point for new docker users. It contrasts heavily from a virtual machine which will live forever, sitting idly doing nothing until instructed.

For example, a contianer may run a command line application that fetches the weather from your local area. Once the container has finished this process it will stop and cease to draw any resources until manually called again by the user or application.

## Docker commands

`docker container run -d -p 8080:80 --name myapache httpd`
`docker container run -it -p 8080:80 --name myapache httpd`

- ps
- image ls
- container ls
- pull
- run
- flags
- rm
- prune
- exec `docker exec -it myapache bash`

Short running container

```sh 
docker run \    # run docker 
--rm \          # remove container once done
-it \           # interactive mode
-v $(pwd):/src python:3 \ # volume to mount with container
python hello.py # process to run; python and script located in the volume - /src
```

Long running container

```sh 
docker run \    # run docker 
--rm \          # remove container once done
-it \           # interactive mode 
# replace -it with -d to detach and run in background
-v $(pwd):/usr/share/nginx/html \ # volume to mount with container
-p 8080:80  \  # port host:container
nginx:latest \ # image to run with tag (latest version)
```

bind routes `-v` put area on local to area on container. `docker run -d -p 8080:80 --name nginx-bind -v $(pwd):/var/share/html nginx`

## Dockerfile

- FROM nginx:latest
- COPY . .
`docker build`
```
[COMMAND] [Argument]
FROM ubuntu:latest # must contain a FROM
# Reads sequentially as layers
```

To allow user input 
> The ENTRYPOINT specifies a command that will always be executed when the container starts. The CMD specifies arguments that will be fed to the ENTRYPOINT
```
FROM debian:wheezy
ENTRYPOINT ["/bin/ping"] # executes when container is run
CMD ["localhost"]        # executes unless user input supplied
```

## docker networks

create a network of `mysql` and `node`: `docker network create <name>`

creates DNS between containers

```sh
# create mysql on docker network named 'test'
docker run --rm -d --net test --name test_mysql -e MYSQL_ROOT_PASSWORD='root' mysql:5.6
# create node linking to the 'test' network and spawn a shell
docker run --rm -it --net test --name test_node node:8 /bin/bash
# once they've been built it will drop into a shell
# test the network connection by trying to ping test_mysql
ping test_mysql # from node /bin/bash shell
# unless explicitly denied, this gives access to internet
ping 1.1.1.1
```

## layers

In a Dockerfile, each line is a layer. 

- layers cannot shrink the size of the image 
- layers are cached (checks the Dockerfile lines for differences)
- caching means it does not need to be rebuilt, only new layers
- container is a layer, one that exists only while the container lives:
    - the build layers are read only - cannot edit them and are shared across many containers
    - if the image needs updating, all layers must be rebuilt

- container layer:
    - read write
    - modifications of source, or volumes within the container exist only within __that__ container unless rebuilt into a new image

- image layer:
    - read only


## docker volumes

`docker volume create <name_of_volume>`

[1]: https://www.norvig.com/21-days.html
