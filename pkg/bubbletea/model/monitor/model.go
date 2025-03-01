package monitor

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/client"
	"github.com/funtimecoding/go-library/pkg/monitor/item"
)

type Model struct {
	table     *table.Model
	width     int
	height    int
	topBar    string
	bottomBar string
	second    int

	client   *client.Client
	connect  bool
	user     string
	hostname string

	items []*item.Item
}
