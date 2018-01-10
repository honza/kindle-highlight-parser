package src

import (
	"bytes"
	"encoding/json"
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

type JsonFormatter interface {
	json() (string, error)
}

func (h Highlight) json() (string, error) {
	out, err := json.MarshalIndent(h, "", "  ")

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (hs Highlights) json() (string, error) {
	out, err := json.MarshalIndent(hs, "", "  ")

	if err != nil {
		return "", err

	}

	return string(out), nil
}

func (hs NewHighlights) json() (string, error) {
	out, err := json.MarshalIndent(hs, "", "  ")

	if err != nil {
		return "", err

	}

	return string(out), nil
}

// https://gist.github.com/kennwhite/306317d81ab4a885a965e25aa835b8ef
func word_wrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped

}

func SingleEmitMarkdown(buf *bytes.Buffer, single Single) error {
	buf.WriteString(word_wrap(single.Content, 79))
	buf.WriteString("\n\n")
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

	buf.WriteString(single.Timestamp.String())
	buf.WriteString("\n\n")

	return nil
}

func AuthorEmitMarkdown(buf *bytes.Buffer, author string, h NewAuthor) error {
	buf.WriteString(author)
	buf.WriteString("\n")
	buf.WriteString(strings.Repeat("=", len(author)))
	buf.WriteString("\n\n")

	for bookName, book := range h {
		buf.WriteString(bookName)
		buf.WriteString("\n")
		buf.WriteString(strings.Repeat("-", len(bookName)))
		buf.WriteString("\n\n")

		for _, single := range book {
			SingleEmitMarkdown(buf, single)
		}

	}

	return nil
}

func EmitMarkdown(hs NewHighlights) (string, error) {
	buf := bytes.NewBufferString("")

	for author := range hs {
		err := AuthorEmitMarkdown(buf, author, hs[author])
		if err != nil {
			return "", err
		}
		buf.WriteString("\n")

	}

	return buf.String(), nil
}
