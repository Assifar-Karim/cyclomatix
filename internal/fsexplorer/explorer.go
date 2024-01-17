package fsexplorer

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	fsinfo "github.com/Assifar-Karim/cyclomatix/internal/fctinfo"
)

type FileHandler interface {
	HandleFile(path string, fctTable *[]fsinfo.FctInfo)
	ComputeComplexities(fctTable *[]fsinfo.FctInfo)
}

type FileList struct {
	paths       []string
	fctTable    *[]fsinfo.FctInfo
	fileHandler FileHandler
}

func (fl FileList) Handle() {
	for _, path := range fl.paths {
		info, err := os.Stat(path)
		if err != nil {
			log.Printf("[INFO] Couldn't get %q's path information: %s\n", path, err)
			continue
		}
		if info.IsDir() {
			fl.handleDir(path)
		} else {
			fl.handleFile(path)
		}
	}
}

func (fl FileList) handleDir(path string) {
	filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() && (d.Name() == "vendor" || d.Name() == "test" || d.Name() == ".git") {
			return filepath.SkipDir
		}
		if err == nil && !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {
			fl.handleFile(p)
		}
		return err
	})
}

func (fl FileList) handleFile(path string) {
	fl.fileHandler.HandleFile(path, fl.fctTable)
}

func NewFileList(paths []string, fctTable *[]fsinfo.FctInfo, fileHandler FileHandler) FileList {
	return FileList{
		paths:       paths,
		fctTable:    fctTable,
		fileHandler: fileHandler,
	}
}
