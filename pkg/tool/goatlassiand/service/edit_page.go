package service

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
)

func (s *Service) EditPage(
	identifier string,
	oldText string,
	newText string,
	title string,
	message string,
) (*page.Page, error) {
	current, e := s.confluence.Page(identifier)

	if e != nil {
		return nil, e
	}

	markdown := page.ToMarkdown(current.Raw.Body.Storage.Value)
	newMarkdown, f := replaceUnique(markdown, oldText, newText)

	if f != nil {
		return nil, f
	}

	if title == "" {
		title = current.Name
	}

	return s.confluence.PutPage(
		identifier,
		title,
		page.ToStorage(newMarkdown),
		current.Raw.Version.Number+1,
		message,
		constant.CurrentStatus,
	)
}
