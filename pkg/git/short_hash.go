package git

func ShortHash(path string) string {
	return Head(Open(path)).Hash().String()[:8]
}
