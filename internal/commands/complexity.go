package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(complexityCmd)
}

var indirectionLvl int32
var files []string

func init() {
	complexityCmd.Flags().Int32VarP(&indirectionLvl, "indirection-lvl", "i", 4, "Sets the maximum allowed level of indirection")
	complexityCmd.Flags().StringArrayVarP(&files, "files", "f", []string{}, "Defines the files/directory to analyze")
}

var complexityCmd *cobra.Command = &cobra.Command{
	Use:   "complexity",
	Short: "List the cyclomatic complexity of all functions in the input files",
	Run: func(cmd *cobra.Command, args []string) {
		// This is where the complexity computation code will be injected
	},
}
