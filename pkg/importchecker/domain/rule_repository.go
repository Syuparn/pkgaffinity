package domain

// AntiAffinityRuleRepository provides AntiAffinityRule as a query
//
//go:generate go run github.com/matryer/moq -fmt goimports -out zz_generated_moq_rule_repository.go . AntiAffinityRuleRepository
type AntiAffinityRuleRepository interface {
	ListByPath(Path) ([]AntiAffinityRule, error)
}
