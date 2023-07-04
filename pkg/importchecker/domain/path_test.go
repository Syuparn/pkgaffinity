package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected Path
	}{
		{
			name:     "create a path",
			path:     "foo/bar",
			expected: Path("foo/bar"),
		},
		{
			name:     "suffix .test is ignored",
			path:     "foo/bar.test",
			expected: Path("foo/bar"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewPath(tt.path)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPathPrefixContains(t *testing.T) {
	tests := []struct {
		name     string
		prefix   PathPrefix
		path     Path
		expected bool
	}{
		{
			name:     "path is same as pathPrefix",
			prefix:   "foo/bar",
			path:     "foo/bar",
			expected: true,
		},
		{
			name:     "path is under pathPrefix",
			prefix:   "foo/bar",
			path:     "foo/bar/baz",
			expected: true,
		},
		{
			name:     "path is not in pathPrefix",
			prefix:   "foo/bar",
			path:     "foo/hoge",
			expected: false,
		},
		{
			name:     "path matches to pathPrefix literally but not in pathPrefix",
			prefix:   "foo/bar",
			path:     "foo/barbara",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.prefix.Contains(tt.path)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
