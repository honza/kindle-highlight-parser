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
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
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

type Single struct {
	Location  Location  `json:"location"`
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
}
type NewBook []Single
type NewAuthor map[string]NewBook
type Highlights map[string]NewAuthor

func (b NewBook) Len() int {
	return len(b)
}

func (b NewBook) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b NewBook) Less(i, j int) bool {
	return b[i].Timestamp.Before(b[j].Timestamp)
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

func TruncateBookTitle(s string) string {
	if len(s) < MARKDOWN_WIDTH {
		return s
	}

	suffix := " [...]"
	return s[0:MARKDOWN_WIDTH-len(suffix)] + suffix

}

func ValidateOutputFormat(output string) error {
	validOuputs := []string{"markdown", "json", "org"}

	for _, v := range validOuputs {
		if output == v {
			return nil
		}
	}

	return errors.New("Invalid output type")

}

func byteArrayToString(b []byte) string {
	return string(b[:])
}

func removeEmptyStrings(s []string) []string {
	if len(s) == 0 {
		return s
	}

	res := make([]string, 0)

	for _, v := range s {
		if v == "" {
			continue
		}

		if v == "\n" {
			continue
		}

		if len(v) == 1 {
			continue
		}

		res = append(res, v)
	}

	return res

}

func trim(s string) string {
	return strings.Trim(s, " ()\ufeff\n\r")

}

func splitAndRemove(s string, sep string) []string {
	return removeEmptyStrings(strings.Split(s, sep))
}

func last(s []string) string {
	return s[len(s)-1]
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)

	if err != nil {
		return 0
	} else {
		return int(i)
	}
}

func ParseTimestamp(date string) time.Time {
	date = strings.Replace(date, " Added on ", "", 1)
	date = trim(date)

	layout := "Monday, January 2, 2006 3:04:05 PM"

	timeObject, err := time.Parse(layout, date)

	if err != nil {
		return time.Now()
	}

	return timeObject

}

func ParseHighlight(line string) (Highlight, error) {
	sublines := splitAndRemove(line, "\n")

	if len(sublines) == 0 {
		return Highlight{}, nil
	}

	isBookmark := strings.HasPrefix(sublines[1], "- Your Bookmark ")

	if isBookmark {
		return Highlight{}, nil
	}

	if len(sublines) != 3 {
		return Highlight{}, errors.New("Unable to parse line: '" + line + "'")
	}

	titleAndAuthorLine := sublines[0]
	// TODO Rename `parts` to something more intelligent
	parts := splitAndRemove(titleAndAuthorLine, "(")

	title := trim(parts[0])
	author := Author{Name: trim(last(parts))}
	book := Book{Title: title, Author: author}

	locationAndTimestampLine := sublines[1]
	parts = strings.Split(locationAndTimestampLine, "|")

	var location Location
	var timestamp time.Time

	switch len(parts) {
	case 3:

		pageStr := last(strings.Split(trim(parts[0]), " "))
		page := parseInt(pageStr)

		locationStr := last(strings.Split(trim(parts[1]), " "))
		locParts := strings.Split(locationStr, "-")
		locStart := parseInt(locParts[0])

		var locEnd int

		if len(locParts) == 2 {
			locEnd = parseInt(locParts[1])
		} else {
			locEnd = locStart
		}

		location = Location{Start: locStart, End: locEnd, Page: page}
		timestamp = ParseTimestamp(parts[2])
	case 2:
		// There is no page, just a location
		locationStr := last(strings.Split(trim(parts[1]), " "))
		locParts := strings.Split(locationStr, "-")
		locStart := parseInt(locParts[0])

		var locEnd int

		if len(locParts) == 2 {
			locEnd = parseInt(locParts[1])
		} else {
			locEnd = locStart
		}
		timestamp = ParseTimestamp(parts[1])
		location = Location{Start: locStart, End: locEnd, Page: 0}
	default:
		location = Location{Start: 0, End: 0, Page: 0}
	}

	return Highlight{Book: book, Location: location, Timestamp: timestamp, Content: sublines[2]}, nil

}

func Parse(fileContents []byte, since time.Time) (Highlights, error) {
	fileContentsString := byteArrayToString(fileContents)
	lines := splitAndRemove(fileContentsString, "==========")
	highlights := Highlights{}

	for _, line := range lines {
		highlight, err := ParseHighlight(line)

		if err != nil {
			return Highlights{}, err
		}

		if highlight.Timestamp.Before(since) {
			continue
		}

		name := highlight.Book.Author.Name
		title := highlight.Book.Title

		// Empty highlight (bookmark or similar)
		if name == "" && title == "" {
			continue
		}

		single := Single{Location: highlight.Location, Timestamp: highlight.Timestamp, Content: highlight.Content}

		existing, present := highlights[name]

		if present {
			existingTitle, presentTitle := existing[title]

			if presentTitle {
				existingTitle = append(existingTitle, single)
				sort.Sort(existingTitle)
			} else {
				existingTitle = []Single{single}
			}
			existing[title] = existingTitle
		} else {
			existing = map[string]NewBook{}
			existing[title] = []Single{single}
		}

		highlights[name] = existing

	}

	return highlights, nil
}

func Format(w io.Writer, data Highlights, format string) error {
	switch format {
	case "json":
		err := EmitJson(w, data)
		if err != nil {
			return err
		}
	case "markdown":
		err := EmitMarkdown(w, data)
		if err != nil {
			return err
		}
	case "org":
		err := EmitOrg(w, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func RunParse(w io.Writer, filename, output, since string) error {
	v := ValidateOutputFormat(output)
	if v != nil {
		return v
	}

	var sinceObj time.Time
	if since == "" {
		sinceObj = time.Unix(1, 1)

	} else {
		var err error
		sinceObj, err = time.Parse("2006-01-02", since)

		if err != nil {
			return err
		}

	}

	fileContents, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	data, err := Parse(fileContents, sinceObj)

	if err != nil {
		return err
	}

	err = Format(w, data, output)

	if err != nil {
		return err
	}

	return nil
}
