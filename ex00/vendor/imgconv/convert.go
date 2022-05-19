/*
画像変換関連の関数をまとめたパッケージです
*/
package imgconv

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"os"
)

type filePath struct {
	in string
	out string
}

// 画像の拡張子を変換する関数
func Convert(path, iExt, oExt string) (err error) {
	var fp filePath
	fp.in = path

	if err := CheckExt(fp.in, iExt); err != nil {
		return err
	}

	// 変換元ファイルを開く
	f, err := os.Open(fp.in)
	if err != nil {
		return err
	}
	defer f.Close()

	// 変換元ファイルを画像オブジェクトに変換
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	// 変換先のファイルパスを取得
	fp.out = replaceExt(fp.in, iExt, oExt)

	// 変換先のファイルを作成
	f2, err := os.Create(fp.out)
	if err != nil {
		return err
	}
	defer func () {
		if e := f2.Close(); e != nil {
			err = e
		}
	}()

	// 画像オブジェクトをエンコード
	switch oExt {
	case ".jpeg", ".jpg":
		err = jpeg.Encode(f2, img, nil)
		if err != nil {
			return err
		}
	case ".gif":
		err = gif.Encode(f2, img, nil)
		if err != nil {
			return err
		}
	case ".png":
		err = png.Encode(f2, img)
		if err != nil {
			return err
		}
	default:
		return errors.New("error: " + oExt + "is not a valid extension")
	}

	return nil
}

func replaceExt(path, iExt, oExt string) string {
	return path[:len(path)-len(iExt)] + oExt
}
