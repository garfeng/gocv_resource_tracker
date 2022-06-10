
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
    "image/color"
)

type AKAZE struct {
    *gocv.AKAZE
	ResourceTracker *GoCVResourceTracker
}
func (a *AKAZE) Close(){}

type AgastFeatureDetector struct {
    *gocv.AgastFeatureDetector
	ResourceTracker *GoCVResourceTracker
}
func (a *AgastFeatureDetector) Close(){}

type BRISK struct {
    *gocv.BRISK
	ResourceTracker *GoCVResourceTracker
}
func (b *BRISK) Close(){}

type FastFeatureDetectorType = gocv.FastFeatureDetectorType
type FastFeatureDetector struct {
    *gocv.FastFeatureDetector
	ResourceTracker *GoCVResourceTracker
}
func (f *FastFeatureDetector) Close(){}

type GFTTDetector struct {
    *gocv.GFTTDetector
	ResourceTracker *GoCVResourceTracker
}
func (g *GFTTDetector) Close(){}

type KAZE struct {
    *gocv.KAZE
	ResourceTracker *GoCVResourceTracker
}
func (k *KAZE) Close(){}

type MSER struct {
    *gocv.MSER
	ResourceTracker *GoCVResourceTracker
}
func (m *MSER) Close(){}

type ORB struct {
    *gocv.ORB
	ResourceTracker *GoCVResourceTracker
}
func (o *ORB) Close(){}

type ORBScoreType = gocv.ORBScoreType
type SimpleBlobDetector struct {
    *gocv.SimpleBlobDetector
	ResourceTracker *GoCVResourceTracker
}
func (s *SimpleBlobDetector) Close(){}

type SimpleBlobDetectorParams = gocv.SimpleBlobDetectorParams
type BFMatcher struct {
    *gocv.BFMatcher
	ResourceTracker *GoCVResourceTracker
}
func (b *BFMatcher) Close(){}

type FlannBasedMatcher struct {
    *gocv.FlannBasedMatcher
	ResourceTracker *GoCVResourceTracker
}
func (f *FlannBasedMatcher) Close(){}

type DrawMatchesFlag = gocv.DrawMatchesFlag
type SIFT struct {
    *gocv.SIFT
	ResourceTracker *GoCVResourceTracker
}
func (s *SIFT) Close(){}


