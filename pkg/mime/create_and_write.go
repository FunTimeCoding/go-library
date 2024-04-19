package mime

import "mime/multipart"

func CreateAndWrite(
	w *multipart.Writer,
	fileType string,
	fileName string,
	b []byte,
) {
	Write(Create(w, fileType, fileName), b)
}
