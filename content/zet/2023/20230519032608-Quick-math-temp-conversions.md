+++
title = "Quick math temp conversions"
categories = ["zet"]
tags = ["zet"]
slug = "Quick-math-temp-conversions"
date = "2023-05-19 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Quick math temp conversions

Commit to memory; converting to and from Celsius and Farenheit.

Close enough approximations which are good enough for field work.

C to F
```shell
(Cx2) + 30

10 celsius X 2 = 20
20 + 30 = 50

Exact is formula is:
(10*1.8) + 32

so comparing the two methods the difference is
quick: 50
exact: 50

If the temp is 30C
the delta between the methods grows slightly
quick: 90
exact: 86
```

F to C
```shell
(F-30)/2

55 - 30 = 25 
25/2 = 12.5

Exact formula is:
(F-32)/1.8

comparing them:
quick: 12.5
exact: 12.7
```

This cowboy math is basically all you need in life.

Calculatingusing decimals and fractions is dumb when you can approximate
it using year 3 math skills.

Hardest part is remembering the formula. 

Tags: 

    #math #til #skills

