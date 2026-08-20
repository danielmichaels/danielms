+++
title = "I don't think we are ready to NOT read the output"
categories = ["zet"]
tags = ["zet"]
slug = "i-don't-think-we-are-ready-to-not-read-the-output"
date = "2026-08-20 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# I don't think we are ready to NOT read the output

I read the code for things that are important. 

Personal projects, non-critical, funsies - sure I'll allow it through and take the output at face value. 

But for real work - I just can't see it. I think 20-40% of the code needs to reworked, always, from agents. 

Common issues:

- Not reuses; duplicating
- Atrocious comments (this is a Claude problem mostly)
- Over testing
- Stupid function/method names 
- Security holes (paradoxically very good at security at the same time?)
- In SQL apps, looooves a migration. Like will generate 5 migrations in a single session instead of 1 migration file with everything in it
    - Not afraid of throwing in some N+1's too

That said, they're still a game changing and I don't think we should go back to "hand-coding". But they are not ready for blind trust.

*Trust, but verify*

Tags:

    #ai
