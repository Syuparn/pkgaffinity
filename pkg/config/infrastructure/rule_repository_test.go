package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syuparn/pkgaffinity/pkg/config/domain"
)

func TestAntiAffinityGroupRuleRepositoryListByPath(t *testing.T) {
	tests := []struct {
		name     string
		repo     *antiAffinityGroupRuleRepository
		path     domain.Path
		expected []*domain.AntiAffinityGroupRule
	}{
		{
			name: "get rules only whose group contains path",
			repo: &antiAffinityGroupRuleRepository{
				rules: []*domain.AntiAffinityGroupRule{
					{Group: "foo/bar"},
					{Group: "hoge/piyo"},
					{Group: "foo/bar/baz"},
				},
			},
			path: "foo/bar/baz/hoge",
			expected: []*domain.AntiAffinityGroupRule{
				{Group: "foo/bar"},
				{Group: "foo/bar/baz"},
			},
		},
		{
			name: "no rules found",
			repo: &antiAffinityGroupRuleRepository{
				rules: []*domain.AntiAffinityGroupRule{
					{Group: "hoge/piyo"},
					{Group: "other"},
				},
			},
			path:     "foo/bar/baz/hoge",
			expected: []*domain.AntiAffinityGroupRule{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.repo.ListByPath(tt.path)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
