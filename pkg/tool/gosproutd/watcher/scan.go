package watcher

func (w *Watcher) scan() {
	files := w.walkDirectory()
	var paths []string

	for _, f := range files {
		paths = append(paths, f.path)
		w.store.UpsertSeed(f.name, f.path, f.contentHash, f.content)
	}

	w.store.RemoveMissing(paths)
	w.notifier.Notify()
}
