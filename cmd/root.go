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
	"bufio"
	"fmt"
	"github.com/honza/kindle-highlight-parser/src"
	"github.com/spf13/cobra"
	"os"
)

const Version = "0.3.0"

var OutputType string
var Since string
var Filename string

var RootCmd = &cobra.Command{
	Use:     "kindle-highlight-parser <input file>",
	Short:   "kindle-highlight-parser",
	Args:    cobra.ExactArgs(1),
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		var writer *bufio.Writer
		var file *os.File

		if Filename != "" {
			_, err := os.Stat(Filename)

			if err == nil {
				fmt.Println("ERROR: File exists")
				os.Exit(1)
			}

			file, err = os.Create(Filename)

			if err != nil {
				fmt.Println("ERROR:", err)
				os.Exit(1)
			}
		} else {
			file = os.Stdout
		}

		writer = bufio.NewWriter(file)

		defer file.Close()
		defer writer.Flush()

		result := src.RunParse(writer, args[0], OutputType, Since)

		if result != nil {
			fmt.Println("ERROR:", result)
			os.Exit(1)
		}

	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&OutputType, "output", "o",
		"markdown", `output format: "org", "markdown", or "json"`)
	RootCmd.PersistentFlags().StringVarP(&Since, "since", "s",
		"", "only output highlights since date (e.g. \"2019-03-21\")")
	RootCmd.PersistentFlags().StringVarP(&Filename, "filename", "f",
		"", `save output to a file`)

}
