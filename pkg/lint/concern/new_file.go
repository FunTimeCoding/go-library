package concern

func NewFile(
	key string,
	text string,
	path string,
	fixed bool,
) *Concern {
	return &Concern{
		Key:   key,
		Text:  text,
		Path:  path,
		Type:  File,
		Fixed: fixed,
	}
}
