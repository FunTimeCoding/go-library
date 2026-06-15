package event_metadata

type EventMetadata struct {
	EventIdentifier uint   `gorm:"primaryKey;column:event_identifier"`
	Key             string `gorm:"primaryKey;column:key"`
	Value           string `gorm:"column:value"`
}
