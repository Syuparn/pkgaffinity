package domain

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewAntiAffinityGroupRule(t *testing.T) {
	tests := []struct {
		name     string
		self     Path
		group    PathPrefix
		expected *AntiAffinityGroupRule
	}{
		{
			name:  "self is directly under selfPath",
			self:  "foo/bar/baz",
			group: "foo/bar",
			expected: &AntiAffinityGroupRule{
				selfPath:        "foo/bar/baz",
				groupPathPrefix: "foo/bar",
				allowNames:      []Name{"baz"},
			},
		},
		{
			name:  "self is not directly under selfPath",
			self:  "foo/bar/baz/quux",
			group: "foo/bar",
			expected: &AntiAffinityGroupRule{
				selfPath:        "foo/bar/baz/quux",
				groupPathPrefix: "foo/bar",
				allowNames:      []Name{"baz"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewAntiAffinityGroupRule(tt.self, tt.group)
			assert.Equal(t, tt.expected, actual)
			assert.NoError(t, err)
		})
	}
}

func TestNewAntiAffinityGroupRuleError(t *testing.T) {
	tests := []struct {
		name     string
		self     Path
		group    PathPrefix
		expected string
	}{
		{
			name:     "self is not in group",
			self:     "hoge/fuga",
			group:    "foo/bar",
			expected: "self `hoge/fuga` must be in group `foo/bar`",
		},
		{
			name:     "self is same as group",
			self:     "foo/bar",
			group:    "foo/bar",
			expected: "self `foo/bar` must be in group `foo/bar`",
		},
		{
			name:     "self has literally group prefix but refers different path",
			self:     "foo/barbara",
			group:    "foo/bar",
			expected: "self `foo/barbara` must be in group `foo/bar`",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAntiAffinityGroupRule(tt.self, tt.group)
			assert.Error(t, err)
			assert.EqualError(t, err, tt.expected)
		})
	}
}

func TestAntiAffinityGroupRuleCheck(t *testing.T) {
	tests := []struct {
		name string
		rule *AntiAffinityGroupRule
		path Path
	}{
		{
			name: "path is in selfPathPrefix",
			rule: lo.Must(NewAntiAffinityGroupRule(
				Path("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
			)),
			path: "foo/bar/baz/hoge",
		},
		{
			name: "path is not in group",
			rule: lo.Must(NewAntiAffinityGroupRule(
				Path("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
			)),
			path: "fuga/piyo",
		},
		{
			name: "path is group itself",
			rule: lo.Must(NewAntiAffinityGroupRule(
				Path("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
			)),
			path: "foo/bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Check(tt.path)
			assert.NoError(t, err)
		})
	}
}

func TestAntiAffinityGroupRuleCheckError(t *testing.T) {
	tests := []struct {
		name     string
		rule     *AntiAffinityGroupRule
		path     Path
		expected string
	}{
		{
			name: "path is in groupPathPrefix but not in selfPathPrefix",
			rule: lo.Must(NewAntiAffinityGroupRule(
				Path("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
			)),
			path:     "foo/bar/hoge/piyo",
			expected: "import `foo/bar/hoge/piyo` in package `foo/bar/baz/quux` breaks anti-affinity group rule `foo/bar`",
		},
		{
			name: "path is in groupPathPrefix but not in selfPathPrefix (path is in the same hierarchy as selfPathPrefix)",
			rule: lo.Must(NewAntiAffinityGroupRule(
				Path("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
			)),
			path:     "foo/bar/hoge",
			expected: "import `foo/bar/hoge` in package `foo/bar/baz/quux` breaks anti-affinity group rule `foo/bar`",
		},
		{
			name: "path has prefix selfPathPrefix literally but not in selfPathPrefix",
			rule: lo.Must(NewAntiAffinityGroupRule(
				Path("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
			)),
			path:     "foo/bar/baz123",
			expected: "import `foo/bar/baz123` in package `foo/bar/baz/quux` breaks anti-affinity group rule `foo/bar`",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Check(tt.path)
			assert.Error(t, err)
			assert.EqualError(t, err, tt.expected)
		})
	}
}
