package generate2

import (
	"fmt"
	"reflect"
	"testing"

	"gocv.io/x/gocv"

	"github.com/garfeng/gocv_resource_tracker/generate/gocv_tracker"
)

func TestGenerateTrackerPackage(t *testing.T) {
	err := GenerateTrackerPackage(new(gocv_tracker.GoCVTracker), new(gocv_tracker.GoCVTrackerType), "../gocv_tracker", "../../")
	if err != nil {
		fmt.Println(err)
	}
}

func TestReadFuncsFromGoFile(t *testing.T) {
	tp := reflect.TypeOf(new(gocv_tracker.GoCVTracker))
	tp2 := reflect.TypeOf(new(gocv_tracker.GoCVTrackerType))
	funcList, types, _, prefix, err := ReadFuncsFromGoFile(tp, tp2, "../gocv_tracker/core.go")
	if err != nil {
		fmt.Println("Err:", err)
	}
	fmt.Println(len(funcList), len(types), prefix)

	for _, v := range types {
		fmt.Println(v)
	}

	/*

		fmt.Println(len(types), types)

		for _, v := range funcList {
			fmt.Println(v.Filename, v.FuncName, v.In)
		}
	*/
}

func TestGetFuncLineStrFromGoCVRoot(t *testing.T) {
	p, err := GetFuncLineStrFromGoCVRoot("Mat", "Region")
	fmt.Println(p, err)
}

func TestParseTypesMethod(t *testing.T) {
	m := gocv.NewMat()
	tp := reflect.TypeOf(m)

	res := ParseTypesMethod(tp, nil)
	for _, v := range res {
		fmt.Println(v)
		fmt.Println("")
	}
}
