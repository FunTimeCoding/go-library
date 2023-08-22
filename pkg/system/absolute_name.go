package system

func AbsoluteName() string {
	return Run("hostname", "-f")
}
