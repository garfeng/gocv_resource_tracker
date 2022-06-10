package generate1

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerateAFile(t *testing.T) {
	os.Mkdir("gocv_tracker", 0777)
	err := GenerateAFile("../vendor/gocv.io/x/gocv/core.go", "gocv_tracker/core.go")
	if err != nil {
		fmt.Println(err)
	}
}

func TestGenerateGoCV(t *testing.T) {
	err := GenerateGoCV("../../vendor/gocv.io/x/gocv/", "../gocv_tracker/")
	if err != nil {
		fmt.Println(err)
	}
}
