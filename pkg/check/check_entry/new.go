package check_entry

import (
	"log"
	"slices"
	"time"
)

func New(
	level string,
	text string,
) *Entry {
	if !slices.Contains(Levels, level) {
		log.Panicf("unexpected level: %s", level)
	}

	return &Entry{Time: time.Now(), Level: level, Text: text}
}
