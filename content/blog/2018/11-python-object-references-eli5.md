+++
title = "Python Object References"
slug = "Python-Object-References"
date = "2018-07-09"
tags = ["ELI5"]
categories = ["python", "programming"]
+++

# Python Object Reference ELI5
----------------------------

I first learnt about variables through the analogy of "variables are
boxes" and that we assign things to those boxes Turns out, this isn't
particularly helpful in objected orientated programming. This post is
about how python treats object assignment and some of the hidden
gotcha's that can cause unintended errors along the way. Instead of
"boxes" it is better to think of variables as "labels" that we
attach to objects. And, as everything in python is an object its
important to remember that all objects have three things; identity, type
and values. Values are the only things that change once an object is
created, and it values that we often care about, and hence label.

### Labels not boxes

Extending the "labels" metaphor a little we look at the assignment of
variables.

```python
a = 2 # we label the integer 2 as 'a'
b = a # 'a' is now labelled as 'b'
c = b # and 'b' is now labelled as 'c'
```

Above we can see that the object `2`} is assigned to the
variable 'a'. Each subsequent assignment thereafter is simply a
reference to the same object. When viewed through this lense you can
start to see how objects have labels. It is not feasible that the
`2` can exist in three different boxes rather we visualise
`2` having three sticky notes attached to it. If we changed
`a` like this `a = 20` then it is just a
matter of peeling off the sticky note with `a` written on
it from `2` and attaching it to `20`. To
further aid in this thinking, always read assignments from right to
left. The right side is where the object is created or retrieved and the
left is what we bind to it (the label.. Enough you get it already!)

When an object like `2` has many labels we called this
*aliasing*. Aliasing is an important concept to grasp, and to illustrate
why we will examine the identity of `a`, `b`,
and `c`.

```python 
print(f'a id: {id(a)}') # original object')  # a id: 139886603774600
print(f'b id: {id(b)}') # alias of a')       # b id: 139886603774600
print(f'c id: {id(c)}') # copy of a')        # c id: 139886603774600
```

All aliases of `a` have the same identity which in python
is unique integer representing its C memory address. If any change were
to be made the identity integer would also change to reflect that.

### When is == true?

Let's check out Equality and Identity (and aliases, too)

An object's identity never changes once it has been created. However
its values might, and generally this is what we care about more. Python
gives us the option to check either like so:

```python
a == b # compares the values
a is b # compares the identities
```

Lets extend this using a more complex example using some dictionaries.

```python
batman = {'name': 'Bruce Wayne', 'job': 'crime fighter'}
bruce = batman
print(batman == bruce)  # True
print(batman is bruce)  # True
```

Both `batman` and `bruce` are equal in
identity, and their values. Suppose we have a vigilante crime fighter
out there pretending to be `batman`, named
`manbat`, does he have the same equality?

```python 
manbat = {'name': 'Bruce Wayne', 'job': 'crime fighter'}
print(batman == manbat) # True
print(batman is manbat) # False
```

In this case, both `manbat` and `batman` share
equal values but not the same identity. `manbat` is not an
alias of `bruce` or `batman`, and thus has his
own unique identity. This is because we created an entirely new identity
albeit with the same values as batman.

Much of the time we care mostly about the values an object holds not its
identity but you will see `is` in a lot during conditionals
such as:

```python 
if x is None:
  do something
if x is not None:
  do something else
```

### Alias Issues

Something I didn't realise until it came back to haunt me much later is
that aliases can have unintended side effects with mutable types. Let's
say we have two lists, the original and its alias. The alias will have
items added to it but we want the original untouched for whatever
reason.

```python 
orig = [10, 20, 30, [100, 200]]
new = orig
```

Looks good, we can now make changes to `new`.

```python 
new.append('FizzBuzz')
print(orig) # [10, 20, 30, [100, 200], 'FizzBuzz']
print(new)  # [10, 20, 30, [100, 200], 'FizzBuzz']
```

After appending to `new` it becomes apparent that this
change has affected both lists. This happens because the alias works two
way with mutable types. I think this is really important to know -
aliases are not copies!

### Copies

If aliases aren't copies then how do we copy?

```python 
orig = [10, 20, 30, [100, 200]]
new = list(orig)
# dict(x) also works this way
print('orig id:', id(orig)) # orig id: 140443406513496
print('new id:', id(new))   # new id:  140443402343535
```

By using the `list()` class we successfully create two new
objects. Now if we append or remove items from either list it does not
propagate through. Except, it does sometimes.

In this case we are only making a new copy of the overall object but not
any **mutable** nested types within the copy. So while any changes made
within the first layer of the object are contained within the copy, any
mutable objects nested more deeply will be aliases.

Confused, an example.

```python 
orig = [10, 20, 30, [100, 200]]
new = list(orig)
new.append('not nested')
print(orig) # [10, 20, 30, [100, 200]]
print(new)  # [10, 20, 30, [100, 200], 'not nested']
# first layer is not affected as it is a copy, not an alias
orig[-1].append('i am aliased to orig')
print(orig) # [10, 20, 30, [100, 200, 'i am aliased to a']]
print(new)  # [10, 20, 30, [100, 200, 'i am aliased to a'], 'not nested']
```

While the `orig` and `new` are independent of
each other when making changes to the first layer of abstraction, any
mutable types within that are simply aliases of the copies source.

Another example to check this out.

```python 
# before we started making alterations to the lists
print(id(orig))     # 140443390926984
print(id(new))      # 140443392352593
print(id(orig[-1])) # 140443395483400
print(id(new[-1]))  # 140443395483400
```

Inspecting the identities reveals that only the overall object\'s were
initialised as new objects but the nested types within were bound to the
original nested type - an alias!

This is something to take into consideration when passing variables
around that have nested types. To circumvent this immutable types such
as tuples can be used in place.

Python can do deep copies which will take care of this issue, but it has
its own drawbacks. Of which we not be discussed here as this post is
already quite long. See [Dan
Bader\'s](https://realpython.com/copying-python-objects/) excellent post
for more information.

### Wrapping Up

In python all objects have a type, identity and values. Only the values
can change after it is created and knowing a little bit more about how
this works can help us prevent unintended bugs.

**Notes:**

-   assignment does not create copies
-   nested mutable types within shallow copies are aliases
-   equality has two different checks; identity, and values
