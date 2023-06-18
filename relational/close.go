package relational

func (d *Database) Close() {
	d.client.Close()
}
