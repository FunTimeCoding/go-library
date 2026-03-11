package vfs

func New() *VFS {
	return &VFS{files: make(map[string]string)}
}
