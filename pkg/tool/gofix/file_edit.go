package gofix

type fileEdit struct {
	offset  int
	length  int
	newText string
}
