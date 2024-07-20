package gw2

import (
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
)

func WrapLogs(l []*log_manager.Log) []*log.Log {
	var result []*log.Log

	for _, element := range l {
		result = append(result, log.New(element))
	}

	return result
}
