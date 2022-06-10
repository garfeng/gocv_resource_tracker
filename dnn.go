
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
    "image"
)

type Net struct {
    *gocv.Net
	ResourceTracker *GoCVResourceTracker
}
func (n *Net) Close(){}

type NetBackendType = gocv.NetBackendType
type NetTargetType = gocv.NetTargetType
type Layer struct {
    *gocv.Layer
	ResourceTracker *GoCVResourceTracker
}
func (l *Layer) Close(){}


func (g *GoCVResourceTracker) ParseNetBackend(backend string) NetBackendType {
    rs0 := gocv.ParseNetBackend(backend)
    return rs0
}


func (g *GoCVResourceTracker) ParseNetTarget(target string) NetTargetType {
    rs0 := gocv.ParseNetTarget(target)
    return rs0
}


func (g *GoCVResourceTracker) ReadNet(model string, config string) Net {
    rs0 := gocv.ReadNet(model, config)
    g.TrackCloseError(&rs0)
    pkg0 := Net{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) ReadNetBytes(framework string, model []byte, config []byte) (Net, error) {
    rs0, rs1 := gocv.ReadNetBytes(framework, model, config)
    g.TrackCloseError(&rs0)
    pkg0 := Net{
	    &rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) ReadNetFromCaffe(prototxt string, caffeModel string) Net {
    rs0 := gocv.ReadNetFromCaffe(prototxt, caffeModel)
    g.TrackCloseError(&rs0)
    pkg0 := Net{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) ReadNetFromCaffeBytes(prototxt []byte, caffeModel []byte) (Net, error) {
    rs0, rs1 := gocv.ReadNetFromCaffeBytes(prototxt, caffeModel)
    g.TrackCloseError(&rs0)
    pkg0 := Net{
	    &rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) ReadNetFromTensorflow(model string) Net {
    rs0 := gocv.ReadNetFromTensorflow(model)
    g.TrackCloseError(&rs0)
    pkg0 := Net{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) ReadNetFromTensorflowBytes(model []byte) (Net, error) {
    rs0, rs1 := gocv.ReadNetFromTensorflowBytes(model)
    g.TrackCloseError(&rs0)
    pkg0 := Net{
	    &rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) ReadNetFromTorch(model string) Net {
    rs0 := gocv.ReadNetFromTorch(model)
    g.TrackCloseError(&rs0)
    pkg0 := Net{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) ReadNetFromONNX(model string) Net {
    rs0 := gocv.ReadNetFromONNX(model)
    g.TrackCloseError(&rs0)
    pkg0 := Net{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) ReadNetFromONNXBytes(model []byte) (Net, error) {
    rs0, rs1 := gocv.ReadNetFromONNXBytes(model)
    g.TrackCloseError(&rs0)
    pkg0 := Net{
	    &rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) BlobFromImage(img Mat, scaleFactor float64, size image.Point, mean Scalar,	swapRB bool, crop bool) Mat {
    rs0 := gocv.BlobFromImage(*(img.Mat), scaleFactor, size, mean, 	swapRB, crop)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) BlobFromImages(imgs []Mat, blob *Mat, scaleFactor float64, size image.Point, mean Scalar,	swapRB bool, crop bool, ddepth MatType) {
    gocv.BlobFromImages(SliceToGoCVCloser(imgs), (blob.Mat), scaleFactor, size, mean, 	swapRB, crop, ddepth)
}


func (g *GoCVResourceTracker) ImagesFromBlob(blob Mat, imgs []Mat) {
    gocv.ImagesFromBlob(*(blob.Mat), SliceToGoCVCloser(imgs))
}


func (g *GoCVResourceTracker) GetBlobChannel(blob Mat, imgidx int, chnidx int) Mat {
    rs0 := gocv.GetBlobChannel(*(blob.Mat), imgidx, chnidx)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) GetBlobSize(blob Mat) Scalar {
    rs0 := gocv.GetBlobSize(*(blob.Mat))
    return rs0
}


func (g *GoCVResourceTracker) NMSBoxes(bboxes []image.Rectangle, scores []float32, scoreThreshold float32, nmsThreshold float32, indices []int) {
    gocv.NMSBoxes(bboxes, scores, scoreThreshold, nmsThreshold, indices)
}


func (g *GoCVResourceTracker) NMSBoxesWithParams(bboxes []image.Rectangle, scores []float32, scoreThreshold float32, nmsThreshold float32, indices []int, eta float32, topK int) {
    gocv.NMSBoxesWithParams(bboxes, scores, scoreThreshold, nmsThreshold, indices, eta, topK)
}
func (net *Net) Forward(outputName string) Mat {
    rs0 := net.Net.Forward(outputName)

    net.ResourceTracker.TrackCloseError(&rs0)
    return Mat{&rs0, net.ResourceTracker}
}

func (net *Net) GetLayer(layer int) Layer {
    rs0 := net.Net.GetLayer(layer)

    net.ResourceTracker.TrackCloseError(&rs0)
    return Layer{&rs0, net.ResourceTracker}
}

func (net *Net) SetInput(blob Mat, name string) {
    net.Net.SetInput(*(blob.Mat), name)
}