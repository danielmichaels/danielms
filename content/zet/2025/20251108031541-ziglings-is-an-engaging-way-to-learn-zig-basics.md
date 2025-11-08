+++
title = "Ziglings is an engaging way to learn zig basics"
categories = ["zet"]
tags = ["zet"]
slug = "ziglings-is-an-engaging-way-to-learn-zig-basics"
date = "2025-11-08 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Ziglings is an engaging way to learn zig basics

I'm learning `zig` (enough to get by) and it's an engaging and simple way to get started.

To learn a new language I'd typically find a blog or series of articles and follow along. This can be a little tedious because I like to type things out and run the code myself.

Often I'd sick of the tedium of writing the boilerplate, or find articles that miss all the context or supporting code the scaffolds the learning point.

Ziglings (like Rustlings, which I haven't tried) instead sets up all of that for you in a single repo.

Each "zigling" is an intentionally broken program with comments on how to get it working. This is great because you can use whatever IDE you're comfortable
with and test things quickly.

I found a couple exercises didn't explain clearly for me which is **great** because it forces me to research why. For me that improves my learning outcomes.

Anyway I'm 30+ `ziglings` in and the language has some cool features. So far I think my favourite is `errdefer` which is like `defer` but only executes on an error.

Snippet from the exercise:

```zig
fn makeNumber() MyErr!u32 {
    std.debug.print("Getting number...", .{});

    // Please make the "failed" message print ONLY if the makeNumber()
    // function exits with an error:
    errdefer std.debug.print("failed!\n", .{});

    var num = try getNumber(); // <-- This could fail!

    num = try increaseNumber(num); // <-- This could ALSO fail!

    std.debug.print("got {}. ", .{num});

    return num;
}
```

Outputs:

```bash
Getting number...got 5. Getting number...failed!
```

Whilst `zig` is potentially more low level than I typically need I am interested in learning it because of that. I need a reason to go deeper in my 
computer science fundamentals (I am self-taught).

After the exercises I think I'll write a simple JWT decoder CLI - it's something I use a lot in my `$dayjob`.

Tags:

    #zig #learning

