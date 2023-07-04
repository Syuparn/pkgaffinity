package domain

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewAntiAffinityGroupRule(t *testing.T) {
	tests := []struct {
		name       string
		self       Path
		group      PathPrefix
		allowNames []Name
		expected   *AntiAffinityGroupRule
	}{
		{
			name:       "self is directly under selfPath",
			self:       "foo/bar/baz",
			group:      "foo/bar",
			allowNames: []Name{},
			expected: &AntiAffinityGroupRule{
				selfPath:        "foo/bar/baz",
				groupPathPrefix: "foo/bar",
				allowNames:      []Name{"baz"},
			},
		},
		{
			name:       "self is not directly under selfPath",
			self:       "foo/bar/baz/quux",
			group:      "foo/bar",
			allowNames: []Name{},
			expected: &AntiAffinityGroupRule{
				selfPath:        "foo/bar/baz/quux",
				groupPathPrefix: "foo/bar",
				allowNames:      []Name{"baz"},
			},
		},
		{
			name:       "with allowNames",
			self:       "foo/bar/baz/quux",
			group:      "foo/bar",
			allowNames: []Name{"hoge", "fuga"},
			expected: &AntiAffinityGroupRule{
				selfPath:        "foo/bar/baz/quux",
				groupPathPrefix: "foo/bar",
				allowNames:      []Name{"hoge", "fuga", "baz"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewAntiAffinityGroupRule(tt.self, tt.group, tt.allowNames)
			assert.Equal(t, tt.expected, actual)
			assert.NoError(t, err)
		})
	}
}

func TestNewAntiAffinityGroupRuleError(t *testing.T) {
	tests := []struct {
		name       string
		self       Path
		group      PathPrefix
		allowNames []Name
		expected   string
	}{
		{
			name:       "self is not in group",
			self:       "hoge/fuga",
			group:      "foo/bar",
			allowNames: []Name{},
			expected:   "self `hoge/fuga` must be in group `foo/bar`",
		},
		{
			name:       "self is same as group",
			self:       "foo/bar",
			group:      "foo/bar",
			allowNames: []Name{},
			expected:   "self `foo/bar` must be in group `foo/bar`",
		},
		{
			name:       "self has literally group prefix but refers different path",
			self:       "foo/barbara",
			group:      "foo/bar",
			allowNames: []Name{},
			expected:   "self `foo/barbara` must be in group `foo/bar`",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAntiAffinityGroupRule(tt.self, tt.group, tt.allowNames)
			assert.Error(t, err)
			assert.EqualError(t, err, tt.expected)
		})
	}
}

