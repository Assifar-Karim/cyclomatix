package fsexplorer

import (
	fsinfo "github.com/Assifar-Karim/cyclomatix/internal/fctinfo"
)

type GoFileHandler struct {
	indirectionLvl int32
}

func (fh GoFileHandler) HandleFile(path string, fctTable []fsinfo.FctInfo) {
	// This is where the complexity computation mechanism will be implemented
}

func NewGoFileHandler(indirectionLvl int32) GoFileHandler {
	return GoFileHandler{
		indirectionLvl: indirectionLvl,
	}
}
