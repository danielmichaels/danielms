+++
title = "I need to track contract time"
categories = ["zet"]
tags = ["zet"]
slug = "I-need-to-track-contract-time"
date = "2022-10-16 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
+++

# I need to track contract time

In a given contract year I am entitled to work 1920 hours. I cannot go over but
I can go under. Going under is money lost for me - my contract owner gets the
total allotment of their share.

I've done the math and over the last two and I'll be pulling up short
by ~300 hours (probably 295). This equates to several tens of thousands 
of dollars in income. Income I've left on the table due to my own ignorance
of my worked hours.

I'm lazy - literally the epitome of Just-In-Time everything. Learning, 
chores, presents for family any topic you can think of. So whilst I 
track every hour I do using an app on my phone I rarely do any tracking
long term. Meaning only today (with 32 working days left on my current contract)
am I figuring out how far behind I am.

```shell
1920 / 8
> 240 # days at 8 hrs per day to complete contract hours at zero
```

So what can I do about this? Build a tool to prompt me every day of course.

`contract-timetracker` (terrible placeholder name). It should:

- report hours remaining on contract
- report mean hours per day required to achieve zero balance
- calculate number of possible working days left in contract - taking public holidays in to account
- produce a end of month report with total hours worked and number of days (possibly with each day's
  total)
- (optional) write to my contract timesheet excel file with the datetimes and email to myself
- run using OpenFaaS on a cronjob so I don't need to manage the service

Considering how JIT I am this seems like a decent 80% solution for keeping me on track for my 
next contract which starts in November. I got a pay rise and want to capitalise on it and not 
leave money I'm entitled to on the table.

Tags:

    #project #contract #openfaas
