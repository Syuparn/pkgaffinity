package infrastructure

import (
	"github.com/samber/lo"

	"github.com/syuparn/pkgaffinity/pkg/config/domain"
)

type antiAffinityGroupRuleRepository struct {
	rules []*domain.AntiAffinityGroupRule
}

// impl check
var _ domain.AntiAffinityGroupRuleRepository = &antiAffinityGroupRuleRepository{}

func (r *antiAffinityGroupRuleRepository) ListByPath(path domain.Path) ([]*domain.AntiAffinityGroupRule, error) {
	rules := lo.Filter(r.rules, func(rule *domain.AntiAffinityGroupRule, _ int) bool {
		return rule.Contains(path)
	})

	return rules, nil
}
