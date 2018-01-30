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
	"testing"
)

func TestTrue(t *testing.T) {

}

func TestSplitAndRemove(t *testing.T) {
	s := "\none\n\ntwo\n\n"
	a := splitAndRemove(s, "\n")

	if len(a) != 2 {
		t.Error("splitAndRemove returns more than it should")
	}

	s = "one"
	a = splitAndRemove(s, "\n")

	if len(a) != 1 {
		t.Error("splitAndRemove is wrong")
	}

	s = "\n\n\n"
	a = splitAndRemove(s, "\n")

	if len(a) != 0 {
		t.Error("splitAndRemove is wrong")
	}

	s = "Name's name (Author)"
	a = splitAndRemove(s, "(")

	if len(a) != 2 {
		t.Error("splitAndRemove is wrong")
	}

}

func TestLast(t *testing.T) {
	s := []string{"1", "2"}
	a := last(s)

	if a != "2" {
		t.Error("TestLast fail")
	}

	s = []string{"1"}
	a = last(s)

	if a != "1" {
		t.Error("TestLast fail")
	}
}

func TestTrim(t *testing.T) {
	s := "(Name)"
	a := trim(s)

	if a != "Name" {
		t.Error("trim doesn't remove ()")
	}

	s = "Name)"
	a = trim(s)

	if a != "Name" {
		t.Error("trim doesn't remove )")
	}

	s = "   Name  "
	a = trim(s)

	if a != "Name" {
		t.Error("trim doesn't remove spaces")
	}

	s = "   (Name)  "
	a = trim(s)

	if a != "Name" {
		t.Error("trim doesn't remove spaces or ()")
	}

	s = "Name)\n"
	a = trim(s)

	if a != "Name" {
		t.Error("trim doesn't remove ) and \n")
	}
}
