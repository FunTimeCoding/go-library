package relational

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var skipNotFoundLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	},
)

func FirstSkipNotFound(
	m *gorm.DB,
	result any,
	condition any,
) any {
	past := m.Logger
	m.Logger = skipNotFoundLogger

	if e := m.First(result, condition).Error; e != nil {
		if NotFound(e) {
			m.Logger = past

			return result
		} else {
			m.Logger = past

			panic(e)
		}
	}

	m.Logger = past

	return result
}
