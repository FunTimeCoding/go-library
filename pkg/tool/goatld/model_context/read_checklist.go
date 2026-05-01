package model_context

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
)

func (s *Server) readChecklist(
	key string,
) ([]convert.ChecklistItem, error) {
	i, e := s.jira.Issue(key)

	if e != nil {
		return nil, e
	}

	value := i.CustomValue(constant.ChecklistField)

	if value == "" ||
		value == issue.NilValue ||
		value == issue.UnknownField ||
		value == issue.UnknownValue {
		return nil, nil
	}

	return ParseChecklist(value), nil
}
