package utils

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/congz666/congzblog/dao/db"
	"github.com/gin-gonic/gin"
)

//GetRootDir ...
func GetRootDir() (rootPath string) {
	exePath := os.Args[0]
	rootPath = filepath.Dir(exePath)
	return rootPath
}

//Session ...
func Session(c *gin.Context) (err error) {
	cookie, err := c.Request.Cookie("_cookie")
	if err == nil {
		if ok, _ := db.Check(cookie.Value); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

//Encrypt ...
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
