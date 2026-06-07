package concern

func NewPackage(
	key string,
	text string,
	path string,
) *Concern {
	return &Concern{
		Key:  key,
		Text: text,
		Path: path,
		Type: Package,
	}
}
