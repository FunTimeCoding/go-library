package assert

import "testing"

func TestMapHasKey(t *testing.T) {
	m := map[string]any{"height": 180}
	MapHasKey(t, m, "height")
}
