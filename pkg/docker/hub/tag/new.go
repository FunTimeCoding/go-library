package tag

func New(v *response) *Tag {
	return &Tag{Name: v.Name, LastUpdated: v.LastUpdated}
}
