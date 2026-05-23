package usage_snapshot

import "time"

func New(
	sessionPercent int,
	weeklyPercent int,
	sonnetPercent int,
	creditPercent int,
	createdAt time.Time,
) *Snapshot {
	return &Snapshot{
		SessionPercent: sessionPercent,
		WeeklyPercent:  weeklyPercent,
		SonnetPercent:  sonnetPercent,
		CreditPercent:  creditPercent,
		CreatedAt:      createdAt,
	}
}
