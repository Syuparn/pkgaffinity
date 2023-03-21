package domain

import "strings"

type Path string
type PathPrefix string

type AntiAffinityGroupRule struct {
	Group PathPrefix
}

func (r *AntiAffinityGroupRule) Contains(path Path) bool {
	// NOTE: add "/" at the end to remove paths such as
	// 1. group itself
	// 2. literally has prefix but refers another path (ex: group: `foo/bar`, path: `foo/barbara`)
	return strings.HasPrefix(string(path), string(r.Group)+"/")
}
