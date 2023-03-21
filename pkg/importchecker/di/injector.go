package di

import (
	"io"
	"os"

	"github.com/samber/do"

	"github.com/syuparn/pkgaffinity/interfaces"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/adapter"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/domain"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/infrastructure"
	"github.com/syuparn/pkgaffinity/pkg/importchecker/usecase"
)

func NewInjector(
	configController interfaces.Config,
) *do.Injector {
	injector := do.New()

	// settings
	do.Provide(injector, func(i *do.Injector) (io.Writer, error) {
		return os.Stdout, nil
	})
	do.Provide(injector, func(i *do.Injector) (interfaces.Config, error) {
		return configController, nil
	})
	// domain
	do.Provide(injector, func(i *do.Injector) (domain.AntiAffinityRuleRepository, error) {
		controller := do.MustInvoke[interfaces.Config](i)
		return infrastructure.NewAntiAffinityRuleRepository(controller), nil
	})
	// usecase
	do.Provide(injector, func(i *do.Injector) (usecase.CheckImportsInputPort, error) {
		out := do.MustInvoke[usecase.CheckImportsOutputPort](i)
		ruleRepository := do.MustInvoke[domain.AntiAffinityRuleRepository](i)
		return usecase.NewCheckImportsInputPort(out, ruleRepository), nil
	})
	do.Provide(injector, func(i *do.Injector) (usecase.CheckImportsOutputPort, error) {
		writer := do.MustInvoke[io.Writer](i)
		return adapter.NewCheckImportsOutputPort(writer), nil
	})
	// adapter
	do.Provide(injector, func(i *do.Injector) (interfaces.ImportChecker, error) {
		in := do.MustInvoke[usecase.CheckImportsInputPort](i)
		return adapter.NewController(in), nil
	})

	return injector
}
