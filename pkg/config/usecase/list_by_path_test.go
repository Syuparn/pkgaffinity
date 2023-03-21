package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syuparn/pkgaffinity/pkg/config/domain"
)

func TestListByPathInputPortExec(t *testing.T) {
	tests := []struct {
		name           string
		in             *ListByPathInputData
		mockGroupRules []*domain.AntiAffinityGroupRule
		expected       *ListByPathOutputData
	}{
		{
			name: "list group rules",
			in: &ListByPathInputData{
				PackagePath: "foo/bar/baz/hoge",
			},
			mockGroupRules: []*domain.AntiAffinityGroupRule{
				{Group: "foo/bar"},
				{Group: "foo/bar/baz"},
			},
			expected: &ListByPathOutputData{
				AntiAffinityGroupRules: []*domain.AntiAffinityGroupRule{
					{Group: "foo/bar"},
					{Group: "foo/bar/baz"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			antiAffinityGroupRuleRepositoryMock := &domain.AntiAffinityGroupRuleRepositoryMock{
				ListByPathFunc: func(_ domain.Path) ([]*domain.AntiAffinityGroupRule, error) {
					return tt.mockGroupRules, nil
				},
			}
			interactor := NewListByPathInputPort(antiAffinityGroupRuleRepositoryMock)

			actual, err := interactor.Exec(tt.in)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
