package usecase

import (
	"fmt"

	"github.com/syuparn/pkgaffinity/pkg/config/domain"
)

type ListByPathInputData struct {
	PackagePath string
}

type ListByPathOutputData struct {
	AntiAffinityGroupRules []*domain.AntiAffinityGroupRule
	AntiAffinityListRules  []*domain.AntiAffinityListRule
}

// HACK: return OutputData directly so that controller cat return the outputdata to callee.
type ListByPathInputPort interface {
	Exec(*ListByPathInputData) (*ListByPathOutputData, error)
}

type listByPathInteractor struct {
	antiAffinityGroupRuleRepository domain.AntiAffinityGroupRuleRepository
	antiAffinityListRuleRepository  domain.AntiAffinityListRuleRepository
}

// check impl
var _ ListByPathInputPort = &listByPathInteractor{}

func NewListByPathInputPort(
	antiAffinityGroupRuleRepository domain.AntiAffinityGroupRuleRepository,
	antiAffinityListRuleRepository domain.AntiAffinityListRuleRepository,
) ListByPathInputPort {
	return &listByPathInteractor{
		antiAffinityGroupRuleRepository: antiAffinityGroupRuleRepository,
		antiAffinityListRuleRepository:  antiAffinityListRuleRepository,
	}
}

func (it *listByPathInteractor) Exec(in *ListByPathInputData) (*ListByPathOutputData, error) {
	path := domain.Path(in.PackagePath)
	groupRules, err := it.antiAffinityGroupRuleRepository.ListByPath(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get anti-affinity group rule: %w", err)
	}
	listRules, err := it.antiAffinityListRuleRepository.ListByPath(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get anti-affinity list rule: %w", err)
	}

	return &ListByPathOutputData{
		AntiAffinityGroupRules: groupRules,
		AntiAffinityListRules:  listRules,
	}, nil
}
