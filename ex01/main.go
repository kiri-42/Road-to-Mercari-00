package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"flag"

	"imgconv"
)

func main() {
	var (
		ext = flag.String("e", ".jpeg", "Specify the extension")
	)
	flag.Parse()

	args := flag.Args()
	// 引数チェック
	if err := checkArg(args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	dir := args[0]
	// 画像ファイルのpathを取得
	paths, err := getPaths(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	// 画像変換処理
	for _, path := range paths {
		if err := imgconv.Convert(path, filepath.Ext(path), *ext); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}

func checkArg(args []string) error {
	if len(args) != 1 {
		return errors.New("error: invalid argument")
	}

	dir := args[0]
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
