package file

func New(b *[]byte) *File {
	return &File{source: b}
}
