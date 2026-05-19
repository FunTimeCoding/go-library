package alias

type Alias struct {
	SessionIdentifier string `gorm:"primaryKey;column:session_identifier"`
	Name              string `gorm:"column:name"`
	Description       string `gorm:"column:description"`
}
