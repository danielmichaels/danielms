+++
title = "restish is a great API CLI tool"
categories = ["zet"]
tags = ["zet"]
slug = "restish-is-a-great-api-cli-tool"
date = "2025-09-29 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# restish is a great API CLI tool

I really like [huma](https://huma.rocks). I think every API no matter how small should aim to have a OpenAPI document. It future proofs you, and with tools like Huma or [goa](https://goa.design) its built in. 

In the Huma docs the creator recommends a tool called [restish](http://rest.sh) - of which he is also the author. I played with it but to be honest didn't really explore well enough to "get" it. Now I do.

If you have a public OpenAPI spec/doc you can add it to a `restish` namespace. Doing this saves the spec and lets `restish` "see" all of it including its request and response bodies.

Heres an example of Synadia Clouds API (it's long):

```shell
API for Synadia Cloud

Usage:
  restish prod [flags]
  restish prod [command]

Signing Key Groups Commands:
  copy-account-sk-group                    Copy Account SK Group
  create-account-sk-group                  Create Account Signing Key Group
  delete-account-sk                        Delete Account Signing
  delete-account-sk-group                  Delete Account Signing Key Group
  get-account-sk                           Get Account Signing
  get-account-sk-group                     Get Account Signing Key Group
  list-account-sk-group                    List Account Signing Key Groups
  list-account-sk-group-keys               List Signing Keys
  rotate-account-sk                        Roate Active Signing Key
  update-account-sk                        Update Account Signing
  update-account-sk-group                  Update Account Signing Key Group

NATS Users Commands:
  assign-nats-user-team-app-user           Assign Team App User to NATS User
  copy-nats-user                           Copy nats user
  create-user                              Create NATS User
  delete-nats-user                         Delete NATS User
  download-nats-user-bearer-jwt            Get Bearer JWT
  download-nats-user-creds                 Get Creds
  download-nats-user-http-gw-token         Get HTTP Gateway Token
  get-nats-user                            Get NATS User
  list-account-sk-group-users              List NATS Users
  list-nats-user-connections               List NATs User Connections
  list-nats-user-issuances                 List nats user issuances
  list-nats-user-team-app-users            List Team App Users
  list-users                               List NATS Users
  rotate-nats-user                         Rotate nats user nkey and seed
  un-assign-nats-user-team-app-user        Unassign Team App User from NATS User
  update-nats-user                         Update NATS User

Accounts Commands:
  create-account                           Create Account
  create-or-update-nats-user-revocation    Create or Update Revocation for a NATS User NKey
  delete-account                           Delete Account
  delete-nats-user-revocation              Delete a for a NATS User NKey
  get-account                              Get Account
  get-account-export                       Export Account Seeds
  get-account-info                         Get Account Info
  get-account-metrics                      Get Account Metrics
  get-nats-user-revocation                 Get Revocation for a NATS User NKey
  list-account-connections                 List Account Connections
  list-accounts                            List Accounts
  list-accounts-overview-metrics           List Accounts overview metrics
  list-agent-tokens                        List Agent Tokens
  nats-core-websocket-viewer               Subscribe to a NATS Core subject over websockets
  unmanage-account                         Unmanage Account
  update-account                           Update Account

Alerts Commands:
  acknowledge-alert                        Acknowledge Alert
  create-alert-rule                        Create Account Alert Rule
  create-system-alert-rule                 Create System Alert Rule
  delete-alert-rule                        Delete Account Alert Rule
  delete-system-alert-rule                 Delete System Alert Rule
  get-alert                                Get Alert
  get-alert-rule                           Get Account Alert Rule
  get-system-alert-rule                    Get System Alert Rule
  list-alert-rules                         List Account Alert Rules
  list-alerts                              List Alerts
  list-system-alert-rules                  List System Alert Rules
  run-alert-rule                           Run Account Alert Rule
  run-system-alert-rule                    Run System Alert Rule
  update-alert-rule                        Update Account Alert Rule
  update-system-alert-rule                 Update System Alert Rules

App Users Commands:
  assign-account-team-app-user             Assign Team App User to Account
  assign-system-team-app-user              Assign Team App User to System
  assign-team-app-user                     Assign App User to Team
  create-app-user                          Create App User
  delete-app-user                          Delete App User
  get-app-user                             Get App User
  get-current-agent-token                  Get Current Agent Token
  list-account-team-app-users              List Account Team App Users
  list-app-user-roles                      List Roles
  list-app-users                           List App Users
  list-system-team-app-users               List System Team App Users
  list-team-app-users                      List App Users
  un-assign-account-team-app-user          Unassign Team App User from Account
  un-assign-system-team-app-user           Unassign Team App User from System
  un-assign-team-app-user                  Unassign App User from Team
  update-app-user                          Update App User
  update-team-app-user                     Update App User Team Assignment

JetStream Commands:
  delete-pull-consumer                     Delete Pull Consumer
  delete-push-consumer                     Delete Push Consumer
  get-jet-stream-placement-options         Get JetStream Placement Options
  get-pull-consumer-info                   Get Pull Consumer
  get-push-consumer-info                   Get Push Consumer
  list-jet-stream-assets                   List JetStream Assets
  update-pull-consumer                     Update Pull Consumer
  update-push-consumer                     Update Push Consumer

JetStream: KV Buckets Commands:
  create-kv-bucket                         Create KV Bucket
  create-kv-pull-consumer                  Create Pull Consumer
  create-kv-push-consumer                  Create Push Consumer
  delete-kv-bucket                         Delete KV Bucket
  get-kv-bucket                            Get KV Bucket
  list-kv-buckets                          List KV buckets
  list-kv-consumers                        List Consumers
  update-kv-bucket                         Update KV Bucket

JetStream: Mirrors Commands:
  create-mirror                            Create Mirror
  create-mirror-pull-consumer              Create Pull Consumer
  create-mirror-push-consumer              Create Push consumer
  delete-mirror                            Delete Mirror
  get-mirror                               Get Mirror
  list-mirror-consumers                    List Consumers
  list-mirrors                             List Mirrors
  update-mirror                            Update Mirror

JetStream: Object Buckets Commands:
  create-obj-pull-consumer                 Create Pull Consumer
  create-obj-push-consumer                 Create Push Consumer
  create-object-bucket                     Create Object Bucket
  delete-object-bucket                     Delete Object Bucket
  get-object-bucket                        Get Object Bucket
  list-obj-consumers                       List Consumers
  list-object-buckets                      List Object buckets
  update-object-bucket                     Update Object Bucket

JetStream: Streams Commands:
  create-pull-consumer                     Create Pull Consumer
  create-push-consumer                     Create Push Consumer
  create-stream                            Create Stream
  delete-stream                            Delete Stream
  get-stream-info                          Get Stream
  list-streams                             List Streams
  purge-kv-bucket                          Purge KV Bucket
  purge-mirror                             Purge Mirror
  purge-obj-bucket                         Purge Object Bucket
  purge-stream                             Purge Stream
  update-stream                            Update Stream

Sharing Commands:
  create-stream-export                     Create Stream Export
  create-stream-import                     Create Stream Import
  create-stream-shares                     Create Stream Shares
  create-subject-export                    Create Subject Export
  create-subject-import                    Create Subject Import
  create-subject-shares                    Create Subject Shares
  delete-stream-export                     Delete Stream Export
  delete-stream-import                     Delete Stream Import
  delete-stream-share                      Delete Stream Share
  delete-subject-export                    Delete Subject Export
  delete-subject-import                    Delete Subject Import
  delete-subject-share                     Delete Subject Share
  get-stream-export                        Get Stream Export
  get-stream-import                        Get Stream Import
  get-subject-export                       Get Subject Export
  get-subject-import                       Get Subject Import
  list-stream-exports                      List Stream Exports
  list-stream-exports-shared               List Shared Stream Exports
  list-stream-imports                      List Stream Imports
  list-stream-shares                       List Stream Shares
  list-subject-exports                     List Subject Exports
  list-subject-exports-shared              List Shared Subject Exports
  list-subject-imports                     List Subject Imports
  list-subject-shares                      List Subject Shares
  update-subject-export                    Update Subject Export
  update-subject-import                    Update Subject Import

Agent Tokens Commands:
  delete-agent-token                       Delete Agent Token

Auth Callout Commands:
  add-auth-callout-target-account          Configure Target Account
  add-auth-callout-user                    Create Auth Callout User
  delete-auth-callout                      Delete Auth Callout Config
  delete-auth-callout-target-account       Delete Target Account
  delete-auth-callout-user                 Delete Control Account User
  get-auth-callout                         Auth Callout Config
  list-auth-callout-authenticators         Get List of Available Authenticators
  list-auth-callout-target-accounts        Get Target Account List
  list-auth-callout-users                  Get Target Account List

Authorization Commands:
  check                                    Check Authz Decisions
  list-policies                            Get Policy List
  list-roles                               Get Authz roles List

Component Versions Commands:
  get-available-component-versions         Get Available Component Versions
  get-available-component-versions-by-type Get Available Component Versions By Type

Invitations Commands:
  decide-invitation                        Accept or reject team invitation
  invite-app-user                          Invite App Users
  list-invitations                         List of pending invitations

JetStream: Consumers Commands:
  list-consumers                           List Consumers

NATS User Issuances Commands:
  get-nats-user-issuance                   Get nats user issuance

Session Commands:
  accept-terms                             Accept terms
  create-app-user-access-token             Create Personal Access Token
  delete-personal-access-token             Delete Personal Access Token
  get-personal-access-token                Get Personal Access Token
  get-version                              Get Version
  leave-team                               Leave Team
  list-app-user-access-tokens              List Personal Access Tokens
  update-personal-access-token             Update Personal Access Token

Platform Components Commands:
  platform-component-connect               Connect a Platform Component

App Service Accounts Commands:
  create-app-service-account               Create App Service Account
  create-app-service-account-token         Create Acess Token for App Service Account
  delete-app-service-account               Delete App Service Account
  delete-app-service-account-token         Delete App Service Account Token
  get-app-service-account                  Get App Service Account
  get-app-service-account-token            Get App Service Account Token
  list-app-service-account-tokens          List Access Tokens for App Service Account
  list-app-service-accounts                List App Service Accounts
  update-app-service-account               Update App Service Account
  update-app-service-account-token         Update App Service Account Token

Team Service Accounts Commands:
  create-team-service-account-token        Create Acess Token for Team Service Account
  delete-team-service-account              Delete Team Service Account
  delete-team-service-account-token        Delete Team Service Account Token
  get-team-service-account                 Get Team Service Account
  get-team-service-account-token           Get Team Service Account Token
  list-team-service-account-tokens         List Access Tokens for Team Service Account
  update-team-service-account              Update Team Service Account
  update-team-service-account-token        Update Team Service Account Token

Systems Commands:
  bulk-share                               Share assets across accounts
  create-system                            Create System
  delete-system                            Delete System
  enable-auth-callout                      Enable Auth Callout For System
  get-component-token                      Get a component access token
  get-system                               Get System
  get-system-audit-job-result              Get Audit Job Result
  get-system-audit-job-status              Get Audit Job Status
  get-system-export                        Export System Seeds
  get-system-limits                        Get System Limits
  get-system-prometheus-metrics            Get Prometheus Metrics
  import-account                           Import Account
  import-user                              Import User
  list-auth-callout-configs                List Auth Callout Configs
  list-clusters                            List Clusters
  list-connections                         List Connections
  list-servers                             List Servers
  list-system-audit-jobs                   List Audit Jobs
  list-system-info-accounts                List System Accounts Info
  list-system-info-servers                 List System Servers info
  list-team-systems                        List Systems
  rotate-agent-token                       Rotate Agent Token
  run-system-audit-check                   Run System Audit Check
  system-jwt-sync                          Re-sync JWTs of all accounts in this system
  unmanage-system                          Unmanage System
  update-platform-components               Update Platform Components for System
  update-system                            Update System

Teams Commands:
  create-team                              Create Team
  create-team-service-account              Create Team Service Account
  delete-team                              Delete Team
  get-team                                 Get Team
  get-team-limits                          Get Team Limits
  import-system                            Import a System
  list-team-info-app-users                 List info of App Users in Team
  list-team-service-accounts               List Team Service Accounts
  list-teams                               List Teams
  update-team                              Update Team

Workloads Commands:
  get-workload-limits                      Get Workloads limits
```

Yep, that is every single endpoint from the doc. 

If I want to list all consumers but don't remember what the request body is?

Just use `restish <namespace> list-consumers` and it'll print this.

```shell
List Consumers
## Argument Schema:

{
  stream-id: (string)
}


## Response 200 (application/json)

Success


{
  items*: [
    allOf{
      {
        ack_floor*: {
          consumer_seq*: (integer min:0 max:1.8446744073709552e+19)
          last_active: (string nullable:true)
          stream_seq*: (integer min:0 max:1.8446744073709552e+19)
        }
        cluster: allOf{
          {
            leader: (string)
            name: (string)
            replicas: [
              allOf{
                {
                  active*: (integer format:int64)
                  current*: (boolean)
                  lag: (integer min:0 max:1.8446744073709552e+19)
                  name*: (string)
                  offline: (boolean)
                }
              }
            ]
          }
        }
        config*: {
          ack_policy*: (string enum:none,all,explicit) enums have been changed to match UnmarshalJSON in https://github.com/nats-io/jsm.go/blob/main/api/consumers.go
          ack_wait: (integer format:int64)
          backoff: [
            (integer format:int64)
          ]
          deliver_group: (string)
          deliver_policy*: (string enum:all,last,new,by_start_sequence,by_start_time,last_per_subject)
          deliver_subject: (string)
          description: (string)
          direct: (boolean) Don't add to general clients.
          durable_name: (string)
          filter_subject: (string)
          filter_subjects: [
            (string)
          ]
          flow_control: (boolean)
          headers_only: (boolean)
          idle_heartbeat: (integer format:int64)
          inactive_threshold: (integer format:int64)
          max_ack_pending: (integer)
          max_batch: (integer)
          max_bytes: (integer)
          max_deliver: (integer)
          max_expires: (integer format:int64)
          max_waiting: (integer)
          mem_storage: (boolean)
          metadata: {
            <any>: (string)
          }
          name: (string)
          num_replicas*: (integer)
          opt_start_seq: (integer min:0 max:1.8446744073709552e+19)
          opt_start_time: (string nullable:true)
          rate_limit_bps: (integer min:0 max:1.8446744073709552e+19)
          replay_policy*: (string enum:instant,original)
          sample_freq: (string)
        }
        created*: (string format:date-time)
        delivered*: {
          consumer_seq*: (integer min:0 max:1.8446744073709552e+19)
          last_active: (string nullable:true)
          stream_seq*: (integer min:0 max:1.8446744073709552e+19)
        }
        name*: (string)
        num_ack_pending*: (integer)
        num_pending*: (integer min:0 max:1.8446744073709552e+19)
        num_redelivered*: (integer)
        num_waiting*: (integer)
        push_bound: (boolean)
        stream_name*: (string)
        ts*: (string format:date-time)
      }
    }
  ]
}


## Responses 400/401/403/404
# truncated...
```

To me this makes using API's using a CLI **actually** possible without consulting the docs.

## Walkthrough 

Let's walkthrough how to set all this up by using the `restish` example API.

```shell
restish api configure example
# next steps are interactive
restish api configure example https://api.rest.sh
> ? Select option Save and exit
```

We've told `restish` to associate <https://api.rest.sh> with the namespace of `example`. This is how we can segresgate many different
API's from one another. 

Now we can auto detect what endpoints exist.

`restish example`

Which returns,

```shell
Usage:
  restish example [flags]
  restish example [command]

Echo Commands:
  delete-echo       
  get-echo          
  patch-echo        
  post-echo         
  put-echo          

Books Commands:
  delete-book       
  get-book          
  list-books        
  patch-book        Patch book
  put-book          

Caching Commands:
  get-cached        

Example Commands:
  get-example       

Images Commands:
  get-image         
  list-images       

Status Commands:
  get-status        

Types Commands:
  get-types-example 
  put-types-example 
```

I want to `put-book` but don't know what I need to provide in the request body.

```
# restish example put-book -h
# truncating some output for brevity

## Argument Schema:

{
  book-id: (string)
}

## Option Schema:

{
  --if-match: [
    (string)
  ]
  --if-none-match: [
    (string)
  ]
  --if-modified-since: (string format:date-time-http) Succeeds if the server's
resource date is more recent than the passed date.
  --if-unmodified-since: (string format:date-time-http) Succeeds if the server's
resource date is older or the same as the passed date.
}

## Input Example

{
  "author": "string",
  "published": "2020-05-14T23:44:51-07:00",
  "rating_average": 1,
  "ratings": 1,
  "recent_ratings": [
    {
      "date": "2020-05-14T23:44:51-07:00",
      "rating": 1
    }
  ],
  "title": "string"
}
```

Now that we know what the request requires we can put a book into the collection

```shell
 echo '{
    "author": "Someone",
    "published": "1900-05-14T00:00:00-12:00",
    "rating_average": 4,
    "ratings": 30,
    "title": "My book is good"
    }' | restish example put-book good-book

# returns HTTP/2.0 204 No Content
```

And `restish example list-books` shows it in the collection

```json
[
  {
    modified: 2025-09-29T03:36:11.74478507Z
    url: "/books/the-demon-haunted-world"
    version: "yWWo/Pq0gMgjMNgWw+K+Wg=="
  }
  {
    modified: 2025-09-29T03:36:11.74478507Z
    url: "/books/the-fabric-of-the-cosmos"
    version: "qhTFCaoJg92NnpRTMJzsEA=="
  }
  {
    modified: 2025-09-29T03:42:34.560280084Z
    url: "/books/good-book"
    version: "3Bvek8xG1BB3i+g4kvDzFg=="
  }
]
```

---

Of course you can also just use it like you would `curl`, here's an example:

```shell
 restish api.github.com/users/rs | jq '{name, company, location, login, bio}'
> {
>   "name": "Olivier Poitrey",
>   "company": "Netflix",
>   "location": "Silicon Valley, California, USA",
>   "login": "rs",
>   "bio": "Director of Engineering at Netflix\r\nCo-Founder & ex-CTO of Dailymotion\r\nCo-Founder of NextDNS"
> }
```

There is a lot more to cover but thats some of the parts I really like. 

`restish` is a tool that makes it much easier to stay in the terminal whilst interacting with a unfamiliar or complex API *without* needing the docs.

Tags:

    #restish #api
