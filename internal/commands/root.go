package commands

import (
	"fmt"
	"os"

	fsinfo "github.com/Assifar-Karim/cyclomatix/internal/fctinfo"
	"github.com/Assifar-Karim/cyclomatix/internal/fsexplorer"
	"github.com/spf13/cobra"
)

var files []string
var fileExplorer FileExplorer
var functionTable []fsinfo.FctInfo
var fileHandler fsexplorer.FileHandler

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "cyclo",
	Short: "A static code analysis tool that computes the cyclomatic complexity of functions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func printBanner() {
	fmt.Println("   ________  __________    ____  __  ______  ___________  __")
	fmt.Println("  / ____/\\ \\/ / ____/ /   / __ \\/  |/  /   |/_  __/  _/ |/ /")
	fmt.Println(" / /      \\  / /   / /   / / / / /|_/ / /| | / /  / / |   / ")
	fmt.Println("/ /___    / / /___/ /___/ /_/ / /  / / ___ |/ / _/ / /   |  ")
	fmt.Println("\\____/   /_/\\____/_____/\\____/_/  /_/_/  |_/_/ /___//_/|_|  ")
	fmt.Println()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
