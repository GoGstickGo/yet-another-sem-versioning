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
		{"good", args{commit: "fsdfds #major"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkCommitMsg(tt.args.commit); (err != nil) != tt.wantErr {
				t.Errorf("checkCommitMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_modifyVersionFile(t *testing.T) {
	type args struct {
		nV string
		f  string
		vf bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"#1", args{nV: "1.0.0", f: "VERISON", vf: true}, true},
		{"#2", args{nV: "1.0.0", f: "testFiles/VERSION", vf: true}, false},
		{"#3", args{nV: "2.0.0", f: "testFiles/VERSION_2", vf: false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := modifyVersionFile(tt.args.nV, tt.args.f, tt.args.vf); (err != nil) != tt.wantErr {
				t.Errorf("modifyVersionFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createVersionFile(t *testing.T) {
	type args struct {
		f string
		m string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"#1", args{f: "testFiles/NEW", m: "sdfdsfdsfdsf  #major adsfdsfds"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createVersionFile(tt.args.f, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("createVersionFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_buildVersion(t *testing.T) {
	type args struct {
		cV string
		m  string
	}
	tests := []struct {
		name   string
		args   args
		wantNV string
	}{
		{"major", args{cV: "1.0.0", m: "#major"}, "2.0.0"},
		{"minor", args{cV: "8.2.3", m: "#minor"}, "8.3.3"},
		{"patch", args{cV: "12.2.1", m: "#patch"}, "12.2.2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNV := buildVersion(tt.args.cV, tt.args.m); gotNV != tt.wantNV {
				t.Errorf("buildVersion() = %v, want %v", gotNV, tt.wantNV)
			}
		})
	}
}
