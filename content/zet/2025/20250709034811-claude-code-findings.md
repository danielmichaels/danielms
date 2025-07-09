+++
title = "Claude Code findings"
categories = ["zet"]
tags = ["zet"]
slug = "claude-code-findings"
date = "2025-07-09 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Claude Code findings

After using it for a few days, some initial thoughts and observations.

- It uses `*.bak` files despite having git. Obviously it's learnt this practice from us, failed humans. Often it forgets to clean them up.
- Sometimes you really have to fight it to get it to do what you want no matter how much "planning" or "context engineering" you do.
    - E.g. using NATS JetStream legacy over the new standard. Several times it reverted from "new" to "legacy". In the end I took over and fixed it after much wasted tokens.
- AI is really a "shovels and picks" moment IMO - the model vendors are absolutely killing it with things like claude code. Forget MAX they are chasing enterprise pay-as-you-go. I, rather cynically, believe it's possible they don't want models that are too smart else they'll lose money on credits; akin to why we'll never get a cure for cancer - no money in it. 
- It's insanely good for scaffolding projects but mediocre at filling in the complex centre of most big projects.
- Very easy to lose your way if you let it auto-edit - I can't use it for work and hand on heart I fully understand it; actually better off writing the code myself and asking AI questions about my intent rather than letting it take the wheel (so no change from status quo for me).
- Reviewing code is boring and you'll do A LOT of that using these tools.
- Asking questions, clarifying how things work and talking through your thought process are amazing quality of life improvements from these tools.
- It's pretty average at logging what it's done in any markdown file; often does the work then I have to request it updates documents.
- Super easy to get lazy.
- Voice prompting is really cool because I can ramble on and it'll understand me - sometimes when I start talking I realise my question/prompt is silly or I refine it more than I would just typing alone. Rubber ducky, if you will.

Now, could I do better at "context engineering"? Most likely. Do I think throwing shade on anyone who has anything negative to say via a "probably not prompting it right" is a very bad pattern, yes I do. IMO, it's another form of character assassination we've seen time and time again in the last few years; "trust the science", "right winger", "MAGA" and so on as a quick way to dismiss someone who is not "on your side". So yeah, I think we do ourselves an injustice here by immediately dismissing valuable feedback by labeling someone by virtue of "not prompting correctly"! Bit of a philosophical rant but I can't help but feel it in this current era of AI. 

Regardless, I can absolutely see how these tools are a real boon for developers. Do I think running 6 agents at the same time on a single codebase is going to yield amazing results - no I don't, at least not long term. 

Tags:

    #ai #claude #rant
