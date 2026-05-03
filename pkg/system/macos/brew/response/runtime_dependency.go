package response

type RuntimeDependency struct {
	FullName         string `json:"full_name"`
	Version          string `json:"version"`
	Revision         int    `json:"revision,omitempty"`
	BottleRebuild    int    `json:"bottle_rebuild,omitempty"`
	PackageVersion   string `json:"pkg_version,omitempty"`
	DeclaredDirectly bool   `json:"declared_directly"`
}
