package util

import (
	"os"
)

func IsLink(path string) (bool, error) {
	f, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	if f.Mode()&os.ModeSymlink != 0 {
		return true, nil
	}

	return false, nil
}
