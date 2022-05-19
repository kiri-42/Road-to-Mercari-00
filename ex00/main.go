package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"imgconv"
)

func main() {
	args := os.Args
	// 引数チェック
	if err := checkArg(args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	dir := args[1]
	// 画像ファイルのpathを取得
	paths, err := getPaths(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	// 画像変換処理
	for _, path := range paths {
		if err := imgconv.ConvertJpgToPng(path); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
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
		return errors.New("error: " + dir + ": no such dir")
	}
	return nil
}

func getPaths(dir string) ([]string, error) {
	var paths []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		paths = append(paths, path)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}
