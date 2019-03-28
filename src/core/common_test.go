package core

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestIsItRunInPosix(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Should expose valid operation system, I love Windows :)", runtime.GOOS == Windows},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsItRunInPosix(); got != tt.want {
				t.Errorf("IsItRunInPosix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileValidation(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	tmpfile, err := ioutil.TempFile(os.TempDir(), "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	fullPath, err := filepath.Abs(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Should find an good existed file", args{fullPath}, true},
		{"Should find a relative existed file", args{tmpfile.Name()}, true},
		{"Should not find any file by a directory path", args{filepath.Dir(fullPath)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileValidation(tt.args.filePath); got != tt.want {
				t.Errorf("FileValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegExpStore_RegExpValidation(t *testing.T) {
	type args struct {
		source string
	}
	storePosix := PosixPathRegex
	storeWin := WinPathRegex

	tests := []struct {
		name   string
		rStore *RegExpStore
		args   args
		want   bool
	}{
		{"Passing a valid absolute linux path should return true", &storePosix, args{"/home/user/Documents/foo.log"}, true},
		{"Passing a valid relative linux path should return true", &storePosix, args{"./foo.log"}, true},
		{"Passing an incorrect linux path should return false", &storePosix, args{"this is just a string /home/user/Documents/foo.log"}, false},
		{"Passing a valid absolute windows path should return true", &storeWin, args{`C:\Soft\logstash-6.4.1\Gemfile.lock`}, true},
		{"Passing an incorrect windows path should return false", &storeWin, args{`this is just a string C:\Soft\logstash-6.4.1\Gemfile.lock`}, false},
		{"Passing a valid relative windows path should return true", &storeWin, args{`.\Gemfile.lock`}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rStore.RegExpValidation(tt.args.source); got != tt.want {
				t.Errorf("RegExpStore.RegExpValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}
