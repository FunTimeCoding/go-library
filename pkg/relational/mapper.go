package relational

import "gorm.io/gorm"

func (d *Database) Mapper() *gorm.DB {
	return d.mapper
}
