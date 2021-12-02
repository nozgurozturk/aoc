<div>

<div>

# [Advent of Code](/)

-   [\[About\]](/2021/about)
-   [\[Events\]](/2021/events)
-   [\[Shop\]](https://teespring.com/stores/advent-of-code)
-   [\[Log In\]](/2021/auth/login)

</div>

<div>

#    <span class="title-event-wrap">$year=</span>[2021](/2021)<span class="title-event-wrap">;</span>

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

[Infi](https://aoc.infi.nl/?utm_source=aoc&utm_campaign=aoc2021&utm_placement=pagina) -
Santa wil zijn elfjes speelgoed laten maken, maar hij heeft zijn
administratie nog niet op orde. Kan jij hem helpen met het uitzoeken?

</div>

</div>

</div>

<div role="main">

## --- Day 2: Dive! ---

Now, you need to figure out how to <span
title="Tank, I need a pilot program for a B212 helicopter.">pilot this
thing</span>.

It seems like the submarine can take a series of commands like
`forward 1`, `down 2`, or `up 3`:

-   `forward X` increases the horizontal position by `X` units.
-   `down X` *increases* the depth by `X` units.
-   `up X` *decreases* the depth by `X` units.

Note that since you're on a submarine, `down` and `up` affect your
*depth*, and so they have the opposite result of what you might expect.

The submarine seems to already have a planned course (your puzzle
input). You should probably figure out where it's going. For example:

    forward 5
    down 5
    forward 8
    up 3
    down 8
    forward 2

Your horizontal position and depth both start at `0`. The steps above
would then modify them as follows:

-   `forward 5` adds `5` to your horizontal position, a total of `5`.
-   `down 5` adds `5` to your depth, resulting in a value of `5`.
-   `forward 8` adds `8` to your horizontal position, a total of `13`.
-   `up 3` decreases your depth by `3`, resulting in a value of `2`.
-   `down 8` adds `8` to your depth, resulting in a value of `10`.
-   `forward 2` adds `2` to your horizontal position, a total of `15`.

After following these instructions, you would have a horizontal position
of `15` and a depth of `10`. (Multiplying these together produces
`150`.)

Calculate the horizontal position and depth you would have after
following the planned course. *What do you get if you multiply your
final horizontal position by your final depth?*

To play, please identify yourself via one of these services:

[\[GitHub\]](/auth/github) [\[Google\]](/auth/google)
[\[Twitter\]](/auth/twitter) [\[Reddit\]](/auth/reddit) <span
class="quiet">- [\[How Does Auth Work?\]](/about#faq_auth)</span>

</div>
