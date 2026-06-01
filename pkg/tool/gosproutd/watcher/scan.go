package watcher

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type scannedFile struct {
	name        string
	path        string
	contentHash string
	content     string
}

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

func (w *Watcher) walkDirectory() []scannedFile {
	var result []scannedFile

	filepath.Walk(w.seedDirectory, func(
		path string,
		i os.FileInfo,
		e error,
	) error {
		if e != nil {
			return nil
		}

		if i.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".md") {
			return nil
		}

		b, f := os.ReadFile(path)

		if f != nil {
			return nil
		}

		relative, g := filepath.Rel(w.seedDirectory, path)

		if g != nil {
			return nil
		}

		result = append(result, scannedFile{
			name:        strings.TrimSuffix(filepath.Base(path), ".md"),
			path:        relative,
			contentHash: fmt.Sprintf("%x", sha256.Sum256(b)),
			content:     string(b),
		})

		return nil
	})

	return result
}
