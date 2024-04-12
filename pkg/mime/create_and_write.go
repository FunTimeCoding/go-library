package mime

import "mime/multipart"

func CreateAndWrite(
	w *multipart.Writer,
	f *File,
) {
	Write(Create(w, f), f)
}
