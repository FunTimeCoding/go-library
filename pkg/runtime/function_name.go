package runtime

import (
	"reflect"
	"runtime"
)

func FunctionName(i any) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
