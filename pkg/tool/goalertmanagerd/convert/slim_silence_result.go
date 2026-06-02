package convert

type SlimSilenceResult struct {
	Silences []*SlimSilence `json:"silences"`
	Total    int            `json:"total"`
}
