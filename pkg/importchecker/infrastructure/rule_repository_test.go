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
		// TODO: list rules
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
