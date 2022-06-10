package gocv_resource_tracker

import (
	"github.com/garfeng/gocv_resource_tracker/tracker"
	"gocv.io/x/gocv"
)

type GoCVResourceTracker struct {
	*tracker.ResourceTracker
}

// NewTracker Returns a GoCVResourceTracker with runtime.SetFinalizer,
// Close() is not required.
func NewTracker() *GoCVResourceTracker {
	return &GoCVResourceTracker{
		ResourceTracker: tracker.NewResourceTracker(),
	}
}

func SliceToGoCVCloser(slice []Mat) []gocv.Mat {
	res := make([]gocv.Mat, len(slice))

	for i, v := range slice {
		res[i] = *v.Mat
	}
	return res
}

func GoCVCloserToSlice(slice []gocv.Mat, g *GoCVResourceTracker) []Mat {
	res := make([]Mat, len(slice))

	for i, v := range slice {
		g.TrackCloseError(&v)
		res[i] = Mat{
			Mat:             &v,
			ResourceTracker: g,
		}
	}
	return res
}
