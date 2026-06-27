package service

import "github.com/funtimecoding/go-library/pkg/tool/goatlassiand/face"

func New(c face.ConfluenceSource) *Service {
	return &Service{confluence: c}
}
