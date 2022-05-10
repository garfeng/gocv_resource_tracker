package gocv_resource_tracker

import (
	"fmt"
	"image"
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

func Example() {
	rt := NewTracker()
	defer rt.Close() // only one time close here

	mat := rt.NewMat()
	// ...
	// do something with mat
	kernel := rt.GetStructuringElement(gocv.MorphRect, image.Pt(10, 10))
	dst := rt.NewMat()

	gocv.Erode(mat, &dst, kernel)
}
