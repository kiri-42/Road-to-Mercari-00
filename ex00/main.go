package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"convert"
)

func main() {
	args := os.Args
	// 引数チェック
	if err := checkArg(args); err != nil {
		println("flag2")
		println(err.Error())
		return
	}

	dir := args[1]
	// 画像ファイルのpathを取得
	paths, err := getPaths(dir)
	if err != nil {
		println(err.Error())
		return
	}

	// 画像変換処理
	for _, path := range paths {
		if err := convert.ConvertJpgToPng(path); err != nil {
			println(err.Error())
			return
		}
	}
}

func checkArg(args []string) error {
	if len(args) != 2 {
		return errors.New("error: invalid argument")
	}

	dir := args[1]
	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		return errors.New("error: " + dir + ": no such file")
	}
	return nil
}

func getPaths(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.New("error")
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			return nil, errors.New("error")
		}
		path := filepath.Join(dir, file.Name())
		if err := checkExtension(path); err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}
	return paths, nil
}

func checkExtension(path string) error {
	if strings.HasSuffix(path, "jpg") {
		return nil
	}
	return errors.New("error: " + path + " is not a valid file")
}
