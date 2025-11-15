# 05

I had some issues with part two. I think it is unclear that the rules will
completely specify the order of the pages. For instance, the following seems
like a valid input:

```
61|13
29|13

61,13,29
```

In this case, either 61 or 29 could be the middle, and this would impact the
sum. I was also unsure if every number would be related to all the others. You
could imagine there being numbers left out:

```
61|13

61,13,29
```

There are 3 possible solutions here, because 29 could be in the middle or at the
start or end. Similarly, You could have disjoint sets:

```
61|29
45|13
13|34

61,13,34,29,45
```

There are 10 solutions here, and in this case any value could be the middle
value.

However, I get the right solution by assuming that the ordering is completely
specified, and that seems to be the impression of other people on Reddit. But I
still don't see that in the text of the challenge. The closest is "very specific
order", but that isn't very specific.

