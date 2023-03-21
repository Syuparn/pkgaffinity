package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAntiAffinityGroupRuleContains(t *testing.T) {
	tests := []struct {
		name     string
		rule     *AntiAffinityGroupRule
		path     Path
		expected bool
	}{
		{
			name: "group contains path (directly under group)",
			rule: &AntiAffinityGroupRule{
				Group: "foo/bar",
			},
			path:     "foo/bar/baz",
			expected: true,
		},
		{
			name: "group contains path",
			rule: &AntiAffinityGroupRule{
				Group: "foo/bar",
			},
			path:     "foo/bar/baz/hoge",
			expected: true,
		},
		{
			name: "path is different location",
			rule: &AntiAffinityGroupRule{
				Group: "foo/bar",
			},
			path:     "foo/baz",
			expected: false,
		},
		{
			name: "group is same as path",
			rule: &AntiAffinityGroupRule{
				Group: "foo/bar",
			},
			path:     "foo/bar",
			expected: false,
		},
		{
			name: "path has prefix group literally but not in group",
			rule: &AntiAffinityGroupRule{
				Group: "foo/bar",
			},
			path:     "foo/barbara",
			expected: false,
		},
		{
			name: "group checks only prefix match",
			rule: &AntiAffinityGroupRule{
				Group: "foo/bar",
			},
			path:     "baz/foo/bar",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.rule.Contains(tt.path)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
