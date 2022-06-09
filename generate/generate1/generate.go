package generate1

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/*
grep -rP '^func( +)[A-Z].*' *.go | grep -v Test > ../../../gocv_func.txt
grep -rP '^type( +)[A-Z].*' *.go | grep -v Test
*/

func GenerateGoCV(packagePath string, dstPath string) error {
	os.MkdirAll(dstPath, 0777)

	b, err := os.Open(packagePath)
	if err != nil {
		return err
	}
	defer b.Close()

	names, err := b.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, v := range names {
		v = strings.ToLower(v)
		ext := filepath.Ext(v)
		if ext == ".go" && !strings.Contains(v, "_test.go") {
			src := filepath.Join(packagePath, v)
			dst := filepath.Join(dstPath, v)
			err = GenerateAFile(src, dst)
			if err != nil {
				return err
			}
		}
	}

	mainFile := filepath.Join(dstPath, "gocv_tracker.go")
	mainCode := `
package gocv_tracker
type GoCVTracker struct {}
type GoCVTrackerType struct{}
`
	ioutil.WriteFile(mainFile, []byte(mainCode), 0755)

	return nil
}

var (
	regexpOfFunc = regexp.MustCompile(`func( +)[A-Z]([^\{]+)\{`)
	regexpOfType = regexp.MustCompile(`^type( +)[A-Z]([^ ]*)`)
	//regexpOfFuncParam = regexp.MustCompile(`\([^\)]*\)`)
)

func GenerateAFile(name string, dstName string) error {
	buff, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}

	data1 := string(buff)
	data1 = strings.ReplaceAll(data1, "\n", "")
	data1 = strings.ReplaceAll(data1, "\r", "")
	data1 = strings.ReplaceAll(data1, "interface{}", "interface##")

	data2 := string(buff)
	data2 = strings.ReplaceAll(data2, "interface{}", "interface##")

	data2List := strings.Split(data2, "\n")

	funcs := regexpOfFunc.FindAllString(data1, -1)
	types := regexpOfType.FindAllString(data2, -1)

	newFuncs := []string{}
	newTypes := []string{}

	fmt.Println(len(types), len(funcs))

	prefixs := strings.Split(string(buff), "package")
	prefix := prefixs[0]

	for _, v := range funcs {
		newV := strings.Replace(v, "func ", "func (g *GoCVTracker) ", 1)
		newV += "\r\n    panic(\"implement me\")\r\n"
		newV += "}"
		newV = strings.ReplaceAll(newV, "interface##", "interface{}")
		newFuncs = append(newFuncs, newV)
	}

	for _, one := range data2List {
		res := regexpOfType.FindAllString(one, -1)
		if len(res) > 0 {
			v := res[0]
			typeName := strings.ReplaceAll(v, "type", "")
			typeName = strings.Trim(typeName, " ")
			typeNameOnly := typeName
			typeName = "gocv." + typeName
			typeName = strings.ReplaceAll(typeName, " ", "")

			newV := v + " = " + typeName
			newV = strings.ReplaceAll(newV, "interface##", "interface{}")

			newTypes = append(newTypes, newV)

			newV2 := "func (g *GoCVTrackerType) " + typeNameOnly + "() " + typeNameOnly + " {\r\n    panic(\"\") \r\n}"

			newTypes = append(newTypes, newV2)

			fmt.Println(v)
		}
	}

	resTypes := strings.Join(newTypes, "\r\n")
	resFuncs := strings.Join(newFuncs, "\r\n\r\n")

	res := prefix + "\r\npackage gocv_tracker\r\n\r\n"
	res += `
import (
`
	pb := resTypes + "\r\n\r\n" + resFuncs
	/*
		"image"
		"image/color"
	*/

	if strings.Contains(pb, "gocv.") {
		res += `    "gocv.io/x/gocv"`
		res += "\r\n"
	}

	if strings.Contains(pb, "image.") {
		res += `    "image"`
		res += "\r\n"
	}
	if strings.Contains(pb, "color.") {
		res += `    "image/color"`
		res += "\r\n"
	}
	res += ")\r\n"

	res += resTypes
	res += "\r\n\r\n"
	res += resFuncs
	return ioutil.WriteFile(dstName, []byte(res), 0755)
}
