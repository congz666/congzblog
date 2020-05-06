package service

import (
	"github.com/congz666/congzblog/dao/db"
)

//JudgeLogin ...
func JudgeLogin(named string) (password string) {
	password = db.GetPassword(named)
	return
}
