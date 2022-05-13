package imgconv

import (
	"errors"
	"strings"
)

func CheckPng(path string) error {
	if strings.HasSuffix(path, "jpg") {
		return nil
	}
	return errors.New("error: " + path + " is not a valid file")
}
