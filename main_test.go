package main

import (
	"testing"
)

func Test_checkCommitMsg(t *testing.T) {
	type args struct {
		commit string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"bad", args{commit: "sometghing"}, true},
		{"good", args{commit: "fsdfds major"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkCommitMsg(tt.args.commit); (err != nil) != tt.wantErr {
				t.Errorf("checkCommitMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_modifyVersion(t *testing.T) {
	t1 := []string{"1", "0", "0"}
	type args struct {
		a int
		b int
		c int
		s []string
	}
	tests := []struct {
		name         string
		args         args
		wantMVersion string
	}{
		{"major", args{a: 0, b: 1, c: 2, s: t1}, "2.0.0"},
		{"minor", args{a: 1, b: 0, c: 2, s: t1}, "1.1.0"},
		{"patch", args{a: 2, b: 0, c: 1, s: t1}, "1.0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMVersion := modifyVersion(tt.args.a, tt.args.b, tt.args.c, tt.args.s); gotMVersion != tt.wantMVersion {
				t.Errorf("modifyVersion() = %v, want %v", gotMVersion, tt.wantMVersion)
			}
		})
	}
}

func Test_buildNewVersion(t *testing.T) {
	type args struct {
		cVersion string
		change   string
	}
	tests := []struct {
		name         string
		args         args
		wantNVersion string
	}{
		{"good", args{cVersion: "2.0.0", change: "major"}, "3.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNVersion := buildNewVersion(tt.args.cVersion, tt.args.change); gotNVersion != tt.wantNVersion {
				t.Errorf("buildNewVersion() = %v, want %v", gotNVersion, tt.wantNVersion)
			}
		})
	}
}
