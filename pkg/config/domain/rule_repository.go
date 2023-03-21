package domain

//go:generate go run github.com/matryer/moq -fmt goimports -out zz_generated_moq_rule_repository.go . AntiAffinityGroupRuleRepository
type AntiAffinityGroupRuleRepository interface {
	ListByPath(path Path) ([]*AntiAffinityGroupRule, error)
}
