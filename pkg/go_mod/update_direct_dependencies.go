package go_mod

func UpdateDirectDependencies() {
	for _, dep := range Read().Require {
		if dep.Indirect {
			continue
		}

		Update(dep.Mod.Path)
	}
}
