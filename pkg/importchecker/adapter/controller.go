package adapter

import (
	"github.com/syuparn/pkgaffinity/interfaces"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/usecase"
)

type controller struct {
	checkImportsInputPort usecase.CheckImportsInputPort
}

func NewController(checkImportsInputPort usecase.CheckImportsInputPort) interfaces.ImportChecker {
	return &controller{
		checkImportsInputPort: checkImportsInputPort,
	}
}

// impl check
var _ interfaces.ImportChecker = &controller{}

func (c *controller) CheckImports(req *interfaces.CheckImportRequest) error {
	in := &usecase.CheckImportsInputData{
		PackagePath: req.PackagePath,
		ImportPaths: req.ImportPaths,
	}

	return c.checkImportsInputPort.Exec(in)
}
