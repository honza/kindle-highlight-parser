package src

import (
	"time"
)

type Author struct {
	name string
}

type Book struct {
	title  string
	author Author
}

type Location struct {
	start int
	end   int
	page  int
}

type Highlight struct {
	book      Book
	location  Location
	timestamp time.Time
	content   string
}

type Highlights []Highlight
