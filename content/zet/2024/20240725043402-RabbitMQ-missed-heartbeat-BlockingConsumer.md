+++
title = "RabbitMQ missed heartbeat BlockingConsumer"
categories = ["zet"]
tags = ["zet"]
slug = "RabbitMQ-missed-heartbeat-BlockingConsumer"
date = "2024-07-25 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# RabbitMQ missed heartbeat BlockingConsumer

Fix missed heartbeats from client, add the following line of code:

    connection.process_data_events()

This is for when using a `basic_consume`, `start_consuming`.

ref: https://anands.me/blog/pika-missed-heartbeats-rabbitmq
