package testutils

import (
	"reflect"
	"regexp"
	"runtime"
)

var nameRe = regexp.MustCompile("[^.]*$")

func FuncName(i interface{}) string {
	full := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return nameRe.FindString(full)
}
