+++
title = "Running a business with OpenFaaS, Stripe and Sendgrid"
categories = ["golang", "openfaas", "stripe"]
tags = ["business", "golang"]
slug = "openfaas-stripe-sendgrid-a-business-model"
date = "2023-05-08"
draft = "false"
ShowToc = "true"
+++

I did some work recently to set up a payments and notification backend for a [Carrd] website.
The website integrates with some booking services which are linked to payment providers.

The frontend needed two more services added that could not integrate with the booking service, 
and that got me thinking about how to level up the site without too much effort.

I have some experience using [OpenFaaS] and its smaller cousin, [Faasd].

## Faasd? OpenFaaS?

[OpenFaaS] is an easy to host Functions-as-a-Service platform. It is a very capable platform 
which lets you write functions in almost any language, on any cloud without any vendor lock-in.
It has great tooling including its own CLI and UI. It is often deployed to Kubernetes which I've
used to achieve interesting things in the past.

This project needed something a lot smaller and lightweight enough to run on a small and cheap VPS.
That's where [faasd] comes in.

It's a single Go binary with enough firepower to easily run this app. Here's the minimum 
[system requirements](https://docs.openfaas.com/deployment/faasd/#deployment):

- 512MB-1GB RAM
- 1-4 vCPU cores
- 10-25GB of disk space

## Functions

This modest website isn't *web scale*, it's serving a minimal number of requests per day. None of 
the requests are system intensive and for the most part only serve to integrate with other third-party 
platforms.

The way I think about functions is as handlers for requests. Each function serves a *route* and 
that's about it. Though I don't split functions up by HTTP method, only by the url they're 
responsible for.

How it works is pretty impressive and well detailed across all the OpenFaaS blogs and documentation.

But, I'll do my best to give a run down on how OpenFaaS works. 

OpenFaaS takes a lower level concept like container orchestration and kubernetes and abstracts away 
all the complexity by exposing an application surface. This application surface is in the form of
*functions*. Each function is its own container and for the most part accepts something from `stdin`
and outputs `stdout` or `stderr` to the caller. Functions can be invoked via HTTP, the CLI or 
through event bridges such as Kafka, NATS, MQTT and so on. It also has a [cron] trigger which is a
great builtin feature.

Since each function is its own container OpenFaaS makes it easier to scale up your functions. just
add more containers for that function. Isolation between functions is also a benefit. One bad function 
won't crash your entire app. 

Being that everything is containers means I can have functions in Go, python, node, C# or even bash. 
I typically use Go these days but have use a python function to augment my Go functions because 
it was the right choice. Had that project been a server written in Go more than likely I would 
of had to stand up a microservice and then setup NATS or HTTP endpoints for them to interact. 
Using OpenFaaS made it a cinch.

Although it's not perfect. Sometimes function codebases can have a lot of repeated code. 
As an example, both my functions use the same logger and database modules which are identical. It's
not so bad because I have scripts that autogenerate new functions and handle these things. 
I will point out that you can create your own [function template][ft] which could include 
any required modules effectively removing this hiccup.

Another beautiful addition is prometheus and Alert Manager out of the box. Whilst I am not using 
these for this project it is nice knowing I can easily hook up [grafana] in the future.

For a really succinct summary of OpenFaaS with great explanatory images please read 
[Ivan Velichko][ivan]'s [blog post].

### Function: Stripe

The primary function is the Stripe webhook handler which serves as the most common entrypoint.

When a customer makes a purchase the webhook handler consumes the response and based on what it is,
acts on it.

As an example, when a `checkout.session.complete` is received, it is marshalled into a struct. Then
inserted into a database, in this case [MongoDB Atlas](https://www.mongodb.com/atlas). The function
then calls another function passing along the key for the recently inserted data.

I've found this works excellently. It quickly writes to the database and send the POST to the next 
function - both done in goroutines.

### Function: Sendgrid/Email

Once a payment intent is received and dealt with in the Stripe function the email function is 
triggered. In my opinion this is one of the powerful parts of OpenFaaS; function chaining.

Once the payload is received the function checks which product was purchased and triggers a series 
of emails to the customer and website administrators.

Leveraging SendGrid here made sense because it has a generous free tier of 100 emails per day. 
Even then once you start paying it is still quite reasonable. It also has a decent email template
generator which is used to deliver specific emails based on the product purchased.

After all that, the functions, third-party integrations and how its invoked look something like 
this.

![](/images/faasd-overview.png "faasd carrd overview")

## Final thoughts

To integrate Stripe with Carrd I used [payment links](https://stripe.com/en-au/payments/payment-links)
It was my first time using them, and they are **fantastic**. Hooking up the URL
to Carrd is as easy as creating a button which directs the customer to the payment link.

Initially, I thought Carrd was only good for a simple landing page but with payment links you
can really stretch it quite far.

It was the realisation that this project was really just the confluence of several other API's that
made me think about a different way to deal with this problem. 

Is it better than a long-running server? I think it depends on what you want or need out of it. 
Do Functions-as-a-Service replace servers? No. 

It's right tool for the right job and I think this is the right tool for now. Knowing that when 
this gets bigger I can lift and shift into kubernetes is nice. Nicer still, I don't have to 
write any deployment manifests to make it all work because OpenFaaS handles everything under the hood.

I think in the future I will start all designs with a simple question; do I really need a 
backend, or do I just need some functions?

[faasd]: https://docs.openfaas.com/deployment/faasd/
[openfaas]: https://openfaas.com
[carrd]: https://carrd.co
[ft]: https://docs.openfaas.com/cli/templates/#customise-a-template
[blog post]: https://iximiuz.com/en/posts/openfaas-case-study/
[ivan]: https://twitter.com/iximiuz
[grafana]: https://docs.openfaas.com/openfaas-pro/grafana-dashboards/
[cron]: https://docs.openfaas.com/reference/cron/
