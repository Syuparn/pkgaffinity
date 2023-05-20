package infrastructure

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syuparn/pkgaffinity/pkg/config/domain"
)

func TestAntiAffinityGroupRuleRepositoryFactoryCreate(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		expected *antiAffinityGroupRuleRepository
	}{
		{
			name:     "one anti-affinity group rule",
			fileName: "one_group",
			expected: &antiAffinityGroupRuleRepository{
				rules: []*domain.AntiAffinityGroupRule{
					{Group: "foo/bar", AllowNames: []domain.Name{}, IgnorePaths: []domain.Path{}},
				},
			},
		},
		{
			name:     "two anti-affinity group rules",
			fileName: "two_groups",
			expected: &antiAffinityGroupRuleRepository{
				rules: []*domain.AntiAffinityGroupRule{
					{Group: "foo/bar", AllowNames: []domain.Name{}, IgnorePaths: []domain.Path{}},
					{Group: "baz", AllowNames: []domain.Name{}, IgnorePaths: []domain.Path{}},
				},
			},
		},
		{
			name:     "no anti-affinity group rules",
			fileName: "no_groups",
			expected: &antiAffinityGroupRuleRepository{
				rules: []*domain.AntiAffinityGroupRule{},
			},
		},
		{
			name:     "anti-affinity group rule with allow names",
			fileName: "group_allow_names",
			expected: &antiAffinityGroupRuleRepository{
				rules: []*domain.AntiAffinityGroupRule{
					{Group: "foo/bar", AllowNames: []domain.Name{"baz", "quux"}, IgnorePaths: []domain.Path{}},
				},
			},
		},
		{
			name:     "anti-affinity group rule with ignore paths",
			fileName: "group_ignore_paths",
			expected: &antiAffinityGroupRuleRepository{
				rules: []*domain.AntiAffinityGroupRule{
					{Group: "foo/bar", AllowNames: []domain.Name{}, IgnorePaths: []domain.Path{"foo/bar/baz/ignore1", "foo/bar/quux/ignore2"}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := fmt.Sprintf("testdata/%s.yaml", tt.fileName)
			factory := NewAntiAffinityGroupRuleRepositoryFactory(filePath)

			actual, err := factory.Create()

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestAntiAffinityGroupRuleRepositoryFactoryCreateError(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		expected string
	}{
		{
			name:     "file is not found",
			fileName: "404",
			expected: "failed to open config file testdata/404.yaml: open testdata/404.yaml: no such file or directory",
		},
		{
			name:     "file is invalid yaml",
			fileName: "invalid",
			expected: "failed to parse config file testdata/invalid.yaml: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `This is...` into schema.ConfigSchema",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := fmt.Sprintf("testdata/%s.yaml", tt.fileName)
			factory := NewAntiAffinityGroupRuleRepositoryFactory(filePath)

			_, err := factory.Create()

			assert.Error(t, err)
			assert.EqualError(t, err, tt.expected)
		})
	}
}

func TestAntiAffinityListRuleRepositoryFactoryCreate(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		expected *antiAffinityListRuleRepository
	}{
		{
			name:     "one anti-affinity list rule",
			fileName: "one_list",
			expected: &antiAffinityListRuleRepository{
				rules: []*domain.AntiAffinityListRule{
					{Label: "rule1", Prefixes: []domain.PathPrefix{"foo/bar", "baz/quux"}},
				},
			},
		},
		{
			name:     "two anti-affinity list rules",
			fileName: "two_lists",
			expected: &antiAffinityListRuleRepository{
				rules: []*domain.AntiAffinityListRule{
					{Label: "rule1", Prefixes: []domain.PathPrefix{"foo/bar", "baz/quux"}},
					{Label: "rule2", Prefixes: []domain.PathPrefix{"hoge/fuga/piyo", "a/b"}},
				},
			},
		},
		{
			name:     "no anti-affinity group rules",
			fileName: "no_lists",
			expected: &antiAffinityListRuleRepository{
				rules: []*domain.AntiAffinityListRule{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := fmt.Sprintf("testdata/%s.yaml", tt.fileName)
			factory := NewAntiAffinityListRuleRepositoryFactory(filePath)

			actual, err := factory.Create()

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestAntiAffinityListRuleRepositoryFactoryCreateError(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		expected string
	}{
		{
			name:     "file is not found",
			fileName: "404",
			expected: "failed to open config file testdata/404.yaml: open testdata/404.yaml: no such file or directory",
		},
		{
			name:     "file is invalid yaml",
			fileName: "invalid",
			expected: "failed to parse config file testdata/invalid.yaml: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `This is...` into schema.ConfigSchema",
		},
		{
			name:     "label has no names",
			fileName: "list_no_name",
			expected: "anti-affinity list rule must have a label: testdata/list_no_name.yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := fmt.Sprintf("testdata/%s.yaml", tt.fileName)
			factory := NewAntiAffinityListRuleRepositoryFactory(filePath)

			_, err := factory.Create()

			assert.Error(t, err)
			assert.EqualError(t, err, tt.expected)
		})
	}
}
