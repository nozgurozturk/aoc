<div>

<div>

# [Advent of Code](/)

-   [\[About\]](/2020/about)
-   [\[Events\]](/2020/events)
-   [\[Shop\]](https://teespring.com/stores/advent-of-code)
-   [\[Log In\]](/2020/auth/login)

</div>

<div>

#    <span class="title-event-wrap">sub y{</span>[2020](/2020)<span class="title-event-wrap">}</span>

-   [\[Calendar\]](/2020)
-   [\[AoC++\]](/2020/support)
-   [\[Sponsors\]](/2020/sponsors)
-   [\[Leaderboard\]](/2020/leaderboard)
-   [\[Stats\]](/2020/stats)

</div>

</div>

<div id="sidebar">

<div id="sponsor">

<div class="quiet">

Our [sponsors](/2020/sponsors) help make Advent of Code possible:

</div>

<div class="sponsor">

[Digit](https://digit.co/) - Help people automate their money and become
financially healthy! Interested in fintech? Email us at aoc@digit.co

</div>

</div>

</div>

<div role="main">

## --- Day 3: Toboggan Trajectory ---

With the toboggan login problems resolved, you set off toward the
airport. While travel by toboggan might be easy, it's certainly not
safe: there's <span
title="It looks like the toboggan steering system even runs on Intcode! Good thing you don't have to modify it.">very
minimal steering</span> and the area is covered in trees. You'll need to
see which angles will take you near the fewest trees.

Due to the local geology, trees in this area only grow on exact integer
coordinates in a grid. You make a map (your puzzle input) of the open
squares (`.`) and trees (`#`) you can see. For example:

    ..##.......
    #...#...#..
    .#....#..#.
    ..#.#...#.#
    .#...##..#.
    ..#.##.....
    .#.#.#....#
    .#........#
    #.##...#...
    #...##....#
    .#..#...#.#

These aren't the only trees, though; due to something you read about
once involving arboreal genetics and biome stability, the same pattern
repeats to the right many times:

    ..##.........##.........##.........##.........##.........##.......  --->
    #...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
    .#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
    ..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
    .#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
    ..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....  --->
    .#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
    .#........#.#........#.#........#.#........#.#........#.#........#
    #.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
    #...##....##...##....##...##....##...##....##...##....##...##....#
    .#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#  --->

You start on the open square (`.`) in the top-left corner and need to
reach the bottom (below the bottom-most row on your map).

The toboggan can only follow a few specific slopes (you opted for a
cheaper model that prefers rational numbers); start by *counting all the
trees* you would encounter for the slope *right 3, down 1*:

From your starting position at the top-left, check the position that is
right 3 and down 1. Then, check the position that is right 3 and down 1
from there, and so on until you go past the bottom of the map.

The locations you'd check in the above example are marked here with `O`
where there was an open square and `X` where there was a tree:

    ..##.........##.........##.........##.........##.........##.......  --->
    #..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
    .#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
    ..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
    .#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
    ..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->
    .#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
    .#........#.#........X.#........#.#........#.#........#.#........#
    #.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...
    #...##....##...##....##...#X....##...##....##...##....##...##....#
    .#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#  --->

In this example, traversing the map using this slope would cause you to
encounter `7` trees.

Starting at the top-left corner of your map and following a slope of
right 3 and down 1, *how many trees would you encounter?*

To play, please identify yourself via one of these services:

[\[GitHub\]](/auth/github) [\[Google\]](/auth/google)
[\[Twitter\]](/auth/twitter) [\[Reddit\]](/auth/reddit) <span
class="quiet">- [\[How Does Auth Work?\]](/about#faq_auth)</span>

</div>
