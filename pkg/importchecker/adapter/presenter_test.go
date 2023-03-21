package adapter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/domain"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/usecase"
)

func TestCheckImportsPresenter(t *testing.T) {
	tests := []struct {
		name     string
		out      *usecase.CheckImportsOutputData
		expected []string
	}{
		{
			name: "two violations",
			out: &usecase.CheckImportsOutputData{
				Violations: []*domain.Violation{
					{
						ImportPath:  "foo/fuga",
						PackagePath: "foo/bar/baz/hoge",
						RuleName:    "anti-affinity group rule `foo`",
					},
					{
						ImportPath:  "foo/bar/piyo",
						PackagePath: "foo/bar/baz/hoge",
						RuleName:    "anti-affinity group rule `foo/bar`",
					},
				},
			},
			expected: []string{
				"package foo/bar/baz/hoge: import \"foo/fuga\" breaks anti-affinity group rule `foo`",
				"package foo/bar/baz/hoge: import \"foo/bar/piyo\" breaks anti-affinity group rule `foo/bar`",
				"", // break line
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w bytes.Buffer
			presenter := NewCheckImportsOutputPort(&w)
			presenter.Present(tt.out)

			assert.Equal(t, strings.Join(tt.expected, "\n"), w.String())
		})
	}
}
