package usecase

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	"github.com/syuparn/pkgaffinity/pkg/importchecker/domain"
)

func TestCheckImportsInputPortExec(t *testing.T) {
	tests := []struct {
		name      string
		in        *CheckImportsInputData
		mockRules []domain.AntiAffinityRule
		expected  *CheckImportsOutputData
	}{
		{
			name: "check importpaths by one rule",
			in: &CheckImportsInputData{
				PackagePath: "foo/bar/baz/hoge",
				ImportPaths: []string{
					"foo/fuga",
					"foo/bar/piyo",
				},
			},
			mockRules: []domain.AntiAffinityRule{
				lo.Must(domain.NewAntiAffinityGroupRule("foo/bar/baz/hoge", "foo/bar", []domain.Name{})),
			},
			expected: &CheckImportsOutputData{
				Violations: []*domain.Violation{
					{
						ImportPath:  "foo/bar/piyo",
						PackagePath: "foo/bar/baz/hoge",
						RuleLabel:   "anti-affinity group rule `foo/bar`",
					},
				},
			},
		},
		{
			name: "check importpaths by two rules",
			in: &CheckImportsInputData{
				PackagePath: "foo/bar/baz/hoge",
				ImportPaths: []string{
					"foo/fuga",
					"foo/bar/piyo",
				},
			},
			mockRules: []domain.AntiAffinityRule{
				lo.Must(domain.NewAntiAffinityGroupRule("foo/bar/baz/hoge", "foo/bar", []domain.Name{})),
				lo.Must(domain.NewAntiAffinityGroupRule("foo/bar/baz/hoge", "foo", []domain.Name{})),
			},
			expected: &CheckImportsOutputData{
				Violations: []*domain.Violation{
					{
						ImportPath:  "foo/fuga",
						PackagePath: "foo/bar/baz/hoge",
						RuleLabel:   "anti-affinity group rule `foo`",
					},
					{
						ImportPath:  "foo/bar/piyo",
						PackagePath: "foo/bar/baz/hoge",
						RuleLabel:   "anti-affinity group rule `foo/bar`",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual *CheckImportsOutputData

			presenterMock := &CheckImportsOutputPortMock{
				PresentFunc: func(out *CheckImportsOutputData) error {
					actual = out // capture
					return nil
				},
			}
			antiAffinityRuleRepositoryMock := &domain.AntiAffinityRuleRepositoryMock{
				ListByPathFunc: func(_ domain.Path) ([]domain.AntiAffinityRule, error) {
					return tt.mockRules, nil
				},
			}
			interactor := NewCheckImportsInputPort(presenterMock, antiAffinityRuleRepositoryMock)

			err := interactor.Exec(tt.in)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
