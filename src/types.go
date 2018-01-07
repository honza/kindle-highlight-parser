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

type JsonFormatter interface {
	json() (string, error)
}

type OutputEmitter func(hs Highlights) (string, error)

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

func HighlightEmitMarkdown(h Highlight) (string, error) {
	buf := bytes.NewBufferString("")

	buf.WriteString(h.Book.Title)
	buf.WriteString("\n")
	buf.WriteString(strings.Repeat("=", len(h.Book.Title)-3))
	buf.WriteString("\n")
	buf.WriteString("\n*")
	buf.WriteString(h.Book.Author.Name)
	buf.WriteString("*\n\n")
	buf.WriteString(word_wrap(h.Content, 79))
	buf.WriteString("\n\nPublished: ")
	buf.WriteString(h.Timestamp.String())
	buf.WriteString("\n\n")

	return buf.String(), nil
}

func EmitMarkdown(hs Highlights) (string, error) {
	result := make([]string, 0)

	for _, h := range hs {
		out, _ := HighlightEmitMarkdown(h)
		result = append(result, out)
	}

	return strings.Join(result, "\n"), nil
}
