package integers32

import "strconv"

func ToString(number int32) string {
	return strconv.FormatInt(int64(number), 10)
}
