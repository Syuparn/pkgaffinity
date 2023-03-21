package pkgaffinity

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/gostaticanalysis/testutil"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	tests := []struct {
		name        string
		packagePath string
		expectedOut []string
	}{
		{
			name:        "package breaks anti-affinity rule",
			packagePath: "a/foo/baz",
			expectedOut: []string{
				"package a/foo/baz: import \"a/foo/bar\" breaks anti-affinity group rule `a/foo`",
				"", // break line
			},
		},
		{
			name:        "package meets anti-affinity rule",
			packagePath: "a/foo/bar",
			expectedOut: []string{
				"", // break line
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// HACK: collect output explicitly because the analyzer does not use reporter
			var out bytes.Buffer
			teardownW := patchWriter(&out)
			defer teardownW()

			teardownC := patchConfigPath("testdata/.pkgaffinity.yaml")
			defer teardownC()

			analyzer := NewAnalyzer()
			testdata := testutil.WithModules(t, analysistest.TestData(), nil)
			analysistest.Run(t, testdata, analyzer, tt.packagePath)

			assert.Equal(t, strings.Join(tt.expectedOut, "\n"), out.String())
		})
	}
}

func patchWriter(out io.Writer) func() {
	original := do.MustInvoke[io.Writer](importcheckerInjector)

	do.Override(importcheckerInjector, func(i *do.Injector) (io.Writer, error) {
		return out, nil
	})

	teardown := func() {
		do.Override(importcheckerInjector, func(i *do.Injector) (io.Writer, error) {
			return original, nil
		})
	}
	return teardown
}

func patchConfigPath(configFilePath string) func() {
	original := do.MustInvoke[string](configInjector)

	do.Override(configInjector, func(i *do.Injector) (string, error) {
		return configFilePath, nil
	})

	teardown := func() {
		do.Override(configInjector, func(i *do.Injector) (string, error) {
			return original, nil
		})
	}
	return teardown
}
