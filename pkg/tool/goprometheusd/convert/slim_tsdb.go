package convert

type SlimTSDB struct {
	NumSeries                   int        `json:"num_series"`
	NumLabelPairs               int        `json:"num_label_pairs"`
	ChunkCount                  int        `json:"chunk_count"`
	SeriesCountByMetricName     []SlimStat `json:"series_count_by_metric_name,omitempty"`
	LabelValueCountByLabelName  []SlimStat `json:"label_value_count_by_label_name,omitempty"`
	MemoryInBytesByLabelName    []SlimStat `json:"memory_in_bytes_by_label_name,omitempty"`
	SeriesCountByLabelValuePair []SlimStat `json:"series_count_by_label_value_pair,omitempty"`
}
