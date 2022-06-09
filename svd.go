
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
)



func (g *GoCVResourceTracker) SVDCompute(src Mat, w, u, vt *Mat) {
    gocv.SVDCompute(src.Mat, &(w.Mat), &(u.Mat), &(vt.Mat))
}