func TestAntiAffinityGroupRuleCheckOK(t *testing.T) {
	tests := []struct {
		name string
		rule *AntiAffinityGroupRule
		path Path
	}{
		{
			name: "path is in selfPathPrefix",
			rule: lo.Must(NewAntiAffinityGroupRule(
				NewPath("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
				[]Name{},
			)),
			path: NewPath("foo/bar/baz/hoge"),
		},
		{
			name: "path is not in group",
			rule: lo.Must(NewAntiAffinityGroupRule(
				NewPath("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
				[]Name{},
			)),
			path: NewPath("fuga/piyo"),
		},
		{
			name: "path is group itself",
			rule: lo.Must(NewAntiAffinityGroupRule(
				NewPath("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
				[]Name{},
			)),
			path: NewPath("foo/bar"),
		},
		{
			name: "path is not in selfPathPrefix but in allowNames",
			rule: lo.Must(NewAntiAffinityGroupRule(
				NewPath("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
				[]Name{"hoge"},
			)),
			path: NewPath("foo/bar/hoge/piyo"),
		},
		{
			name: "package suffix .test is ignored",
			rule: lo.Must(NewAntiAffinityGroupRule(
				NewPath("foo/bar/baz"),
				PathPrefix("foo/bar"),
				[]Name{""},
			)),
			path: NewPath("foo/bar/baz.test"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violation := tt.rule.Check(tt.path)
			assert.Nil(t, violation)
		})
	}
}

func TestAntiAffinityGroupRuleCheckNG(t *testing.T) {
	tests := []struct {
		name     string
		rule     *AntiAffinityGroupRule
		path     Path
		expected *Violation
	}{
		{
			name: "path is in groupPathPrefix but not in selfPathPrefix",
			rule: lo.Must(NewAntiAffinityGroupRule(
				NewPath("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
				[]Name{},
			)),
			path: NewPath("foo/bar/hoge/piyo"),
			expected: &Violation{
				ImportPath:  "foo/bar/hoge/piyo",
				PackagePath: "foo/bar/baz/quux",
				RuleLabel:   "anti-affinity group rule `foo/bar`",
			},
		},
		{
			name: "path is in groupPathPrefix but not in selfPathPrefix (path is in the same hierarchy as selfPathPrefix)",
			rule: lo.Must(NewAntiAffinityGroupRule(
				NewPath("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
				[]Name{},
			)),
			path: NewPath("foo/bar/hoge"),
			expected: &Violation{
				ImportPath:  "foo/bar/hoge",
				PackagePath: "foo/bar/baz/quux",
				RuleLabel:   "anti-affinity group rule `foo/bar`",
			},
		},
		{
			name: "path has prefix selfPathPrefix literally but not in selfPathPrefix",
			rule: lo.Must(NewAntiAffinityGroupRule(
				NewPath("foo/bar/baz/quux"),
				PathPrefix("foo/bar"),
				[]Name{},
			)),
			path: NewPath("foo/bar/baz123"),
			expected: &Violation{
				ImportPath:  "foo/bar/baz123",
				PackagePath: "foo/bar/baz/quux",
				RuleLabel:   "anti-affinity group rule `foo/bar`",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violation := tt.rule.Check(tt.path)
			assert.NotNil(t, violation)
			assert.Equal(t, violation, tt.expected)
		})
	}
}

func TestNewAntiAffinityListRule(t *testing.T) {
	tests := []struct {
		name     string
		label    RuleLabel
		self     Path
		prefixes []PathPrefix
		expected *AntiAffinityListRule
	}{
		{
			name:  "self is under a prefix",
			label: "rule1",
			self:  NewPath("foo/bar/baz/hoge"),
			prefixes: []PathPrefix{
				"foo/bar",
				"fuga/piyo",
			},
			expected: &AntiAffinityListRule{
				selfPath: "foo/bar/baz/hoge",
				pathPrefixes: []PathPrefix{
					"fuga/piyo",
					// "foo/bar" is removed because this is where selfPath belongs to
				},
				label: "rule1",
			},
		},
		{
			name:  "self is same as a prefix",
			label: "rule1",
			self:  NewPath("foo/bar"),
			prefixes: []PathPrefix{
				"foo/bar",
				"fuga/piyo",
			},
			expected: &AntiAffinityListRule{
				selfPath: "foo/bar",
				pathPrefixes: []PathPrefix{
					"fuga/piyo",
					// "foo/bar" is removed because this is where selfPath belongs to
				},
				label: "rule1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewAntiAffinityListRule(tt.self, tt.prefixes, tt.label)
			assert.Equal(t, tt.expected, actual)
			assert.NoError(t, err)
		})
	}
}

func TestAntiAffinityListRuleCheckOK(t *testing.T) {
	tests := []struct {
		name string
		rule *AntiAffinityListRule
		path Path
	}{
		{
			name: "path does not match to any prefixes",
			rule: lo.Must(NewAntiAffinityListRule(
				NewPath("foo/bar/baz/quux"),
				[]PathPrefix{
					"foo/bar",
					"fuga/piyo",
				},
				"rule1",
			)),
			path: NewPath("quux"),
		},
		{
			name: "path matches a prefix but the prefix contains selfPath",
			rule: lo.Must(NewAntiAffinityListRule(
				NewPath("foo/bar/baz/quux"),
				[]PathPrefix{
					"foo/bar",
					"fuga/piyo",
				},
				"rule1",
			)),
			path: NewPath("foo/bar/quux"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violation := tt.rule.Check(tt.path)
			assert.Nil(t, violation)
		})
	}
}

func TestAntiAffinityListRuleCheckNG(t *testing.T) {
	tests := []struct {
		name     string
		rule     *AntiAffinityListRule
		path     Path
		expected *Violation
	}{
		{
			name: "path matches to a prefix",
			rule: lo.Must(NewAntiAffinityListRule(
				NewPath("foo/bar/baz/quux"),
				[]PathPrefix{
					"foo/bar",
					"fuga/piyo",
				},
				"rule1",
			)),
			path: NewPath("fuga/piyo/hoge"),
			expected: &Violation{
				ImportPath:  "fuga/piyo/hoge",
				PackagePath: "foo/bar/baz/quux",
				RuleLabel:   "rule1",
			},
		},
		{
			name: "path matches to a prefix (path is same as the prefix)",
			rule: lo.Must(NewAntiAffinityListRule(
				NewPath("foo/bar/baz/quux"),
				[]PathPrefix{
					"foo/bar",
					"fuga/piyo",
				},
				"rule1",
			)),
			path: NewPath("fuga/piyo"),
			expected: &Violation{
				ImportPath:  "fuga/piyo",
				PackagePath: "foo/bar/baz/quux",
				RuleLabel:   "rule1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			violation := tt.rule.Check(tt.path)
			assert.NotNil(t, violation)
			assert.Equal(t, violation, tt.expected)
		})
	}
}
