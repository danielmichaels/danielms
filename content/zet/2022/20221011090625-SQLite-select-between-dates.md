+++
title = "SQLite select between dates"
categories = ["zet"]
tags = ["zet"]
slug = "SQLite-select-between-dates"
date = "2022-10-11 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# SQLite select between dates

SQLite has spartan types and today I learnt how to select between a date
range.

```sql
SELECT date1,date2,amount,ROUND((JULIANDAY(date2) - JULIANDAY(date1))*86400) 
AS diff 
FROM times 
WHERE date1 between '2022-09-01' AND '2022-09-30';
```

Outputs

```shell
# .mode column
# .headers on
date1             date2             amount      diff
----------------  ----------------  ----------  ----------
2022-09-02 08:25  2022-09-02 17:11  1008.17     31560.0
2022-09-05 08:22  2022-09-05 17:05  1002.42     31380.0
2022-09-06 08:56  2022-09-06 16:30  870.17      27240.0
2022-09-07 08:34  2022-09-07 17:05  979.42      30660.0
```

Using `JULIANDAY` and multiplying by `86400` returns the difference in
seconds between `date1` and `date2`. 

**Last seven days**
```sql
SELECT date1,date2,amount,ROUND((JULIANDAY(date2)-JULIANDAY(date1))*86400)
AS diff
FROM times
WHERE date1 > datetime('now', '-7 days');
```

**Since start of the month**

```sql
SELECT date1,date2,amount,ROUND((JULIANDAY(date2)-JULIANDAY(date1))*86400)
AS diff
FROM times
WHERE date1
BETWEEN datetime('now', 'start of month') AND datetime('now', 'localtime');`
```

Tags:

    #sql
