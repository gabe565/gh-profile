package util

import "testing"

func TestUpperFirst(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simple", args{"test string"}, "Test string"},
		{"single rune", args{"a"}, "A"},
		{"empty", args{""}, ""},
		{"multi byte", args{"ñ"}, "Ñ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperFirst(tt.args.s); got != tt.want {
				t.Errorf("UpperFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
