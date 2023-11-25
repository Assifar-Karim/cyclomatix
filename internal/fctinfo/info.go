package fsinfo

import "fmt"

type FctInfo struct {
	pkgName         string
	fctName         string
	filename        string
	cyclomaticCmplx int32
	// cfg attribute
}

func (f FctInfo) Print() {
	fmt.Printf("%s %s %s", f.pkgName, f.fctName, f.filename)
}
