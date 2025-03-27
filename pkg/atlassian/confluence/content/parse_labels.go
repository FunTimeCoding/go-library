package content

import kaos "github.com/essentialkaos/go-confluence/v6"

func parseLabels(v *kaos.Content) []string {
	var labels []string

	if v.Metadata != nil && v.Metadata.Labels != nil {
		for _, l := range v.Metadata.Labels.Result {
			labels = append(labels, l.Name)
		}
	}

	return labels
}
