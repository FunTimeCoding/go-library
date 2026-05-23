package label

type Label struct {
	Identifier        uint   `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string `gorm:"uniqueIndex:session_key;column:session_identifier"`
	Key               string `gorm:"uniqueIndex:session_key;column:key"`
	Value             string `gorm:"column:value"`
}
