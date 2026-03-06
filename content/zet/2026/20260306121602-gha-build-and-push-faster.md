+++
title = "GHA build and push faster"
categories = ["zet"]
tags = ["zet"]
slug = "gha-build-and-push-faster"
date = "2026-03-06 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# GHA build and push faster

I've been using the same GHA build and push for my local/private/personal projects for years and they take forever.

I asked claude to make it faster and knocked a 15m build to 4m with this:

```yaml
name: Build and Push to GHCR

on:
  push:
    branches: [main]
    tags: ["v*.*.*"]
  workflow_dispatch:

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  build:
    name: Build ${{ matrix.platform }}
    runs-on: ${{ matrix.runner }}
    permissions:
      contents: read
      packages: write
    strategy:
      fail-fast: false
      matrix:
        include:
          - platform: linux/amd64
            runner: ubuntu-latest
          - platform: linux/arm64
            runner: ubuntu-24.04-arm

    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Log in to GHCR
        uses: docker/login-action@v3
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Extract metadata
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}

      - name: Build and push by digest
        id: build
        uses: docker/build-push-action@v6
        with:
          context: .
          platforms: ${{ matrix.platform }}
          push: true
          labels: ${{ steps.meta.outputs.labels }}
          outputs: type=image,name=${{ env.REGISTRY }}/${{ env.IMAGE_NAME }},push-by-digest=true,name-canonical=true
          cache-from: type=gha,scope=${{ matrix.platform }}
          cache-to: type=gha,mode=max,scope=${{ matrix.platform }}

      - name: Export digest
        run: |
          mkdir -p /tmp/digests
          echo "${{ steps.build.outputs.digest }}" > /tmp/digests/${{ matrix.platform == 'linux/amd64' && 'amd64' || 'arm64' }}

      - name: Upload digest
        uses: actions/upload-artifact@v4
        with:
          name: digest-${{ matrix.platform == 'linux/amd64' && 'amd64' || 'arm64' }}
          path: /tmp/digests/
          if-no-files-found: error
          retention-days: 1

  merge:
    name: Merge manifests
    runs-on: ubuntu-latest
    needs: build
    permissions:
      contents: read
      packages: write

    steps:
      - name: Download digests
        uses: actions/download-artifact@v4
        with:
          path: /tmp/digests
          pattern: digest-*
          merge-multiple: true

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Log in to GHCR
        uses: docker/login-action@v3
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Extract metadata
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
          tags: |
            type=ref,event=branch
            type=semver,pattern={{version}}
            type=semver,pattern={{major}}.{{minor}}
            type=sha,prefix=sha-,format=short

      - name: Create and push manifest
        run: |
          DIGESTS=$(cat /tmp/digests/amd64 /tmp/digests/arm64 | xargs -I{} echo "${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}@{}")
          TAGS=$(echo "${{ steps.meta.outputs.tags }}" | xargs -I{} echo "--tag {}")
          docker buildx imagetools create $TAGS $DIGESTS
```

Instead of using QEMU.

Tags:

    #til #cicd
