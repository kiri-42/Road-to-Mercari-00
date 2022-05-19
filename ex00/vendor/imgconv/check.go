package imgconv

import (
	"errors"
	// "net/http"
	// "os"
	"strings"
)

// func checkFileContentType(path string) (bool, error) {
// 	f, err := os.Open(path)
// 	if err != nil {
// 		return false, err
// 	}
// 	defer f.Close()

// 	// 最初の512bytesを読む
// 	buffer := make([]byte, 512)
// 	size, err := f.Read(buffer)
// 	if size == 0 {
// 		return false, nil
// 	}
// 	if err != nil {
// 		return false, err
// 	}

// 	// 有効なMIMEタイプでない場合は「application/octet-stream」が返る
// 	if http.DetectContentType(buffer) == "application/octet-stream" {
// 		return false, nil
// 	}

// 	return true, nil
// }

// 引数の拡張子がjpgか確認する関数
func CheckJpg(path string) error {
	if !strings.HasSuffix(path, "jpg") {
		return errors.New("error: " + path + " is not a valid file")
	}

	// ok, err := checkFileContentType(path)
	// if err != nil {
	// 	return err
	// }
	// if !ok {
	// 	return errors.New("error: " + path + " is not a valid file")
	// }

	return nil
}
