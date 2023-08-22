package relational

func (d *Database) Delete(a any) {
	d.DeleteWhere(a, "TRUE")
}
