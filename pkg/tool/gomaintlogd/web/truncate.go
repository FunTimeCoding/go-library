package web

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}

	return s[:n-3] + "..."
}
