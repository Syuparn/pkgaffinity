package schema

// Config defines config file yaml schema.
type ConfigSchema struct {
	Version           string            `yaml:"version"`
	AntiAffinityRules AntiAffinityRules `yaml:"antiAffinityRules"`
}

type AntiAffinityRules struct {
	Groups []*AntiAffinityGroupRule `yaml:"groups"`
	Lists  []*AntiAffinityListRule  `yaml:"lists"`
}

type AntiAffinityGroupRule struct {
	PathPrefix  string   `yaml:"pathPrefix"`
	AllowNames  []string `yaml:"allowNames"`
	IgnorePaths []string `yaml:"ignorePaths"`
}

type AntiAffinityListRule struct {
	Label        string   `yaml:"label"`
	PathPrefixes []string `yaml:"pathPrefixes"`
}
