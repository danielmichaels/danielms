+++
title = "Microservice contracts can be coupled"
categories = ["zet"]
tags = ["zet"]
slug = "microservice-contracts-can-be-coupled"
date = "2022-05-02 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Microservice contracts can be coupled

Microservices aren't evil or the saviour, but they are often
touted as the solution to a lot of problems. One of the biggest
microservice misconceptions is *decoupling*.

A service which sends stuff to a queue to be consumed by 
another service is not magically decoupled. It takes work
and consistency to do that. Lately, I've been seeing a lot
of service code with references to things that are only a
concern for the consuming services. I.e. they're coupled.

Further, if service A needs more data from another service, now I need
to touch two code bases to make that change. If I deploy one without 
updating the other, it could result in a crash (though not always). 


I like writing services when it makes sense but I do not agree its the
panacea of all coding. For a lot of things, monolith's or large 
*services* are more efficient. This holds especially true if your team
size is small.

Tags:

    #rant #microservices #design
