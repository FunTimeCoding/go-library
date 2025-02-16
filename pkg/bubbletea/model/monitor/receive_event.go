package monitor

import (
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/receive"
	"log"
)

func (m *Model) receiveEvent(s receive.Message) {
	for _, line := range s.Text {
		log.Printf("receive event: %s\n", line)
	}
}
