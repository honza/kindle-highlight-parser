// kindle-highlight-parser
// Copyright (C) 2018  Honza Pokorny <me@honza.ca>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package src

import (
	"fmt"
	"io"
	"sort"
	"time"
)

const ORG_WIDTH = 79

func SingleEmitOrg(w io.Writer, single Single) error {
	fmt.Fprint(w, "\n")
	fmt.Fprint(w, word_wrap(single.Content, ORG_WIDTH, ""))
	fmt.Fprint(w, "\n\n")
	fmt.Fprintf(w, "%s", FormatLocation(single.Location))
	EmitTimestamp(w, single.Timestamp)
	fmt.Fprint(w, "\n\n")

	return nil
}

func AuthorEmitOrg(w io.Writer, author string, h NewAuthor) error {
	fmt.Fprintf(w, "* %s", author)
	fmt.Fprint(w, "\n")

	books := make([]string, 0, len(h))

	for book := range h {
		books = append(books, book)
	}

	sort.Strings(books)

	for _, book := range books {
		bookTitle := TruncateBookTitle(book)
		fmt.Fprintf(w, "** %s", bookTitle)
		fmt.Fprint(w, "\n")

		for _, single := range h[book] {
			SingleEmitOrg(w, single)
		}

	}

	return nil
}

func EmitOrg(w io.Writer, hs Highlights) error {
	authors := make([]string, 0, len(hs))

	for author := range hs {
		authors = append(authors, author)
	}

	sort.Strings(authors)

	for _, author := range authors {
		err := AuthorEmitOrg(w, author, hs[author])
		if err != nil {
			return err
		}
		fmt.Fprint(w, "\n")

	}

	return nil
}

func EmitTimestamp(w io.Writer, stamp time.Time) error {
	f := stamp.Format("2006-01-02 Mon 15:04:05")
	fmt.Fprintf(w, "<%s>", f)
	return nil
}
