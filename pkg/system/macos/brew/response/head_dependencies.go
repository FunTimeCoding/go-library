package response

type HeadDependencies struct {
	BuildDependencies       []string     `json:"build_dependencies"`
	Dependencies            []string     `json:"dependencies"`
	TestDependencies        []string     `json:"test_dependencies"`
	RecommendedDependencies []any        `json:"recommended_dependencies"`
	OptionalDependencies    []any        `json:"optional_dependencies"`
	UsesFromMacos           []any        `json:"uses_from_macos"`
	UsesFromMacosBounds     []MacosBound `json:"uses_from_macos_bounds"`
}
