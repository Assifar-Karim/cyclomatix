package fsexplorer

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileList struct {
	paths []string
}

func (fl FileList) Handle() {
	for _, path := range fl.paths {
		info, err := os.Stat(path)
		if err != nil {
			log.Printf("[INFO] Couldn't get %q's path information: %s\n", path, err)
			continue
		}
		if info.IsDir() {
			handleDir(path)
		} else {
			handleFile(path)
		}
	}
}

func handleDir(path string) {
	filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() && (d.Name() == "vendor" || d.Name() == "test" || d.Name() == ".git") {
			return filepath.SkipDir
		}
		if err == nil && !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {
			handleFile(p)
		}
		return err
	})
}

func handleFile(path string) {
	// This is where the function handling logic will be injected
}

func NewFileList(paths []string) FileList {
	return FileList{
		paths: paths,
	}
}
