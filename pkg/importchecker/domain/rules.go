package domain

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

type RuleLabel string

// Violation represents why affinity rule is not fulfilled.
type Violation struct {
	ImportPath  Path
	PackagePath Path
	RuleLabel   RuleLabel
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

func NewAntiAffinityGroupRule(self Path, group PathPrefix, allowNames []Name) (*AntiAffinityGroupRule, error) {
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
		allowNames: append(
			allowNames,
			// group rule allows path which belongs to same package prefix as selfPath
			Name(relativeSelfElems[0]),
		),
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
		RuleLabel:   r.label(),
	}
}

func (r *AntiAffinityGroupRule) label() RuleLabel {
	return RuleLabel(fmt.Sprintf("anti-affinity group rule `%s`", r.groupPathPrefix))
}

// AntiAffinityListRule defines import path anti-affinity list of a package.
// any packages inside a path prefix cannot be imported.
// (ex: when prefix `foo/bar` and `baz` are defined, `foo/bar/*` and `baz/*` cannot be imported)
type AntiAffinityListRule struct {
	// selfPath is the path of the package
	// ex: `"foo/bar/baz/hoge"`
	selfPath Path
	// pathPrefix defines prefixes that selfPath cannot import
	// ex: `[]string{"foo/bar/quux"}`
	pathPrefixes []PathPrefix
	// label is only used as metadata to distinguish rules
	label RuleLabel
}

// impl check
var _ AntiAffinityRule = &AntiAffinityListRule{}

func NewAntiAffinityListRule(self Path, prefixes []PathPrefix, label RuleLabel) (*AntiAffinityListRule, error) {
	// NOTE: remove prefixes which contain selfPath because they don't break anti-affinity
	requiredPrefixes := lo.Reject(prefixes, func(p PathPrefix, _ int) bool {
		return p.Contains(self)
	})

	rule := &AntiAffinityListRule{
		selfPath:     self,
		pathPrefixes: requiredPrefixes,
		label:        label,
	}
	return rule, nil
}

func (r *AntiAffinityListRule) Check(path Path) *Violation {
	for _, prefix := range r.pathPrefixes {
		if prefix.Contains(path) {
			return &Violation{
				ImportPath:  path,
				PackagePath: r.selfPath,
				RuleLabel:   r.label,
			}
		}
	}

	return nil
}
