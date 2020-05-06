package service

import (
	"io/ioutil"
	"os"

	"github.com/congz666/congzblog/utils/logging"
)

//ReadAll ...
//读取单个文件
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
