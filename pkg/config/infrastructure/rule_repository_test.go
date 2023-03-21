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

func TestAntiAffinityListRuleRepositoryListByPath(t *testing.T) {
	tests := []struct {
		name     string
		repo     *antiAffinityListRuleRepository
		path     domain.Path
		expected []*domain.AntiAffinityListRule
	}{
		{
			name: "get only rules that contain path",
			repo: &antiAffinityListRuleRepository{
				rules: []*domain.AntiAffinityListRule{
					{
						Label: "rule1",
						Prefixes: []domain.PathPrefix{
							"foo/bar",
							"foo/hoge",
						},
					},
					{
						Label: "rule2",
						Prefixes: []domain.PathPrefix{
							"hoge/piyo",
						},
					},
					{
						Label: "rule3",
						Prefixes: []domain.PathPrefix{
							"foo/bar/baz",
						},
					},
				},
			},
			path: "foo/bar/baz/hoge",
			expected: []*domain.AntiAffinityListRule{
				{
					Label: "rule1",
					Prefixes: []domain.PathPrefix{
						"foo/bar",
						"foo/hoge",
					},
				},
				{
					Label: "rule3",
					Prefixes: []domain.PathPrefix{
						"foo/bar/baz",
					},
				},
			},
		},
		{
			name: "no rules found",
			repo: &antiAffinityListRuleRepository{
				rules: []*domain.AntiAffinityListRule{
					{
						Label: "rule2",
						Prefixes: []domain.PathPrefix{
							"hoge/piyo",
						},
					},
				},
			},
			path:     "foo/bar/baz/hoge",
			expected: []*domain.AntiAffinityListRule{},
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
