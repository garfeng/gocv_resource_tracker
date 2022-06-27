package main

import (
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"

	tracker "github.com/garfeng/gocv_resource_tracker"
)

func main() {
	rt := tracker.NewTracker()
	defer rt.Close()

	mat := rt.NewMatWithSize(1000, 1000, gocv.MatTypeCV8UC1)
	for i := 0; i < 10; i++ {
		rt.Rectangle(mat, image.Rect(i*100, i*100, i*100+50, i*100+50),
			color.RGBA{255, 255, 255, 255}, -1)
	}

	kernel := rt.GetStructuringElement(gocv.MorphEllipse, image.Pt(5, 5))
	erode := rt.NewMat()
	rt.Erode(mat, erode, kernel)
	rt.IMWrite("hello.png", erode)

	thres := rt.NewMat()
	rt.Threshold(erode, thres, 0, 255, gocv.ThresholdOtsu)

	contours := rt.FindContours(erode, gocv.RetrievalList, gocv.ChainApproxNone)

	number := contours.Size()
	for i := 0; i < number; i++ {
		c := contours.At(i)
		fmt.Println("contourSize :", c.Size())
		//c.Close()
	}
	fmt.Println("contourNumber: ", contours.Size())

	v := rt.NewPointVectorFromPoints([]image.Point{{1, 2},
		{5, 5}, {1, 2}})
	contours.Append(v)
	fmt.Println("contourNumber2: ", contours.Size())

	c := contours.At(4)
	fmt.Println(c.Size())

}

type ABC struct {
	P P2
}

type P2 struct {
	ID int
}
