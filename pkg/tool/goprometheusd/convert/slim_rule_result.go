package convert

type SlimRuleResult struct {
	Rules []*SlimRule `json:"rules"`
	Total int         `json:"total"`
}
