package rule_parser

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"

func (p *Parser) Documentation(name string) string {
	if p.list == nil {
		return constant.None
	}

	result := p.list.FindDocumentation(name)

	if result == "" {
		result = constant.None
	}

	return result
}
