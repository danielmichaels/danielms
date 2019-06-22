#!/bin/bash
# Script that generates Hugo metadata.
# Takes one CLI  argument of FILENAME.

DATA="
+++
title = \"\"
categories = [\"\"]
tags = [\"\"]
slug = \"\"
date = \"$(date +'%Y-%m-%d')\"
draft = \"true\"
+++
"

DATE=$(date +"%d%m%Y")
FILENAME=$DATE'-'$1.md
FILE=$FILENAME

if [ $# -eq 0 ]
  then
    echo "Missing filename argument"
    exit
fi

if [ -e $1.md ]; then
  echo "File Aleady Exists"
else
  touch $FILENAME
  echo "$DATA" > $FILE
  echo "Created $FILE"
fi
