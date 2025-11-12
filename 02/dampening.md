# Dampening

The brute force solution to dampening would be to literally re-execute the check
while holding out a value. There is a super easy improvement right of the bat,
which is to find the bad index first and only check its removal.

You could also have an index that you skip, as in don't make a new slice. You
could also make it so that you can resume from an index, so you end up with a
single scan across the array.

If you don't mind modifying the slice, I do have a clever solution: swap the
invalid values one time. Imagine you have a slice [... x y z ...]. If you find
that x to y breaks the sequence, you want to check if x to z is valid and carry
on from z. You could do this by swapping x and y, giving you [... y x z ...].
The next natural iteration of the loop would check x to z, and the carry on from
z.

## Testing

This change is somewhat annoying because it changes the core definition that I
have been testing against. Basically all tests would change if isReportSafe
changes to be dampened. Instead of changing it, I will add more funcs.

## Solution

While swapping is kinda neat, I think I will try to approach this with the
resumable checks. In order to make the checks resumable, I think I need two
things: I need to know where the error was found, and I need to be able to
specify where to start. There will be a "dampened" method that coordinates
everything, and it will need to be able to execute a single check (skipping the
invalid value) when an error is found. This means I will need to pull out the
specific check logic as well.

