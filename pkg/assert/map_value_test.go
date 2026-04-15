package assert

import "testing"

func TestMapValue(t *testing.T) {
	m := map[string]any{"height": 180, "name": "Fern"}
	MapValue(t, 180, m, "height")
	MapValue(t, "Fern", m, "name")
}
