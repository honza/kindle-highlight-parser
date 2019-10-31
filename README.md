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

License
-------

GPLv3


[1]: https://github.com/honza/kindle-highlight-parser/releases
