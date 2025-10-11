package constant

const File = "dictionary.dic"

var (
	Skip       = map[string]bool{".git": true, "tmp": true}
	Extensions = map[string]bool{
		".conf":      true,
		".envrc":     true,
		".gitignore": true,
		".go":        true,
		".json":      true,
		".md":        true,
		".rego":      true,
		".sh":        true,
		".txt":       true,
		".xml":       true,
		".yaml":      true,
		".yml":       true,
	}
	NoExtension = map[string]bool{
		"Containerfile": true,
		"Dockerfile":    true,
	}
)
