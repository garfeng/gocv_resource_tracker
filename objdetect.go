
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
    "image"
)

type CascadeClassifier struct {
    gocv.CascadeClassifier
	ResourceTracker *GoCVResourceTracker
}
type HOGDescriptor struct {
    gocv.HOGDescriptor
	ResourceTracker *GoCVResourceTracker
}
type QRCodeDetector struct {
    gocv.QRCodeDetector
	ResourceTracker *GoCVResourceTracker
}

func (g *GoCVResourceTracker) NewCascadeClassifier() CascadeClassifier {
    rs0 := gocv.NewCascadeClassifier()
    g.TrackCloseError(&rs0)
    pkg0 := CascadeClassifier{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) NewHOGDescriptor() HOGDescriptor {
    rs0 := gocv.NewHOGDescriptor()
    g.TrackCloseError(&rs0)
    pkg0 := HOGDescriptor{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) HOGDefaultPeopleDetector() Mat {
    rs0 := gocv.HOGDefaultPeopleDetector()
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) GroupRectangles(rects []image.Rectangle, groupThreshold int, eps float64) []image.Rectangle {
    rs0 := gocv.GroupRectangles(rects, groupThreshold, eps)
    return rs0
}


func (g *GoCVResourceTracker) NewQRCodeDetector() QRCodeDetector {
    rs0 := gocv.NewQRCodeDetector()
    g.TrackCloseError(&rs0)
    pkg0 := QRCodeDetector{
	    rs0,
	    g,
    }
    return pkg0
}
func (c *CascadeClassifier) DetectMultiScale(img Mat) []image.Rectangle {
    rs0 := c.CascadeClassifier.DetectMultiScale(img.Mat)
    return rs0
}

func (c *CascadeClassifier) DetectMultiScaleWithParams(img Mat, scale float64,
	minNeighbors, flags int, minSize, maxSize image.Point) []image.Rectangle {
    rs0 := c.CascadeClassifier.DetectMultiScaleWithParams(img.Mat, scale, 
	minNeighbors, flags, minSize, maxSize)
    return rs0
}
func (h *HOGDescriptor) DetectMultiScale(img Mat) []image.Rectangle {
    rs0 := h.HOGDescriptor.DetectMultiScale(img.Mat)
    return rs0
}

func (h *HOGDescriptor) DetectMultiScaleWithParams(img Mat, hitThresh float64,
	winStride, padding image.Point, scale, finalThreshold float64, useMeanshiftGrouping bool) []image.Rectangle {
    rs0 := h.HOGDescriptor.DetectMultiScaleWithParams(img.Mat, hitThresh, 
	winStride, padding, scale, finalThreshold, useMeanshiftGrouping)
    return rs0
}

func (h *HOGDescriptor) SetSVMDetector(det Mat) error {
    rs0 := h.HOGDescriptor.SetSVMDetector(det.Mat)
    return rs0
}
func (a *QRCodeDetector) Decode(input Mat, points Mat, straight_qrcode *Mat) string {
    rs0 := a.QRCodeDetector.Decode(input.Mat, points.Mat, &(straight_qrcode.Mat))
    return rs0
}

func (a *QRCodeDetector) Detect(input Mat, points *Mat) bool {
    rs0 := a.QRCodeDetector.Detect(input.Mat, &(points.Mat))
    return rs0
}

func (a *QRCodeDetector) DetectAndDecode(input Mat, points *Mat, straight_qrcode *Mat) string {
    rs0 := a.QRCodeDetector.DetectAndDecode(input.Mat, &(points.Mat), &(straight_qrcode.Mat))
    return rs0
}

func (a *QRCodeDetector) DetectAndDecodeMulti(input Mat, decoded *[]string, points *Mat, qrCodes *[]Mat) bool {
    param4 := SliceToGoCVCloser(*qrCodes)
    rs0 := a.QRCodeDetector.DetectAndDecodeMulti(input.Mat, decoded, &(points.Mat), &param4)
    return rs0
}

func (a *QRCodeDetector) DetectMulti(input Mat, points *Mat) bool {
    rs0 := a.QRCodeDetector.DetectMulti(input.Mat, &(points.Mat))
    return rs0
}