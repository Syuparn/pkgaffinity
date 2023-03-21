package pkgaffinity

import (
	"github.com/samber/do"

	"github.com/syuparn/pkgaffinity/interfaces"
	configdi "github.com/syuparn/pkgaffinity/pkg/config/di"
	importcheckerdi "github.com/syuparn/pkgaffinity/pkg/importchecker/di"
)

var configInjector = configdi.NewInjector()
var importcheckerInjector = importcheckerdi.NewInjector()

func getImportChecker() interfaces.ImportChecker {
	do.Override(importcheckerInjector, func(i *do.Injector) (interfaces.Config, error) {
		return do.Invoke[interfaces.Config](configInjector)
	})

	return do.MustInvoke[interfaces.ImportChecker](importcheckerInjector)
}
