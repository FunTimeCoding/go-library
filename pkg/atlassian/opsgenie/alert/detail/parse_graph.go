package detail

import "strings"

func (p *Prometheus) ParseGraph() string {
	if p.Graph == "" {
		return ""
	}

	result := p.Graph
	result = strings.ReplaceAll(result, ">", "%3E")
	result = strings.ReplaceAll(result, "<", "%3C")

	return shortenGraph(result)
}
