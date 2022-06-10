
package gocv_resource_tracker
import (
    "gocv.io/x/gocv"
    "image"
)



func (g *GoCVResourceTracker) FP16BlobFromImage(img Mat, scaleFactor float32, size image.Point, mean float32,	swapRB bool, crop bool) []byte {
    rs0 := gocv.FP16BlobFromImage(*(img.Mat), scaleFactor, size, mean, 	swapRB, crop)
    return rs0
}
