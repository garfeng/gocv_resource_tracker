package gocv_resource_tracker

import "gocv.io/x/gocv"

func (g *GoCVResourceTracker) IMRead(name string, flag gocv.IMReadFlag) gocv.Mat {
	mat := gocv.IMRead(name, flag)
	g.TrackCloseError(&mat)
	return mat
}

func (g *GoCVResourceTracker) IMDecode(buf []byte, flags gocv.IMReadFlag) (gocv.Mat, error) {
	mat, err := gocv.IMDecode(buf, flags)
	g.TrackCloseError(&mat)
	return mat, err
}
