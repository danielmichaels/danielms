+++
title = "docker run helper"
categories = ["zet"]
tags = ["zet"]
slug = "docker-run-helper"
date = "2023-01-24 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# docker run helper

I get tired of running `docker run --rm -it <image>:<tag> sh`

```shell
PROGNAME="dockerr"

usage="${PROGNAME} <image> [-h] [-t] [-c] -- docker run helper 

where:
  -h, --help          Show this help text
  -t, --tag           Image tag to use (default: latest)
  -c, --command       Command to execute when container runs (default: none)
      --entrypoint    Override containers entrypoint. This is mutually exclusive with
      --command and will remove it if applied

examples:
  ${PROGNAME} python
  ${PROGNAME} python -t buster -c bash
  ${PROGNAME} hashicorp/vault -t dev --entrypoint sh
```

tag=latest 
executable=""
entrypoint=""

if [[ ${i} != -* ]]
then
  image="${i}"
fi

if [ "$#" -ne 0 ]; then 
  while [ "$#" -gt 0 ]; then 
  do
    case "$i" in
      -h|--help)
        usage
        ;;
      -t|--tag)
        tag="$2"
        ;;
      -c|--command)
        executable="$2"
        ;;
      --entrypoint)
        entrypoint="--entrypoint $2"
        executable=""
        ;;
      --)
        break
        ;;
      -*)
        echo "Invalid option '$1'. Use --help to see arguments" >&2
        exit 1
        ;;
      *)
        ;;
      esac
      shift
    done
  else
    echo "$usage"
    exit 1
fi 

echo "running: docker run --rm -it $entrypoint $image:$tag $executable"
docker run --rm -it $entrypoint $image:$tag $executable
```

This saves me a lot of needless typing. It doesn't handle <image>:<tag> 
but that can be achieved. This works good enough.

Tags:

    #bash #docker 

