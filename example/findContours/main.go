package main

import (
	"fmt"
	"image"

	tracker "github.com/garfeng/gocv_resource_tracker"
)

func main() {
	tr := tracker.NewTracker()
	defer tr.Close()

	contours := tr.NewPointsVectorFromPoints([][]image.Point{
		{{1, 1}, {1, 2}, {1, 3}, {1, 1}},
		{{5, 5}, {5, 2}, {5, 3}, {5, 5}},
		{{5, 5}, {5, 2}, {5, 3}, {5, 5}},
		{{5, 5}, {5, 2}, {5, 3}, {5, 1}},
	})
	number := contours.Size()
	for i := 0; i < number; i++ {
		c := contours.At(i)
		fmt.Println(c.Size())
		//c.Close()
	}
	//gocv.NewMat().Close()
	//contours.Close()

}

type ABC struct {
	P P2
}

type P2 struct {
	ID int
}
