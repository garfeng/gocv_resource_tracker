
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
    "image"
)

type Window struct {
    gocv.Window
	ResourceTracker *GoCVResourceTracker
}
type WindowFlag = gocv.WindowFlag
type WindowPropertyFlag = gocv.WindowPropertyFlag
type Trackbar = gocv.Trackbar

func (g *GoCVResourceTracker) NewWindow(name string) *Window {
    rs0 := gocv.NewWindow(name)
    g.TrackCloseError(rs0)
    pkg0 := &Window{
	    *rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) SelectROI(name string, img Mat) image.Rectangle {
    rs0 := gocv.SelectROI(name, img.Mat)
    return rs0
}


func (g *GoCVResourceTracker) SelectROIs(name string, img Mat) []image.Rectangle {
    rs0 := gocv.SelectROIs(name, img.Mat)
    return rs0
}


func (g *GoCVResourceTracker) WaitKey(delay int) int {
    rs0 := gocv.WaitKey(delay)
    return rs0
}
func (w *Window) IMShow(img Mat) {
    w.Window.IMShow(img.Mat)
}

func (w *Window) SelectROI(img Mat) image.Rectangle {
    rs0 := w.Window.SelectROI(img.Mat)
    return rs0
}

func (w *Window) SelectROIs(img Mat) []image.Rectangle {
    rs0 := w.Window.SelectROIs(img.Mat)
    return rs0
}