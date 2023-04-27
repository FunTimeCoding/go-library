package normalize_change

// Integer
//
//	If now is at maximum and change is positive, change is 0
//	If now is below maximum and change exceeds it, change is the remaining amount
//	Same for the minimum
func Integer(
	now int,
	change int,
	minimum int,
	maximum int,
) int {
	if maximum > minimum && now+change > maximum {
		return maximum - now
	}

	if now+change < minimum {
		return (now - minimum) * -1
	}

	return change
}
