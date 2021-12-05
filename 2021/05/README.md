<div>

<div>

# [Advent of Code](/)

-   [\[About\]](/2021/about)
-   [\[Events\]](/2021/events)
-   [\[Shop\]](https://teespring.com/stores/advent-of-code)
-   [\[Log In\]](/2021/auth/login)

</div>

<div>

#           <span class="title-event-wrap"></span>[2021](/2021)<span class="title-event-wrap"></span>

-   [\[Calendar\]](/2021)
-   [\[AoC++\]](/2021/support)
-   [\[Sponsors\]](/2021/sponsors)
-   [\[Leaderboard\]](/2021/leaderboard)
-   [\[Stats\]](/2021/stats)

</div>

</div>

<div id="sidebar">

<div id="sponsor">

<div class="quiet">

Our [sponsors](/2021/sponsors) help make Advent of Code possible:

</div>

<div class="sponsor">

[JetBrains](https://jb.gg/AoC2021tips) - Get ready to jingle with Advent
of Code in Kotlin! Have fun, learn new things, and win prizes. Believe
in magic with Kotlin. Happy holidays! https://jb.gg/AoC

</div>

</div>

</div>

<div role="main">

## --- Day 5: Hydrothermal Venture ---

You come across a field of [hydrothermal
vents](https://en.wikipedia.org/wiki/Hydrothermal_vent) on the ocean
floor! These vents constantly produce large, opaque clouds, so it would
be best to avoid them if possible.

They tend to form in *lines*; the submarine helpfully produces a list of
nearby <span title="Maybe they're Bresenham vents.">lines of
vents</span> (your puzzle input) for you to review. For example:

    0,9 -> 5,9
    8,0 -> 0,8
    9,4 -> 3,4
    2,2 -> 2,1
    7,0 -> 7,4
    6,4 -> 2,0
    0,9 -> 2,9
    3,4 -> 1,4
    0,0 -> 8,8
    5,5 -> 8,2

Each line of vents is given as a line segment in the format
`x1,y1 -> x2,y2` where `x1`,`y1` are the coordinates of one end the line
segment and `x2`,`y2` are the coordinates of the other end. These line
segments include the points at both ends. In other words:

-   An entry like `1,1 -> 1,3` covers points `1,1`, `1,2`, and `1,3`.
-   An entry like `9,7 -> 7,7` covers points `9,7`, `8,7`, and `7,7`.

For now, *only consider horizontal and vertical lines*: lines where
either `x1 = x2` or `y1 = y2`.

So, the horizontal and vertical lines from the above list would produce
the following diagram:

    .......1..
    ..1....1..
    ..1....1..
    .......1..
    .112111211
    ..........
    ..........
    ..........
    ..........
    222111....

In this diagram, the top left corner is `0,0` and the bottom right
corner is `9,9`. Each position is shown as *the number of lines which
cover that point* or `.` if no line covers that point. The top-left pair
of `1`s, for example, comes from `2,2 -> 2,1`; the very bottom row is
formed by the overlapping lines `0,9 -> 5,9` and `0,9 -> 2,9`.

To avoid the most dangerous areas, you need to determine *the number of
points where at least two lines overlap*. In the above example, this is
anywhere in the diagram with a `2` or larger - a total of `5` points.

Consider only horizontal and vertical lines. *At how many points do at
least two lines overlap?*

To play, please identify yourself via one of these services:

[\[GitHub\]](/auth/github) [\[Google\]](/auth/google)
[\[Twitter\]](/auth/twitter) [\[Reddit\]](/auth/reddit) <span
class="quiet">- [\[How Does Auth Work?\]](/about#faq_auth)</span>

</div>