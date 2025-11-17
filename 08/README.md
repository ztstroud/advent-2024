# 08

## Approaches

I see two broad strategies. One is to look at the spaces and see if they are an
antinode by looking at the antennas. The second is to look at antennas, find
antinodes for pairs, and dedupe the results.

I think the first solution seems better. You can exit early the first time you
find an antinode. Both will require you to search over many antennas, and would
benefit from speed up by collecting lists of antennas at different frequencies.
I am going to start from this approach.

## By Groups is Better

Looking at candidates has O(NM) runtime in the worst case, and the worst
case is if all the tiles have any flag on them, which behaves like
O(N^2).

Looking at groups of antennas gives something different, O(N+M^2). This
still could behave as O(N^2), but only if ALL antennas are in the same
group.

