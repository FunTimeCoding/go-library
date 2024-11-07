package system

func FindUnusedPort(startPort int) int {
	endPort := startPort + 100

	for port := startPort; port < endPort; port++ {
		if IsPortUnused(port) {
			return port
		}
	}

	return -1
}
