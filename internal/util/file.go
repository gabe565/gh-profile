package util

import (
	"os"
)

func IsLink(path string) (bool, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	if fi.Mode()&os.ModeSymlink != 0 {
		return true, nil
	}

	return false, nil
}
