package infrastructure

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	"github.com/syuparn/pkgaffinity/interfaces"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/domain"
)

func TestAntiAffinityRuleRepositoryListByPath(t *testing.T) {
	tests := []struct {
		name         string
		packagePath  domain.Path
		mockResponse *interfaces.ListRulesByPathResponse
		expected     []domain.AntiAffinityRule
	}{
		{
			name:        "get one group rule",
			packagePath: "foo/bar/baz/hoge",
			mockResponse: &interfaces.ListRulesByPathResponse{
				AntiAffinityGroupRules: []*interfaces.AntiAffinityGroupRule{
					{GroupPathPrefix: "foo/bar"},
				},
			},
			expected: []domain.AntiAffinityRule{
				lo.Must(domain.NewAntiAffinityGroupRule("foo/bar/baz/hoge", "foo/bar", []domain.Name{})),
			},
		},
		{
			name:        "get two group rules",
			packagePath: "foo/bar/baz/hoge",
			mockResponse: &interfaces.ListRulesByPathResponse{
				AntiAffinityGroupRules: []*interfaces.AntiAffinityGroupRule{
					{GroupPathPrefix: "foo/bar"},
					{GroupPathPrefix: "foo/bar/baz"},
				},
			},
			expected: []domain.AntiAffinityRule{
				lo.Must(domain.NewAntiAffinityGroupRule("foo/bar/baz/hoge", "foo/bar", []domain.Name{})),
				lo.Must(domain.NewAntiAffinityGroupRule("foo/bar/baz/hoge", "foo/bar/baz", []domain.Name{})),
			},
		},
		{
			name:        "get one group rule with allowNames",
			packagePath: "foo/bar/baz/hoge",
			mockResponse: &interfaces.ListRulesByPathResponse{
				AntiAffinityGroupRules: []*interfaces.AntiAffinityGroupRule{
					{GroupPathPrefix: "foo/bar", AllowNames: []string{"fuga", "piyo"}},
				},
			},
			expected: []domain.AntiAffinityRule{
				lo.Must(domain.NewAntiAffinityGroupRule("foo/bar/baz/hoge", "foo/bar", []domain.Name{"fuga", "piyo"})),
			},
		},
		{
			name:        "get one group rule with ignorePaths",
			packagePath: "foo/bar/baz/hoge",
			mockResponse: &interfaces.ListRulesByPathResponse{
				AntiAffinityGroupRules: []*interfaces.AntiAffinityGroupRule{
					{GroupPathPrefix: "foo/bar", IgnorePaths: []string{"foo/bar/baz/ignored", "foo/bar/baz/other"}},
				},
			},
			expected: []domain.AntiAffinityRule{
				lo.Must(domain.NewAntiAffinityGroupRule("foo/bar/baz/hoge", "foo/bar", []domain.Name{})),
			},
		},
		{
			name:        "get one group rule but matched to ignorePaths",
			packagePath: "foo/bar/baz/hoge",
			mockResponse: &interfaces.ListRulesByPathResponse{
				AntiAffinityGroupRules: []*interfaces.AntiAffinityGroupRule{
					{GroupPathPrefix: "foo/bar", IgnorePaths: []string{"foo/bar/baz/other", "foo/bar/baz/hoge"}},
				},
			},
			expected: []domain.AntiAffinityRule{},
		},
		{
			name:        "get two list rules",
			packagePath: "foo/bar/baz/hoge",
			mockResponse: &interfaces.ListRulesByPathResponse{
				AntiAffinityListRules: []*interfaces.AntiAffinityListRule{
					{Label: "rule1", PathPrefixes: []string{"foo/bar", "fuga/piyo"}},
					{Label: "rule2", PathPrefixes: []string{"foo/bar/baz", "quux"}},
				},
			},
			expected: []domain.AntiAffinityRule{
				lo.Must(domain.NewAntiAffinityListRule("foo/bar/baz/hoge", []domain.PathPrefix{"foo/bar", "fuga/piyo"}, "rule1")),
				lo.Must(domain.NewAntiAffinityListRule("foo/bar/baz/hoge", []domain.PathPrefix{"foo/bar/baz", "quux"}, "rule2")),
			},
		},
		{
			name:        "get a group rule and a list rule",
			packagePath: "foo/bar/baz/hoge",
			mockResponse: &interfaces.ListRulesByPathResponse{
				AntiAffinityGroupRules: []*interfaces.AntiAffinityGroupRule{
					{GroupPathPrefix: "foo/bar"},
				},
				AntiAffinityListRules: []*interfaces.AntiAffinityListRule{
					{Label: "rule1", PathPrefixes: []string{"foo/bar", "fuga/piyo"}},
				},
			},
			expected: []domain.AntiAffinityRule{
				lo.Must(domain.NewAntiAffinityGroupRule("foo/bar/baz/hoge", "foo/bar", []domain.Name{})),
				lo.Must(domain.NewAntiAffinityListRule("foo/bar/baz/hoge", []domain.PathPrefix{"foo/bar", "fuga/piyo"}, "rule1")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configMock := &interfaces.ConfigMock{
				ListRulesByPathFunc: func(_ *interfaces.ListRulesByPathRequest) (*interfaces.ListRulesByPathResponse, error) {
					return tt.mockResponse, nil
				},
			}
			repository := NewAntiAffinityRuleRepository(configMock)

			actual, err := repository.ListByPath(tt.packagePath)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
