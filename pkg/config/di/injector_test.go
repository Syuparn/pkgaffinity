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
	t.Setenv("PKGAFFINITY_CONFIG_PATH", "testdata/.pkgaffinity.yaml")

	controller, err := do.Invoke[interfaces.Config](injector)

	assert.NoError(t, err)
	assert.NotEmpty(t, controller)
}
