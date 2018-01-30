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
	"strings"
)

const MARKDOWN_WIDTH = 79

func SingleEmitMarkdown(w io.Writer, single Single) error {
	fmt.Fprint(w, word_wrap(single.Content, MARKDOWN_WIDTH, "> "))
	fmt.Fprint(w, "\n\n")
	fmt.Fprint(w, FormatLocation(single.Location))
	fmt.Fprint(w, single.Timestamp.String())
	fmt.Fprint(w, "\n\n")

	return nil
}

func AuthorEmitMarkdown(w io.Writer, author string, h NewAuthor) error {
	fmt.Fprint(w, author)
	fmt.Fprint(w, "\n")
	fmt.Fprint(w, strings.Repeat("=", len(author)))
	fmt.Fprint(w, "\n\n")

	books := make([]string, 0, len(h))

	for book := range h {
		books = append(books, book)
	}

	sort.Strings(books)

	for _, book := range books {
		bookTitle := TruncateBookTitle(book)
		fmt.Fprint(w, bookTitle)
		fmt.Fprint(w, "\n")
		fmt.Fprint(w, strings.Repeat("-", len(bookTitle)))
		fmt.Fprint(w, "\n\n")

		for _, single := range h[book] {
			SingleEmitMarkdown(w, single)
		}

	}

	return nil
}

func EmitMarkdown(w io.Writer, hs Highlights) error {
	authors := make([]string, 0, len(hs))

	for author := range hs {
		authors = append(authors, author)
	}

	sort.Strings(authors)

	for _, author := range authors {
		err := AuthorEmitMarkdown(w, author, hs[author])
		if err != nil {
			return err
		}
		fmt.Fprint(w, "\n")

	}

	return nil
}
