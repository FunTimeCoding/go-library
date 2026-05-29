package store

import "github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/store/mute_rule"

func (s *Store) DeleteMuteRule(identifier uint) error {
	return s.database.Delete(mute_rule.Stub(), identifier).Error
}
