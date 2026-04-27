package variable_naming

func scopePriority(k variableKind) int {
	switch k {
	case kindReceiver:
		return 0
	case kindParameter:
		return 1
	default:
		return 2
	}
}
