package response

type FormulaInstallation struct {
	Version               string              `json:"version"`
	UsedOptions           []any               `json:"used_options"`
	BuiltAsBottle         bool                `json:"built_as_bottle"`
	PouredFromBottle      bool                `json:"poured_from_bottle"`
	Time                  int                 `json:"time"`
	RuntimeDependencies   []RuntimeDependency `json:"runtime_dependencies"`
	InstalledAsDependency bool                `json:"installed_as_dependency"`
	InstalledOnRequest    bool                `json:"installed_on_request"`
}
