package service

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/site"
	"time"
)

func (s *Service) pollUsage() {
	defer func() {
		if r := recover(); r != nil {
			s.logger.Structured(
				"usage poll failed",
				"error",
				r,
			)
		}
	}()
	needRefresh := s.usageLastPoll.IsZero() ||
		time.Since(s.usageLastPoll) >= 5*time.Minute
	browser := site.New()

	if needRefresh {
		browser.ClickRefresh()
		time.Sleep(2 * time.Second)
	}

	result := browser.ReadUsage()

	if result == nil {
		return
	}

	s.usageMutex.Lock()
	s.usage = result
	s.usageLastPoll = time.Now()
	s.usageMutex.Unlock()
	s.store.SaveUsageSnapshot(
		result.SessionPercent,
		result.WeeklyAllPercent,
		result.WeeklySonnetPercent,
		result.CreditPercent,
	)
	s.store.TrimUsageSnapshots()
}
