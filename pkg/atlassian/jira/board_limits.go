package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/board_limit"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"log"
)

func (c *Client) BoardLimits(identifier int) []*board_limit.Limit {
	var result []*board_limit.Limit

	b := c.Board(identifier)
	o := c.BoardConfiguration(identifier)

	if o.ColumnConfig.ConstraintType != constant.IssueCountType {
		log.Panicf(
			"unexpected constraint type: %s",
			o.ColumnConfig.ConstraintType,
		)
	}

	for _, column := range o.ColumnConfig.Columns {
		if column.Min == 0 && column.Max == 0 {
			continue
		}

		result = append(
			result,
			&board_limit.Limit{
				Board:   b.Name,
				Column:  column.Name,
				Maximum: column.Max,
			},
		)
	}

	return result
}
