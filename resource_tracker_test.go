package gocv_resource_tracker

import (
	"fmt"
	"image"
	"runtime"
	"testing"
	"time"

	"gocv.io/x/gocv"
)

func CreateOneTrackerAndMat() Mat {
	rt := NewTracker()
	mat := rt.NewMatWithSize(100, 100, gocv.MatTypeCV8UC1)

	fmt.Println("create", mat.Rows(), mat.Cols())

	return mat
}

func TestNewTracker(t *testing.T) {
	matList := []Mat{}
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

	rt.Erode(mat, &dst, kernel)
}

func CreateThenAutoGC() {
	rt := NewTracker()

	for i := 0; i < 15; i++ {
		rt.NewMatWithSize(5120, 1024, gocv.MatTypeCV8UC1)
	}
}

func TestAutoGC(t *testing.T) {
	for i := 0; i < 100; i++ {
		CreateThenAutoGC()
		<-time.After(time.Second * 5)
	}
}

func TestPanic(t *testing.T) {
	/*
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()

	*/
	f := NewTracker()

	//mat := f.NewMatWithSize(100, 100, gocv.MatTypeCV8UC1)

	//mat.Region(image.Rect(-1, -1, 100, 100))

	//region := f.CloneRegion(mat, image.Rect(0, 0, 200, 200))

	//fmt.Println(region.Cols())

	//f.GetStructuringElement(gocv.MorphRect, image.Pt(0, 0))
	f.NewMatFromBytes(100, 100, gocv.MatTypeCV8UC1, make([]byte, 10))

}
