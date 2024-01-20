package fsinfo

import (
	"fmt"

	"github.com/Assifar-Karim/cyclomatix/internal/utils"
)

type FctInfo struct {
	pkgName         string
	fctName         string
	filename        string
	cyclomaticCmplx int32
	cfg             utils.Graph
}

func (f FctInfo) Print() {
	fmt.Printf("%s %s %s\n", f.pkgName, f.fctName, f.filename)
}

func (f FctInfo) GetCfg() utils.Graph {
	return f.cfg
}

func (f *FctInfo) SetCycloCmplx(value int32) {
	f.cyclomaticCmplx = value
}

func NewFctInfo(pkgName string, fctName string, filename string, cfg utils.Graph) FctInfo {
	return FctInfo{
		pkgName:  pkgName,
		fctName:  fctName,
		filename: filename,
		cfg:      cfg,
	}
}
