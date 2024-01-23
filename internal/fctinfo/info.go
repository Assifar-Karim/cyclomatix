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
	CallList        map[string]int
	IsVisited       bool
}

func (f FctInfo) Print() {
	fmt.Printf("%-10s %-10s %-20s %v\n", f.pkgName, f.fctName, f.filename, f.cyclomaticCmplx)
}

func (f FctInfo) GetCfg() utils.Graph {
	return f.cfg
}

func GetFctByNameAndPkg(fctTable *[]FctInfo, name string, pkg string) (*FctInfo, error) {
	for _, fct := range *fctTable {
		if name == fct.fctName && pkg == fct.pkgName {
			return &fct, nil
		}
	}
	return nil, fmt.Errorf("function not found")
}

func (f *FctInfo) SetAsVisited() {
	f.IsVisited = true
}

func (f FctInfo) GetCycloCmplx() int32 {
	return f.cyclomaticCmplx
}

func (f *FctInfo) SetCycloCmplx(value int32) {
	f.cyclomaticCmplx = value
}

func NewFctInfo(pkgName string, fctName string, filename string, cfg utils.Graph, callList map[string]int) FctInfo {
	return FctInfo{
		pkgName:   pkgName,
		fctName:   fctName,
		filename:  filename,
		cfg:       cfg,
		CallList:  callList,
		IsVisited: false,
	}
}
