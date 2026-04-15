package assert

import "testing"

func TestMapNotHasKey(t *testing.T) {
	m := map[string]any{"height": 180}
	MapNotHasKey(t, m, "weight")
}
