# Resource Tracker for GOCV

The repository is aimed to free you from calling `Close` every time after create a C Object in [hybridgroup/gocv](https://github.com/hybridgroup/gocv).

### Example

#### In GoCV

``` Go
func GOCVExample() {
	mat := gocv.IMRead("test.png", gocv.IMReadGrayScale)
	defer mat.Close() // 1

	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(5, 5))
	defer kernel.Close() // 2

	erode := gocv.NewMat()
	defer erode.Close() // 3

	gocv.Erode(mat, &erode, kernel)

	thres := gocv.NewMat()
	defer thres.Close() // 4

	gocv.Threshold(erode, &thres, 0, 255, gocv.ThresholdOtsu)

	contours := gocv.FindContours(thres, gocv.RetrievalList, gocv.ChainApproxNone)
	defer contours.Close() // 5

	contoursNum := contours.Size()
	for i := 0; i < contoursNum; i++ {
		c := contours.At(i)
		area := gocv.ContourArea(c)
		if area > 10 {
			// do something
		}
	}
}
```

#### By Resource Tracker

``` go
func ResourceTrackerExample() {
	rt := gocv_resource_tracker.NewTracker()
	defer rt.Close()

	mat := rt.IMRead("test.png", gocv.IMReadGrayScale)

	kernel := rt.GetStructuringElement(gocv.MorphEllipse, image.Pt(5, 5))
	erode := rt.NewMat()
	gocv.Erode(mat, &erode, kernel)

	thres := rt.NewMat()
	gocv.Threshold(erode, &thres, 0, 255, gocv.ThresholdOtsu)

	contours := rt.FindContours(thres, gocv.RetrievalList, gocv.ChainApproxNone)

	contoursNum := contours.Size()
	for i := 0; i < contoursNum; i++ {
		c := contours.At(i)
		area := gocv.ContourArea(c)
		if area > 10 {
			// do something
		}
	}
}
```





