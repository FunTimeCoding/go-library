package store

import "github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/store/mute_rule"

func (s *Store) CreateMuteRule(
	reason string,
	message string,
	cluster string,
) error {
	return s.database.Create(
		mute_rule.New(reason, message, cluster),
	).Error
}
