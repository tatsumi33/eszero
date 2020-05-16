package main

import (
	"os"
	"runtime"
	"testing"
)

func Test_codeStr(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	type testData struct {
		name        string
		args        []string
		wantS       string
		wantArgExst bool
	}

	tests := []testData{
		// empty args
		{name: `"eszero" & empty args`, args: []string{"eszero"}, wantS: "eszero", wantArgExst: false},
		{name: `full "eszero" & empty args`, args: []string{"/usr/eszero"}, wantS: "eszero", wantArgExst: false},
		{name: `"x" & empty args`, args: []string{"x"}, wantS: "x", wantArgExst: true},
		{name: `"a_1" & empty args`, args: []string{"a_1"}, wantS: "1", wantArgExst: true},
		{name: `"a_1.exe" & empty args`, args: []string{"a_1.exe"}, wantS: "1", wantArgExst: true},
		{name: `"ab_0x1" & empty args`, args: []string{"ab_0x1"}, wantS: "0x1", wantArgExst: true},
		{name: `"ab_0x1.exe" & empty args`, args: []string{"ab_0x1.exe"}, wantS: "0x1", wantArgExst: true},
		{name: `"0_-0b1" & empty args`, args: []string{"0_-0b1"}, wantS: "-0b1", wantArgExst: true},
		{name: `"0_-0b1.exe" & empty args`, args: []string{"0_-0b1.exe"}, wantS: "-0b1", wantArgExst: true},
		{name: `"a__" & empty args`, args: []string{"a__"}, wantS: "", wantArgExst: true},
		{name: `"a__.exe" & empty args`, args: []string{"a__.exe"}, wantS: "", wantArgExst: true},
		{name: `"_" & empty args`, args: []string{"_"}, wantS: "", wantArgExst: true},
		{name: `"_.exe" & empty args`, args: []string{"_.exe"}, wantS: "", wantArgExst: true},
		{name: `"__" & empty args`, args: []string{"__"}, wantS: "", wantArgExst: true},
		{name: `"__.exe" & empty args`, args: []string{"__.exe"}, wantS: "", wantArgExst: true},
		{name: `"_-" & empty args`, args: []string{"_-"}, wantS: "-", wantArgExst: true},
		{name: `"_-.exe" & empty args`, args: []string{"_-.exe"}, wantS: "-", wantArgExst: true},
	}

	if runtime.GOOS == "windows" {
		tests = append(tests,
			testData{name: `"eszero.exe" & empty args`, args: []string{"eszero.exe"}, wantS: "eszero", wantArgExst: false})
		tests = append(tests,
			testData{name: `full "eszero.exe" & empty args`, args: []string{`C:\Users\eszero.exe`}, wantS: "eszero", wantArgExst: false})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			gotS, gotArgExst := codeStr()
			if gotS != tt.wantS {
				t.Errorf("codeStr() s = %v, want %v", gotS, tt.wantS)
			}
			if gotArgExst != tt.wantArgExst {
				t.Errorf("codeStr() argexst = %v, want %v", gotArgExst, tt.wantArgExst)
			}
		})
	}
}
