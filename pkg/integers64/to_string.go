package integers64

import "strconv"

func ToString(number int64) string {
	return strconv.FormatInt(number, 10)
}
