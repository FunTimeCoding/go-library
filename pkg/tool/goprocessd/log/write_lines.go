package log

import (
	"bytes"
	"time"
)

func (l *Logger) writeLines() {
	var tick <-chan time.Time

	for {
		select {
		case w, open := <-l.writes:
			if !open {
				if len(l.buffers) > 0 {
					l.writeBuffers([]byte("\n"))
				}

				return
			}

			buffer := bytes.NewBuffer(w)

			for {
				line, e := buffer.ReadBytes('\n')

				if len(line) > 0 {
					if line[len(line)-1] == '\n' {
						if len(line) != 1 || len(l.buffers) > 0 {
							l.writeBuffers(line)
						}

						tick = nil
					} else {
						l.buffers = append(l.buffers, line)
						tick = time.After(l.timeout)
					}
				}

				if e != nil {
					break
				}
			}

			l.done <- struct{}{}
		case <-tick:
			if len(l.buffers) > 0 {
				l.writeBuffers([]byte("\n"))
			}

			tick = nil
		}
	}
}
