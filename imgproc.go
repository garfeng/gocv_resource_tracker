
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
    "image"
    "image/color"
)

type HistCompMethod = gocv.HistCompMethod
type DistanceTransformLabelTypes = gocv.DistanceTransformLabelTypes
type DistanceTransformMasks = gocv.DistanceTransformMasks
type RetrievalMode = gocv.RetrievalMode
type ContourApproximationMode = gocv.ContourApproximationMode
type RotatedRect = gocv.RotatedRect
type ConnectedComponentsAlgorithmType = gocv.ConnectedComponentsAlgorithmType
type ConnectedComponentsTypes = gocv.ConnectedComponentsTypes
type TemplateMatchMode = gocv.TemplateMatchMode
type MorphShape = gocv.MorphShape
type MorphType = gocv.MorphType
type BorderType = gocv.BorderType
type GrabCutMode = gocv.GrabCutMode
type HoughMode = gocv.HoughMode
type ThresholdType = gocv.ThresholdType
type AdaptiveThresholdType = gocv.AdaptiveThresholdType
type HersheyFont = gocv.HersheyFont
type LineType = gocv.LineType
type InterpolationFlags = gocv.InterpolationFlags
type ColormapTypes = gocv.ColormapTypes
type HomographyMethod = gocv.HomographyMethod
type DistanceTypes = gocv.DistanceTypes
type CLAHE struct {
    gocv.CLAHE
	ResourceTracker *GoCVResourceTracker
}

func (g *GoCVResourceTracker) ArcLength(curve PointVector, isClosed bool) float64 {
    rs0 := gocv.ArcLength(curve.PointVector, isClosed)
    return rs0
}


func (g *GoCVResourceTracker) ApproxPolyDP(curve PointVector, epsilon float64, closed bool) PointVector {
    rs0 := gocv.ApproxPolyDP(curve.PointVector, epsilon, closed)
    g.TrackCloser(rs0)
    pkg0 := PointVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) ConvexHull(points PointVector, hull *Mat, clockwise bool, returnPoints bool) {
    gocv.ConvexHull(points.PointVector, &(hull.Mat), clockwise, returnPoints)
}


func (g *GoCVResourceTracker) ConvexityDefects(contour PointVector, hull Mat, result *Mat) {
    gocv.ConvexityDefects(contour.PointVector, hull.Mat, &(result.Mat))
}


func (g *GoCVResourceTracker) CvtColor(src Mat, dst *Mat, code ColorConversionCode) {
    gocv.CvtColor(src.Mat, &(dst.Mat), code)
}


