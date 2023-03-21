package infrastructure

import (
	"fmt"
	"os"

	"github.com/samber/lo"
	"github.com/syuparn/pkgaffinity/pkg/config/domain"
	"github.com/syuparn/pkgaffinity/pkg/config/infrastructure/schema"
	"gopkg.in/yaml.v3"
)

// AntiAffinityGroupRuleRepositoryFactory creates AntiAffinityGroupRuleRepository from configs.
// NOTE: this is factory for repository (not factory for rule itself!)
type AntiAffinityGroupRuleRepositoryFactory interface {
	Create() (domain.AntiAffinityGroupRuleRepository, error)
}

type antiAffinityGroupRuleRepositoryFactory struct {
	configFilePath string
}

func NewAntiAffinityGroupRuleRepositoryFactory(configFilePath string) AntiAffinityGroupRuleRepositoryFactory {
	return &antiAffinityGroupRuleRepositoryFactory{
		configFilePath: configFilePath,
	}
}

// check impl
var _ AntiAffinityGroupRuleRepositoryFactory = &antiAffinityGroupRuleRepositoryFactory{}

func (f *antiAffinityGroupRuleRepositoryFactory) Create() (domain.AntiAffinityGroupRuleRepository, error) {
	b, err := os.ReadFile(f.configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", f.configFilePath, err)
	}

	var cfg schema.ConfigSchema
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", f.configFilePath, err)
	}

	rules := lo.Map(cfg.AntiAffinityRules.Groups, func(g *schema.AntiAffinityGroupRule, _ int) *domain.AntiAffinityGroupRule {
		return &domain.AntiAffinityGroupRule{
			Group: domain.PathPrefix(g.PathPrefix),
		}
	})

	return &antiAffinityGroupRuleRepository{rules: rules}, nil
}
