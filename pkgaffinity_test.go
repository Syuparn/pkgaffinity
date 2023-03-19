package pkgaffinity_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/syuparn/pkgaffinity"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, pkgaffinity.Analyzer, "a")
}
