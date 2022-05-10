package gocv_resource_tracker

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"gocv.io/x/gocv"
)

func CreateOneTrackerAndMat() gocv.Mat {
	rt := NewTracker()
	mat := rt.NewMatWithSize(100, 100, gocv.MatTypeCV8UC1)

	fmt.Println("create", mat.Rows(), mat.Cols())

	return mat
}

func TestNewTracker(t *testing.T) {
	matList := []gocv.Mat{}
	for i := 0; i < 100; i++ {
		mat := CreateOneTrackerAndMat()
		matList = append(matList, mat)
	}
	runtime.GC()
	<-time.After(time.Second)
}
