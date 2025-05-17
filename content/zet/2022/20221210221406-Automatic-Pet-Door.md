+++
title = "Automatic Pet Door"
categories = ["zet"]
tags = ["zet"]
slug = "Automatic-Pet-Door"
date = "2022-12-10 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Automatic Pet Door

I have a dog and recently got a cat. The cat keeps going out the dog
door which poses a problem; how do I allow the dog to go out freely 
but prevent the cat from going out?

Enter a DIY automatic pet door.

Requirements:

- Dog can go freely
- Cat must not be able to exit freely 
- Anything on animals must be passive i.e. not battery powered
- No circuitry must live on the pet door (a flapper)

Idea: tag based access control:

* the dog door locks when a tag is presented (the cats). 

I think the door being default to unlocked works better than the
inverse. This way if the cat is at the door whilst the dog it outside
the dog would be locked out. This is a design consideration which 
should prevent the cat from tail gating or ambushing the door when
the dog enters or exits.

My current dog door is a flap based system which uses light magnets 
to keep the flap in place. I want to keep using this as its a pet
door which can be used in any sliding door (great for rentals).

The *locking* mechanism would be stronger magnets which are energised
whenever the cat is present preventing the flap from opening.

A few options:
- Acousto Magnetic detection
- RF detection 

Tags:

    #rfid #rf #diy #iot
