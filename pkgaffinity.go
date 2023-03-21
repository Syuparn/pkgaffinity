package pkgaffinity

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"

	"github.com/samber/lo"
	"github.com/syuparn/pkgaffinity/interfaces"
)

const doc = "pkgaffinity checks import statements which will break encapsulations"

func NewAnalyzer() *analysis.Analyzer {
	runner := &analysisRunner{importChecker: getImportChecker()}

	return &analysis.Analyzer{
		Name: "pkgaffinity",
		Doc:  doc,
		Run:  runner.run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}
}

type analysisRunner struct {
	importChecker interfaces.ImportChecker
}

func (r *analysisRunner) run(pass *analysis.Pass) (any, error) {
	// NOTE: due to performance, this does not use inspectors
	// (result is written to stdout instead of reporter)
	req := &interfaces.CheckImportRequest{
		PackagePath: pass.Pkg.Path(),
		ImportPaths: lo.Map(pass.Pkg.Imports(), func(p *types.Package, _ int) string {
			return p.Path()
		}),
	}

	err := r.importChecker.CheckImports(req)
	if err != nil {
		return nil, fmt.Errorf("failed to check package: %+v: %w", req, err)
	}

	return nil, nil
}
