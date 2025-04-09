package page_file

func New(
	name string,
	body string,
) *File {
	return &File{Name: name, Body: body}
}
