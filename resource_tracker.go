package gocv_resource_tracker

import (
	"image"

	"github.com/garfeng/gocv_resource_tracker/tracker"
	"gocv.io/x/gocv"
)

type GoCVResourceTracker struct {
	*tracker.ResourceTracker
}

func NewTracker() *GoCVResourceTracker {
	return &GoCVResourceTracker{
		ResourceTracker: tracker.NewResourceTracker(),
	}
}

func (g *GoCVResourceTracker) NewMat() gocv.Mat {
	mat := gocv.NewMat()
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) NewMatWithSize(rows, cols int, matType gocv.MatType) gocv.Mat {
	mat := gocv.NewMatWithSize(rows, cols, matType)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) NewPointVectorFromPoints(points []image.Point) gocv.PointVector {
	pv := gocv.NewPointVectorFromPoints(points)
	g.TrackCloser(&pv)
	return pv
}

func (g *GoCVResourceTracker) NewPointVectorFromMat(mat gocv.Mat) gocv.PointVector {
	pv := gocv.NewPointVectorFromMat(mat)
	g.TrackCloser(&pv)
	return pv
}

func (g *GoCVResourceTracker) FindContours(src gocv.Mat, mode gocv.RetrievalMode, method gocv.ContourApproximationMode) gocv.PointsVector {
	pvs := gocv.FindContours(src, mode, method)
	g.TrackCloser(&pvs)
	return pvs
}

func (g *GoCVResourceTracker) Clone(src gocv.Mat) gocv.Mat {
	dst := src.Clone()
	g.TrackCloseError(&dst)
	return dst
}

func (g *GoCVResourceTracker) Region(src gocv.Mat, rect image.Rectangle) gocv.Mat {
	region := src.Region(rect)
	g.TrackCloseError(&region)
	return region
}

func (g *GoCVResourceTracker) GetStructuringElement(morphShape gocv.MorphShape, size image.Point) gocv.Mat {
	kernel := gocv.GetStructuringElement(morphShape, size)
	g.TrackCloseError(&kernel)
	return kernel
}
