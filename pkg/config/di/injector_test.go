package di

import (
	"testing"

	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
	"github.com/syuparn/pkgaffinity/interfaces"
)

func TestNewInjector(t *testing.T) {
	injector := NewInjector()

	// HACK: override config file path to right one
	do.Override(injector, func(i *do.Injector) (string, error) {
		return "testdata/.pkgaffinity.yaml", nil
	})

	controller, err := do.Invoke[interfaces.Config](injector)

	assert.NoError(t, err)
	assert.NotEmpty(t, controller)
}
