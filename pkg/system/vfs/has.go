package vfs

func (v *VFS) Has(path string) bool {
	_, ok := v.files[path]

	return ok
}
