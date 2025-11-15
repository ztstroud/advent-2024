# 06

## Approach

My initial thought is to just simulate steps following the given rules until the
robot walks of the edge. I don't think you can do better.

To represent the world, I want to use a 2D array of bytes, each with three
state:
- 0 for empty
- 1 for empty but visited
- 2 for an impassible wall

The simulation keeps track of the robot position separately, and updates 0 to 1
as it goes. In the end, I can just count up the 1s.

## Looping Cases

A cup is never a loop:
```
.#.    .#.
#.#    #.#
... -> .#.
.^.    .^.
```

Even though you could place something to block the exit, it is also the entrance
and blocking that changes the path completely.

Reversing is possible legal in a loop.

```
.#.#...    .#.#...
......#    2.1...#
.....#. -> .....#.
...^...    ...^...
```

This example has a few loops. Placing an object at 1 sends you into a short 3
tile wide loop. Placing an object at 2 reveals something more tricky: even
though you can only place one block, you can introduce multiple other blocks.
This doesn't just happen when turning around.

```
.1...
....#
#..#.
.^...
```

This proves that you cannot rely on just looking for intersections with your old
path.

I think the solution is to augment looking for your old path in the main loop
with a simulation that ONLY looks for the old path. Simulate what would happen
if you turned right, and check if you collide with your path.

```
.#......
......#.
..#.....
......#.
.1......
.^...#..
```

Don't block your own old path. If blocking your old path was useful, it would
have already been discovered.

```
.##.......
......#...
.1........
.....#....
..#.......
#.........
.^..#.....
```

This also means that you will never try the same tile twice, as you always move
into the spot where you simulated placing a wall.

