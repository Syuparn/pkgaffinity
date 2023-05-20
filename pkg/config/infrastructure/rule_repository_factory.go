package infrastructure

import (
	"fmt"
	"os"
	"path"

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
			AllowNames: lo.Map(g.AllowNames, func(n string, _ int) domain.Name {
				return domain.Name(n)
			}),
			IgnorePaths: lo.Map(g.IgnorePaths, func(p string, _ int) domain.Path {
				return domain.Path(path.Join(g.PathPrefix, p))
			}),
		}
	})

	return &antiAffinityGroupRuleRepository{rules: rules}, nil
}

// AntiAffinityListRuleRepositoryFactory creates AntiAffinityListRuleRepository from configs.
// NOTE: this is factory for repository (not factory for rule itself!)
type AntiAffinityListRuleRepositoryFactory interface {
	Create() (domain.AntiAffinityListRuleRepository, error)
}

type antiAffinityListRuleRepositoryFactory struct {
	configFilePath string
}

func NewAntiAffinityListRuleRepositoryFactory(configFilePath string) AntiAffinityListRuleRepositoryFactory {
	return &antiAffinityListRuleRepositoryFactory{
		configFilePath: configFilePath,
	}
}

// check impl
var _ AntiAffinityListRuleRepositoryFactory = &antiAffinityListRuleRepositoryFactory{}

func (f *antiAffinityListRuleRepositoryFactory) Create() (domain.AntiAffinityListRuleRepository, error) {
	b, err := os.ReadFile(f.configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", f.configFilePath, err)
	}

	var cfg schema.ConfigSchema
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", f.configFilePath, err)
	}

	rules := make([]*domain.AntiAffinityListRule, len(cfg.AntiAffinityRules.Lists))
	for i, l := range cfg.AntiAffinityRules.Lists {
		if l.Label == "" {
			return nil, fmt.Errorf("anti-affinity list rule must have a label: %s", f.configFilePath)
		}

		rules[i] = &domain.AntiAffinityListRule{
			Label: domain.RuleLabel(l.Label),
			Prefixes: lo.Map(l.PathPrefixes, func(p string, _ int) domain.PathPrefix {
				return domain.PathPrefix(p)
			}),
		}
	}

	return &antiAffinityListRuleRepository{rules: rules}, nil
}
