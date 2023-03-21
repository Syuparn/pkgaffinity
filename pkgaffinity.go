package pkgaffinity

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/samber/lo"
	"github.com/syuparn/pkgaffinity/interfaces"
)

const doc = "pkgaffinity checks import statements which will break encapsulations"

var Analyzer = &analysis.Analyzer{
	Name: "pkgaffinity",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var runner = analysisRunner{importChecker: getCheckImportController()}

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

func run(pass *analysis.Pass) (any, error) {

	// TODO: use for linter
	fmt.Println(pass.Pkg.Path())
	fmt.Println(lo.Map(pass.Pkg.Imports(), func(p *types.Package, _ int) string { return p.Path() }))

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ImportSpec)(nil),
	}

	// TODO: use only pass.Pkg.Imports() for performance
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ImportSpec:
			if n.Path.Value == `"fmt"` {
				pass.Reportf(n.Pos(), "import is \"fmt\"")
			}
		}
	})

	return nil, nil
}
