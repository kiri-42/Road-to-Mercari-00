package imgconv

import (
	"errors"
	"strings"
)

// 引数の拡張子がjpgか確認する関数
func CheckExt(path, ext string) error {
	if !strings.HasSuffix(path, ext) {
		return errors.New("error: " + path + " is not a valid file")
	}

	return nil
}
