package model

//Admin ...
//管理员
type Admin struct {
	Name     string `db:"name"`
	Password string `db:"password"`
}
