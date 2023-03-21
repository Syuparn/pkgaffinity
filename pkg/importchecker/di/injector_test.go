package di

import (
	"testing"

	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
	"github.com/syuparn/pkgaffinity/interfaces"
)

func TestNewInjector(t *testing.T) {
	configControllerMock := interfaces.ConfigMock{}
	injector := NewInjector(&configControllerMock)

	controller, err := do.Invoke[interfaces.ImportChecker](injector)

	assert.NoError(t, err)
	assert.NotEmpty(t, controller)
}
