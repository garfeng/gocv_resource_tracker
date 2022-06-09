
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
)



func (g *GoCVResourceTracker) Version() string {
    rs0 := gocv.Version()
    return rs0
}


func (g *GoCVResourceTracker) OpenCVVersion() string {
    rs0 := gocv.OpenCVVersion()
    return rs0
}
