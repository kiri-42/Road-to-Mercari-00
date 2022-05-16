package imgconv

import (
	"errors"
	"strings"
)

// 引数の拡張子がpngか確認する関数
func CheckPng(path string) error {
	if strings.HasSuffix(path, "jpg") {
		return nil
	}
	return errors.New("error: " + path + " is not a valid file")
}