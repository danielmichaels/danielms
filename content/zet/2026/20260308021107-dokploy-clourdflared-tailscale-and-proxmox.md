+++
title = "Dokploy Clourdflared Tailscale and Proxmox"
categories = ["zet"]
tags = ["zet"]
slug = "dokploy-clourdflared-tailscale-and-proxmox"
date = "2026-03-08 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Dokploy Clourdflared Tailscale and Proxmox

I am pretty happy with this little experiment:

Can I host Dokploy inside my LAN using Proxmox and have an application auto-deploy on git push from GitHub?

Well, yes.

---

Installing Dokploy isn't covered here but I just following the PVE Helper Scripts and installed it into a Ubuntu LXC (and then snapshotted it).

I hooked up GitHub and a particular repo which has a docker-compose file. Then deployed it using that fille. Note that I don't build this but instead pull from <ghcr.io> - this saves a lot of compute/memory on the host Dokploy and circumvents build issues.

I also setup a Cloudflared container in another Dokploy project using a compose file as well - passing in the API key from Cloudflare. All apples.

The trickiest and more painful part when setting up the tunnels was the subdomain limitations I found on Cloudflare.

These did not work:

`app.dev.domain.com`

These did:

`app-dev.domain.com`

I got the `app.dev.domain.com` to work initially but then I needed `river.dev.domain.com` but it would not allow another subdomain. I removed the `app.dev.` domain and then could using any nested subdomains at all. Instead had to resort to using `$app-dev.$domain` which looks dumb IMO. But, it works so moving on.

Now I got all that working pretty easily. Could access the applications from my LAN on a domain I controlled. But, Dokploy only offers a URL for re-deploying when using a Compose project which means it needs internet access.

Here comes Tailscale and GitHub actions to the rescue.


| Secret | Source |
|---|---|
| `TS_OAUTH_CLIENT_ID` | Tailscale admin → OAuth clients |
| `TS_OAUTH_SECRET` | Tailscale admin → OAuth clients |
| `DOKPLOY_API_KEY` | Dokploy → Settings → Profile → API |
| `DOKPLOY_COMPOSE_ID` | One-time API query above |
| `DOKPLOY_URL` | `http://<tailscale-ip-or-hostname>:3000` |


Github Action snippet:

```yaml

# this step comes after build and push to GHCR
deploy:
  name: Deploy to Dokploy
  runs-on: ubuntu-latest
  needs: merge
  if: success()
  steps:
    - name: Connect to Tailscale
      uses: tailscale/github-action@v2
      with:
        oauth-client-id: ${{ secrets.TS_OAUTH_CLIENT_ID }}
        oauth-secret: ${{ secrets.TS_OAUTH_SECRET }}
        tags: tag:ci

    - name: Trigger Dokploy deployment
      run: |
        curl -fsSL -X POST \
          "${{ secrets.DOKPLOY_URL }}/api/compose.deploy" \
          -H 'Content-Type: application/json' \
          -H "x-api-key: ${{ secrets.DOKPLOY_API_KEY }}" \
          -d "{\"composeId\": \"${{ secrets.DOKPLOY_COMPOSE_ID }}\"}"


I like the OAuth approach as its ephemeral and really wasn't that much extra work to get set up.

So now, instead of deploying Dokploy to a VPS costing me money I now get all benefits (and trade-offs too) of running it in my homelab but having it internet accessible. 

Also it was my first time using Dokploy (I run Coolify in another VPS) and I think its got some nice
features but for my usecase I think they're equally as good. I do like that Dokploy supports multi-instance containers due to its native `docker compose` support. With that also comes the really nice ability to run NATS in an isolated compose setup. 

Tags:

    #dokploy #tailscale #cicd
