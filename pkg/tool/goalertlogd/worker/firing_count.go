package worker

func (w *Worker) FiringCount() int {
	return len(w.firing)
}
