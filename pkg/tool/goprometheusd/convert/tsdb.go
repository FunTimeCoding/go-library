package convert

import "github.com/prometheus/client_golang/api/prometheus/v1"

func TSDB(v v1.TSDBResult) *SlimTSDB {
	result := &SlimTSDB{
		NumSeries:     v.HeadStats.NumSeries,
		NumLabelPairs: v.HeadStats.NumLabelPairs,
		ChunkCount:    v.HeadStats.ChunkCount,
	}

	for _, s := range v.SeriesCountByMetricName {
		result.SeriesCountByMetricName = append(
			result.SeriesCountByMetricName,
			SlimStat{Name: s.Name, Value: s.Value},
		)
	}

	for _, s := range v.LabelValueCountByLabelName {
		result.LabelValueCountByLabelName = append(
			result.LabelValueCountByLabelName,
			SlimStat{Name: s.Name, Value: s.Value},
		)
	}

	for _, s := range v.MemoryInBytesByLabelName {
		result.MemoryInBytesByLabelName = append(
			result.MemoryInBytesByLabelName,
			SlimStat{Name: s.Name, Value: s.Value},
		)
	}

	for _, s := range v.SeriesCountByLabelValuePair {
		result.SeriesCountByLabelValuePair = append(
			result.SeriesCountByLabelValuePair,
			SlimStat{Name: s.Name, Value: s.Value},
		)
	}

	return result
}
