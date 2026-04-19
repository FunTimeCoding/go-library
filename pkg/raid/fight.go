package raid

import "time"

type Fight struct {
	Filename     string    `gorm:"primaryKey" json:"filename"`
	Timestamp    time.Time `json:"timestamp"`
	DurationMS   int       `json:"duration_ms"`
	MapID        int       `json:"map_id"`
	MapName      string    `json:"map_name"`
	AlliedCount  int       `json:"allied_count"`
	AlliedTeamID int       `json:"allied_team_id"`
	EnemyCount   int       `json:"enemy_count"`
	EnemyTeams   *string   `json:"enemy_teams,omitempty"`
	Success      bool      `json:"success"`
	Enriched     bool      `gorm:"not null;default:false" json:"enriched"`
}
