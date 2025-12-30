package maintenance_log

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Store struct {
	db *gorm.DB
}

func NewStore() (*Store, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open in-memory database: %w", err)
	}

	if err := db.AutoMigrate(&Entry{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &Store{db: db}, nil
}

func (s *Store) AddEntry(entry *Entry) error {
	if err := s.db.Create(entry).Error; err != nil {
		return fmt.Errorf("failed to create entry: %w", err)
	}
	return nil
}

func (s *Store) ListEntries(filter *EntryFilter) ([]Entry, error) {
	query := s.db.Model(&Entry{})

	if filter != nil {
		if filter.System != "" {
			query = query.Where("system = ?", filter.System)
		}
		if filter.Service != "" {
			query = query.Where("service = ?", filter.Service)
		}
		if filter.User != "" {
			query = query.Where("user = ?", filter.User)
		}
		if !filter.Since.IsZero() {
			query = query.Where("timestamp >= ?", filter.Since)
		}
		if !filter.Until.IsZero() {
			query = query.Where("timestamp <= ?", filter.Until)
		}
		if filter.Limit > 0 {
			query = query.Limit(filter.Limit)
		}
	}

	var entries []Entry
	if err := query.Order("timestamp DESC").Find(&entries).Error; err != nil {
		return nil, fmt.Errorf("failed to list entries: %w", err)
	}

	return entries, nil
}

type EntryFilter struct {
	System  string
	Service string
	User    string
	Since   time.Time
	Until   time.Time
	Limit   int
}
