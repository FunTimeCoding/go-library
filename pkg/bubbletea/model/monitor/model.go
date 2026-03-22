package monitor

import (
	"charm.land/bubbles/v2/table"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/toast"
	"github.com/funtimecoding/go-library/pkg/monitor/gorilla/client"
	"github.com/funtimecoding/go-library/pkg/monitor/item"
	"time"
)

type Model struct {
	width          int
	height         int
	auto           bool
	topBar         string
	table          *table.Model
	items          []*item.Item
	bottomBar      string
	second         int
	client         *client.Client
	connect        bool
	user           string
	hostname       string
	toast          []*toast.Toast
	nextToast      int
	initialResized bool
	modal          *Modal
	lastFetch      time.Time
}
