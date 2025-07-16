package system

func IsExecutable(path string) bool {
	return Mode(path)&0111 != 0
}
