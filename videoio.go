
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
)

type VideoCaptureAPI = gocv.VideoCaptureAPI
type VideoCaptureProperties = gocv.VideoCaptureProperties
type VideoCapture struct {
    *gocv.VideoCapture
	ResourceTracker *GoCVResourceTracker
}
func (v *VideoCapture) Close(){}

type VideoWriter struct {
    *gocv.VideoWriter
	ResourceTracker *GoCVResourceTracker
}
func (v *VideoWriter) Close(){}


func (g *GoCVResourceTracker) VideoCaptureFile(uri string) (vc *VideoCapture, err error) {
    rs0, rs1 := gocv.VideoCaptureFile(uri)
    g.TrackCloseError(rs0)
    pkg0 := &VideoCapture{
	    &*rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) VideoCaptureFileWithAPI(uri string, apiPreference VideoCaptureAPI) (vc *VideoCapture, err error) {
    rs0, rs1 := gocv.VideoCaptureFileWithAPI(uri, apiPreference)
    g.TrackCloseError(rs0)
    pkg0 := &VideoCapture{
	    &*rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) VideoCaptureDevice(device int) (vc *VideoCapture, err error) {
    rs0, rs1 := gocv.VideoCaptureDevice(device)
    g.TrackCloseError(rs0)
    pkg0 := &VideoCapture{
	    &*rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) VideoCaptureDeviceWithAPI(device int, apiPreference VideoCaptureAPI) (vc *VideoCapture, err error) {
    rs0, rs1 := gocv.VideoCaptureDeviceWithAPI(device, apiPreference)
    g.TrackCloseError(rs0)
    pkg0 := &VideoCapture{
	    &*rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) VideoWriterFile(name string, codec string, fps float64, width int, height int, isColor bool) (vw *VideoWriter, err error) {
    rs0, rs1 := gocv.VideoWriterFile(name, codec, fps, width, height, isColor)
    g.TrackCloseError(rs0)
    pkg0 := &VideoWriter{
	    &*rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) OpenVideoCapture(v interface{}) (*VideoCapture, error) {
    rs0, rs1 := gocv.OpenVideoCapture(v)
    g.TrackCloseError(rs0)
    pkg0 := &VideoCapture{
	    &*rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) OpenVideoCaptureWithAPI(v interface{}, apiPreference VideoCaptureAPI) (*VideoCapture, error) {
    rs0, rs1 := gocv.OpenVideoCaptureWithAPI(v, apiPreference)
    g.TrackCloseError(rs0)
    pkg0 := &VideoCapture{
	    &*rs0,
	    g,
    }
    return pkg0, rs1
}
func (v *VideoCapture) Read(m *Mat) bool {
    rs0 := v.VideoCapture.Read((m.Mat))
    return rs0
}
func (vw *VideoWriter) Write(img Mat) error {
    rs0 := vw.VideoWriter.Write(*(img.Mat))
    return rs0
}