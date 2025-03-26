package pair

func (p *Pair) Match(
	k string,
	v string,
) bool {
	return p.Key == k && p.Value == v
}
