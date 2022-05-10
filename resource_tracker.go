package gocv_resource_tracker

import (
	"github.com/garfeng/gocv_resource_tracker/tracker"
)

type GoCVResourceTracker struct {
	*tracker.ResourceTracker
}

// NewAutoGCTracker returns a GoCVResourceTracker, you should call Close() by manual.
func NewAutoGCTracker() *GoCVResourceTracker {
	return &GoCVResourceTracker{
		ResourceTracker: tracker.NewAutoGCResourceTracker(),
	}
}

// NewTracker Returns a GoCVResourceTracker with runtime.SetFinalizer,
// Close() is not required, but you should be careful when deal with it.
func NewTracker() *GoCVResourceTracker {
	return &GoCVResourceTracker{
		ResourceTracker: tracker.NewResourceTracker(),
	}
}
