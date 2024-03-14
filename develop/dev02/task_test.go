package main

import "testing"

func TestUnpack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{"a4bc2d5e", args{s: "a4bc2d5e"}, "aaaabccddddde", nil},
		{"abcd", args{s: "abcd"}, "abcd", nil},
		{"45", args{s: "45"}, "", ErrIncorrectString},
		{"empty", args{s: ""}, "", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unpack(tt.args.s)
			if err != tt.wantErr {
				t.Errorf("Unpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Unpack() = %v, want %v", got, tt.want)
			}
		})
	}
}
