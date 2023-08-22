package above_below

// Integer
//
//	If i > magnitude, above closure is called
//	If i*-1 > magnitude, below closure is called
func Integer(
	i int,
	magnitude int,
	above func(),
	below func(),
) {
	if i > magnitude {
		above()
	} else if i*-1 > magnitude {
		below()
	}
}
