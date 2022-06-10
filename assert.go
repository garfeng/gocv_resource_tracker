
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
)



func (g *GoCVResourceTracker) Caller(skip int) (name string, line int) {
    rs0, rs1 := gocv.Caller(skip)
    return rs0, rs1
}


func (g *GoCVResourceTracker) PanicWithCaller(skip int, e interface{}) {
    gocv.PanicWithCaller(skip, e)
}


func (g *GoCVResourceTracker) MustEqual(a interface{}, b interface{}, format string, i... interface{}) {
    gocv.MustEqual(a, b, format, i...)
}


func (g *GoCVResourceTracker) Must(a bool, format string, i... interface{}) {
    gocv.Must(a, format, i...)
}


func (g *GoCVResourceTracker) MustNot(a bool, format string, i... interface{}) {
    gocv.MustNot(a, format, i...)
}
