package src

import (
	"errors"
	"fmt"
	"io/ioutil"
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
	return strings.Trim(s, " ()")

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
	author := Author{name: trim(last(parts))}
	book := Book{title: title, author: author}

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

		location = Location{start: locStart, end: locEnd, page: page}
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
		location = Location{start: locStart, end: locEnd, page: 0}
	default:
		location = Location{start: 0, end: 0, page: 0}
	}

	return Highlight{book: book, location: location, timestamp: timestamp,
		content: sublines[2]}, nil

}

func Parse(fileContents []byte) (Highlights, error) {
	fileContentsString := byteArrayToString(fileContents)
	lines := splitAndRemove(fileContentsString, "==========")
	highlights := make([]Highlight, 0)

	for _, line := range lines {
		fmt.Println("===== start")
		highlight, err := ParseHighlight(line)

		if err != nil {
			continue
		}
		highlights = append(highlights, highlight)
	}

	return highlights, nil
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

	data, err := Parse(fileContents)
	fmt.Println(" === data ===")
	fmt.Println(data)

	// parse file
	// clean up results
	// print

	return nil
}
