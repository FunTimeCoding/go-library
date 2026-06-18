package service

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
)

func (s *Service) EditSession(
	identifier string,
	a *argument.EditSession,
) error {
	if a.Alias != nil && *a.Alias != "" {
		owner, e := s.store.AliasOwner(*a.Alias)

		if e != nil {
			return e
		}

		if owner != "" && owner != identifier {
			return fmt.Errorf(
				"%w: %q is used by session %s",
				constant.ErrorAliasCollision,
				*a.Alias,
				owner,
			)
		}
	}

	if e := s.store.EditSession(identifier, a); e != nil {
		return e
	}

	s.notify()

	return nil
}
