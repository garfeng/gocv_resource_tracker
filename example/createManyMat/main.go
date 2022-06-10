package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	tracker "github.com/garfeng/gocv_resource_tracker"
	"gocv.io/x/gocv"
)

func createOnce() {
	tr := tracker.NewTracker()
	defer tr.Close()

	mat := tr.NewMatWithSize(5000, 5000, gocv.MatTypeCV64FC3)

	tr.Rectangle(&mat, image.Rect(40, 40, 60, 60), color.RGBA{
		R: 100,
		G: 100,
		B: 100,
		A: 255,
	}, 5)
	tr.IMWrite("hello.png", mat)

	mat.DivideFloat(1)
}

func main() {
	N := 100
	for i := 0; i < N; i++ {
		createOnce()
		fmt.Println(fmt.Sprintf("%d/%d", i, N))
	}

	for i := 0; i < 30; i++ {
		<-time.After(time.Second)
		fmt.Println(30 - i)
	}
}
