package repository

import (
	"github.com/google/go-github/v85/github"
	"time"
)

type Repository struct {
	Identifier int64
	Owner      string
	Name       string
	FullName   string
	CreatedAt  time.Time
	Raw        *github.Repository
}
