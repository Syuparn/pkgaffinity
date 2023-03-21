package di

import (
	"github.com/samber/do"

	"github.com/syuparn/pkgaffinity/interfaces"
	"github.com/syuparn/pkgaffinity/pkg/config/adapter"
	"github.com/syuparn/pkgaffinity/pkg/config/domain"
	"github.com/syuparn/pkgaffinity/pkg/config/infrastructure"
	"github.com/syuparn/pkgaffinity/pkg/config/usecase"
)

func NewInjector() *do.Injector {
	injector := do.New()

	// settings
	do.Provide(injector, func(i *do.Injector) (string, error) {
		// default config file path
		return ".pkgaffinity.yaml", nil
	})
	// domain
	do.Provide(injector, func(i *do.Injector) (infrastructure.AntiAffinityGroupRuleRepositoryFactory, error) {
		configFilePath := do.MustInvoke[string](i)
		return infrastructure.NewAntiAffinityGroupRuleRepositoryFactory(configFilePath), nil
	})
	do.Provide(injector, func(i *do.Injector) (domain.AntiAffinityGroupRuleRepository, error) {
		factory := do.MustInvoke[infrastructure.AntiAffinityGroupRuleRepositoryFactory](i)
		return factory.Create()
	})
	// usecase
	do.Provide(injector, func(i *do.Injector) (usecase.ListByPathInputPort, error) {
		groupRuleRepository := do.MustInvoke[domain.AntiAffinityGroupRuleRepository](i)
		return usecase.NewListByPathInputPort(groupRuleRepository), nil
	})
	// adapter
	do.Provide(injector, func(i *do.Injector) (interfaces.Config, error) {
		in := do.MustInvoke[usecase.ListByPathInputPort](i)
		return adapter.NewController(in), nil
	})

	return injector
}
