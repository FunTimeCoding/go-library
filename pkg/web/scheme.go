package web

func Scheme(secure bool) string {
	if secure {
		return SecureScheme
	}

	return InsecureSchemePrefix
}
