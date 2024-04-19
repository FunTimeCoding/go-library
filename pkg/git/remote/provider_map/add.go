package provider_map

func (m *Map) Add(
	host string,
	provider string,
) {
	m.Known[host] = provider
}
