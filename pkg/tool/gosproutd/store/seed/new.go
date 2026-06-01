package seed

func New(
	name string,
	path string,
	contentHash string,
	content string,
	position int,
) *Seed {
	return &Seed{
		Name:        name,
		Path:        path,
		ContentHash: contentHash,
		Content:     content,
		Position:    position,
	}
}
