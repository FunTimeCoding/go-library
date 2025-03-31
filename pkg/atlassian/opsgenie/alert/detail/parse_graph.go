package detail

import "strings"

func (p *Prometheus) ParseGraph() string {
	result := p.Graph
	result = strings.ReplaceAll(result, ">", "%3E")
	result = strings.ReplaceAll(result, "<", "%3C")

	return shortenGraph(result)
}
