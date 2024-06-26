package web

func SchemePrefix(secure bool) string {
	if secure {
		return SecureSchemePrefix
	}

	return InsecureSchemePrefix
}
