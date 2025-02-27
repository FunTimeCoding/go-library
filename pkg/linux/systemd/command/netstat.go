package command

func Netstat() []string {
	return []string{
		"ip",
		"-all",
		"netns",
		"exec",
		"netstat",
		"--program",
		"--listening",
		"--tcp",
		"--numeric",
	}
}
