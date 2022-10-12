package util

import (
	"errors"
	"os"
	"syscall"
)

func IsLink(path string) (bool, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	if fi.Mode()&os.ModeSymlink != 0 {
		return true, nil
	}

	s, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return false, errors.New("cannot convert stat value to syscall.Stat_t")
	}

	return s.Nlink > 1, nil
}