func (g *GoCVResourceTracker) EqualizeHist(src Mat, dst *Mat) {
    gocv.EqualizeHist(src.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) CalcHist(src []Mat, channels []int, mask Mat, hist *Mat, size []int, ranges []float64, acc bool) {
    gocv.CalcHist(SliceToGoCVCloser(src), channels, mask.Mat, &(hist.Mat), size, ranges, acc)
}


func (g *GoCVResourceTracker) CalcBackProject(src []Mat, channels []int, hist Mat, backProject *Mat, ranges []float64, uniform bool) {
    gocv.CalcBackProject(SliceToGoCVCloser(src), channels, hist.Mat, &(backProject.Mat), ranges, uniform)
}


func (g *GoCVResourceTracker) CompareHist(hist1 Mat, hist2 Mat, method HistCompMethod) float32 {
    rs0 := gocv.CompareHist(hist1.Mat, hist2.Mat, method)
    return rs0
}


func (g *GoCVResourceTracker) ClipLine(imgSize image.Point, pt1 image.Point, pt2 image.Point) bool {
    rs0 := gocv.ClipLine(imgSize, pt1, pt2)
    return rs0
}


func (g *GoCVResourceTracker) BilateralFilter(src Mat, dst *Mat, diameter int, sigmaColor float64, sigmaSpace float64) {
    gocv.BilateralFilter(src.Mat, &(dst.Mat), diameter, sigmaColor, sigmaSpace)
}


func (g *GoCVResourceTracker) Blur(src Mat, dst *Mat, ksize image.Point) {
    gocv.Blur(src.Mat, &(dst.Mat), ksize)
}


func (g *GoCVResourceTracker) BoxFilter(src Mat, dst *Mat, depth int, ksize image.Point) {
    gocv.BoxFilter(src.Mat, &(dst.Mat), depth, ksize)
}


func (g *GoCVResourceTracker) SqBoxFilter(src Mat, dst *Mat, depth int, ksize image.Point) {
    gocv.SqBoxFilter(src.Mat, &(dst.Mat), depth, ksize)
}


func (g *GoCVResourceTracker) Dilate(src Mat, dst *Mat, kernel Mat) {
    gocv.Dilate(src.Mat, &(dst.Mat), kernel.Mat)
}


func (g *GoCVResourceTracker) DilateWithParams(src Mat, dst *Mat, kernel Mat, anchor image.Point, iterations, borderType BorderType, borderValue color.RGBA) {
    gocv.DilateWithParams(src.Mat, &(dst.Mat), kernel.Mat, anchor, iterations, borderType, borderValue)
}


func (g *GoCVResourceTracker) DistanceTransform(src Mat, dst *Mat, labels *Mat, distType DistanceTypes, maskSize DistanceTransformMasks, labelType DistanceTransformLabelTypes) {
    gocv.DistanceTransform(src.Mat, &(dst.Mat), &(labels.Mat), distType, maskSize, labelType)
}


func (g *GoCVResourceTracker) Erode(src Mat, dst *Mat, kernel Mat) {
    gocv.Erode(src.Mat, &(dst.Mat), kernel.Mat)
}


func (g *GoCVResourceTracker) ErodeWithParams(src Mat, dst *Mat, kernel Mat, anchor image.Point, iterations, borderType int) {
    gocv.ErodeWithParams(src.Mat, &(dst.Mat), kernel.Mat, anchor, iterations, borderType)
}


func (g *GoCVResourceTracker) BoundingRect(contour PointVector) image.Rectangle {
    rs0 := gocv.BoundingRect(contour.PointVector)
    return rs0
}


func (g *GoCVResourceTracker) BoxPoints(rect RotatedRect, pts *Mat) {
    gocv.BoxPoints(rect, &(pts.Mat))
}


func (g *GoCVResourceTracker) ContourArea(contour PointVector) float64 {
    rs0 := gocv.ContourArea(contour.PointVector)
    return rs0
}


func (g *GoCVResourceTracker) MinAreaRect(points PointVector) RotatedRect {
    rs0 := gocv.MinAreaRect(points.PointVector)
    return rs0
}


func (g *GoCVResourceTracker) FitEllipse(pts PointVector) RotatedRect {
    rs0 := gocv.FitEllipse(pts.PointVector)
    return rs0
}


func (g *GoCVResourceTracker) MinEnclosingCircle(pts PointVector) (x, y, radius float32) {
    rs0, rs1, rs2 := gocv.MinEnclosingCircle(pts.PointVector)
    return rs0, rs1, rs2
}


func (g *GoCVResourceTracker) FindContours(src Mat, mode RetrievalMode, method ContourApproximationMode) PointsVector {
    rs0 := gocv.FindContours(src.Mat, mode, method)
    g.TrackCloser(rs0)
    pkg0 := PointsVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) FindContoursWithParams(src Mat, hierarchy *Mat, mode RetrievalMode, method ContourApproximationMode) PointsVector {
    rs0 := gocv.FindContoursWithParams(src.Mat, &(hierarchy.Mat), mode, method)
    g.TrackCloser(rs0)
    pkg0 := PointsVector{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) PointPolygonTest(pts PointVector, pt image.Point, measureDist bool) float64 {
    rs0 := gocv.PointPolygonTest(pts.PointVector, pt, measureDist)
    return rs0
}


func (g *GoCVResourceTracker) ConnectedComponents(src Mat, labels *Mat) int {
    rs0 := gocv.ConnectedComponents(src.Mat, &(labels.Mat))
    return rs0
}


func (g *GoCVResourceTracker) ConnectedComponentsWithParams(src Mat, labels *Mat, conn int, ltype MatType,	ccltype ConnectedComponentsAlgorithmType) int {
    rs0 := gocv.ConnectedComponentsWithParams(src.Mat, &(labels.Mat), conn, ltype, 	ccltype)
    return rs0
}


func (g *GoCVResourceTracker) ConnectedComponentsWithStats(src Mat, labels *Mat, stats *Mat, centroids *Mat) int {
    rs0 := gocv.ConnectedComponentsWithStats(src.Mat, &(labels.Mat), &(stats.Mat), &(centroids.Mat))
    return rs0
}


func (g *GoCVResourceTracker) ConnectedComponentsWithStatsWithParams(src Mat, labels *Mat, stats *Mat, centroids *Mat,	conn int, ltype MatType, ccltype ConnectedComponentsAlgorithmType) int {
    rs0 := gocv.ConnectedComponentsWithStatsWithParams(src.Mat, &(labels.Mat), &(stats.Mat), &(centroids.Mat), 	conn, ltype, ccltype)
    return rs0
}


func (g *GoCVResourceTracker) MatchTemplate(image Mat, templ Mat, result *Mat, method TemplateMatchMode, mask Mat) {
    gocv.MatchTemplate(image.Mat, templ.Mat, &(result.Mat), method, mask.Mat)
}


func (g *GoCVResourceTracker) Moments(src Mat, binaryImage bool) map[string]float64 {
    rs0 := gocv.Moments(src.Mat, binaryImage)
    return rs0
}


func (g *GoCVResourceTracker) PyrDown(src Mat, dst *Mat, ksize image.Point, borderType BorderType) {
    gocv.PyrDown(src.Mat, &(dst.Mat), ksize, borderType)
}


func (g *GoCVResourceTracker) PyrUp(src Mat, dst *Mat, ksize image.Point, borderType BorderType) {
    gocv.PyrUp(src.Mat, &(dst.Mat), ksize, borderType)
}


func (g *GoCVResourceTracker) MorphologyDefaultBorderValue() Scalar {
    rs0 := gocv.MorphologyDefaultBorderValue()
    return rs0
}


func (g *GoCVResourceTracker) MorphologyEx(src Mat, dst *Mat, op MorphType, kernel Mat) {
    gocv.MorphologyEx(src.Mat, &(dst.Mat), op, kernel.Mat)
}


func (g *GoCVResourceTracker) MorphologyExWithParams(src Mat, dst *Mat, op MorphType, kernel Mat, iterations int, borderType BorderType) {
    gocv.MorphologyExWithParams(src.Mat, &(dst.Mat), op, kernel.Mat, iterations, borderType)
}


func (g *GoCVResourceTracker) GetStructuringElement(shape MorphShape, ksize image.Point) Mat {
    rs0 := gocv.GetStructuringElement(shape, ksize)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) GaussianBlur(src Mat, dst *Mat, ksize image.Point, sigmaX float64,	sigmaY float64, borderType BorderType) {
    gocv.GaussianBlur(src.Mat, &(dst.Mat), ksize, sigmaX, 	sigmaY, borderType)
}


func (g *GoCVResourceTracker) GetGaussianKernel(ksize int, sigma float64) Mat {
    rs0 := gocv.GetGaussianKernel(ksize, sigma)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) GetGaussianKernelWithParams(ksize int, sigma float64, ktype MatType) Mat {
    rs0 := gocv.GetGaussianKernelWithParams(ksize, sigma, ktype)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) Sobel(src Mat, dst *Mat, ddepth MatType, dx, dy, ksize int, scale, delta float64, borderType BorderType) {
    gocv.Sobel(src.Mat, &(dst.Mat), ddepth, dx, dy, ksize, scale, delta, borderType)
}


func (g *GoCVResourceTracker) SpatialGradient(src Mat, dx, dy *Mat, ksize MatType, borderType BorderType) {
    gocv.SpatialGradient(src.Mat, &(dx.Mat), &(dy.Mat), ksize, borderType)
}


func (g *GoCVResourceTracker) Laplacian(src Mat, dst *Mat, dDepth MatType, size int, scale float64,	delta float64, borderType BorderType) {
    gocv.Laplacian(src.Mat, &(dst.Mat), dDepth, size, scale, 	delta, borderType)
}


func (g *GoCVResourceTracker) Scharr(src Mat, dst *Mat, dDepth MatType, dx int, dy int, scale float64,	delta float64, borderType BorderType) {
    gocv.Scharr(src.Mat, &(dst.Mat), dDepth, dx, dy, scale, 	delta, borderType)
}


func (g *GoCVResourceTracker) MedianBlur(src Mat, dst *Mat, ksize int) {
    gocv.MedianBlur(src.Mat, &(dst.Mat), ksize)
}


func (g *GoCVResourceTracker) Canny(src Mat, edges *Mat, t1 float32, t2 float32) {
    gocv.Canny(src.Mat, &(edges.Mat), t1, t2)
}


func (g *GoCVResourceTracker) CornerSubPix(img Mat, corners *Mat, winSize image.Point, zeroZone image.Point, criteria TermCriteria) {
    gocv.CornerSubPix(img.Mat, &(corners.Mat), winSize, zeroZone, criteria)
}


func (g *GoCVResourceTracker) GoodFeaturesToTrack(img Mat, corners *Mat, maxCorners int, quality float64, minDist float64) {
    gocv.GoodFeaturesToTrack(img.Mat, &(corners.Mat), maxCorners, quality, minDist)
}


func (g *GoCVResourceTracker) GrabCut(img Mat, mask *Mat, r image.Rectangle, bgdModel *Mat, fgdModel *Mat, iterCount int, mode GrabCutMode) {
    gocv.GrabCut(img.Mat, &(mask.Mat), r, &(bgdModel.Mat), &(fgdModel.Mat), iterCount, mode)
}


func (g *GoCVResourceTracker) HoughCircles(src Mat, circles *Mat, method HoughMode, dp, minDist float64) {
    gocv.HoughCircles(src.Mat, &(circles.Mat), method, dp, minDist)
}


func (g *GoCVResourceTracker) HoughCirclesWithParams(src Mat, circles *Mat, method HoughMode, dp, minDist, param1, param2 float64, minRadius, maxRadius int) {
    gocv.HoughCirclesWithParams(src.Mat, &(circles.Mat), method, dp, minDist, param1, param2, minRadius, maxRadius)
}


func (g *GoCVResourceTracker) HoughLines(src Mat, lines *Mat, rho float32, theta float32, threshold int) {
    gocv.HoughLines(src.Mat, &(lines.Mat), rho, theta, threshold)
}


func (g *GoCVResourceTracker) HoughLinesP(src Mat, lines *Mat, rho float32, theta float32, threshold int) {
    gocv.HoughLinesP(src.Mat, &(lines.Mat), rho, theta, threshold)
}


func (g *GoCVResourceTracker) HoughLinesPWithParams(src Mat, lines *Mat, rho float32, theta float32, threshold int, minLineLength float32, maxLineGap float32) {
    gocv.HoughLinesPWithParams(src.Mat, &(lines.Mat), rho, theta, threshold, minLineLength, maxLineGap)
}


func (g *GoCVResourceTracker) HoughLinesPointSet(points Mat, lines *Mat, linesMax int, threshold int,	minRho float32, maxRho float32, rhoStep float32,	minTheta float32, maxTheta float32, thetaStep float32) {
    gocv.HoughLinesPointSet(points.Mat, &(lines.Mat), linesMax, threshold, 	minRho, maxRho, rhoStep, 	minTheta, maxTheta, thetaStep)
}


func (g *GoCVResourceTracker) Integral(src Mat, sum *Mat, sqsum *Mat, tilted *Mat) {
    gocv.Integral(src.Mat, &(sum.Mat), &(sqsum.Mat), &(tilted.Mat))
}


func (g *GoCVResourceTracker) Threshold(src Mat, dst *Mat, thresh float32, maxvalue float32, typ ThresholdType) (threshold float32) {
    rs0 := gocv.Threshold(src.Mat, &(dst.Mat), thresh, maxvalue, typ)
    return rs0
}


func (g *GoCVResourceTracker) AdaptiveThreshold(src Mat, dst *Mat, maxValue float32, adaptiveTyp AdaptiveThresholdType, typ ThresholdType, blockSize int, c float32) {
    gocv.AdaptiveThreshold(src.Mat, &(dst.Mat), maxValue, adaptiveTyp, typ, blockSize, c)
}


func (g *GoCVResourceTracker) ArrowedLine(img *Mat, pt1 image.Point, pt2 image.Point, c color.RGBA, thickness int) {
    gocv.ArrowedLine(&(img.Mat), pt1, pt2, c, thickness)
}


func (g *GoCVResourceTracker) Circle(img *Mat, center image.Point, radius int, c color.RGBA, thickness int) {
    gocv.Circle(&(img.Mat), center, radius, c, thickness)
}


func (g *GoCVResourceTracker) CircleWithParams(img *Mat, center image.Point, radius int, c color.RGBA, thickness int, lineType LineType, shift int) {
    gocv.CircleWithParams(&(img.Mat), center, radius, c, thickness, lineType, shift)
}


func (g *GoCVResourceTracker) Ellipse(img *Mat, center, axes image.Point, angle, startAngle, endAngle float64, c color.RGBA, thickness int) {
    gocv.Ellipse(&(img.Mat), center, axes, angle, startAngle, endAngle, c, thickness)
}


func (g *GoCVResourceTracker) EllipseWithParams(img *Mat, center, axes image.Point, angle, startAngle, endAngle float64, c color.RGBA, thickness int, lineType LineType, shift int) {
    gocv.EllipseWithParams(&(img.Mat), center, axes, angle, startAngle, endAngle, c, thickness, lineType, shift)
}


func (g *GoCVResourceTracker) Line(img *Mat, pt1 image.Point, pt2 image.Point, c color.RGBA, thickness int) {
    gocv.Line(&(img.Mat), pt1, pt2, c, thickness)
}


func (g *GoCVResourceTracker) Rectangle(img *Mat, r image.Rectangle, c color.RGBA, thickness int) {
    gocv.Rectangle(&(img.Mat), r, c, thickness)
}


func (g *GoCVResourceTracker) RectangleWithParams(img *Mat, r image.Rectangle, c color.RGBA, thickness int, lineType LineType, shift int) {
    gocv.RectangleWithParams(&(img.Mat), r, c, thickness, lineType, shift)
}


func (g *GoCVResourceTracker) FillPoly(img *Mat, pts PointsVector, c color.RGBA) {
    gocv.FillPoly(&(img.Mat), pts.PointsVector, c)
}


func (g *GoCVResourceTracker) FillPolyWithParams(img *Mat, pts PointsVector, c color.RGBA, lineType LineType, shift int, offset image.Point) {
    gocv.FillPolyWithParams(&(img.Mat), pts.PointsVector, c, lineType, shift, offset)
}


func (g *GoCVResourceTracker) Polylines(img *Mat, pts PointsVector, isClosed bool, c color.RGBA, thickness int) {
    gocv.Polylines(&(img.Mat), pts.PointsVector, isClosed, c, thickness)
}


func (g *GoCVResourceTracker) GetTextSize(text string, fontFace HersheyFont, fontScale float64, thickness int) image.Point {
    rs0 := gocv.GetTextSize(text, fontFace, fontScale, thickness)
    return rs0
}


func (g *GoCVResourceTracker) GetTextSizeWithBaseline(text string, fontFace HersheyFont, fontScale float64, thickness int) (image.Point, int) {
    rs0, rs1 := gocv.GetTextSizeWithBaseline(text, fontFace, fontScale, thickness)
    return rs0, rs1
}


func (g *GoCVResourceTracker) PutText(img *Mat, text string, org image.Point, fontFace HersheyFont, fontScale float64, c color.RGBA, thickness int) {
    gocv.PutText(&(img.Mat), text, org, fontFace, fontScale, c, thickness)
}


func (g *GoCVResourceTracker) PutTextWithParams(img *Mat, text string, org image.Point, fontFace HersheyFont, fontScale float64, c color.RGBA, thickness int, lineType LineType, bottomLeftOrigin bool) {
    gocv.PutTextWithParams(&(img.Mat), text, org, fontFace, fontScale, c, thickness, lineType, bottomLeftOrigin)
}


func (g *GoCVResourceTracker) Resize(src Mat, dst *Mat, sz image.Point, fx, fy float64, interp InterpolationFlags) {
    gocv.Resize(src.Mat, &(dst.Mat), sz, fx, fy, interp)
}


func (g *GoCVResourceTracker) GetRectSubPix(src Mat, patchSize image.Point, center image.Point, dst *Mat) {
    gocv.GetRectSubPix(src.Mat, patchSize, center, &(dst.Mat))
}


func (g *GoCVResourceTracker) GetRotationMatrix2D(center image.Point, angle, scale float64) Mat {
    rs0 := gocv.GetRotationMatrix2D(center, angle, scale)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) WarpAffine(src Mat, dst *Mat, m Mat, sz image.Point) {
    gocv.WarpAffine(src.Mat, &(dst.Mat), m.Mat, sz)
}


func (g *GoCVResourceTracker) WarpAffineWithParams(src Mat, dst *Mat, m Mat, sz image.Point, flags InterpolationFlags, borderType BorderType, borderValue color.RGBA) {
    gocv.WarpAffineWithParams(src.Mat, &(dst.Mat), m.Mat, sz, flags, borderType, borderValue)
}


func (g *GoCVResourceTracker) WarpPerspective(src Mat, dst *Mat, m Mat, sz image.Point) {
    gocv.WarpPerspective(src.Mat, &(dst.Mat), m.Mat, sz)
}


func (g *GoCVResourceTracker) Watershed(image Mat, markers *Mat) {
    gocv.Watershed(image.Mat, &(markers.Mat))
}


func (g *GoCVResourceTracker) ApplyColorMap(src Mat, dst *Mat, colormapType ColormapTypes) {
    gocv.ApplyColorMap(src.Mat, &(dst.Mat), colormapType)
}


func (g *GoCVResourceTracker) ApplyCustomColorMap(src Mat, dst *Mat, customColormap Mat) {
    gocv.ApplyCustomColorMap(src.Mat, &(dst.Mat), customColormap.Mat)
}


func (g *GoCVResourceTracker) GetPerspectiveTransform(src, dst PointVector) Mat {
    rs0 := gocv.GetPerspectiveTransform(src.PointVector, dst.PointVector)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) GetPerspectiveTransform2f(src, dst Point2fVector) Mat {
    rs0 := gocv.GetPerspectiveTransform2f(src.Point2fVector, dst.Point2fVector)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) GetAffineTransform(src, dst PointVector) Mat {
    rs0 := gocv.GetAffineTransform(src.PointVector, dst.PointVector)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) GetAffineTransform2f(src, dst Point2fVector) Mat {
    rs0 := gocv.GetAffineTransform2f(src.Point2fVector, dst.Point2fVector)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) FindHomography(srcPoints Mat, dstPoints *Mat, method HomographyMethod, ransacReprojThreshold float64, mask *Mat, maxIters int, confidence float64) Mat {
    rs0 := gocv.FindHomography(srcPoints.Mat, &(dstPoints.Mat), method, ransacReprojThreshold, &(mask.Mat), maxIters, confidence)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) DrawContours(img *Mat, contours PointsVector, contourIdx int, c color.RGBA, thickness int) {
    gocv.DrawContours(&(img.Mat), contours.PointsVector, contourIdx, c, thickness)
}


func (g *GoCVResourceTracker) Remap(src Mat, dst, map1, map2 *Mat, interpolation InterpolationFlags, borderMode BorderType, borderValue color.RGBA) {
    gocv.Remap(src.Mat, &(dst.Mat), &(map1.Mat), &(map2.Mat), interpolation, borderMode, borderValue)
}


func (g *GoCVResourceTracker) Filter2D(src Mat, dst *Mat, ddepth MatType, kernel Mat, anchor image.Point, delta float64, borderType BorderType) {
    gocv.Filter2D(src.Mat, &(dst.Mat), ddepth, kernel.Mat, anchor, delta, borderType)
}


func (g *GoCVResourceTracker) SepFilter2D(src Mat, dst *Mat, ddepth MatType, kernelX, kernelY Mat, anchor image.Point, delta float64, borderType BorderType) {
    gocv.SepFilter2D(src.Mat, &(dst.Mat), ddepth, kernelX.Mat, kernelY.Mat, anchor, delta, borderType)
}


func (g *GoCVResourceTracker) LogPolar(src Mat, dst *Mat, center image.Point, m float64, flags InterpolationFlags) {
    gocv.LogPolar(src.Mat, &(dst.Mat), center, m, flags)
}


func (g *GoCVResourceTracker) LinearPolar(src Mat, dst *Mat, center image.Point, maxRadius float64, flags InterpolationFlags) {
    gocv.LinearPolar(src.Mat, &(dst.Mat), center, maxRadius, flags)
}


func (g *GoCVResourceTracker) FitLine(pts PointVector, line *Mat, distType DistanceTypes, param, reps, aeps float64) {
    gocv.FitLine(pts.PointVector, &(line.Mat), distType, param, reps, aeps)
}


func (g *GoCVResourceTracker) NewCLAHE() CLAHE {
    rs0 := gocv.NewCLAHE()
    g.TrackCloseError(&rs0)
    pkg0 := CLAHE{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewCLAHEWithParams(clipLimit float64, tileGridSize image.Point) CLAHE {
    rs0 := gocv.NewCLAHEWithParams(clipLimit, tileGridSize)
    g.TrackCloseError(&rs0)
    pkg0 := CLAHE{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) InvertAffineTransform(src Mat, dst *Mat) {
    gocv.InvertAffineTransform(src.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) PhaseCorrelate(src1, src2, window Mat) (phaseShift Point2f, response float64) {
    rs0, rs1 := gocv.PhaseCorrelate(src1.Mat, src2.Mat, window.Mat)
    return rs0, rs1
}


func (g *GoCVResourceTracker) ImageToMatRGBA(img image.Image) (Mat, error) {
    rs0, rs1 := gocv.ImageToMatRGBA(img)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) ImageToMatRGB(img image.Image) (Mat, error) {
    rs0, rs1 := gocv.ImageToMatRGB(img)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) ImageGrayToMatGray(img *image.Gray) (Mat, error) {
    rs0, rs1 := gocv.ImageGrayToMatGray(img)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) Accumulate(src Mat, dst *Mat) {
    gocv.Accumulate(src.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) AccumulateWithMask(src Mat, dst *Mat, mask Mat) {
    gocv.AccumulateWithMask(src.Mat, &(dst.Mat), mask.Mat)
}


func (g *GoCVResourceTracker) AccumulateSquare(src Mat, dst *Mat) {
    gocv.AccumulateSquare(src.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) AccumulateSquareWithMask(src Mat, dst *Mat, mask Mat) {
    gocv.AccumulateSquareWithMask(src.Mat, &(dst.Mat), mask.Mat)
}


func (g *GoCVResourceTracker) AccumulateProduct(src1 Mat, src2 Mat, dst *Mat) {
    gocv.AccumulateProduct(src1.Mat, src2.Mat, &(dst.Mat))
}


func (g *GoCVResourceTracker) AccumulateProductWithMask(src1 Mat, src2 Mat, dst *Mat, mask Mat) {
    gocv.AccumulateProductWithMask(src1.Mat, src2.Mat, &(dst.Mat), mask.Mat)
}


func (g *GoCVResourceTracker) AccumulatedWeighted(src Mat, dst *Mat, alpha float64) {
    gocv.AccumulatedWeighted(src.Mat, &(dst.Mat), alpha)
}


func (g *GoCVResourceTracker) AccumulatedWeightedWithMask(src Mat, dst *Mat, alpha float64, mask Mat) {
    gocv.AccumulatedWeightedWithMask(src.Mat, &(dst.Mat), alpha, mask.Mat)
}
func (c *CLAHE) Apply(src Mat, dst *Mat) {
c.CLAHE.Apply(src.Mat, &(dst.Mat))
}