+++
title = "Deploy to Netlify using Github Actions"
categories = ["zet"]
tags = ["zet"]
slug = "Deploy-to-Netlify-using-Github-Actions"
date = "2024-10-22 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Deploy to Netlify using Github Actions

After switching some projects from my personal account into a new GitHub
organisation my Netlify automated deploys started failing.

To use a private organisation you have to pay $19/USD per month. Considering its
the 22nd and I've used 5 of my 300 free build minutes they can get stuffed.

So I had to switch from automated to triggering the deployment using their CLI
in GitHub actions.

Honestly I just ask CodyAI (sourcegraph's AI tool) to convert my `hugo.toml`
file into a GitHub actions template and it got me 90% done. I had to tweak a
couple things (switch to pnpm for example)

Here's the file,

```yaml
name: Netlify Deploy

on:
  workflow_dispatch:
  pull_request:
  push:
    branches:
      - main

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Repository
        uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: 20

      - name: Install pnpm
        uses: pnpm/action-setup@v2
        with:
          version: 8

      - name: Get pnpm store directory
        shell: bash
        run: |
          echo "STORE_PATH=$(pnpm store path --silent)" >> $GITHUB_ENV

      - name: Setup pnpm cache
        uses: actions/cache@v3
        with:
          path: ${{ env.STORE_PATH }}
          key: ${{ runner.os }}-pnpm-store-${{ hashFiles('**/pnpm-lock.yaml') }}
          restore-keys: |
            ${{ runner.os }}-pnpm-store-

      - name: Setup Hugo
        uses: peaceiris/actions-hugo@v2
        with:
          hugo-version: "0.134.3"
          extended: true

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: "1.23"

      - name: Install Dependencies
        run: pnpm install

      - name: Project Setup and Build
        run: |
          pnpm project-setup
          pnpm build

      - name: Deploy to Netlify
        uses: nwtgck/actions-netlify@v2.0
        with:
          publish-dir: "./public"
          production-branch: main
          github-token: ${{ secrets.NETLIFY_GITHUB_TOKEN }}
          deploy-message: "Deploy from GitHub Actions"
          enable-pull-request-comment: false
          enable-commit-comment: true
          overwrites-pull-request-comment: true
        env:
          NETLIFY_AUTH_TOKEN: ${{ secrets.NETLIFY_AUTH_TOKEN }}
          NETLIFY_SITE_ID: ${{ secrets.NETLIFY_SITE_ID }}
        timeout-minutes: 1
```

And its working perfectly!

Tags:

    #netlify #github #cicd
