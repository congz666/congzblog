package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	//Cfg ...
	Cfg *ini.File
	//RunMode ...
	RunMode string
	//HTTPPort ...
	HTTPPort int
	//ReadTimeout ...
	ReadTimeout time.Duration
	//WriteTimeout ...
	WriteTimeout time.Duration
	//PageSize ...
	PageSize int
	//JwtSecret ...
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
}

//LoadBase ...
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

//LoadServer ...
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}