package domain

import "strings"

type Path string
type Name string
type PathPrefix string

func (p PathPrefix) Contains(path Path) bool {
	return string(path) == string(p) || strings.HasPrefix(string(path), string(p)+"/")
}
