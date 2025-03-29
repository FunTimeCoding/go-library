package run

func CommandExists(name string) bool {
	r := New()
	r.Panic = false
	r.Start("which", name)

	return r.Exit == 0
}
