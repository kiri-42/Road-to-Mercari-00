package imgconv

import (
	"errors"
	"strings"
)

// 引数のpathの拡張子が正しいか確認する関数
func CheckExt(path, ext string) error {
	switch ext {
	case ".jpeg", ".jpg", ".gif", ".png":
	default:
		return errors.New("error: " + ext + " is not a valid extension")
	}

	if !strings.HasSuffix(path, ext) {
		return errors.New("error: " + path + " is not a valid file")
	}

	return nil
}
