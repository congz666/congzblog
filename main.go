package main

import (
	"fmt"
	"net/http"

	"github.com/congz666/congzblog/dao/db"
	"github.com/congz666/congzblog/pkg/setting"
	"github.com/congz666/congzblog/routers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
	router := routers.Init()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
