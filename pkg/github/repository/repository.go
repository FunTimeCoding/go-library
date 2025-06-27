package repository

import (
	"github.com/google/go-github/v70/github"
	"time"
)

type Repository struct {
	Identifier int64
	Owner      string
	Name       string
	CreatedAt  time.Time

	Raw *github.Repository
}
