+++
title = "Contract time countdown and hours remaining"
categories = ["zet"]
tags = ["zet"]
slug = "Contract-time-countdown-and-hours-remaining"
date = "2022-10-11 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# Contract time countdown and hours remaining

I am useless and tracking my remaining hours for my contract year. Typically,
I do my timesheets at the end of the month which isn't helpful for forecasting
how many hours I need to do over the course of the month to keep pace with my
contract. I would also like a way to see how many hours I could achieve in a month
and/or how many I need to do per day to achieve a set number, such as 160.

Example 1: Remaining hours in contract

This would calculate how many hours I have done since contract start date
until today and then output how many I have remaining

Example 2: Given a target output how I have remaining for the month

Subtract hours done from target each day, returning a mean for each working
day remaining in the month.

I think this could be achieved via my current timesheet accounting app. I
can export the database each day and upload it to the cloud. A cron could
fire and do all the calculations. A daily email, or notification could be
sent with the figures presented.

Extra bonus: apply the timesheet data into my official excell spreadsheet
which I need to submit to management each month. This is a timesink I could
easily automate.

Tags:

    #accounting #project
