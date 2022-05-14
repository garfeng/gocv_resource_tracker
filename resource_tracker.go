package gocv_resource_tracker

import (
	"github.com/garfeng/gocv_resource_tracker/tracker"
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
