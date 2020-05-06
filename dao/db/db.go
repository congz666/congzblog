package db

import (
	"crypto/rand"
	"fmt"

	"github.com/congz666/congzblog/pkg/setting"
	"github.com/congz666/congzblog/utils/logging"
	"github.com/jmoiron/sqlx"
)

//DB ...
var DB *sqlx.DB

//Init ...
func Init() (err error) {
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		logging.Error(err, "Fail to get section 'database'")
	}
	//获取数据库配置
	dbType := sec.Key("TYPE").String()
	dbName := sec.Key("NAME").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()
	DB, err = sqlx.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True",
		user,
		password,
		host,
		dbName))
	if err != nil {
		logging.Error(err, "open database failed")
		return
	}
	err = DB.Ping()
	if err != nil {
		logging.Error(err, "database didn't work")
		return
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		logging.Error(err, "Cannot generate UUID")
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
