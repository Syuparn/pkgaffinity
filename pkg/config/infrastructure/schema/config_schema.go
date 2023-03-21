package schema

// Config defines config file yaml schema.
type ConfigSchema struct {
	Version           string            `yaml:"version"`
	AntiAffinityRules AntiAffinityRules `yaml:"antiAffinityRules"`
}

type AntiAffinityRules struct {
	Groups []*AntiAffinityGroupRule `yaml:"groups"`
	// TODO: add lists
}

type AntiAffinityGroupRule struct {
	PathPrefix string   `yaml:"pathPrefix"`
	AllowNames []string `yaml:"allowNames"`
}
