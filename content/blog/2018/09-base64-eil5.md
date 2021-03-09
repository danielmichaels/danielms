+++
title = "Base64 ELI5"
slug = "Base64-ELI5"
date = "2018-05-22"
categories = ["linux", "ELI5"]
tags = ["ELI5"]
+++

Base64
------

Base64 is a common encoding scheme. But what is it, why 64 bits and how does it work?

### TL;DR

> Base64 is a binary to ASCII encoding that takes a byte (8 bits) and chunks its down into segments of 6 bits - six ones in binary equates to 64 which is where it derives its name.

## first, a demo

Let's look at an example using the three letter word "The".

This tabulated data below shows how the word 'The' is represented in both binary and base64. Each letter corresponds with an integer. For instance, 'T' is mapped to the number 84. This is then converted into binary, and the entire word 'The' can be expressed in three bytes.

```bash
#ascii       T (84)        h (104)       e (101) <-- ASCII plus base10 number
#binary   01010100      01101000     01100101    <-- three bytes
#base64   010101    000110    100001   100101    <-- four segments of six bits
```

While binary takes 8 bits, base64 only accepts 6. This means that the last two bits are clipped from the end and become the first two bits of the next segment.

Now that we have base64 binary we can calculate the base10 integer. This is done so we can map that integer to its base64 character.

```bash
#base64   010101    000110    100001   100101
#total      21        06        33       37
```

Or, a graphical representation for mobile users.

![Base 64 THE](/images/base64the.png "The in base64 image")

Character to value conversion chart for base64. By cross referencing the binary representation of the four segments (21, 06, 33, 37) we get the ASCII characters; `VGhl`.

![Base64 Chart](/img/base64.png "base64 char to int conversion chart")



#### Gotcha's

In the example we used a three byte word which can neatly be broken down
into four segments of six. What if the word is four characters equalling 32
bits. Dividing 32 by 6 gives us 5 with a remaining two bits.

We can see this clearly using `python` (because, why not.)

```python
32 / 6
>> 5.33333333333
32 % 6
>> 2
```

In that case we need to pad out the base64 to indicate that the last
segment is not complete.

The decimal representation of "Them" is `84 104 101 109` or `VGhlbQ==` in base64. As you can see there are two `=` symbols tacked on to the end. This indicates that there is two empty segments of 6 bits at the end of the encoding.

```bash
#ascii       T (84)        h (104)       e (101)     m (109)
#binary   01010100      01101000     01100101      01101101
#base64   010101    000110    100001   100101   001101  010000
#total      21        06        33        37      13      16
```

Looks like that works out perfectly, so where does the extra two
segments of zero's come in that generates the `=`'s?

The segmenting of eight bits into six bits cannot delimit or end half way
between a byte. Meaning the base64 encoding must continue until there is
no remainder. Let's explore this again using just one character.

In text format

```bash
#ascii       M (77)
#binary   01011101   00000000   00000000
#base64   010111  010000  000000  000000
#total      19      16      00      00
```

A pretty picture, too.

![Base 64 M](/images/base64M.png "M in base64 image")

I chose the letter 'M' as it is easier to explain than 'T'. Looking at it you can see that even with only one byte (read: character), it takes three bytes for base64's segments of six to equally divide - (18 / 6 = 3). This is why the letter 'M' in base64 would be padded with `==`.

This example uses two equals signs but base64 can also be padded with just one as seen below.

![Base 64 Ya](/images/base64Ya.png "Ya in base64 image")

If you are wondering about the `-1` in the columns, ignore them, its just how the JavaScript that powers the encoder has been developed. See [Ty Lewis]' CodePen to see it in action.


Base64 increases the size of the data transmission as more packets are required to be sent. So why use it? Base64's genesis was in MIME (Multipurpose Internet Mail Extensions) where large streams of text are used to send emails and attachments.

In the early days of email, everything was plain text but the inclusion of HTML and attachments created problems. Email is sent as a stream and streaming binary, or hexadecimal could of had unintended consequences as some systems or programs may interpret binary incorrectly. For instance, a null byte in some systems may indicate that the message has ended when in fact it has not.

Today base64 is used everywhere on the net, mostly as a obfuscation technique. A poor one at that. Here is some more information on base64 in the modern web skewed towards infosec.

- [LiveOverFlow] highlighting base64 patterns in hiding JSON data.
- [SANS] whitepaper on how base64 can get you pwned.
- [MalwareBytes] looking at malware obfuscation techniques.

### More information?

As always [Wikipedia](https://en.wikipedia.org/wiki/Base64) has an
excellent page on Base64.

Shout out to [Ty Lewis] for his nice encoder which I have used for demonstration.

[Ty Lewis]: https://codepen.io/lewistg/pen/MEQbmB
[LiveOverFlow]: https://www.youtube.com/watch?v=Jpaq0QkepgA
[SANS]: https://www.sans.org/reading-room/whitepapers/detection/base64-pwned-33759
[MalwareBytes]: https://blog.malwarebytes.com/threat-analysis/2013/03/obfuscation-malwares-best-friend/
