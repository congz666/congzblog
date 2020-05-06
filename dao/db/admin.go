package db

//GetPassword ...
func GetPassword(named string) (password string) {
	sqlstr := "select password from administrator where name=?"
	_ = DB.Get(&password, sqlstr, named)

	return
}
