package constant

func ValidKind(kind string) bool {
	switch kind {
	case Baseline, Domain:
		return true
	default:
		return false
	}
}
