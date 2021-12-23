<div>

<div>

# [Advent of Code](/)

-   [\[About\]](/2021/about)
-   [\[Events\]](/2021/events)
-   [\[Shop\]](https://teespring.com/stores/advent-of-code)
-   [\[Log In\]](/2021/auth/login)

</div>

<div>

#       <span class="title-event-wrap">/\*</span>[2021](/2021)<span class="title-event-wrap">\*/</span>

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

[Panthera Investment](https://pantherainvestment.com) - We solve
Financial Markets. You want to join?

</div>

</div>

</div>

<div role="main">

## --- Day 23: Amphipod ---

A group of [amphipods](https://en.wikipedia.org/wiki/Amphipoda) notice
your fancy submarine and flag you down. "With such an impressive shell,"
one amphipod <span
title="What? You didn't know amphipods can talk?">says</span>, "surely
you can help us with a question that has stumped our best scientists."

They go on to explain that a group of timid, stubborn amphipods live in
a nearby burrow. Four types of amphipods live there: *Amber* (`A`),
*Bronze* (`B`), *Copper* (`C`), and *Desert* (`D`). They live in a
burrow that consists of a *hallway* and four *side rooms*. The side
rooms are initially full of amphipods, and the hallway is initially
empty.

They give you a *diagram of the situation* (your puzzle input),
including locations of each amphipod (`A`, `B`, `C`, or `D`, each of
which is occupying an otherwise open space), walls (`#`), and open space
(`.`).

For example:

    #############
    #...........#
    ###B#C#B#D###
      #A#D#C#A#
      #########

The amphipods would like a method to organize every amphipod into side
rooms so that each side room contains one type of amphipod and the types
are sorted `A`-`D` going left to right, like this:

    #############
    #...........#
    ###A#B#C#D###
      #A#B#C#D#
      #########

Amphipods can move up, down, left, or right so long as they are moving
into an unoccupied open space. Each type of amphipod requires a
different amount of *energy* to move one step: Amber amphipods require
`1` energy per step, Bronze amphipods require `10` energy, Copper
amphipods require `100`, and Desert ones require `1000`. The amphipods
would like you to find a way to organize the amphipods that requires the
*least total energy*.

However, because they are timid and stubborn, the amphipods have some
extra rules:

-   Amphipods will never *stop on the space immediately outside any
    room*. They can move into that space so long as they immediately
    continue moving. (Specifically, this refers to the four open spaces
    in the hallway that are directly above an amphipod starting
    position.)
-   Amphipods will never *move from the hallway into a room* unless that
    room is their destination room *and* that room contains no amphipods
    which do not also have that room as their own destination. If an
    amphipod's starting room is not its destination room, it can stay in
    that room until it leaves the room. (For example, an Amber amphipod
    will not move from the hallway into the right three rooms, and will
    only move into the leftmost room if that room is empty or if it only
    contains other Amber amphipods.)
-   Once an amphipod stops moving in the hallway, *it will stay in that
    spot until it can move into a room*. (That is, once any amphipod
    starts moving, any other amphipods currently in the hallway are
    locked in place and will not move again until they can move fully
    into a room.)

In the above example, the amphipods can be organized using a minimum of
`12521` energy. One way to do this is shown below.

Starting configuration:

    #############
    #...........#
    ###B#C#B#D###
      #A#D#C#A#
      #########

One Bronze amphipod moves into the hallway, taking 4 steps and using
`40` energy:

    #############
    #...B.......#
    ###B#C#.#D###
      #A#D#C#A#
      #########

The only Copper amphipod not in its side room moves there, taking 4
steps and using `400` energy:

    #############
    #...B.......#
    ###B#.#C#D###
      #A#D#C#A#
      #########

A Desert amphipod moves out of the way, taking 3 steps and using `3000`
energy, and then the Bronze amphipod takes its place, taking 3 steps and
using `30` energy:

    #############
    #.....D.....#
    ###B#.#C#D###
      #A#B#C#A#
      #########

The leftmost Bronze amphipod moves to its room using `40` energy:

    #############
    #.....D.....#
    ###.#B#C#D###
      #A#B#C#A#
      #########

Both amphipods in the rightmost room move into the hallway, using `2003`
energy in total:

    #############
    #.....D.D.A.#
    ###.#B#C#.###
      #A#B#C#.#
      #########

Both Desert amphipods move into the rightmost room using `7000` energy:

    #############
    #.........A.#
    ###.#B#C#D###
      #A#B#C#D#
      #########

Finally, the last Amber amphipod moves into its room, using `8` energy:

    #############
    #...........#
    ###A#B#C#D###
      #A#B#C#D#
      #########

*What is the least energy required to organize the amphipods?*

To play, please identify yourself via one of these services:

[\[GitHub\]](/auth/github) [\[Google\]](/auth/google)
[\[Twitter\]](/auth/twitter) [\[Reddit\]](/auth/reddit) <span
class="quiet">- [\[How Does Auth Work?\]](/about#faq_auth)</span>

</div>
