package system

func EnsureFileDeleted(path string) {
	if FileExists(path) {
		DeleteFile(path)
	}
}
