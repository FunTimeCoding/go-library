package log

import "github.com/funtimecoding/go-library/pkg/gw2/log_manager"

func NewSlice(l []*log_manager.Log) []*Log {
	var result []*Log

	for _, element := range l {
		result = append(result, New(element))
	}

	SortByTime(result)

	return result
}
