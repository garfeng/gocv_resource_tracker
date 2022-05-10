package gocv_resource_tracker

import (
	"fmt"
	"image"
	"testing"

	"gocv.io/x/gocv"
)

func TestDoesPointsVectorAtNeedClose(t *testing.T) {
	pvs := gocv.NewPointsVectorFromPoints([][]image.Point{
		{
			{1, 1}, {2, 2},
			{3, 3},
		},
		{
			{4, 4}, {4, 5},
		},
	})
	/*
		c1 := pvs.At(0)
		c1.Close()
		c2 := pvs.At(1)
		c2.Close()
	*/
	// Result: do not need

	for i := 0; i < pvs.Size(); i++ {
		fmt.Println(pvs.At(i))
	}
}
