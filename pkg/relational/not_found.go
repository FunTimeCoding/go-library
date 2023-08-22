package relational

import (
	"errors"
	"gorm.io/gorm"
)

func NotFound(e error) bool {
	return errors.Is(e, gorm.ErrRecordNotFound)
}
