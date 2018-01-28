package cmd

import (
	"bytes"
	"fmt"
	"github.com/honza/kindle-highlight-parser/src"
	"github.com/spf13/cobra"
	"os"
)

var OutputType string

var RootCmd = &cobra.Command{
	Use:   "kindle-highlight-parser <input file>",
	Short: "kindle-highlight-parser",
	Args:  cobra.ExactArgs(1),
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
