+++
title = "Public Key Crypto: The Basics"
tags = ["linux"]
date = "2018-06-14"
categories = ["ELI5"]
+++

# Public Key Cryptography

My simplest explanation
------------------------

**TL;DR:** Public key cryptography is a method of using a different key
for the encryption and decryption of a file, message or other medium.

Public Key Cryptography
-----------------------

*What problem does it solve?*

If two parties wish to exchange information to one another but the
medium which they use to communicate is compromised, how do they share a
secret key to unlock their encrypted messages? Symmetric key
cryptography would make this untenable for the most part. But, if the
parties use two keys; one for encryption and one for decryption they can
circumvent the interception as having only one of the keys is not enough
to unlock the messages contents. This is also known as *asymmetric*
cryptography.

Public key cryptography uses two keys; a public and private key. The
public key can be transferred or published globally without worry, as it
requires the private key to unlock the cipher.

*how does it work?*

We have two actors whom wish to send a secure message; *Bob* and
*Alice*. If Alice wants to send a message to Bob, first she finds Bob's
public key either by looking it up globally, or getting it directly from
Bob. Alice now encrypts her message using Bob's public key. This is safe
because Bob's public key cannot decrypt this message, meaning it does
not matter if anyone else has Bob's public key.

When Bob receives Alice's encrypted message, he uses his private key
-which only Bob has access to - and decrypts the message. Should Bob
wish to reply, he simply repeats the process, now using Alice's public
key to encrypt his message to her.

At it's core, this is the principle of asymmetric or public key
encryption.

Digital Signatures
------------------

In the physical world, a signature is a form of validation. It certifies
that the signing party agrees to the document, and that they are who
they say they are. Unfortunately, it is easy to spoof or refute a
signature over the internet - how does one prove it was *them* and not
someone else?

Using public key cryptography we are able to digitally watermark an
object in a way that verifies the signature is from the party in
question. If Alice wants to digitally verify a document with her name,
she can sign it using public key cryptography.

When Bob receives Alice's encrypted message everyone can decrypt this
message as they can get her public key, but no-one else has her private
key. This means that no-one could of altered the message, or forged it
entirely because Alice is the sole owner of the private key.

This is often referred to *non-repudiation*. Meaning that at a later
time the author cannot refute the fact that they are the owner of the
message.

Can Someone Use My Key?
-----------------------

The entire principle of public key cryptography relies on trusting the
source of the keys. If a third party can steal or intercept the private
and public keys then they can masquerade as the legitimate recipient or
sender of the message.

Two common examples are possible; interception, and theft.

If a third party can intercept the public key before the legitimate
recipient and then forward on their own in its place, they will effect a
man-in-the-middle (MITM) attack. In this case, it would be possible for
the third party to read the messages, and reply on their behalf.

Another possibility is theft. If a belligerent were able to break into a
key-store or server which has the private key stored on file, it is
possible to now read, write and intercept messages. More often this type
of intrusion will use that key to bypass other servers authentication
which in the case of passwordless-ssh may give an attacker access to
another server in the network.

This issue of authentication is outside this post, but it is essential
to understand that asymmetric cryptography provides confidentially,
rather than authentication.
