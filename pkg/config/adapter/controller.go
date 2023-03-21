package adapter

import (
	"fmt"

	"github.com/samber/lo"

	"github.com/syuparn/pkgaffinity/interfaces"
	"github.com/syuparn/pkgaffinity/pkg/config/domain"
	"github.com/syuparn/pkgaffinity/pkg/config/usecase"
)

type controller struct {
	listByPathInputPort usecase.ListByPathInputPort
}

func NewController(listByPathInputPort usecase.ListByPathInputPort) interfaces.Config {
	return &controller{
		listByPathInputPort: listByPathInputPort,
	}
}

// impl check
var _ interfaces.Config = &controller{}

func (c *controller) ListRulesByPath(req *interfaces.ListRulesByPathRequest) (*interfaces.ListRulesByPathResponse, error) {
	in := &usecase.ListByPathInputData{
		PackagePath: req.PackagePath,
	}

	out, err := c.listByPathInputPort.Exec(in)
	if err != nil {
		return nil, fmt.Errorf("ListRulesByPath failed: %w", err)
	}

	groupRules := lo.Map(out.AntiAffinityGroupRules, func(r *domain.AntiAffinityGroupRule, _ int) *interfaces.AntiAffinityGroupRule {
		return &interfaces.AntiAffinityGroupRule{
			GroupPathPrefix: string(r.Group),
			AllowNames: lo.Map(r.AllowNames, func(n domain.Name, _ int) string {
				return string(n)
			}),
		}
	})

	listRules := lo.Map(out.AntiAffinityListRules, func(r *domain.AntiAffinityListRule, _ int) *interfaces.AntiAffinityListRule {
		return &interfaces.AntiAffinityListRule{
			Label: string(r.Label),
			PathPrefixes: lo.Map(r.Prefixes, func(p domain.PathPrefix, _ int) string {
				return string(p)
			}),
		}
	})

	return &interfaces.ListRulesByPathResponse{
		AntiAffinityGroupRules: groupRules,
		AntiAffinityListRules:  listRules,
	}, nil
}
