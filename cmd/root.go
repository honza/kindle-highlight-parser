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

package cmd

import (
	"bytes"
	"fmt"
	"github.com/honza/kindle-highlight-parser/src"
	"github.com/spf13/cobra"
	"os"
)

const Version = "0.2.0"

var OutputType string

var RootCmd = &cobra.Command{
	Use:     "kindle-highlight-parser <input file>",
	Short:   "kindle-highlight-parser",
	Args:    cobra.ExactArgs(1),
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		w := new(bytes.Buffer)
		result := src.RunParse(w, args[0], OutputType)
		if result != nil {
			fmt.Println("ERROR:", result)
			os.Exit(1)
		}

		fmt.Println(w)
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&OutputType, "output", "o",
		"markdown", "output format")

}
