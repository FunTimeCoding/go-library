package face

import "io/fs"

type SnippetClient interface {
	WriteFile(
		path string,
		content []byte,
	)
	ReadFile(path string) []byte
	ListDirectory(path string) []fs.FileInfo
	RemoveFile(path string)
}
