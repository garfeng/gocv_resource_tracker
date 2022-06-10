
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
    "image"
)

type BackgroundSubtractorMOG2 struct {
    *gocv.BackgroundSubtractorMOG2
	ResourceTracker *GoCVResourceTracker
}
func (b *BackgroundSubtractorMOG2) Close(){}

type BackgroundSubtractorKNN struct {
    *gocv.BackgroundSubtractorKNN
	ResourceTracker *GoCVResourceTracker
}
func (b *BackgroundSubtractorKNN) Close(){}

type Tracker struct {
    gocv.Tracker
	ResourceTracker *GoCVResourceTracker
}
func (t *Tracker) Close(){}

type TrackerMIL struct {
    *gocv.TrackerMIL
	ResourceTracker *GoCVResourceTracker
}
func (t *TrackerMIL) Close(){}


func (g *GoCVResourceTracker) NewBackgroundSubtractorMOG2() BackgroundSubtractorMOG2 {
    rs0 := gocv.NewBackgroundSubtractorMOG2()
    g.TrackCloseError(&rs0)
    pkg0 := BackgroundSubtractorMOG2{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewBackgroundSubtractorMOG2WithParams(history int, varThreshold float64, detectShadows bool) BackgroundSubtractorMOG2 {
    rs0 := gocv.NewBackgroundSubtractorMOG2WithParams(history, varThreshold, detectShadows)
    g.TrackCloseError(&rs0)
    pkg0 := BackgroundSubtractorMOG2{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewBackgroundSubtractorKNN() BackgroundSubtractorKNN {
    rs0 := gocv.NewBackgroundSubtractorKNN()
    g.TrackCloseError(&rs0)
    pkg0 := BackgroundSubtractorKNN{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewBackgroundSubtractorKNNWithParams(history int, dist2Threshold float64, detectShadows bool) BackgroundSubtractorKNN {
    rs0 := gocv.NewBackgroundSubtractorKNNWithParams(history, dist2Threshold, detectShadows)
    g.TrackCloseError(&rs0)
    pkg0 := BackgroundSubtractorKNN{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) CalcOpticalFlowFarneback(prevImg Mat, nextImg Mat, flow *Mat, pyrScale float64, levels int, winsize int,	iterations int, polyN int, polySigma float64, flags int) {
    gocv.CalcOpticalFlowFarneback(*(prevImg.Mat), *(nextImg.Mat), (flow.Mat), pyrScale, levels, winsize, 	iterations, polyN, polySigma, flags)
}


func (g *GoCVResourceTracker) CalcOpticalFlowPyrLK(prevImg Mat, nextImg Mat, prevPts Mat, nextPts Mat, status *Mat, err *Mat) {
    gocv.CalcOpticalFlowPyrLK(*(prevImg.Mat), *(nextImg.Mat), *(prevPts.Mat), *(nextPts.Mat), (status.Mat), (err.Mat))
}


func (g *GoCVResourceTracker) CalcOpticalFlowPyrLKWithParams(prevImg Mat, nextImg Mat, prevPts Mat, nextPts Mat, status *Mat, err *Mat,	winSize image.Point, maxLevel int, criteria TermCriteria, flags int, minEigThreshold float64) {
    gocv.CalcOpticalFlowPyrLKWithParams(*(prevImg.Mat), *(nextImg.Mat), *(prevPts.Mat), *(nextPts.Mat), (status.Mat), (err.Mat), 	winSize, maxLevel, criteria, flags, minEigThreshold)
}


func (g *GoCVResourceTracker) FindTransformECC(templateImage Mat, inputImage Mat, warpMatrix *Mat, motionType int, criteria TermCriteria, inputMask Mat, gaussFiltSize int) float64 {
    rs0 := gocv.FindTransformECC(*(templateImage.Mat), *(inputImage.Mat), (warpMatrix.Mat), motionType, criteria, *(inputMask.Mat), gaussFiltSize)
    return rs0
}


func (g *GoCVResourceTracker) NewTrackerMIL() Tracker {
    rs0 := gocv.NewTrackerMIL()
    g.TrackCloseError(rs0)
    pkg0 := Tracker{
	    rs0,
	    g,
    }
    return pkg0
}
func (b *BackgroundSubtractorMOG2) Apply(src Mat, dst *Mat) {
    b.BackgroundSubtractorMOG2.Apply(*(src.Mat), (dst.Mat))
}
func (k *BackgroundSubtractorKNN) Apply(src Mat, dst *Mat) {
    k.BackgroundSubtractorKNN.Apply(*(src.Mat), (dst.Mat))
}
func (trk TrackerMIL) Init(img Mat, boundingBox image.Rectangle) bool {
    rs0 := trk.TrackerMIL.Init(*(img.Mat), boundingBox)
    return rs0
}

func (trk TrackerMIL) Update(img Mat) (image.Rectangle, bool) {
    rs0, rs1 := trk.TrackerMIL.Update(*(img.Mat))
    return rs0, rs1
}