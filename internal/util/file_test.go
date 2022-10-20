package util

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsLink(t *testing.T) {
	// Create tmp dir
	d, err := os.MkdirTemp("", "gh_profile_test")
	if err != nil {
		t.Errorf("failed to create tmpdir: %v", err)
		return
	}
	defer func() {
		_ = os.RemoveAll(d)
	}()

	// Create regular file
	regularPath := filepath.Join(d, "regular")
	f, err := os.Create(regularPath)
	if err != nil {
		t.Errorf("failed to create regular file: %v", err)
		return
	}
	if err := f.Close(); err != nil {
		t.Errorf("failed to close regular file: %v", err)
		return
	}

	// Create symlink
	linkPath := filepath.Join(d, "link")
	if err := os.Symlink(regularPath, linkPath); err != nil {
		t.Errorf("failed to create symlink: %v", err)
		return
	}

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"symlink", args{linkPath}, true, false},
		{"regular", args{regularPath}, false, false},
		{"not found", args{filepath.Join(d, "a")}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsLink(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyFile(t *testing.T) {
	// Create tmp dir
	d, err := os.MkdirTemp("", "gh_profile_test")
	if err != nil {
		t.Errorf("failed to create tmpdir: %v", err)
		return
	}
	defer func() {
		_ = os.RemoveAll(d)
	}()

	srcPath := filepath.Join(d, "src")
	src, err := os.Create(srcPath)
	if err != nil {
		t.Errorf("failed to create src file: %v", err)
		return
	}
	if err := src.Close(); err != nil {
		t.Errorf("failed to close src file: %v", err)
		return
	}

	dstPath := filepath.Join(d, "dst")
	if err := CopyFile(srcPath, dstPath); err != nil {
		t.Errorf("CopyFile() got err: %v", err)
		return
	}

	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"simple", args{srcPath, dstPath}, false},
		{"file not exist", args{srcPath + "_", dstPath}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				_ = os.Remove(dstPath)
			}()

			if err := CopyFile(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if _, err := os.Stat(dstPath); err != nil {
				t.Errorf("dst file error: %v", err)
				return
			}
		})
	}
}
