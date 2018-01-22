package src

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ValidateOutputFormat(output string) error {
	validOuputs := []string{"markdown", "json"}

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

		res = append(res, v)
	}

	return res

}

func trim(s string) string {
	return strings.Trim(s, " ()\ufeff")

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

	if len(sublines) != 3 {
		return Highlight{}, errors.New("Missing content")
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

func Parse(fileContents []byte) (NewHighlights, error) {
	fileContentsString := byteArrayToString(fileContents)
	lines := splitAndRemove(fileContentsString, "==========")
	highlights := NewHighlights{}

	for _, line := range lines {
		highlight, err := ParseHighlight(line)

		if err != nil {
			continue
		}

		name := highlight.Book.Author.Name
		title := highlight.Book.Title

		existing, present := highlights[name]

		if present {
			existingTitle, presentTitle := existing[title]

			single := Single{Location: highlight.Location, Timestamp: highlight.Timestamp, Content: highlight.Content}

			if presentTitle {
				existingTitle = append(existingTitle, single)
				sort.Sort(existingTitle)
			} else {
				existingTitle = []Single{single}
			}
			existing[title] = existingTitle
		} else {
			existing = NewAuthor{title: NewBook{}}
		}

		highlights[name] = existing

	}

	return highlights, nil
}

func Format(w io.Writer, data NewHighlights, format string) error {
	switch format {
	case "json":
		out, err := data.json()
		if err != nil {
			return err
		}
		w.Write(out)
	case "markdown":
		err := EmitMarkdown(w, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func RunParse(filename string, output string) error {
	v := ValidateOutputFormat(output)
	if v != nil {
		return v
	}

	fileContents, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	w := new(bytes.Buffer)

	data, err := Parse(fileContents)
	err = Format(w, data, output)

	if err != nil {
		return err
	}

	fmt.Println(w)

	return nil
}
