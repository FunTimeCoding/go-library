package system

func FileSize(path string) int64 {
	return Stat(path).Size()
}
