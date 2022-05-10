package gocv_resource_tracker

import (
	"image"

	"github.com/garfeng/gocv_resource_tracker/tracker"
	"gocv.io/x/gocv"
)

type GoCVResourceTracker struct {
	*tracker.ResourceTracker
}

// NewAutoGCTracker returns a GoCVResourceTracker, you should call Close() by manual.
func NewAutoGCTracker() *GoCVResourceTracker {
	return &GoCVResourceTracker{
		ResourceTracker: tracker.NewAutoGCResourceTracker(),
	}
}

// NewTracker Returns a GoCVResourceTracker with runtime.SetFinalizer,
// Close() is not required, but you should be careful when deal with it.
func NewTracker() *GoCVResourceTracker {
	return &GoCVResourceTracker{
		ResourceTracker: tracker.NewResourceTracker(),
	}
}

func (g *GoCVResourceTracker) IMRead(name string, flag gocv.IMReadFlag) gocv.Mat {
	mat := gocv.IMRead(name, flag)
	g.TrackCloseError(&mat)
	return mat
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

func (g *GoCVResourceTracker) NewPointVectorFromPoints(pts []image.Point) gocv.PointVector {
	pv := gocv.NewPointVectorFromPoints(pts)
	g.TrackCloser(&pv)
	return pv
}

func (g *GoCVResourceTracker) NewPointVectorFromMat(mat gocv.Mat) gocv.PointVector {
	pv := gocv.NewPointVectorFromMat(mat)
	g.TrackCloser(&pv)
	return pv
}

func (g *GoCVResourceTracker) NewPointsVectorFromPoints(pts [][]image.Point) gocv.PointsVector {
	pvs := gocv.NewPointsVectorFromPoints(pts)
	g.TrackCloser(&pvs)
	return pvs
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
