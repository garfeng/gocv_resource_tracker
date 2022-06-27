package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	dst = flag.String("o", "../../gocv.go", "dst gocv file")
)

func main() {
	flag.Parse()

	if (*dst) == "" {
		flag.PrintDefaults()
		return
	}

	imp := NewImporter("gocv.io/x/gocv", *dst)

	err := imp.Import()
	if err != nil {
		fmt.Println(err)
		return
	}

	w := bytes.NewBuffer(nil)
	err = imp.Encode(w)
	if err != nil {
		fmt.Println(err)
		return
	}

	dir, _ := filepath.Split(*dst)
	os.MkdirAll(dir, 0755)

	err = ioutil.WriteFile(*dst, w.Bytes(), 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
}
