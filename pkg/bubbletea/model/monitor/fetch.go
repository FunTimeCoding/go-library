package monitor

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"log"
)

func fetch() tea.Cmd {
	return func() tea.Msg {
		exists := run.CommandExists(constant.GoSensor)

		if false {
			log.Printf("%s exists: %v\n", constant.GoSensor, exists)
		}

		result := FetchMessage{}

		if exists {
			items := monitor.Run(constant.GoSensor)
			result.Items = items

			if false {
				for _, element := range items {
					log.Printf("Item: %+v\n", element)
				}
			}
		}

		return result
	}
}
