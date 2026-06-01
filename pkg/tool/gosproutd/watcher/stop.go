package watcher

func (w *Watcher) Stop() {
	close(w.stop)
}
