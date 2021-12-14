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

[Ramp](https://ramp.com/?utm_source=advent-of-code) - Have you ever sped
up a real-world job 5x using software? Ramp does that for companies
every day with financial automation. We're hiring ambitious engineers
(Python, Elixir, Typescript) - join us if you like fast growth!

</div>

</div>

</div>

<div role="main">

## --- Day 14: Extended Polymerization ---

The incredible pressures at this depth are starting to put a strain on
your submarine. The submarine has
[polymerization](https://en.wikipedia.org/wiki/Polymerization) equipment
that would produce suitable materials to reinforce the submarine, and
the nearby volcanically-active caves should even have the necessary
input elements in sufficient quantities.

The submarine manual contains <span title="HO

HO -&gt; OH">instructions</span> for finding the optimal polymer
formula; specifically, it offers a *polymer template* and a list of
*pair insertion* rules (your puzzle input). You just need to work out
what polymer would result after repeating the pair insertion process a
few times.

For example:

    NNCB

    CH -> B
    HH -> N
    CB -> H
    NH -> C
    HB -> C
    HC -> B
    HN -> C
    NN -> C
    BH -> H
    NC -> B
    NB -> B
    BN -> B
    BB -> N
    BC -> B
    CC -> N
    CN -> C

The first line is the *polymer template* - this is the starting point of
the process.

The following section defines the *pair insertion* rules. A rule like
`AB -> C` means that when elements `A` and `B` are immediately adjacent,
element `C` should be inserted between them. These insertions all happen
simultaneously.

So, starting with the polymer template `NNCB`, the first step
simultaneously considers all three pairs:

-   The first pair (`NN`) matches the rule `NN -> C`, so element `C` is
    inserted between the first `N` and the second `N`.
-   The second pair (`NC`) matches the rule `NC -> B`, so element `B` is
    inserted between the `N` and the `C`.
-   The third pair (`CB`) matches the rule `CB -> H`, so element `H` is
    inserted between the `C` and the `B`.

Note that these pairs overlap: the second element of one pair is the
first element of the next pair. Also, because all pairs are considered
simultaneously, inserted elements are not considered to be part of a
pair until the next step.

After the first step of this process, the polymer becomes `NCNBCHB`.

Here are the results of a few steps using the above rules:

    Template:     NNCB
    After step 1: NCNBCHB
    After step 2: NBCCNBBBCBHCB
    After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
    After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB

This polymer grows quickly. After step 5, it has length 97; After step
10, it has length 3073. After step 10, `B` occurs 1749 times, `C` occurs
298 times, `H` occurs 191 times, and `N` occurs 865 times; taking the
quantity of the most common element (`B`, 1749) and subtracting the
quantity of the least common element (`H`, 161) produces
`1749 - 161 = 1588`.

Apply 10 steps of pair insertion to the polymer template and find the
most and least common elements in the result. *What do you get if you
take the quantity of the most common element and subtract the quantity
of the least common element?*

To play, please identify yourself via one of these services:

[\[GitHub\]](/auth/github) [\[Google\]](/auth/google)
[\[Twitter\]](/auth/twitter) [\[Reddit\]](/auth/reddit) <span
class="quiet">- [\[How Does Auth Work?\]](/about#faq_auth)</span>

</div>
