package pkgaffinity

import (
	"github.com/samber/do"
	"github.com/syuparn/pkgaffinity/interfaces"
	configdi "github.com/syuparn/pkgaffinity/pkg/config/di"
	importcheckerdi "github.com/syuparn/pkgaffinity/pkg/importchecker/di"
)

func getCheckImportController() interfaces.ImportChecker {
	configInjector := configdi.NewInjector()
	importcheckerInjector := importcheckerdi.NewInjector(do.MustInvoke[interfaces.Config](configInjector))

	return do.MustInvoke[interfaces.ImportChecker](importcheckerInjector)
}
