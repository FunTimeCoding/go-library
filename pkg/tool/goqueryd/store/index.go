package store

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/index"
	"strings"
	"time"
)

func (s *Store) Index(collection string) *index.Index {
	path, pattern := s.collectionPath(collection)
	now := time.Now().UTC().Format(time.RFC3339)
	seen := map[string]bool{}
	r := index.Stub()

	for _, relative := range globFiles(path, pattern) {
		content := string(system.ReadBytes(path, relative))

		if strings.TrimSpace(content) == "" {
			continue
		}

		hash := hashContent(content)
		title := extractTitle(content, relative)
		seen[relative] = true
		existing := s.findActiveDocument(collection, relative)

		if existing != nil {
			if existing.hash == hash {
				if existing.title != title {
					s.updateDocumentTitle(existing.identifier, title, now)
					r.Updated++
				} else {
					r.Unchanged++
				}
			} else {
				s.insertContent(hash, content, now)
				s.updateDocument(existing.identifier, title, hash, now)
				r.Updated++
			}
		} else {
			s.insertContent(hash, content, now)
			s.insertDocument(collection, relative, title, hash, now)
			r.Indexed++
		}
	}

	for _, active := range s.activeDocumentPaths(collection) {
		if !seen[active] {
			s.deactivateDocument(collection, active)
			r.Removed++
		}
	}

	return r
}
