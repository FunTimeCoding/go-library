package worker

func (w *Worker) Stop(_ bool) {
	w.Face.Stop()
}
