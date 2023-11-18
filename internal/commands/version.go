package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd *cobra.Command = &cobra.Command{
	Use:   "version",
	Short: "Print the version of cyclomatix",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cyclomatix version 0.1")
	},
}
