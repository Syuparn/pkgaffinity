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
		mockListRules  []*domain.AntiAffinityListRule
		expected       *ListByPathOutputData
	}{
		{
			name: "list rules",
			in: &ListByPathInputData{
				PackagePath: "foo/bar/baz/hoge",
			},
			mockGroupRules: []*domain.AntiAffinityGroupRule{
				{Group: "foo/bar"},
				{Group: "foo/bar/baz"},
			},
			mockListRules: []*domain.AntiAffinityListRule{
				{Label: "listrule1", Prefixes: []domain.PathPrefix{"hoge/fuga", "piyo"}},
			},
			expected: &ListByPathOutputData{
				AntiAffinityGroupRules: []*domain.AntiAffinityGroupRule{
					{Group: "foo/bar"},
					{Group: "foo/bar/baz"},
				},
				AntiAffinityListRules: []*domain.AntiAffinityListRule{
					{Label: "listrule1", Prefixes: []domain.PathPrefix{"hoge/fuga", "piyo"}},
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
			antiAffinityListRuleRepositoryMock := &domain.AntiAffinityListRuleRepositoryMock{
				ListByPathFunc: func(_ domain.Path) ([]*domain.AntiAffinityListRule, error) {
					return tt.mockListRules, nil
				},
			}
			interactor := NewListByPathInputPort(antiAffinityGroupRuleRepositoryMock, antiAffinityListRuleRepositoryMock)

			actual, err := interactor.Exec(tt.in)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
