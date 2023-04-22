package above_below

// Float
//
//	See Integer
func Float(
	value float64,
	mark float64,
	magnitude float64,
	above func(),
	below func(),
) {
	if value > mark && value > magnitude {
		above()
	} else if value < mark && value*-1 > magnitude {
		below()
	}
}
