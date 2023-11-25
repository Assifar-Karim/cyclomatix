package commands

import (
	fsinfo "github.com/Assifar-Karim/cyclomatix/internal/fctinfo"
	"github.com/Assifar-Karim/cyclomatix/internal/fsexplorer"
	"github.com/spf13/cobra"
)

type FileExplorer interface {
	Handle()
}

func init() {
	rootCmd.AddCommand(complexityCmd)
}

var indirectionLvl int32
var files []string
var fileExplorer FileExplorer
var functionTable []fsinfo.FctInfo
var FileHandler fsexplorer.FileHandler

func init() {
	complexityCmd.Flags().Int32VarP(&indirectionLvl, "indirection-lvl", "i", 4, "Sets the maximum allowed level of indirection")
	complexityCmd.Flags().StringArrayVarP(&files, "files", "f", []string{}, "Defines the files/directory to analyze")
}

var complexityCmd *cobra.Command = &cobra.Command{
	Use:   "complexity",
	Short: "List the cyclomatic complexity of all functions in the input files",
	PreRun: func(cmd *cobra.Command, args []string) {
		functionTable = []fsinfo.FctInfo{}
		fileExplorer = fsexplorer.NewFileList(
			files,
			functionTable,
			fsexplorer.NewGoFileHandler(indirectionLvl),
		)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fileExplorer.Handle()
	},
}
