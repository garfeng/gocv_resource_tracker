
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
)

type IMReadFlag = gocv.IMReadFlag
type FileExt = gocv.FileExt

func (g *GoCVResourceTracker) IMRead(name string, flags IMReadFlag) Mat {
    rs0 := gocv.IMRead(name, flags)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    &rs0,
	    g,
    }
    return pkg0
}


func (g *GoCVResourceTracker) IMWrite(name string, img Mat) bool {
    rs0 := gocv.IMWrite(name, *(img.Mat))
    return rs0
}


func (g *GoCVResourceTracker) IMWriteWithParams(name string, img Mat, params []int) bool {
    rs0 := gocv.IMWriteWithParams(name, *(img.Mat), params)
    return rs0
}


func (g *GoCVResourceTracker) IMEncode(fileExt FileExt, img Mat) (buf *NativeByteBuffer, err error) {
    rs0, rs1 := gocv.IMEncode(fileExt, *(img.Mat))
    g.TrackCloser(rs0)
    pkg0 := &NativeByteBuffer{
	    &*rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) IMEncodeWithParams(fileExt FileExt, img Mat, params []int) (buf *NativeByteBuffer, err error) {
    rs0, rs1 := gocv.IMEncodeWithParams(fileExt, *(img.Mat), params)
    g.TrackCloser(rs0)
    pkg0 := &NativeByteBuffer{
	    &*rs0,
	    g,
    }
    return pkg0, rs1
}


func (g *GoCVResourceTracker) IMDecode(buf []byte, flags IMReadFlag) (Mat, error) {
    rs0, rs1 := gocv.IMDecode(buf, flags)
    g.TrackCloseError(&rs0)
    pkg0 := Mat{
	    &rs0,
	    g,
    }
    return pkg0, rs1
}
