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
	Must(rows > 0 && cols > 0, "rows > 0 && cols > 0")

	mat := gocv.NewMatWithSize(rows, cols, matType)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) Clone(src gocv.Mat) gocv.Mat {
	dst := src.Clone()
	g.TrackCloseError(&dst)
	return dst
}

func (g *GoCVResourceTracker) Region(src gocv.Mat, rect image.Rectangle) gocv.Mat {

	Must(rect.In(image.Rect(0, 0, src.Cols(), src.Rows())),
		"rect inside src.Bounds()")

	region := src.Region(rect)
	g.TrackCloseError(&region)
	return region
}

func (g *GoCVResourceTracker) CloneRegion(src gocv.Mat, rect image.Rectangle) gocv.Mat {
	Must(rect.In(image.Rect(0, 0, src.Cols(), src.Rows())),
		"rect inside src.Bounds()")

	region := src.Region(rect)
	defer region.Close()

	return g.Clone(region)
}

func (g *GoCVResourceTracker) GetStructuringElement(morphShape gocv.MorphShape, size image.Point) gocv.Mat {
	Must(size.X > 0 && size.Y > 0, "size.X > 0 && size.Y > 0")

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
	Must(rows > 0 && cols > 0, "rows > 0 && cols > 0")

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

var (
	matTypeSizeMap = map[gocv.MatType]int{
		gocv.MatTypeCV8U:  1,
		gocv.MatTypeCV8S:  1,
		gocv.MatTypeCV16U: 2,
		gocv.MatTypeCV16S: 2,
		gocv.MatTypeCV32S: 4,
		gocv.MatTypeCV32F: 4,
		gocv.MatTypeCV64F: 8,
	}
)

func sizeOfMatType(mt gocv.MatType) int {
	sz := matTypeSizeMap[mt]
	channelNum := int(mt)/8 + 1
	return sz * channelNum
}

func (g *GoCVResourceTracker) NewMatFromBytes(rows int, cols int, mt gocv.MatType, data []byte) (gocv.Mat, error) {
	Must(rows > 0 && cols > 0, "rows > 0 && cols > 0")
	pixSize := sizeOfMatType(mt)
	MustEqual(pixSize*rows*cols, len(data), "sizeof(%v) * rows * cols != len(data)", mt)

	mat, err := gocv.NewMatFromBytes(rows, cols, mt, data)
	g.TrackCloseError(&mat)
	return mat, err
}

func (g *GoCVResourceTracker) FromPtr(src gocv.Mat, rows int, cols int, mt gocv.MatType, prow int, pcol int) (gocv.Mat, error) {
	mat, err := src.FromPtr(rows, cols, mt, prow, pcol)
	g.TrackCloseError(&mat)
	return mat, err
}

func (g *GoCVResourceTracker) NewPointVector() gocv.PointVector {
	pv := gocv.NewPointVector()
	g.TrackCloser(&pv)
	return pv
}

func (g *GoCVResourceTracker) NewPointVectorFromPoints(pts []image.Point) gocv.PointVector {
	if len(pts) == 0 {
		return g.NewPointVector()
	}

	pv := gocv.NewPointVectorFromPoints(pts)
	g.TrackCloser(&pv)
	return pv
}

func (g *GoCVResourceTracker) NewPointVectorFromMat(mat gocv.Mat) gocv.PointVector {

	if mat.Rows() == 0 {
		return g.NewPointVector()
	}

	pv := gocv.NewPointVectorFromMat(mat)
	g.TrackCloser(&pv)
	return pv
}

func (g *GoCVResourceTracker) FindContours(src gocv.Mat, mode gocv.RetrievalMode, method gocv.ContourApproximationMode) gocv.PointsVector {
	pvs := gocv.FindContours(src, mode, method)
	g.TrackCloser(&pvs)
	return pvs
}

func (g *GoCVResourceTracker) NewPointsVector() gocv.PointsVector {
	pv := gocv.NewPointsVector()
	g.TrackCloser(&pv)
	return pv
}

func (g *GoCVResourceTracker) NewPointsVectorFromPoints(pts [][]image.Point) gocv.PointsVector {
	if len(pts) == 0 {
		return g.NewPointsVector()
	}

	pvs := gocv.NewPointsVectorFromPoints(pts)
	g.TrackCloser(&pvs)
	return pvs
}

func (g *GoCVResourceTracker) NewPoint2fVector() gocv.Point2fVector {
	pvs := gocv.NewPoint2fVector()
	g.TrackCloser(&pvs)
	return pvs
}
func (g *GoCVResourceTracker) NewPoint2fVectorFromPoints(pts []gocv.Point2f) gocv.Point2fVector {
	if len(pts) == 0 {
		return g.NewPoint2fVector()
	}

	pvs := gocv.NewPoint2fVectorFromPoints(pts)
	g.TrackCloser(&pvs)
	return pvs
}
func (g *GoCVResourceTracker) NewPoint2fVectorFromMat(mat gocv.Mat) gocv.Point2fVector {
	if mat.Rows() == 0 {
		return g.NewPoint2fVector()
	}

	pvs := gocv.NewPoint2fVectorFromMat(mat)
	g.TrackCloser(&pvs)
	return pvs
}
