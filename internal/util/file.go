package util

import (
	"io"
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

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(in *os.File) {
		_ = in.Close()
	}(in)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	return out.Close()
}
