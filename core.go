package gocv_resource_tracker

import (
	"image"

	"gocv.io/x/gocv"
)

func (g *GoCVResourceTracker) NewMat() gocv.Mat {
	mat := gocv.NewMat()
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) NewMatWithSizes(sizes []int, mt gocv.MatType) gocv.Mat {
	mat := gocv.NewMatWithSizes(sizes, mt)
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

func (g *GoCVResourceTracker) CloneRegion(src gocv.Mat, rect image.Rectangle) gocv.Mat {
	region := src.Region(rect)
	defer region.Close()
	return g.Clone(region)
}

func (g *GoCVResourceTracker) GetStructuringElement(morphShape gocv.MorphShape, size image.Point) gocv.Mat {
	kernel := gocv.GetStructuringElement(morphShape, size)
	g.TrackCloseError(&kernel)
	return kernel
}

func (g *GoCVResourceTracker) NewMatWithSizesWithScalar(sizes []int, mt gocv.MatType, s gocv.Scalar) gocv.Mat {
	mat := gocv.NewMatWithSizesWithScalar(sizes, mt, s)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) NewMatFromScalar(s gocv.Scalar, mt gocv.MatType) gocv.Mat {
	mat := gocv.NewMatFromScalar(s, mt)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) NewMatWithSizeFromScalar(s gocv.Scalar, rows int, cols int, mt gocv.MatType) gocv.Mat {
	mat := gocv.NewMatWithSizeFromScalar(s, rows, cols, mt)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) Eye(rows int, cols int, mt gocv.MatType) gocv.Mat {
	mat := gocv.Eye(rows, cols, mt)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) Zeros(rows int, cols int, mt gocv.MatType) gocv.Mat {
	mat := gocv.Zeros(rows, cols, mt)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) Ones(rows int, cols int, mt gocv.MatType) gocv.Mat {
	mat := gocv.Ones(rows, cols, mt)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) Reshape(src gocv.Mat, cn int, rows int) gocv.Mat {
	mat := src.Reshape(cn, rows)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) ConvertFp16(src gocv.Mat) gocv.Mat {
	mat := src.ConvertFp16()
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) Sqrt(src gocv.Mat) gocv.Mat {
	mat := src.Sqrt()
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) MultiplyMatrix(src, x gocv.Mat) gocv.Mat {
	mat := src.MultiplyMatrix(x)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) T(src gocv.Mat) gocv.Mat {
	mat := src.T()
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) RowRange(src gocv.Mat, start, end int) gocv.Mat {
	mat := src.RowRange(start, end)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) ColRange(src gocv.Mat, start, end int) gocv.Mat {
	mat := src.ColRange(start, end)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) NewMatWithSizesFromBytes(sizes []int, mt gocv.MatType, data []byte) (gocv.Mat, error) {
	mat, err := gocv.NewMatWithSizesFromBytes(sizes, mt, data)
	g.TrackCloseError(&mat)
	return mat, err
}

func (g *GoCVResourceTracker) NewMatFromBytes(rows int, cols int, mt gocv.MatType, data []byte) (gocv.Mat, error) {
	mat, err := gocv.NewMatFromBytes(rows, cols, mt, data)
	g.TrackCloseError(&mat)
	return mat, err
}

func (g *GoCVResourceTracker) FromPtr(src gocv.Mat, rows int, cols int, mt gocv.MatType, prow int, pcol int) (gocv.Mat, error) {
	mat, err := src.FromPtr(rows, cols, mt, prow, pcol)
	g.TrackCloseError(&mat)
	return mat, err
}
