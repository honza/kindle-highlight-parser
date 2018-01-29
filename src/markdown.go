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
