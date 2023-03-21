package adapter

import (
	"fmt"
	"io"

	"github.com/samber/lo"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/domain"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/usecase"
)

type checkImportsPresenter struct {
	writer io.Writer
}

// impl check
var _ usecase.CheckImportsOutputPort = &checkImportsPresenter{}

func NewCheckImportsOutputPort(writer io.Writer) usecase.CheckImportsOutputPort {
	return &checkImportsPresenter{
		writer: writer,
	}
}

func (p *checkImportsPresenter) Present(out *usecase.CheckImportsOutputData) {
	lo.ForEach(out.Violations, func(v *domain.Violation, _ int) {
		fmt.Fprintf(p.writer, "package %s: import \"%s\" breaks %s\n", v.PackagePath, v.ImportPath, v.RuleName)
	})
}
