package silence

import "github.com/prometheus/alertmanager/api/v2/models"

func NewSlice(v models.GettableSilences) []*Silence {
	var result []*Silence

	for _, s := range v {
		result = append(result, New(s))
	}

	return result
}
