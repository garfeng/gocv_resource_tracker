package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	imp := NewImporter("gocv.io/x/gocv", "../../gocv.go")
	imp.Import()
	w := bytes.NewBuffer(nil)
	err := imp.Encode(w)
	if err != nil {
		fmt.Println(err)
	}
	os.Mkdir("tmp", 0755)
	ioutil.WriteFile("../../gocv.go", w.Bytes(), 0644)
}
