package service

import "fmt"

func safe(f func()) (e error) {
	defer func() {
		if r := recover(); r != nil {
			if v, okay := r.(error); okay {
				e = v
			} else {
				e = fmt.Errorf("%v", r)
			}
		}
	}()
	f()

	return nil
}
