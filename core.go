
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
    "image"
    "image/color"
)

type MatType = gocv.MatType
type CompareType = gocv.CompareType
type Point2f = gocv.Point2f
type Mat struct {
    gocv.Mat
	ResourceTracker *GoCVResourceTracker
}
type CovarFlags = gocv.CovarFlags
type DftFlags = gocv.DftFlags
type RotateFlag = gocv.RotateFlag
type KMeansFlags = gocv.KMeansFlags
type NormType = gocv.NormType
type TermCriteriaType = gocv.TermCriteriaType
type SolveDecompositionFlags = gocv.SolveDecompositionFlags
type ReduceTypes = gocv.ReduceTypes
type SortFlags = gocv.SortFlags
type TermCriteria = gocv.TermCriteria
type Scalar = gocv.Scalar
type KeyPoint = gocv.KeyPoint
type DMatch = gocv.DMatch
type Vecb = gocv.Vecb
type Vecf = gocv.Vecf
type Vecd = gocv.Vecd
type Veci = gocv.Veci
type PointVector struct {
    gocv.PointVector
	ResourceTracker *GoCVResourceTracker
}
type PointsVector struct {
    gocv.PointsVector
	ResourceTracker *GoCVResourceTracker
}
type Point2fVector struct {
    gocv.Point2fVector
	ResourceTracker *GoCVResourceTracker
}
type RNG = gocv.RNG
type RNGDistType = gocv.RNGDistType
type NativeByteBuffer struct {
    gocv.NativeByteBuffer
	ResourceTracker *GoCVResourceTracker
}

