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
		}
	})

	return &interfaces.ListRulesByPathResponse{
		AntiAffinityGroupRules: groupRules,
	}, nil
}
