package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/api"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"
)

func toResponse(records []store.Record) []api.AlertsResponse {
	result := make([]api.AlertsResponse, 0, len(records))

	for _, record := range records {
		entry := api.AlertsResponse{
			Fingerprint: record.Fingerprint,
			Name:        record.Name,
			Severity:    record.Severity,
			Summary:     record.Summary,
			Labels:      record.Labels,
			Start:       record.Start.Format(DateFormat),
			Status:      api.Firing,
		}

		if record.End != nil {
			entry.End = new(record.End.Format(DateFormat))
			entry.Status = api.Resolved
		}

		result = append(result, entry)
	}

	return result
}
