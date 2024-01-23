package commands

import (
	"fmt"

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
var fileHandler fsexplorer.FileHandler

func init() {
	complexityCmd.Flags().Int32VarP(&indirectionLvl, "indirection-lvl", "i", 4, "Sets the maximum allowed level of indirection")
	complexityCmd.Flags().StringArrayVarP(&files, "files", "f", []string{}, "Defines the files/directory to analyze")
}

var complexityCmd *cobra.Command = &cobra.Command{
	Use:   "complexity",
	Short: "List the cyclomatic complexity of all functions in the input files",
	PreRun: func(cmd *cobra.Command, args []string) {
		functionTable = []fsinfo.FctInfo{}
		fileHandler = fsexplorer.NewGoFileHandler(indirectionLvl)
		fileExplorer = fsexplorer.NewFileList(
			files,
			&functionTable,
			fileHandler,
		)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("   ________  __________    ____  __  ______  ___________  __")
		fmt.Println("  / ____/\\ \\/ / ____/ /   / __ \\/  |/  /   |/_  __/  _/ |/ /")
		fmt.Println(" / /      \\  / /   / /   / / / / /|_/ / /| | / /  / / |   / ")
		fmt.Println("/ /___    / / /___/ /___/ /_/ / /  / / ___ |/ / _/ / /   |  ")
		fmt.Println("\\____/   /_/\\____/_____/\\____/_/  /_/_/  |_/_/ /___//_/|_|  ")
		fmt.Println()
		fileExplorer.Handle()
		fileHandler.ComputeComplexities(&functionTable)
	},
}
