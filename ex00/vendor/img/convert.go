// 画像処理関連のパッケージ
package img

import (
	"image"
	"image/png"
	_ "image/jpeg"
	"os"
	"strings"
)

type filePath struct {
	jpg string
	png string
}

func ConvertJpgToPng(jpg string) error {
	fp := new(filePath)
	fp.jpg = jpg

	// 変換元ファイルを開く
	f, err := os.Open(fp.jpg)
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
	fp.png = convertExtension(fp.jpg, "jpg", "png")

	// 変換先のファイルを作成
	f2, err := os.Create(fp.png)
	if err != nil {
		return err
	}
	defer f2.Close()

	// 画像オブジェクトをpng形式にエンコード
	err = png.Encode(f2, img)
	if err != nil {
		return err
	}

	return nil
}

func convertExtension(path, fromExt, toExt string) string {
	r := strings.Replace(reverse(path), reverse(fromExt), reverse(toExt), 1)
	return reverse(r)
}
