package relational

func (d *Database) validateInputs(
	user string,
	password string,
	dbname string,
) {
	d.validateUser(user)
	d.validatePassword(password)
	d.validateDatabase(dbname)
}
