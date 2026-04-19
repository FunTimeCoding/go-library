package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/raid"
	"gorm.io/gorm"
)

func New(
	mapper *gorm.DB,
	logCachePath string,
	eliteInsightsPath string,
) *Store {
	errors.PanicOnError(mapper.AutoMigrate(&raid.Raid{}))
	errors.PanicOnError(mapper.AutoMigrate(&raid.Fight{}))
	errors.PanicOnError(mapper.AutoMigrate(&raid.PlayerFightStat{}))
	s := &Store{
		mapper:            mapper,
		logCachePath:      logCachePath,
		eliteInsightsPath: eliteInsightsPath,
	}
	s.syncLogCache()

	return s
}
