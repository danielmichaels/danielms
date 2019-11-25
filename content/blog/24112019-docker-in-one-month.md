
+++
title = "learn docker in one month"
categories = ["docker", "development"]
tags = ["docker"]
slug = "learn-docker-in-one-month"
date = "2019-11-24"
draft = "true"
+++

# Learning Docker in one month

[!] insert pics of docker in 5 min etc videos [!]

why one month? link to norvig learn prog in ten years

## Docker basics

[!] insert docker vm/container image [!]

- what
- why
- vm v container
- origin containers
- what can I do with it?

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

