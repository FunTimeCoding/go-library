package system

func FileEmpty(path string) bool {
	return Stat(path).Size() == 0
}
