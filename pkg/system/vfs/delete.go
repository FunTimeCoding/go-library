package vfs

func (v *VFS) Delete(path string) {
	delete(v.files, path)
}
