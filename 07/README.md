# 07

## Initial Approach

My initial idea is to solve this with a recursive function. Each call
corresponds to an operator, and tries out the possibilities before passing on to
the rest of the array.

The fact that operators are always evaluated left to right is important for
this.

There is opportunity to cache results, as there are multiple ways to get to the
same input value.

