Kindle highlight parser
=======================

When I read books on my Kindle, I like to highlight interesting sections.  These
highlights are stored in a file called `My Clippings.txt` on the Kindle.  This
project takes that file as input, and transforms it to a JSON or Markdown
output.  Other formats are planned (e.g. org-mode, tex bibliography, sql, etc).

```
kindle-highlight-parser

Usage:
  kindle-highlight-parser <input file> [flags]

Flags:
  -h, --help            help for kindle-highlight-parser
  -o, --output string   output format (default "markdown")
```

Todo:

* Add installation instructions (ie `go get`)
* Remove partial highlights

License
-------

GPLv3
