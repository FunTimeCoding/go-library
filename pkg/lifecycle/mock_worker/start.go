package mock_worker

import "fmt"

func (w *MockWorker) Start() {
	*w.Log = append(*w.Log, fmt.Sprintf("start:%s", w.Tag))
}
