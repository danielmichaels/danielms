+++
title = "minio for local s3 testing"
categories = ["zet"]
tags = ["zet"]
slug = "minio-for-local-s3-testing"
date = "2024-06-27 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# minio for local s3 testing

How to setup minio for use in local development.

```yaml
services:
  minio:
    image: quay.io/minio/minio
    container_name: minio-local
    volumes:
      - minio_data:/data
    ports:
      - "9090:9090"
      - "9000:9000"
    environment:
      MINIO_ROOT_USER: root
      MINIO_ROOT_PASSWORD: password
    command: "server /data --console-address :9090"
volumes:
  minio_data: {}
```

The `minio` profile will point any s3 data to the `s3-specific` service which
uses our local `minio` container as the endpoint_url.

You'll also need to `aws configure --profile minio` and enter the API keys from
`minio`. To create them in the UI go to `Access Keys`.

```
[profile minio]
services = s3-specific
[services s3-specific]
s3 =
  endpoint_url = http://localhost:9000
```

In the code we have to create and load some environment variables so that it
won't use the defaults

```environment
AWS_PROFILE=minio
AWS_ENDPOINT_URL_S3=http://localhost:9000
```

And finally in we need to create a bucket in `minio` called `my-bucket` or
whatever you like. I just do that in the UI.

### AWSWrangler

Once those items have been established, its possible to use the `minio` profile
for `awswrangler`.

The best way to do this without making any code changes is to use environment
variables.

In addition to the variables above also add
`AWS_ENDPOINT_URL=http://localhost:9000`.

All together it will look like:

```environment
AWS_PROFILE=minio
AWS_ENDPOINT_URL=http://localhost:9000
AWS_ENDPOINT_URL_S3=http://localhost:9000
```

This will work on most s3 projects

Tags:

    #s3 #aws #development
