package gocv_resource_tracker

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"

	"gocv.io/x/gocv"
)

func Caller(skip int) (name string, line int) {
	var file string
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return "", 0
	}

	root, name := filepath.Split(file)

	base := filepath.Base(root)

	return base + "/" + name, line
}

func PanicWithCaller(skip int, e interface{}) {
	file, line := Caller(skip + 1)

	msg := fmt.Sprint("[", file, ":", line, "] ", e)
	panic(msg)
}

func MustEqual(a interface{}, b interface{}, format string, i ...interface{}) {
	if reflect.DeepEqual(a, b) {
		return
	}

	if format == "" {
		PanicWithCaller(1, fmt.Sprintf("%v != %v", a, b))
	} else {
		s1 := fmt.Sprintf(format, i...)
		//	s2 := fmt.Sprintf("(%v != %v)", a, b)
		PanicWithCaller(1, s1)
	}
}

func Must(a bool, format string, i ...interface{}) {
	if a {
		return
	}
	PanicWithCaller(1, fmt.Sprintf("("+format+") != true", i...))
}

func MustNot(a bool, format string, i ...interface{}) {
	if !a {
		return
	}
	PanicWithCaller(1, fmt.Sprintf(format+" != false", i...))
}

var (
	matTypeSizeMap = map[gocv.MatType]int{
		gocv.MatTypeCV8U:  1,
		gocv.MatTypeCV8S:  1,
		gocv.MatTypeCV16U: 2,
		gocv.MatTypeCV16S: 2,
		gocv.MatTypeCV32S: 4,
		gocv.MatTypeCV32F: 4,
		gocv.MatTypeCV64F: 8,
	}
)

func sizeOfMatType(mt gocv.MatType) int {
	sz := matTypeSizeMap[mt]
	channelNum := (int(mt) >> 3) + 1
	return sz * channelNum
}
