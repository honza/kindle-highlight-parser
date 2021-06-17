Kindle highlight parser
=======================

When I read books on my Kindle, I like to highlight interesting sections.  These
highlights are stored in a file called `My Clippings.txt` on the Kindle.  This
project takes that file as input, and transforms it to a JSON, Markdown, or
org-mode output.  Other formats are planned (e.g. tex bibliography, sql, etc).

```
kindle-highlight-parser

Usage:
  kindle-highlight-parser <input file> [flags]

Flags:
  -f, --filename string   save output to a file
  -h, --help              help for kindle-highlight-parser
  -o, --output string     output format: "org", "markdown", or "json" (default "markdown")
  -s, --since string      only output highlights since date (e.g. "2019-03-21")
```

Install
-------

```
go get -u github.com/honza/kindle-highlight-parser
```

Or grab a prebuilt binary from [GitHub][1].

Output types
------------

* json
* markdown
* org-mode

Todo
----

* Remove partial highlights

Examples
--------

Given this input:

```
==========
Institutes of the Christian Religion (John Calvin)
- Your Highlight on Location 388-391 | Added on Monday, February 26, 2018 2:20:39 PM

Since the perfection of blessedness consists in the knowledge of God, he has been pleased, in order that none might be excluded from the means of obtaining felicity, not only to deposit in our minds that seed of religion of which we have already spoken, but so to manifest his perfections in the whole structure of the universe, and daily place himself in our view, that we cannot open our eyes without being compelled to behold him.
==========
On the Christian Life (John Calvin)
- Your Highlight on page 13 | Location 189-191 | Added on Monday, December 18, 2017 1:39:43 PM

The great point, then, is, that we are consecrated and dedicated to God, and, therefore, should not henceforth think, speak, design, or act, without a view to his glory.
==========
Atomic Habits: Tiny Changes, Remarkable Results (James Clear)
- Your Highlight on Location 688-688 | Added on Friday, March 12, 2021 10:34:34 PM

Without good learning habits, you will always feel like you’re behind the curve.
==========
```

We can produce the following Markdown output:

``` markdown
James Clear
===========

Atomic Habits: Tiny Changes, Remarkable Results
-----------------------------------------------

> Without good learning habits, you will always feel like you’re behind the
> curve.

2021-03-12 22:34:34 +0000 UTC


John Calvin
===========

Institutes of the Christian Religion
------------------------------------

> Since the perfection of blessedness consists in the knowledge of God, he
> has been pleased, in order that none might be excluded from the means of
> obtaining felicity, not only to deposit in our minds that seed of religion of
> which we have already spoken, but so to manifest his perfections in the whole
> structure of the universe, and daily place himself in our view, that we
> cannot open our eyes without being compelled to behold him.

2018-02-26 14:20:39 +0000 UTC

On the Christian Life
---------------------

> The great point, then, is, that we are consecrated and dedicated to God,
> and, therefore, should not henceforth think, speak, design, or act, without a
> view to his glory.

Location: 189-191, Page: 13, 2017-12-18 13:39:43 +0000 UTC
```

Or, the following org-mode content:

``` org
* James Clear
** Atomic Habits: Tiny Changes, Remarkable Results
*** <2021-03-12 Fri 22:34:34>
Without good learning habits, you will always feel like you’re behind the
curve.
* John Calvin
** Institutes of the Christian Religion
*** <2018-02-26 Mon 14:20:39>
Since the perfection of blessedness consists in the knowledge of God, he has
been pleased, in order that none might be excluded from the means of obtaining
felicity, not only to deposit in our minds that seed of religion of which we
have already spoken, but so to manifest his perfections in the whole structure
of the universe, and daily place himself in our view, that we cannot open our
eyes without being compelled to behold him.
** On the Christian Life
*** <2017-12-18 Mon 13:39:43>
The great point, then, is, that we are consecrated and dedicated to God, and,
therefore, should not henceforth think, speak, design, or act, without a view
to his glory.
```

License
-------

GPLv3


[1]: https://github.com/honza/kindle-highlight-parser/releases
