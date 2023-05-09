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

This modest website isn't *web scale*, it's serving very few requests per day. None of the 
requests are system intensive and for the most part only serve to integrate with other third-party 
platforms.

The way I think about functions is as handlers for requests. Each function serves a *route* and 
that's about it. Though I don't split functions up by HTTP method, only by the url they're 
responsible for.

How it works is pretty impressive and well detailed across all the OpenFaaS blogs and documentation.

But, I'll do my best to give a run down on how it works (todo)
### Stripe

The primary function is the Stripe webhook handler which serves as the most common entrypoint.

When a customer makes a purchase the webhook handler consumes the response and based on what it is,
acts on it.

As an example, when a `checkout.session.complete` is received, it is marshalled into a struct. Then
inserted into a database, in this case [MongoDB Atlas](https://www.mongodb.com/atlas). The function
then calls another function passing along the key for the recently inserted data.

I've found this works excellently. It quickly writes to the database and send the POST to the next 
function - both done in goroutines.

### Sendgrid/Email

Once a payment intent is received and dealt with in the Stripe function the email function is 
triggered. In my opinion this is one of the powerful parts of OpenFaaS; function chaining.

Once the payload is received the function checks which product was purchased and triggers a series 
of emails to the customer and website administrators.

Leveraging SendGrid here made sense because it has a generous free tier of 100 emails per day. 
Even then once you start paying it is still quite reasonable. It also has a decent email template
generator which is used to deliver specific emails based on the product purchased.

## Final thoughts

To integrate Stripe with Carrd I used [payment links](https://stripe.com/en-au/payments/payment-links)
It was my first time using them, and they are **fantastic**. Hooking up the URL
to Carrd is as easy as creating a button which directs the customer to the payment link.

Initially, I thought Carrd was only good for a simple landing page but with payment links you
can really stretch it quite far.

It was the realisation that this project was really just the confluence of several other API's that
made me think about a different way to deal with this problem. 

I think in the future I will start all designs with a simple question; do I really need a 
backend, or do I just need some functions?

[faasd]: https://docs.openfaas.com/deployment/faasd/
[openfaas]: https://openfaas.com
[carrd]: https://carrd.co
