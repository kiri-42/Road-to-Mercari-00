/*
画像変換関連の関数をまとめたパッケージです
*/
package imgconv

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	// "strings"
	// "path/filepath"
)

type filePath struct {
	in string
	out string
}

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
	// fp.png = convertExtension(fp.jpg, "jpg", "png")
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

	// 画像オブジェクトをpng形式にエンコード
	err = png.Encode(f2, img)
	if err != nil {
		return err
	}

	return nil
}

// jpgファイルをpngファイルに変換する関数
// func ConvertJpgToPng(jpg string) (err error) {
// 	var fp filePath
// 	fp.jpg = jpg

// 	if err := CheckJpg(fp.jpg); err != nil {
// 		return err
// 	}

// 	// 変換元ファイルを開く
// 	f, err := os.Open(fp.jpg)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	// 変換元ファイルを画像オブジェクトに変換
// 	img, _, err := image.Decode(f)
// 	if err != nil {
// 		return err
// 	}

// 	// 変換先のファイルパスを取得
// 	// fp.png = convertExtension(fp.jpg, "jpg", "png")
// 	fp.png = replaceExt(fp.jpg, ".jpg", ".png")

// 	// 変換先のファイルを作成
// 	f2, err := os.Create(fp.png)
// 	if err != nil {
// 		return err
// 	}
// 	defer func () {
// 		if e := f2.Close(); e != nil {
// 			err = e
// 		}
// 	}()

// 	// 画像オブジェクトをpng形式にエンコード
// 	err = png.Encode(f2, img)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func replaceExt(path, iExt, oExt string) string {
	return path[:len(path)-len(iExt)] + oExt
}

// func convertExtension(path, fromExt, toExt string) string {
// 	r := strings.Replace(reverse(path), reverse(fromExt), reverse(toExt), 1)
// 	return reverse(r)
// }
