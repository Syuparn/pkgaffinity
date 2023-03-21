package domain

import (
	"strings"
)

type Path string
type PathPrefix string
type Name string

type AntiAffinityGroupRule struct {
	Group      PathPrefix
	AllowNames []Name
}

func (r *AntiAffinityGroupRule) Contains(path Path) bool {
	// NOTE: add "/" at the end to remove paths such as
	// 1. group itself
	// 2. literally has prefix but refers another path (ex: group: `foo/bar`, path: `foo/barbara`)
	return strings.HasPrefix(string(path), string(r.Group)+"/")
}

type RuleLabel string

type AntiAffinityListRule struct {
	Label    RuleLabel
	Prefixes []PathPrefix
}

func (r *AntiAffinityListRule) Contains(path Path) bool {
	for _, prefix := range r.Prefixes {
		// NOTE: strings.HasPrefix(string(path), string(prefix)) is insufficient because it may refer another path
		// (ex: group: `foo/bar`, path: `foo/barbara`)

		if string(path) == string(prefix) {
			return true
		}

		if strings.HasPrefix(string(path), string(prefix)+"/") {
			return true
		}
	}

	return false
}
