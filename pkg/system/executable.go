package system

func Executable(path string) {
	ChangeMode(path, Mode(path)|0111)
}
