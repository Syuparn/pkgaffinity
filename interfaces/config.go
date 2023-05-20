package interfaces

//go:generate go run github.com/matryer/moq -fmt goimports -out zz_generated_moq_config.go . Config
type Config interface {
	ListRulesByPath(*ListRulesByPathRequest) (*ListRulesByPathResponse, error)
}

type ListRulesByPathRequest struct {
	PackagePath string
}

type ListRulesByPathResponse struct {
	AntiAffinityGroupRules []*AntiAffinityGroupRule
	AntiAffinityListRules  []*AntiAffinityListRule
}

type AntiAffinityGroupRule struct {
	GroupPathPrefix string
	AllowNames      []string
	IgnorePaths     []string
}

type AntiAffinityListRule struct {
	Label        string
	PathPrefixes []string
}
