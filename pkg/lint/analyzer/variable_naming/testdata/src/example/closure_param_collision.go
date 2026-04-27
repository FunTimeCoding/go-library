package example

import (
	"sort"
	"strconv"
)

func ClosureParamCollision() {
	values := []string{"3", "1", "2"}
	sort.Slice(
		values,
		func(a, b int) bool {
			x, xFail := strconv.Atoi(values[a]) // want `variable x of type int should be named i` `variable xFail of type error should be named e`
			y, yFail := strconv.Atoi(values[b]) // want `variable y of type int should be named c` `variable yFail of type error should be named f`

			if xFail != nil || yFail != nil {
				return values[a] > values[b]
			}

			return x > y
		},
	)
}
