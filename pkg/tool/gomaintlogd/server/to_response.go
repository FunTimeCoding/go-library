package server

import (
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store/entry"
)

func toResponse(entries []entry.Entry) []generated.EntryResponse {
	result := make([]generated.EntryResponse, len(entries))

	for i, e := range entries {
		r := generated.EntryResponse{
			Id:        int(e.ID),
			Timestamp: e.Timestamp,
			Action:    e.Action,
			User:      e.User,
			CreatedAt: &e.CreatedAt,
		}

		if e.System != "" {
			r.System = &e.System
		}

		if e.Service != "" {
			r.Service = &e.Service
		}

		if e.Description != "" {
			r.Description = &e.Description
		}

		result[i] = r
	}

	return result
}