func (g *GoCVResourceTracker) NewMat() Mat {
    rs0 := gocv.NewMat()
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewMatWithSize(rows int, cols int, mt MatType) Mat {
    rs0 := gocv.NewMatWithSize(rows, cols, mt)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewMatWithSizes(sizes []int, mt MatType) Mat {
    rs0 := gocv.NewMatWithSizes(sizes, mt)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewMatWithSizesWithScalar(sizes []int, mt MatType, s Scalar) Mat {
    rs0 := gocv.NewMatWithSizesWithScalar(sizes, mt, s)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewMatWithSizesFromBytes(sizes []int, mt MatType, data []byte) (Mat, error) {
    rs0, rs1 := gocv.NewMatWithSizesFromBytes(sizes, mt, data)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) NewMatFromScalar(s Scalar, mt MatType) Mat {
    rs0 := gocv.NewMatFromScalar(s, mt)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewMatWithSizeFromScalar(s Scalar, rows int, cols int, mt MatType) Mat {
    rs0 := gocv.NewMatWithSizeFromScalar(s, rows, cols, mt)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewMatFromBytes(rows int, cols int, mt MatType, data []byte) (Mat, error) {
    rs0, rs1 := gocv.NewMatFromBytes(rows, cols, mt, data)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) Eye(rows int, cols int, mt MatType) Mat {
    rs0 := gocv.Eye(rows, cols, mt)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) Zeros(rows int, cols int, mt MatType) Mat {
    rs0 := gocv.Zeros(rows, cols, mt)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) Ones(rows int, cols int, mt MatType) Mat {
    rs0 := gocv.Ones(rows, cols, mt)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) LUT(src, wbLUT Mat, dst *Mat) {
    gocv.LUT(src.Mat, wbLUT.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) AbsDiff(src1, src2 Mat, dst *Mat) {
    gocv.AbsDiff(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) Add(src1, src2 Mat, dst *Mat) {
    gocv.Add(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) AddWeighted(src1 Mat, alpha float64, src2 Mat, beta float64, gamma float64, dst *Mat) {
    gocv.AddWeighted(src1.Mat, alpha, src2.Mat, beta, gamma, &(dst.Mat))
}


func (g *GoCVResourceTracker) BitwiseAnd(src1 Mat, src2 Mat, dst *Mat) {
    gocv.BitwiseAnd(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) BitwiseAndWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
    gocv.BitwiseAndWithMask(src1.Mat, src2.Mat, &(dst.Mat), mask.Mat)
}


func (g *GoCVResourceTracker) BitwiseNot(src1 Mat, dst *Mat) {
    gocv.BitwiseNot(src1.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) BitwiseNotWithMask(src1 Mat, dst *Mat, mask Mat) {
    gocv.BitwiseNotWithMask(src1.Mat, &(dst.Mat), mask.Mat)
}


func (g *GoCVResourceTracker) BitwiseOr(src1 Mat, src2 Mat, dst *Mat) {
    gocv.BitwiseOr(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) BitwiseOrWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
    gocv.BitwiseOrWithMask(src1.Mat, src2.Mat, &(dst.Mat), mask.Mat)
}


func (g *GoCVResourceTracker) BitwiseXor(src1 Mat, src2 Mat, dst *Mat) {
    gocv.BitwiseXor(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) BitwiseXorWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
    gocv.BitwiseXorWithMask(src1.Mat, src2.Mat, &(dst.Mat), mask.Mat)
}


func (g *GoCVResourceTracker) BatchDistance(src1 Mat, src2 Mat, dist Mat, dtype MatType, nidx Mat, normType NormType, K int, mask Mat, update int, crosscheck bool) {
    gocv.BatchDistance(src1.Mat, src2.Mat, dist.Mat, dtype, nidx.Mat, normType, K, mask.Mat, update, crosscheck)
}


func (g *GoCVResourceTracker) BorderInterpolate(p int, len int, borderType CovarFlags) int {
    rs0 := gocv.BorderInterpolate(p, len, borderType)
    return rs0
}


func (g *GoCVResourceTracker) CalcCovarMatrix(samples Mat, covar *Mat, mean *Mat, flags CovarFlags, ctype MatType) {
    gocv.CalcCovarMatrix(samples.Mat, &(covar.Mat), &(mean.Mat), flags, ctype)
}


func (g *GoCVResourceTracker) CartToPolar(x Mat, y Mat, magnitude *Mat, angle *Mat, angleInDegrees bool) {
    gocv.CartToPolar(x.Mat, y.Mat, &(magnitude.Mat), &(angle.Mat), angleInDegrees)
}


func (g *GoCVResourceTracker) CheckRange(src Mat) bool {
    rs0 := gocv.CheckRange(src.Mat)
    return rs0
}


func (g *GoCVResourceTracker) Compare(src1 Mat, src2 Mat, dst *Mat, ct CompareType) {
    gocv.Compare(src1.Mat, src2.Mat, &(dst.Mat), ct)
}


func (g *GoCVResourceTracker) CountNonZero(src Mat) int {
    rs0 := gocv.CountNonZero(src.Mat)
    return rs0
}


func (g *GoCVResourceTracker) CompleteSymm(m Mat, lowerToUpper bool) {
    gocv.CompleteSymm(m.Mat, lowerToUpper)
}


func (g *GoCVResourceTracker) ConvertScaleAbs(src Mat, dst *Mat, alpha float64, beta float64) {
    gocv.ConvertScaleAbs(src.Mat, &(dst.Mat), alpha, beta)
}


func (g *GoCVResourceTracker) CopyMakeBorder(src Mat, dst *Mat, top int, bottom int, left int, right int, bt BorderType, value color.RGBA) {
    gocv.CopyMakeBorder(src.Mat, &(dst.Mat), top, bottom, left, right, bt, value)
}


func (g *GoCVResourceTracker) DCT(src Mat, dst *Mat, flags DftFlags) {
    gocv.DCT(src.Mat, &(dst.Mat), flags)
}


func (g *GoCVResourceTracker) Determinant(src Mat) float64 {
    rs0 := gocv.Determinant(src.Mat)
    return rs0
}


func (g *GoCVResourceTracker) DFT(src Mat, dst *Mat, flags DftFlags) {
    gocv.DFT(src.Mat, &(dst.Mat), flags)
}


func (g *GoCVResourceTracker) Divide(src1 Mat, src2 Mat, dst *Mat) {
    gocv.Divide(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) Eigen(src Mat, eigenvalues *Mat, eigenvectors *Mat) bool {
    rs0 := gocv.Eigen(src.Mat, &(eigenvalues.Mat), &(eigenvectors.Mat))
    return rs0
}


func (g *GoCVResourceTracker) EigenNonSymmetric(src Mat, eigenvalues *Mat, eigenvectors *Mat) {
    gocv.EigenNonSymmetric(src.Mat, &(eigenvalues.Mat), &(eigenvectors.Mat))
}


func (g *GoCVResourceTracker) Exp(src Mat, dst *Mat) {
    gocv.Exp(src.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) ExtractChannel(src Mat, dst *Mat, coi int) {
    gocv.ExtractChannel(src.Mat, &(dst.Mat), coi)
}


func (g *GoCVResourceTracker) FindNonZero(src Mat, idx *Mat) {
    gocv.FindNonZero(src.Mat, &(idx.Mat))
}


func (g *GoCVResourceTracker) Flip(src Mat, dst *Mat, flipCode int) {
    gocv.Flip(src.Mat, &(dst.Mat), flipCode)
}


func (g *GoCVResourceTracker) Gemm(src1, src2 Mat, alpha float64, src3 Mat, beta float64, dst *Mat, flags int) {
    gocv.Gemm(src1.Mat, src2.Mat, alpha, src3.Mat, beta, &(dst.Mat), flags)
}


func (g *GoCVResourceTracker) GetOptimalDFTSize(vecsize int) int {
    rs0 := gocv.GetOptimalDFTSize(vecsize)
    return rs0
}


func (g *GoCVResourceTracker) Hconcat(src1, src2 Mat, dst *Mat) {
    gocv.Hconcat(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) Vconcat(src1, src2 Mat, dst *Mat) {
    gocv.Vconcat(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) Rotate(src Mat, dst *Mat, code RotateFlag) {
    gocv.Rotate(src.Mat, &(dst.Mat), code)
}


func (g *GoCVResourceTracker) IDCT(src Mat, dst *Mat, flags int) {
    gocv.IDCT(src.Mat, &(dst.Mat), flags)
}


func (g *GoCVResourceTracker) IDFT(src Mat, dst *Mat, flags, nonzeroRows int) {
    gocv.IDFT(src.Mat, &(dst.Mat), flags, nonzeroRows)
}


func (g *GoCVResourceTracker) InRange(src, lb, ub Mat, dst *Mat) {
    gocv.InRange(src.Mat, lb.Mat, ub.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) InRangeWithScalar(src Mat, lb, ub Scalar, dst *Mat) {
    gocv.InRangeWithScalar(src.Mat, lb, ub, &(dst.Mat))
}


func (g *GoCVResourceTracker) InsertChannel(src Mat, dst *Mat, coi int) {
    gocv.InsertChannel(src.Mat, &(dst.Mat), coi)
}


func (g *GoCVResourceTracker) Invert(src Mat, dst *Mat, flags SolveDecompositionFlags) float64 {
    rs0 := gocv.Invert(src.Mat, &(dst.Mat), flags)
    return rs0
}


func (g *GoCVResourceTracker) KMeans(data Mat, k int, bestLabels *Mat, criteria TermCriteria, attempts int, flags KMeansFlags, centers *Mat) float64 {
    rs0 := gocv.KMeans(data.Mat, k, &(bestLabels.Mat), criteria, attempts, flags, &(centers.Mat))
    return rs0
}


func (g *GoCVResourceTracker) KMeansPoints(points PointVector, k int, bestLabels *Mat, criteria TermCriteria, attempts int, flags KMeansFlags, centers *Mat) float64 {
    rs0 := gocv.KMeansPoints(points.PointVector, k, &(bestLabels.Mat), criteria, attempts, flags, &(centers.Mat))
    return rs0
}


func (g *GoCVResourceTracker) Log(src Mat, dst *Mat) {
    gocv.Log(src.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) Magnitude(x, y Mat, magnitude *Mat) {
    gocv.Magnitude(x.Mat, y.Mat, &(magnitude.Mat))
}


func (g *GoCVResourceTracker) Max(src1, src2 Mat, dst *Mat) {
    gocv.Max(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) MeanStdDev(src Mat, dst *Mat, dstStdDev *Mat) {
    gocv.MeanStdDev(src.Mat, &(dst.Mat), &(dstStdDev.Mat))
}


func (g *GoCVResourceTracker) Merge(mv []Mat, dst *Mat) {
    gocv.Merge(SliceToGoCVCloser(mv), &(dst.Mat))
}


func (g *GoCVResourceTracker) Min(src1, src2 Mat, dst *Mat) {
    gocv.Min(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) MinMaxIdx(input Mat) (minVal, maxVal float32, minIdx, maxIdx int) {
    rs0, rs1, rs2, rs3 := gocv.MinMaxIdx(input.Mat)
    return rs0, rs1, rs2, rs3
}


func (g *GoCVResourceTracker) MinMaxLoc(input Mat) (minVal, maxVal float32, minLoc, maxLoc image.Point) {
    rs0, rs1, rs2, rs3 := gocv.MinMaxLoc(input.Mat)
    return rs0, rs1, rs2, rs3
}


func (g *GoCVResourceTracker) MixChannels(src []Mat, dst []Mat, fromTo []int) {
    gocv.MixChannels(SliceToGoCVCloser(src), SliceToGoCVCloser(dst), fromTo)
}


func (g *GoCVResourceTracker) MulSpectrums(a Mat, b Mat, dst *Mat, flags DftFlags) {
    gocv.MulSpectrums(a.Mat, b.Mat, &(dst.Mat), flags)
}


func (g *GoCVResourceTracker) Multiply(src1 Mat, src2 Mat, dst *Mat) {
    gocv.Multiply(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) MultiplyWithParams(src1 Mat, src2 Mat, dst *Mat, scale float64, dtype MatType) {
    gocv.MultiplyWithParams(src1.Mat, src2.Mat, &(dst.Mat), scale, dtype)
}


func (g *GoCVResourceTracker) Normalize(src Mat, dst *Mat, alpha float64, beta float64, typ NormType) {
    gocv.Normalize(src.Mat, &(dst.Mat), alpha, beta, typ)
}


func (g *GoCVResourceTracker) Norm(src1 Mat, normType NormType) float64 {
    rs0 := gocv.Norm(src1.Mat, normType)
    return rs0
}


func (g *GoCVResourceTracker) NormWithMats(src1 Mat, src2 Mat, normType NormType) float64 {
    rs0 := gocv.NormWithMats(src1.Mat, src2.Mat, normType)
    return rs0
}


func (g *GoCVResourceTracker) PerspectiveTransform(src Mat, dst *Mat, tm Mat) {
    gocv.PerspectiveTransform(src.Mat, &(dst.Mat), tm.Mat)
}


func (g *GoCVResourceTracker) Solve(src1 Mat, src2 Mat, dst *Mat, flags SolveDecompositionFlags) bool {
    rs0 := gocv.Solve(src1.Mat, src2.Mat, &(dst.Mat), flags)
    return rs0
}


func (g *GoCVResourceTracker) SolveCubic(coeffs Mat, roots *Mat) int {
    rs0 := gocv.SolveCubic(coeffs.Mat, &(roots.Mat))
    return rs0
}


func (g *GoCVResourceTracker) SolvePoly(coeffs Mat, roots *Mat, maxIters int) float64 {
    rs0 := gocv.SolvePoly(coeffs.Mat, &(roots.Mat), maxIters)
    return rs0
}


func (g *GoCVResourceTracker) Reduce(src Mat, dst *Mat, dim int, rType ReduceTypes, dType MatType) {
    gocv.Reduce(src.Mat, &(dst.Mat), dim, rType, dType)
}


func (g *GoCVResourceTracker) Repeat(src Mat, nY int, nX int, dst *Mat) {
    gocv.Repeat(src.Mat, nY, nX, &(dst.Mat))
}


func (g *GoCVResourceTracker) ScaleAdd(src1 Mat, alpha float64, src2 Mat, dst *Mat) {
    gocv.ScaleAdd(src1.Mat, alpha, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) SetIdentity(src Mat, scalar float64) {
    gocv.SetIdentity(src.Mat, scalar)
}


func (g *GoCVResourceTracker) Sort(src Mat, dst *Mat, flags SortFlags) {
    gocv.Sort(src.Mat, &(dst.Mat), flags)
}


func (g *GoCVResourceTracker) SortIdx(src Mat, dst *Mat, flags SortFlags) {
    gocv.SortIdx(src.Mat, &(dst.Mat), flags)
}


func (g *GoCVResourceTracker) Split(src Mat) (mv []Mat) {
    rs0 := gocv.Split(src.Mat)
    return GoCVCloserToSlice(rs0, g)
}


func (g *GoCVResourceTracker) Subtract(src1 Mat, src2 Mat, dst *Mat) {
    gocv.Subtract(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) Trace(src Mat) Scalar {
    rs0 := gocv.Trace(src.Mat)
    return rs0
}


func (g *GoCVResourceTracker) Transform(src Mat, dst *Mat, tm Mat) {
    gocv.Transform(src.Mat, &(dst.Mat), tm.Mat)
}


func (g *GoCVResourceTracker) Transpose(src Mat, dst *Mat) {
    gocv.Transpose(src.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) Pow(src Mat, power float64, dst *Mat) {
    gocv.Pow(src.Mat, power, &(dst.Mat))
}


func (g *GoCVResourceTracker) PolarToCart(magnitude Mat, degree Mat, x *Mat, y *Mat, angleInDegrees bool) {
    gocv.PolarToCart(magnitude.Mat, degree.Mat, &(x.Mat), &(y.Mat), angleInDegrees)
}


func (g *GoCVResourceTracker) Phase(x, y Mat, angle *Mat, angleInDegrees bool) {
    gocv.Phase(x.Mat, y.Mat, &(angle.Mat), angleInDegrees)
}


func (g *GoCVResourceTracker) NewTermCriteria(typ TermCriteriaType, maxCount int, epsilon float64) TermCriteria {
    rs0 := gocv.NewTermCriteria(typ, maxCount, epsilon)
    return rs0
}


func (g *GoCVResourceTracker) NewScalar(v1 float64, v2 float64, v3 float64, v4 float64) Scalar {
    rs0 := gocv.NewScalar(v1, v2, v3, v4)
    return rs0
}


func (g *GoCVResourceTracker) NewPointVector() PointVector {
    rs0 := gocv.NewPointVector()
    g.TrackCloser(rs0)
    pkg0 := PointVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewPointVectorFromPoints(pts []image.Point) PointVector {
    rs0 := gocv.NewPointVectorFromPoints(pts)
    g.TrackCloser(rs0)
    pkg0 := PointVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewPointVectorFromMat(mat Mat) PointVector {
    rs0 := gocv.NewPointVectorFromMat(mat.Mat)
    g.TrackCloser(rs0)
    pkg0 := PointVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewPointsVector() PointsVector {
    rs0 := gocv.NewPointsVector()
    g.TrackCloser(rs0)
    pkg0 := PointsVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewPointsVectorFromPoints(pts [][]image.Point) PointsVector {
    rs0 := gocv.NewPointsVectorFromPoints(pts)
    g.TrackCloser(rs0)
    pkg0 := PointsVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewPoint2fVector() Point2fVector {
    rs0 := gocv.NewPoint2fVector()
    g.TrackCloser(rs0)
    pkg0 := Point2fVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewPoint2fVectorFromPoints(pts []Point2f) Point2fVector {
    rs0 := gocv.NewPoint2fVectorFromPoints(pts)
    g.TrackCloser(rs0)
    pkg0 := Point2fVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewPoint2fVectorFromMat(mat Mat) Point2fVector {
    rs0 := gocv.NewPoint2fVectorFromMat(mat.Mat)
    g.TrackCloser(rs0)
    pkg0 := Point2fVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) GetTickCount() float64 {
    rs0 := gocv.GetTickCount()
    return rs0
}


func (g *GoCVResourceTracker) GetTickFrequency() float64 {
    rs0 := gocv.GetTickFrequency()
    return rs0
}


func (g *GoCVResourceTracker) TheRNG() RNG {
    rs0 := gocv.TheRNG()
    return rs0
}


func (g *GoCVResourceTracker) SetRNGSeed(seed int) {
    gocv.SetRNGSeed(seed)
}


func (g *GoCVResourceTracker) RandN(mat *Mat, mean, stddev Scalar) {
    gocv.RandN(&(mat.Mat), mean, stddev)
}


func (g *GoCVResourceTracker) RandShuffle(mat *Mat) {
    gocv.RandShuffle(&(mat.Mat))
}


func (g *GoCVResourceTracker) RandShuffleWithParams(mat *Mat, iterFactor float64, rng RNG) {
    gocv.RandShuffleWithParams(&(mat.Mat), iterFactor, rng)
}


func (g *GoCVResourceTracker) RandU(mat *Mat, low, high Scalar) {
    gocv.RandU(&(mat.Mat), low, high)
}
func (m *Mat) Clone() Mat {
    rs0 := m.Mat.Clone()

    m.ResourceTracker.TrackCloseError(&rs0)
    return Mat{rs0, m.ResourceTracker}
}

func (m *Mat) ColRange(start, end int) Mat {
    rs0 := m.Mat.ColRange(start, end)

    m.ResourceTracker.TrackCloseError(&rs0)
    return Mat{rs0, m.ResourceTracker}
}

func (m *Mat) ConvertFp16() Mat {
    rs0 := m.Mat.ConvertFp16()

    m.ResourceTracker.TrackCloseError(&rs0)
    return Mat{rs0, m.ResourceTracker}
}

func (m *Mat) ConvertTo(dst *Mat, mt MatType) {
    m.Mat.ConvertTo(&(dst.Mat), mt)
}

func (m *Mat) ConvertToWithParams(dst *Mat, mt MatType, alpha, beta float32) {
    m.Mat.ConvertToWithParams(&(dst.Mat), mt, alpha, beta)
}

func (m *Mat) CopyTo(dst *Mat) {
    m.Mat.CopyTo(&(dst.Mat))
}

func (m *Mat) CopyToWithMask(dst *Mat, mask Mat) {
    m.Mat.CopyToWithMask(&(dst.Mat), mask.Mat)
}

func (m *Mat) FromPtr(rows int, cols int, mt MatType, prow int, pcol int) (Mat, error) {
    rs0, rs1 := m.Mat.FromPtr(rows, cols, mt, prow, pcol)

    m.ResourceTracker.TrackCloseError(&rs0)
    return Mat{rs0, m.ResourceTracker}, rs1
}

func (m *Mat) MeanWithMask(mask Mat) Scalar {
    rs0 := m.Mat.MeanWithMask(mask.Mat)
    return rs0
}

func (m *Mat) MultiplyMatrix(x Mat) Mat {
    rs0 := m.Mat.MultiplyMatrix(x.Mat)

    m.ResourceTracker.TrackCloseError(&rs0)
    return Mat{rs0, m.ResourceTracker}
}

func (m *Mat) Region(rio image.Rectangle) Mat {
    rs0 := m.Mat.Region(rio)

    m.ResourceTracker.TrackCloseError(&rs0)
    return Mat{rs0, m.ResourceTracker}
}

func (m *Mat) Reshape(cn int, rows int) Mat {
    rs0 := m.Mat.Reshape(cn, rows)

    m.ResourceTracker.TrackCloseError(&rs0)
    return Mat{rs0, m.ResourceTracker}
}

func (m *Mat) RowRange(start, end int) Mat {
    rs0 := m.Mat.RowRange(start, end)

    m.ResourceTracker.TrackCloseError(&rs0)
    return Mat{rs0, m.ResourceTracker}
}

func (m *Mat) Sqrt() Mat {
    rs0 := m.Mat.Sqrt()

    m.ResourceTracker.TrackCloseError(&rs0)
    return Mat{rs0, m.ResourceTracker}
}

func (m *Mat) T() Mat {
    rs0 := m.Mat.T()

    m.ResourceTracker.TrackCloseError(&rs0)
    return Mat{rs0, m.ResourceTracker}
}
func (pvs PointsVector) Append(pv PointVector) {
    pvs.PointsVector.Append(pv.PointVector)
}

func (pvs PointsVector) At(idx int) PointVector {
    rs0 := pvs.PointsVector.At(idx)

    pvs.ResourceTracker.TrackCloser(rs0)
    return PointVector{rs0, pvs.ResourceTracker}
}