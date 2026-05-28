package log

import "time"

func New(
	name string,
	colorIndex int,
	maxNameWidth int,
) *Logger {
	l := &Logger{
		colorIndex:   colorIndex % len(colors),
		name:         name,
		writes:       make(chan []byte),
		done:         make(chan struct{}),
		timeout:      2 * time.Millisecond,
		maxNameWidth: maxNameWidth,
	}
	go l.writeLines()

	return l
}
