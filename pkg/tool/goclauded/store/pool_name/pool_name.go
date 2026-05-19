package pool_name

type PoolName struct {
	Identifier uint   `gorm:"primaryKey;column:identifier"`
	Name       string `gorm:"column:name;uniqueIndex"`
}
