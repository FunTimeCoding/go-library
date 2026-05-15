package model_context_tester

func (t *Tester) Close() {
	t.server.Close()
}
