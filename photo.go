
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
type AlignMTB struct {
    *gocv.AlignMTB
	ResourceTracker *GoCVResourceTracker
}
type EdgeFilter = gocv.EdgeFilter

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


func (g *GoCVResourceTracker) FastNlMeansDenoising(src Mat, dst *Mat) {
    gocv.FastNlMeansDenoising(*(src.Mat), (dst.Mat))
}


func (g *GoCVResourceTracker) FastNlMeansDenoisingWithParams(src Mat, dst *Mat, h float32, templateWindowSize int, searchWindowSize int) {
    gocv.FastNlMeansDenoisingWithParams(*(src.Mat), (dst.Mat), h, templateWindowSize, searchWindowSize)
}


func (g *GoCVResourceTracker) FastNlMeansDenoisingColored(src Mat, dst *Mat) {
    gocv.FastNlMeansDenoisingColored(*(src.Mat), (dst.Mat))
}


func (g *GoCVResourceTracker) FastNlMeansDenoisingColoredWithParams(src Mat, dst *Mat, h float32, hColor float32, templateWindowSize int, searchWindowSize int) {
    gocv.FastNlMeansDenoisingColoredWithParams(*(src.Mat), (dst.Mat), h, hColor, templateWindowSize, searchWindowSize)
}


func (g *GoCVResourceTracker) DetailEnhance(src Mat, dst *Mat, sigma_s, sigma_r float32) {
    gocv.DetailEnhance(*(src.Mat), (dst.Mat), sigma_s, sigma_r)
}


func (g *GoCVResourceTracker) EdgePreservingFilter(src Mat, dst *Mat, filter EdgeFilter, sigma_s, sigma_r float32) {
    gocv.EdgePreservingFilter(*(src.Mat), (dst.Mat), filter, sigma_s, sigma_r)
}


func (g *GoCVResourceTracker) PencilSketch(src Mat, dst1, dst2 *Mat, sigma_s, sigma_r, shade_factor float32) {
    gocv.PencilSketch(*(src.Mat), (dst1.Mat), (dst2.Mat), sigma_s, sigma_r, shade_factor)
}


func (g *GoCVResourceTracker) Stylization(src Mat, dst *Mat, sigma_s, sigma_r float32) {
    gocv.Stylization(*(src.Mat), (dst.Mat), sigma_s, sigma_r)
}
func (b *MergeMertens) Process(src []Mat, dst *Mat) {
    b.MergeMertens.Process(SliceToGoCVCloser(src), (dst.Mat))
}