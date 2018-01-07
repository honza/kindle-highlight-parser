package src

import (
	"encoding/json"
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
