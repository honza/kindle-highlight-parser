package src

import (
	"fmt"
	"io"
	"sort"
)

const ORG_WIDTH = 79

func SingleEmitOrg(w io.Writer, single Single) error {
	fmt.Fprint(w, "\n")
	fmt.Fprint(w, word_wrap(single.Content, ORG_WIDTH, "    "))
	fmt.Fprint(w, "\n\n")
	fmt.Fprintf(w, "    %s", FormatLocation(single.Location))
	fmt.Fprint(w, single.Timestamp.String())
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
