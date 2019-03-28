package core

import (
	"os"
	"regexp"
	"runtime"
)

func (rStore *RegExpStore) RegExpValidation(source string) bool {
	var r *regexp.Regexp

	var err error
	if r, err = regexp.Compile(string(*rStore)); err != nil {
		panic(err)
	}
	return r.Match([]byte(source))
}

func FileValidation(filePath string) bool {
	fi, err := os.Stat(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		panic(err.Error())
	} else {
		switch mode := fi.Mode(); {
		case mode.IsDir():
			return false
		}
	}

	return true
}

func IsItRunInPosix() bool {
	return runtime.GOOS == Windows
}
