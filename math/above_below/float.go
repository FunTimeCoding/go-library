package above_below

// Float
//
//	See Integer
func Float(
	f float64,
	magnitude float64,
	above func(),
	below func(),
) {
	if f > magnitude {
		above()
	} else if f*-1 > magnitude {
		below()
	}
}
