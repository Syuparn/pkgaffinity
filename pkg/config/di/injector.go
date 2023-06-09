package di

import (
	"os"

	"github.com/samber/do"

	"github.com/syuparn/pkgaffinity/interfaces"
	"github.com/syuparn/pkgaffinity/interfaces/consts"
	"github.com/syuparn/pkgaffinity/pkg/config/adapter"
	"github.com/syuparn/pkgaffinity/pkg/config/domain"
	"github.com/syuparn/pkgaffinity/pkg/config/infrastructure"
	"github.com/syuparn/pkgaffinity/pkg/config/usecase"
)

func NewInjector() *do.Injector {
	injector := do.New()

	// settings
	do.Provide(injector, func(i *do.Injector) (string, error) {
		filePath := os.Getenv(consts.ConfigFilePathEnvKey)
		if filePath == "" {
			// default config file path
			return ".pkgaffinity.yaml", nil
		}

		return filePath, nil
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
	do.Provide(injector, func(i *do.Injector) (infrastructure.AntiAffinityListRuleRepositoryFactory, error) {
		configFilePath := do.MustInvoke[string](i)
		return infrastructure.NewAntiAffinityListRuleRepositoryFactory(configFilePath), nil
	})
	do.Provide(injector, func(i *do.Injector) (domain.AntiAffinityListRuleRepository, error) {
		factory := do.MustInvoke[infrastructure.AntiAffinityListRuleRepositoryFactory](i)
		return factory.Create()
	})
	// usecase
	do.Provide(injector, func(i *do.Injector) (usecase.ListByPathInputPort, error) {
		groupRuleRepository := do.MustInvoke[domain.AntiAffinityGroupRuleRepository](i)
		listRuleRepository := do.MustInvoke[domain.AntiAffinityListRuleRepository](i)
		return usecase.NewListByPathInputPort(groupRuleRepository, listRuleRepository), nil
	})
	// adapter
	do.Provide(injector, func(i *do.Injector) (interfaces.Config, error) {
		in := do.MustInvoke[usecase.ListByPathInputPort](i)
		return adapter.NewController(in), nil
	})

	return injector
}
