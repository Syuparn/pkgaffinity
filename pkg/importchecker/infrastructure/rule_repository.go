package infrastructure

import (
	"fmt"

	"github.com/syuparn/pkgaffinity/interfaces"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/domain"
)

type antiAffinityRuleRepository struct {
	configController interfaces.Config
}

func NewAntiAffinityRuleRepository(configController interfaces.Config) domain.AntiAffinityRuleRepository {
	return &antiAffinityRuleRepository{
		configController: configController,
	}
}

// impl check
var _ domain.AntiAffinityRuleRepository = &antiAffinityRuleRepository{}

func (r *antiAffinityRuleRepository) ListByPath(packagePath domain.Path) ([]domain.AntiAffinityRule, error) {
	res, err := r.configController.ListRulesByPath(&interfaces.ListRulesByPathRequest{
		PackagePath: string(packagePath),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get rules from config: %w", err)
	}

	groupRules := make([]domain.AntiAffinityRule, len(res.AntiAffinityGroupRules))
	for i, gr := range res.AntiAffinityGroupRules {
		rule, err := domain.NewAntiAffinityGroupRule(packagePath, domain.PathPrefix(gr.GroupPathPrefix))
		if err != nil {
			return nil, fmt.Errorf("failed to parse group anti affinity rule %+v (package %s): %w", gr, packagePath, err)
		}
		groupRules[i] = rule
	}

	return groupRules, nil
}
