package constant

func ValidSurface(surface string) bool {
	switch surface {
	case ModelContext, CommandLine, Web, WebService:
		return true
	default:
		return false
	}
}
