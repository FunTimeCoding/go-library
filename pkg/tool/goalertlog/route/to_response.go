package route

import "github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"

func toResponse(records []store.Record) []AlertsResponse {
	result := make([]AlertsResponse, 0, len(records))

	for _, record := range records {
		entry := AlertsResponse{
			Fingerprint: record.Fingerprint,
			Name:        record.Name,
			Severity:    record.Severity,
			Summary:     record.Summary,
			Labels:      record.Labels,
			Start:       record.Start.Format(DateFormat),
			Status:      Firing,
		}

		if record.End != nil {
			entry.End = record.End.Format(DateFormat)
			entry.Status = Resolved
		}

		result = append(result, entry)
	}

	return result
}
