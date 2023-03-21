package domain

//go:generate go run github.com/matryer/moq -fmt goimports -out zz_generated_moq_rule_repository.go . AntiAffinityGroupRuleRepository AntiAffinityListRuleRepository
type AntiAffinityGroupRuleRepository interface {
	ListByPath(path Path) ([]*AntiAffinityGroupRule, error)
}

type AntiAffinityListRuleRepository interface {
	ListByPath(path Path) ([]*AntiAffinityListRule, error)
}
