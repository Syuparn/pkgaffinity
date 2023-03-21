package di

import (
	"testing"

	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
	"github.com/syuparn/pkgaffinity/interfaces"
)

func TestNewInjector(t *testing.T) {
	injector := NewInjector()
	// override controller
	do.Override(injector, func(i *do.Injector) (interfaces.Config, error) {
		return &interfaces.ConfigMock{}, nil
	})

	controller, err := do.Invoke[interfaces.ImportChecker](injector)

	assert.NoError(t, err)
	assert.NotEmpty(t, controller)
}

func TestNewInjectorError(t *testing.T) {
	injector := NewInjector()
	// controller must be overridden
	_, err := do.Invoke[interfaces.ImportChecker](injector)

	assert.Error(t, err)
}
