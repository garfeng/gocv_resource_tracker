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
	fmt.Println("import finished")
	w := bytes.NewBuffer(nil)
	err := imp.Encode(w)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("encode finished")
	os.Mkdir("tmp", 0755)
	ioutil.WriteFile("../../gocv.go", w.Bytes(), 0644)
}
