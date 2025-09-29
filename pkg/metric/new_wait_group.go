package metric

import "sync"

func NewWaitGroup() *sync.WaitGroup {
	return &sync.WaitGroup{}
}
