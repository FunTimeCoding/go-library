package provider_map

func New() *Map {
	return &Map{Known: make(map[string]string)}
}
