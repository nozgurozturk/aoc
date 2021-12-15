<div>

<div>

# [Advent of Code](/)

-   [\[About\]](/2021/about)
-   [\[Events\]](/2021/events)
-   [\[Shop\]](https://teespring.com/stores/advent-of-code)
-   [\[Log In\]](/2021/auth/login)

</div>

<div>

#  <span class="title-event-wrap">{'year':</span>[2021](/2021)<span class="title-event-wrap">}</span>

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

[Replit](https://2021-aoc-templates.util.repl.co/) - Code and host in
your browser with no setup in Python, React, Kaboom.js, Java, C, Nix,
you name it, even Solidity. Happy coding!

</div>

</div>

</div>

<div role="main">

## --- Day 15: Chiton ---

You've almost reached the exit of the cave, but the walls are getting
closer together. Your submarine can barely still fit, though; the main
problem is that the walls of the cave are covered in
[chitons](https://en.wikipedia.org/wiki/Chiton), and it would be best
not to bump any of them.

The cavern is large, but has a very low ceiling, restricting your motion
to two dimensions. The shape of the cavern resembles a square; a quick
scan of chiton density produces a map of *risk level* throughout the
cave (your puzzle input). For example:

    1163751742
    1381373672
    2136511328
    3694931569
    7463417111
    1319128137
    1359912421
    3125421639
    1293138521
    2311944581

You start in the top left position, your destination is the bottom right
position, and you <span
title="Can't go diagonal until we can repair the caterpillar unit. Could be the liquid helium or the superconductors.">cannot
move diagonally</span>. The number at each position is its *risk level*;
to determine the total risk of an entire path, add up the risk levels of
each position you *enter* (that is, don't count the risk level of your
starting position unless you enter it; leaving it adds no risk to your
total).

Your goal is to find a path with the *lowest total risk*. In this
example, a path with the lowest total risk is highlighted here:

    1163751742
    1381373672
    2136511328
    3694931569
    7463417111
    1319128137
    1359912421
    3125421639
    1293138521
    2311944581

The total risk of this path is `40` (the starting position is never
entered, so its risk is not counted).

*What is the lowest total risk of any path from the top left to the
bottom right?*

To play, please identify yourself via one of these services:

[\[GitHub\]](/auth/github) [\[Google\]](/auth/google)
[\[Twitter\]](/auth/twitter) [\[Reddit\]](/auth/reddit) <span
class="quiet">- [\[How Does Auth Work?\]](/about#faq_auth)</span>

</div>
