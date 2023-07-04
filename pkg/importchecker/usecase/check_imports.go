package usecase

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/domain"
)

type CheckImportsInputData struct {
	PackagePath string
	ImportPaths []string
}

type CheckImportsOutputData struct {
	Violations []*domain.Violation
}

type CheckImportsInputPort interface {
	Exec(*CheckImportsInputData) error
}

//go:generate go run github.com/matryer/moq -fmt goimports -out zz_generated_moq_check_imports.go . CheckImportsOutputPort
type CheckImportsOutputPort interface {
	Present(*CheckImportsOutputData) error
}

type checkImportsInteractor struct {
	out                        CheckImportsOutputPort
	antiAffinityRuleRepository domain.AntiAffinityRuleRepository
}

// check impl
var _ CheckImportsInputPort = &checkImportsInteractor{}

func NewCheckImportsInputPort(
	presenter CheckImportsOutputPort,
	antiAffinityRuleRepository domain.AntiAffinityRuleRepository,
) CheckImportsInputPort {
	return &checkImportsInteractor{
		out:                        presenter,
		antiAffinityRuleRepository: antiAffinityRuleRepository,
	}
}

func (it *checkImportsInteractor) Exec(in *CheckImportsInputData) error {
	packagePath := domain.NewPath(in.PackagePath)
	rules, err := it.antiAffinityRuleRepository.ListByPath(packagePath)
	if err != nil {
		return fmt.Errorf("failed to get anti affinity rules of package `%s`: %w", packagePath, err)
	}

	violations := lo.FlatMap(in.ImportPaths, func(importPath string, _ int) []*domain.Violation {
		return lo.Compact(lo.Map(rules, func(rule domain.AntiAffinityRule, _ int) *domain.Violation {
			return rule.Check(domain.NewPath(importPath))
		}))
	})

	return it.out.Present(&CheckImportsOutputData{Violations: violations})
}
