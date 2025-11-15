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

