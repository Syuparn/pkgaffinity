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
					{Group: "foo/bar"},
				},
			},
		},
		{
			name:     "two anti-affinity group rules",
			fileName: "two_groups",
			expected: &antiAffinityGroupRuleRepository{
				rules: []*domain.AntiAffinityGroupRule{
					{Group: "foo/bar"},
					{Group: "baz"},
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
