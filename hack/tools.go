package hack

// HACK: import these packages to fix tools versions explicitly in go.mod
import (
	_ "github.com/matryer/moq/generate" //nolint:typecheck
)
