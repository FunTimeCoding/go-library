package vfs

func (v *VFS) Read(path string) string {
	return v.files[path]
}
