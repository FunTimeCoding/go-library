package mock_worker

import "fmt"

type MockWorker struct {
	Tag string
	Log *[]string
}

func (w *MockWorker) Start() {
	*w.Log = append(*w.Log, fmt.Sprintf("start:%s", w.Tag))
}

func (w *MockWorker) Stop() {
	*w.Log = append(*w.Log, fmt.Sprintf("stop:%s", w.Tag))
}
