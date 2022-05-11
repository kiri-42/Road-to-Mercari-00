package img

import (
	"image"
	"image/png"
	"image/jpeg"
	"os"
	"strings"
)

func ConvertJpgToPng(path string) error {
	// 変換元ファイルを開く
	f, err := os.Open(path)
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
	pngPath := convertExtension(path, "jpg", "png")

	// 変換先のファイルを作成
	f2, err := os.Create(pngPath)
	if err != nil {
		return err
	}
	defer f2.Close()

	if pngPath == "jpg" {
		jpeg.Encode(f2, img, &jpeg.Options{})
	}

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
