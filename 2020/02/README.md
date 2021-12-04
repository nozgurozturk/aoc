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

[Bolt](https://careers.bolt.eu/positions?team=engineering&utm_source=adventofcode&utm_medium=banner&utm_campaign=adventofcode_traffic) -
Fastest growing mobility startup in Europe: distributed systems,
microservices, ML, complex algorithms for demand prediction, and much
more!

</div>

</div>

</div>

<div role="main">

## --- Day 2: Password Philosophy ---

Your flight departs in a few days from the coastal airport; the easiest
way down to the coast from here is via
[toboggan](https://en.wikipedia.org/wiki/Toboggan).

The shopkeeper at the North Pole Toboggan Rental Shop is having a bad
day. "Something's wrong with our computers; we can't log in!" You ask if
you can take a look.

Their password database seems to be a little corrupted: some of the
passwords wouldn't have been allowed by the <span
title="To ensure your safety, your password must be the following string...">Official
Toboggan Corporate Policy</span> that was in effect when they were
chosen.

To try to debug the problem, they have created a list (your puzzle
input) of *passwords* (according to the corrupted database) and *the
corporate policy when that password was set*.

For example, suppose you have the following list:

    1-3 a: abcde
    1-3 b: cdefg
    2-9 c: ccccccccc

Each line gives the password policy and then the password. The password
policy indicates the lowest and highest number of times a given letter
must appear for the password to be valid. For example, `1-3 a` means
that the password must contain `a` at least `1` time and at most `3`
times.

In the above example, `2` passwords are valid. The middle password,
`cdefg`, is not; it contains no instances of `b`, but needs at least
`1`. The first and third passwords are valid: they contain one `a` or
nine `c`, both within the limits of their respective policies.

*How many passwords are valid* according to their policies?

To play, please identify yourself via one of these services:

[\[GitHub\]](/auth/github) [\[Google\]](/auth/google)
[\[Twitter\]](/auth/twitter) [\[Reddit\]](/auth/reddit) <span
class="quiet">- [\[How Does Auth Work?\]](/about#faq_auth)</span>

</div>
