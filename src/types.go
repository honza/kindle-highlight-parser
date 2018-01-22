package src

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

type Author struct {
	Name string `json:"name"`
}

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Location struct {
	Start int `json:"start"`
	End   int `json:"end"`
	Page  int `json:"page"`
}

type Highlight struct {
	Book      Book      `json:"book"`
	Location  Location  `json:"location"`
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
}

type Highlights []Highlight

type Single struct {
	Location  Location  `json:"location"`
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
}
type NewBook []Single
type NewAuthor map[string]NewBook
type NewHighlights map[string]NewAuthor

func (b NewBook) Len() int {
	return len(b)
}

func (b NewBook) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b NewBook) Less(i, j int) bool {
	return b[i].Timestamp.Before(b[j].Timestamp)
}

func (hs NewHighlights) json() ([]byte, error) {
	out, err := json.MarshalIndent(hs, "", "  ")

	if err != nil {
		return []byte{}, err

	}

	return out, nil
}

// https://gist.github.com/kennwhite/306317d81ab4a885a965e25aa835b8ef
func word_wrap(text string, lineWidth int, prefix string) string {
	words := strings.Fields(strings.TrimSpace(text))

	lineWidth = lineWidth - len(prefix)
	if len(words) == 0 {
		return text
	}
	wrapped := prefix + words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + prefix + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped

}

func FormatLocation(location Location) string {
	if location.Start != 0 && location.End != 0 && location.Page != 0 {
		return fmt.Sprintf("Location: %d-%d, Page: %d, ", location.Start, location.End, location.Page)
	}

	if location.Start != 0 && location.End != 0 {
		return fmt.Sprintf("Location: %d-%d, ", location.Start, location.End)
	}
	if location.Start != 0 {
		return fmt.Sprintf("Location: %d, ", location.Start)
	}
	if location.Page != 0 {
		return fmt.Sprintf("Page: %d, ", location.Page)
	}
	return fmt.Sprintf("")
}

func SingleEmitMarkdown(w io.Writer, single Single) error {
	fmt.Fprint(w, word_wrap(single.Content, 79, "> "))
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
		fmt.Fprint(w, book)
		fmt.Fprint(w, "\n")
		fmt.Fprint(w, strings.Repeat("-", len(book)))
		fmt.Fprint(w, "\n\n")

		for _, single := range h[book] {
			SingleEmitMarkdown(w, single)
		}

	}

	return nil
}

func EmitMarkdown(w io.Writer, hs NewHighlights) error {
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
