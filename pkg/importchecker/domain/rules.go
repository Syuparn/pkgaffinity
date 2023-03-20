package domain

import (
	"fmt"
	"strings"
)

type Path string
type Name string
type PathPrefix string

type AntiAffinityRuleName string

// Violation represents why affinity rule is not fulfilled.
type Violation struct {
	ImportPath  Path
	PackagePath Path
	RuleName    AntiAffinityRuleName
}

type AntiAffinityRule interface {
	Check(Path) *Violation
}

// AntiAffinityGroupRule defines import path anti-affinity group of a package.
// Rule `foo/bar` means any package inside it cannot import any other packages except itself.
// (ex: `foo/bar/baz/*` cannot import `foo/bar/quux/*` and can only import `foo/bar/baz/*`)
type AntiAffinityGroupRule struct {
	// selfPath is the path of the package
	// ex: `"foo/bar/baz/hoge"`
	selfPath Path
	// groupPathPrefix defines anti-affinity group path prefix.
	// ex: `"foo/bar"`
	groupPathPrefix PathPrefix
	// allowNames is the list of package names directly under groupPathPrefix which the specified import path can contain.
	// ex: `[]string{"foo/bar/baz"}`
	allowNames []Name
}

// impl check
var _ AntiAffinityRule = &AntiAffinityGroupRule{}

func NewAntiAffinityGroupRule(self Path, group PathPrefix) (*AntiAffinityGroupRule, error) {
	// NOTE: strings.HasPrefix(string(self), string(group)) is insufficient!
	// ex1. different path but has same prefix:
	//     self == `foo/barbara`, group == `foo/bar` (invalid)
	// ex2. self is same as group (because group rule defines for group subpaths, not group itself):
	//     self == `foo/bar`, group == `foo/bar`
	if !strings.HasPrefix(string(self), string(group)+"/") {
		return nil, fmt.Errorf("self `%s` must be in group `%s`", self, group)
	}

	relativeSelfPath := strings.TrimPrefix(string(self), string(group)+"/")
	relativeSelfElems := strings.Split(relativeSelfPath, "/")

	return &AntiAffinityGroupRule{
		selfPath:        self,
		groupPathPrefix: group,
		allowNames: []Name{
			// group rule allows path which belongs to same package prefix as selfPath
			Name(relativeSelfElems[0]),
		},
	}, nil
}

func (r *AntiAffinityGroupRule) Check(path Path) *Violation {
	if !strings.HasPrefix(string(path), string(r.groupPathPrefix)+"/") {
		return nil
	}

	relativePath := strings.TrimPrefix(string(path), string(r.groupPathPrefix)+"/")
	name, _, _ := strings.Cut(relativePath, "/")
	for _, allowName := range r.allowNames {
		if Name(name) == allowName {
			return nil
		}
	}

	return &Violation{
		ImportPath:  path,
		PackagePath: r.selfPath,
		RuleName:    r.name(),
	}
}

func (r *AntiAffinityGroupRule) name() AntiAffinityRuleName {
	return AntiAffinityRuleName(fmt.Sprintf("anti-affinity group rule `%s`", r.groupPathPrefix))
}

// TODO: make AntiAffinityListRule
