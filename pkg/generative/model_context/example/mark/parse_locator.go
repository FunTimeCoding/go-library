package mark

func parseLocator(locator string) string {
	// Format: user://{id}/profile
	l := len(locator)
	prefix := len(UserPrefix)
	suffix := len(ProfileSuffix)

	if l > prefix+suffix {
		return locator[prefix : l-suffix]
	}

	return ""
}
