package domain

import "strings"

type Path string

func NewPath(path string) Path {
	// NOTE: path may have suffix .test if the file is *_test.go and package is `main`.
	// Since this suffix is nothing to do with the affinity rules, we ignore it.
	return Path(strings.TrimSuffix(path, ".test"))
}

type Name string
type PathPrefix string

func (p PathPrefix) Contains(path Path) bool {
	return string(path) == string(p) || strings.HasPrefix(string(path), string(p)+"/")
}
