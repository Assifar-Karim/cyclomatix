package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	fsinfo "github.com/Assifar-Karim/cyclomatix/internal/fctinfo"
	"github.com/Assifar-Karim/cyclomatix/internal/fsexplorer"
	"github.com/dominikbraun/graph/draw"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cfgVizCommand)
}

var outputDir string

func init() {
	cfgVizCommand.Flags().StringVarP(&outputDir, "output", "o", ".", "Defines the output directory"+
		" where the cfg dot files will be generated")
	cfgVizCommand.Flags().StringArrayVarP(&files, "files", "f", []string{}, "Defines the files/directory to analyze")
}

var cfgVizCommand *cobra.Command = &cobra.Command{
	Use:   "cfg",
	Short: "Generate the control flow graphs of all functions in the input files as dot files",
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
		printBanner()
		fileExplorer.Handle()
		if outputDir != "." {
			os.Mkdir(outputDir, 0755)
		}
		for _, fct := range functionTable {
			fctPath := filepath.Join(outputDir, fct.GetPkg())
			os.Mkdir(fctPath, 0755)
			dotGraph := fct.GetCfg().GenerateDot()
			fpath := filepath.Join(fctPath, fct.GetName()+".gv")
			file, _ := os.Create(fpath)
			draw.DOT[string, string](dotGraph, file)
			fmt.Printf("[INFO]: The DOT Graph for %v was generated.\n", fct.GetPkg()+"/"+fct.GetName())
			cmd := exec.Command("dot", "-Tpng", "-O", fpath)
			err := cmd.Run()
			if err != nil {
				log.Println(err)
			}
		}
	},
}
