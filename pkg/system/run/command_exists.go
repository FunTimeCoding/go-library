package run

func CommandExists(name string) bool {
	r := New()
	r.Panic = false
	r.Start("which", name)

	return r.ExitCode == 0
}
