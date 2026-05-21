package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/resolve_result"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

func (s *Store) ResolveSessionIdentifier(query string) *resolve_result.Result {
	seen := map[string]bool{}
	var matches []*resolve_result.Match
	var byCallsign session.Session

	if r := s.database.Where(
		"callsign = ?",
		query,
	).Limit(1).Find(&byCallsign); r.RowsAffected > 0 {
		seen[byCallsign.Identifier] = true
		matches = append(matches, matchFrom(&byCallsign, constant.Callsign))
	}

	var byIdentifier session.Session

	if r := s.database.Where(
		"identifier = ?",
		query,
	).Limit(1).Find(&byIdentifier); r.RowsAffected > 0 && !seen[byIdentifier.Identifier] {
		seen[byIdentifier.Identifier] = true
		matches = append(matches, matchFrom(&byIdentifier, constant.Identifier))
	}

	var byAlias session.Session

	if r := s.database.Where(
		"alias = ?",
		query,
	).Limit(1).Find(&byAlias); r.RowsAffected > 0 && !seen[byAlias.Identifier] {
		seen[byAlias.Identifier] = true
		matches = append(matches, matchFrom(&byAlias, constant.Alias))
	}

	var byPrefix []session.Session
	s.database.Where(
		"identifier LIKE ?",
		fmt.Sprintf("%s%%", query),
	).Find(&byPrefix)

	for i := range byPrefix {
		if !seen[byPrefix[i].Identifier] {
			seen[byPrefix[i].Identifier] = true
			matches = append(matches, matchFrom(&byPrefix[i], "prefix"))
		}
	}

	return resolve_result.New(matches)
}
