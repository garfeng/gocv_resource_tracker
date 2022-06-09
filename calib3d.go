
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
    "image"
)

type CalibFlag = gocv.CalibFlag
type CalibCBFlag = gocv.CalibCBFlag

func (g *GoCVResourceTracker) FisheyeUndistortImage(distorted Mat, undistorted *Mat, k, d Mat) {
    gocv.FisheyeUndistortImage(distorted.Mat, &(undistorted.Mat), k.Mat, d.Mat)
}


func (g *GoCVResourceTracker) FisheyeUndistortImageWithParams(distorted Mat, undistorted *Mat, k, d, knew Mat, size image.Point) {
    gocv.FisheyeUndistortImageWithParams(distorted.Mat, &(undistorted.Mat), k.Mat, d.Mat, knew.Mat, size)
}


func (g *GoCVResourceTracker) FisheyeUndistortPoints(distorted Mat, undistorted *Mat, k, d, r, p Mat) {
    gocv.FisheyeUndistortPoints(distorted.Mat, &(undistorted.Mat), k.Mat, d.Mat, r.Mat, p.Mat)
}


func (g *GoCVResourceTracker) EstimateNewCameraMatrixForUndistortRectify(k, d Mat, imgSize image.Point, r Mat, p *Mat, balance float64, newSize image.Point, fovScale float64) {
    gocv.EstimateNewCameraMatrixForUndistortRectify(k.Mat, d.Mat, imgSize, r.Mat, &(p.Mat), balance, newSize, fovScale)
}


func (g *GoCVResourceTracker) InitUndistortRectifyMap(cameraMatrix Mat, distCoeffs Mat, r Mat, newCameraMatrix Mat, size image.Point, m1type int, map1 Mat, map2 Mat) {
    gocv.InitUndistortRectifyMap(cameraMatrix.Mat, distCoeffs.Mat, r.Mat, newCameraMatrix.Mat, size, m1type, map1.Mat, map2.Mat)
}


func (g *GoCVResourceTracker) GetOptimalNewCameraMatrixWithParams(cameraMatrix Mat, distCoeffs Mat, imageSize image.Point, alpha float64, newImgSize image.Point, centerPrincipalPoint bool) (Mat, image.Rectangle) {
    rs0, rs1 := gocv.GetOptimalNewCameraMatrixWithParams(cameraMatrix.Mat, distCoeffs.Mat, imageSize, alpha, newImgSize, centerPrincipalPoint)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) Undistort(src Mat, dst *Mat, cameraMatrix Mat, distCoeffs Mat, newCameraMatrix Mat) {
    gocv.Undistort(src.Mat, &(dst.Mat), cameraMatrix.Mat, distCoeffs.Mat, newCameraMatrix.Mat)
}


func (g *GoCVResourceTracker) UndistortPoints(src Mat, dst *Mat, cameraMatrix, distCoeffs, rectificationTransform, newCameraMatrix Mat) {
    gocv.UndistortPoints(src.Mat, &(dst.Mat), cameraMatrix.Mat, distCoeffs.Mat, rectificationTransform.Mat, newCameraMatrix.Mat)
}


func (g *GoCVResourceTracker) FindChessboardCorners(image Mat, patternSize image.Point, corners *Mat, flags CalibCBFlag) bool {
    rs0 := gocv.FindChessboardCorners(image.Mat, patternSize, &(corners.Mat), flags)
    return rs0
}


func (g *GoCVResourceTracker) DrawChessboardCorners(image *Mat, patternSize image.Point, corners Mat, patternWasFound bool) {
    gocv.DrawChessboardCorners(&(image.Mat), patternSize, corners.Mat, patternWasFound)
}


func (g *GoCVResourceTracker) EstimateAffinePartial2D(from, to Point2fVector) Mat {
    rs0 := gocv.EstimateAffinePartial2D(from.Point2fVector, to.Point2fVector)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}
