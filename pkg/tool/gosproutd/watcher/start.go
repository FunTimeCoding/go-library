package watcher

func (w *Watcher) Start() {
	go w.run()
}
