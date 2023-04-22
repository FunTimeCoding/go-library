package above_below

// Integer
//
//	If value is above mark, above closure is called
//	If value is below mark, below closure is called
func Integer(
	value int,
	mark int,
	above func(),
	below func(),
) {
	if value > mark {
		above()
	} else if value < mark {
		below()
	}
}
