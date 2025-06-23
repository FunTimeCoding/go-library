package protocol

func (p *Protocol) Close() {
	p.client.Close()
}