func (g *GoCVResourceTracker) NewAKAZE() AKAZE {
    rs0 := gocv.NewAKAZE()
    g.TrackCloseError(&rs0)
    pkg0 := AKAZE{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewAgastFeatureDetector() AgastFeatureDetector {
    rs0 := gocv.NewAgastFeatureDetector()
    g.TrackCloseError(&rs0)
    pkg0 := AgastFeatureDetector{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewBRISK() BRISK {
    rs0 := gocv.NewBRISK()
    g.TrackCloseError(&rs0)
    pkg0 := BRISK{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewFastFeatureDetector() FastFeatureDetector {
    rs0 := gocv.NewFastFeatureDetector()
    g.TrackCloseError(&rs0)
    pkg0 := FastFeatureDetector{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewFastFeatureDetectorWithParams(threshold int, nonmaxSuppression bool, typ FastFeatureDetectorType) FastFeatureDetector {
    rs0 := gocv.NewFastFeatureDetectorWithParams(threshold, nonmaxSuppression, typ)
    g.TrackCloseError(&rs0)
    pkg0 := FastFeatureDetector{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewGFTTDetector() GFTTDetector {
    rs0 := gocv.NewGFTTDetector()
    g.TrackCloseError(&rs0)
    pkg0 := GFTTDetector{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewKAZE() KAZE {
    rs0 := gocv.NewKAZE()
    g.TrackCloseError(&rs0)
    pkg0 := KAZE{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewMSER() MSER {
    rs0 := gocv.NewMSER()
    g.TrackCloseError(&rs0)
    pkg0 := MSER{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewORB() ORB {
    rs0 := gocv.NewORB()
    g.TrackCloseError(&rs0)
    pkg0 := ORB{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewORBWithParams(nFeatures int, scaleFactor float32, nLevels int, edgeThreshold int, firstLevel int, WTAK int, scoreType ORBScoreType, patchSize int, fastThreshold int) ORB {
    rs0 := gocv.NewORBWithParams(nFeatures, scaleFactor, nLevels, edgeThreshold, firstLevel, WTAK, scoreType, patchSize, fastThreshold)
    g.TrackCloseError(&rs0)
    pkg0 := ORB{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewSimpleBlobDetector() SimpleBlobDetector {
    rs0 := gocv.NewSimpleBlobDetector()
    g.TrackCloseError(&rs0)
    pkg0 := SimpleBlobDetector{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewSimpleBlobDetectorWithParams(params SimpleBlobDetectorParams) SimpleBlobDetector {
    rs0 := gocv.NewSimpleBlobDetectorWithParams(params)
    g.TrackCloseError(&rs0)
    pkg0 := SimpleBlobDetector{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewSimpleBlobDetectorParams() SimpleBlobDetectorParams {
    rs0 := gocv.NewSimpleBlobDetectorParams()
    return rs0
}


func (g *GoCVResourceTracker) NewBFMatcher() BFMatcher {
    rs0 := gocv.NewBFMatcher()
    g.TrackCloseError(&rs0)
    pkg0 := BFMatcher{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewBFMatcherWithParams(normType NormType, crossCheck bool) BFMatcher {
    rs0 := gocv.NewBFMatcherWithParams(normType, crossCheck)
    g.TrackCloseError(&rs0)
    pkg0 := BFMatcher{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewFlannBasedMatcher() FlannBasedMatcher {
    rs0 := gocv.NewFlannBasedMatcher()
    g.TrackCloseError(&rs0)
    pkg0 := FlannBasedMatcher{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) DrawKeyPoints(src Mat, keyPoints []KeyPoint, dst *Mat, color color.RGBA, flag DrawMatchesFlag) {
    gocv.DrawKeyPoints(*(src.Mat), keyPoints, (dst.Mat), color, flag)
}


func (g *GoCVResourceTracker) NewSIFT() SIFT {
    rs0 := gocv.NewSIFT()
    g.TrackCloseError(&rs0)
    pkg0 := SIFT{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) DrawMatches(img1 Mat, kp1 []KeyPoint, img2 Mat, kp2 []KeyPoint, matches1to2 []DMatch, outImg *Mat, matchColor color.RGBA, singlePointColor color.RGBA, matchesMask []byte, flags DrawMatchesFlag) {
    gocv.DrawMatches(*(img1.Mat), kp1, *(img2.Mat), kp2, matches1to2, (outImg.Mat), matchColor, singlePointColor, matchesMask, flags)
}
func (a *AKAZE) Detect(src Mat) []KeyPoint {
    rs0 := a.AKAZE.Detect(*(src.Mat))
    return rs0
}

func (a *AKAZE) DetectAndCompute(src Mat, mask Mat) ([]KeyPoint, Mat) {
    rs0, rs1 := a.AKAZE.DetectAndCompute(*(src.Mat), *(mask.Mat))

    a.ResourceTracker.TrackCloseError(&rs1)
    return rs0, Mat{&rs1, a.ResourceTracker}
}
func (a *AgastFeatureDetector) Detect(src Mat) []KeyPoint {
    rs0 := a.AgastFeatureDetector.Detect(*(src.Mat))
    return rs0
}
func (b *BRISK) Detect(src Mat) []KeyPoint {
    rs0 := b.BRISK.Detect(*(src.Mat))
    return rs0
}

func (b *BRISK) DetectAndCompute(src Mat, mask Mat) ([]KeyPoint, Mat) {
    rs0, rs1 := b.BRISK.DetectAndCompute(*(src.Mat), *(mask.Mat))

    b.ResourceTracker.TrackCloseError(&rs1)
    return rs0, Mat{&rs1, b.ResourceTracker}
}
func (f *FastFeatureDetector) Detect(src Mat) []KeyPoint {
    rs0 := f.FastFeatureDetector.Detect(*(src.Mat))
    return rs0
}
func (a *GFTTDetector) Detect(src Mat) []KeyPoint {
    rs0 := a.GFTTDetector.Detect(*(src.Mat))
    return rs0
}
func (a *KAZE) Detect(src Mat) []KeyPoint {
    rs0 := a.KAZE.Detect(*(src.Mat))
    return rs0
}

func (a *KAZE) DetectAndCompute(src Mat, mask Mat) ([]KeyPoint, Mat) {
    rs0, rs1 := a.KAZE.DetectAndCompute(*(src.Mat), *(mask.Mat))

    a.ResourceTracker.TrackCloseError(&rs1)
    return rs0, Mat{&rs1, a.ResourceTracker}
}
func (a *MSER) Detect(src Mat) []KeyPoint {
    rs0 := a.MSER.Detect(*(src.Mat))
    return rs0
}
func (o *ORB) Detect(src Mat) []KeyPoint {
    rs0 := o.ORB.Detect(*(src.Mat))
    return rs0
}

func (o *ORB) DetectAndCompute(src Mat, mask Mat) ([]KeyPoint, Mat) {
    rs0, rs1 := o.ORB.DetectAndCompute(*(src.Mat), *(mask.Mat))

    o.ResourceTracker.TrackCloseError(&rs1)
    return rs0, Mat{&rs1, o.ResourceTracker}
}
func (b *SimpleBlobDetector) Detect(src Mat) []KeyPoint {
    rs0 := b.SimpleBlobDetector.Detect(*(src.Mat))
    return rs0
}
func (b *BFMatcher) KnnMatch(query, train Mat, k int) [][]DMatch {
    rs0 := b.BFMatcher.KnnMatch(*(query.Mat), *(train.Mat), k)
    return rs0
}
func (f *FlannBasedMatcher) KnnMatch(query, train Mat, k int) [][]DMatch {
    rs0 := f.FlannBasedMatcher.KnnMatch(*(query.Mat), *(train.Mat), k)
    return rs0
}
func (d *SIFT) Detect(src Mat) []KeyPoint {
    rs0 := d.SIFT.Detect(*(src.Mat))
    return rs0
}

func (d *SIFT) DetectAndCompute(src Mat, mask Mat) ([]KeyPoint, Mat) {
    rs0, rs1 := d.SIFT.DetectAndCompute(*(src.Mat), *(mask.Mat))

    d.ResourceTracker.TrackCloseError(&rs1)
    return rs0, Mat{&rs1, d.ResourceTracker}
}