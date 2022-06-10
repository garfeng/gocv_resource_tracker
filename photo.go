
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
    "image"
)

type SeamlessCloneFlags = gocv.SeamlessCloneFlags
type MergeMertens struct {
    *gocv.MergeMertens
	ResourceTracker *GoCVResourceTracker
}
func (m *MergeMertens) Close(){}

type AlignMTB struct {
    *gocv.AlignMTB
	ResourceTracker *GoCVResourceTracker
}
func (a *AlignMTB) Close(){}


func (g *GoCVResourceTracker) ColorChange(src, mask Mat, dst *Mat, red_mul, green_mul, blue_mul float32) {
    gocv.ColorChange(*(src.Mat), *(mask.Mat), (dst.Mat), red_mul, green_mul, blue_mul)
}


func (g *GoCVResourceTracker) SeamlessClone(src, dst, mask Mat, p image.Point, blend *Mat, flags SeamlessCloneFlags) {
    gocv.SeamlessClone(*(src.Mat), *(dst.Mat), *(mask.Mat), p, (blend.Mat), flags)
}


func (g *GoCVResourceTracker) IlluminationChange(src, mask Mat, dst *Mat, alpha, beta float32) {
    gocv.IlluminationChange(*(src.Mat), *(mask.Mat), (dst.Mat), alpha, beta)
}


func (g *GoCVResourceTracker) TextureFlattening(src, mask Mat, dst *Mat, lowThreshold, highThreshold float32, kernelSize int) {
    gocv.TextureFlattening(*(src.Mat), *(mask.Mat), (dst.Mat), lowThreshold, highThreshold, kernelSize)
}


func (g *GoCVResourceTracker) FastNlMeansDenoisingColoredMulti(src []Mat, dst *Mat, imgToDenoiseIndex int, temporalWindowSize int) {
    gocv.FastNlMeansDenoisingColoredMulti(SliceToGoCVCloser(src), (dst.Mat), imgToDenoiseIndex, temporalWindowSize)
}


func (g *GoCVResourceTracker) FastNlMeansDenoisingColoredMultiWithParams(src []Mat, dst *Mat, imgToDenoiseIndex int, temporalWindowSize int, h float32, hColor float32, templateWindowSize int, searchWindowSize int) {
    gocv.FastNlMeansDenoisingColoredMultiWithParams(SliceToGoCVCloser(src), (dst.Mat), imgToDenoiseIndex, temporalWindowSize, h, hColor, templateWindowSize, searchWindowSize)
}


func (g *GoCVResourceTracker) NewMergeMertens() MergeMertens {
    rs0 := gocv.NewMergeMertens()
    g.TrackCloseError(&rs0)
    pkg0 := MergeMertens{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewMergeMertensWithParams(contrast_weight float32, saturation_weight float32, exposure_weight float32) MergeMertens {
    rs0 := gocv.NewMergeMertensWithParams(contrast_weight, saturation_weight, exposure_weight)
    g.TrackCloseError(&rs0)
    pkg0 := MergeMertens{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewAlignMTB() AlignMTB {
    rs0 := gocv.NewAlignMTB()
    g.TrackCloseError(&rs0)
    pkg0 := AlignMTB{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewAlignMTBWithParams(max_bits int, exclude_range int, cut bool) AlignMTB {
    rs0 := gocv.NewAlignMTBWithParams(max_bits, exclude_range, cut)
    g.TrackCloseError(&rs0)
    pkg0 := AlignMTB{
	    &rs0,
	    g,
    }
    return pkg0
}
func (b *MergeMertens) Process(src []Mat, dst *Mat) {
    b.MergeMertens.Process(SliceToGoCVCloser(src), (dst.Mat))
}