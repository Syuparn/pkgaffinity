package pkgaffinity

import (
	"fmt"
	"go/ast"
	"go/types"

	"github.com/samber/lo"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "pkgaffinity is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "pkgaffinity",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {

	// TODO: use for linter
	fmt.Println(pass.Pkg.Path())
	fmt.Println(lo.Map(pass.Pkg.Imports(), func(p *types.Package, _ int) string { return p.Path() }))

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ImportSpec)(nil),
	}

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