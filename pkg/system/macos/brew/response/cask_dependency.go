package response

type CaskDependency struct {
	Macos   MacosDependency  `json:"macos"`
	Arch    []ArchDependency `json:"arch,omitempty"`
	Cask    []string         `json:"cask,omitempty"`
	Formula []string         `json:"formula,omitempty"`
}
