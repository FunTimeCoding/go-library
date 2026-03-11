package vfs

func (v *VFS) Write(
	path string,
	content string,
) {
	v.files[path] = content
}
