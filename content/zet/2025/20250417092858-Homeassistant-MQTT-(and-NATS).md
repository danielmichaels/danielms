+++
title = "Homeassistant MQTT (and NATS)"
categories = ["zet"]
tags = ["zet"]
slug = "Homeassistant-MQTT-(and-NATS)"
date = "2025-04-17 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Homeassistant MQTT (and NATS)

I have a couple of automations (pool pump being most important) that run over Zigbee2MQTT.

I was using the Homeassistant built in MQTT broker but decided to test out NATS MQTT implementation and see if it was compatible.
It works a treat!

However, when doing this to make it work I didn't realise that I need to update the Z2M configuration and point it at the NATS cluster **and**
reconfigure HA's MQTT service to use it as well. 

When I triggered things from the HA Z2M panel everything worked as expected and I could see messages being routed but my automations wouldn't work.

The fix is `Settings -> Devices & services -> MQTT -> Reconfigure` then change the broker address. Once I did that my automations started working again.

Tags:

    #mqtt #nats #Homeassistant

