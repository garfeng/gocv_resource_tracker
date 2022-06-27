package main

import (
	"embed"
	"fmt"
	"github.com/wzshiming/gotype"
	"io"
	"strings"
	"text/template"
)

//go:embed tmpl
var goCodeTmpl embed.FS

func (v *Value) IsCloser() bool {
	_, find := v.Type.MethodByName("Close")
	if find {
		return find
	}
	if v.Type.Kind() == gotype.Ptr {
		elem := v.Type.Elem()
		_, find = elem.MethodByName("Close")
	}
	return find
}
func (v *Value) IsRealElemCloser() bool {
	v2 := GetRealElement(v.Type)

	_, find := v2.MethodByName("Close")
	if find {
		return find
	}

	return find
}

func (v *Value) IsPtr() bool {
	return v.Type.Kind() == gotype.Ptr
}

func (v *Value) IsSlice() bool {
	return v.Type.Kind() == gotype.Slice
}

func (v *Value) RealTypeName() string {
	v2 := GetRealElementOfPtr(v.Type)
	return v2.String()
}

func GoFuncFormat(m *Method, isTypeMethod bool) string {
	ins := []string{}
	inValues := []string{}
	outs := []string{}
	outValues := []string{}
	outValuesToReturn := []string{}

	for _, v := range m.Ins {
		if v.IsCloser() && (!v.IsPtr()) {
			ins = append(ins, fmt.Sprintf("%s *%s", v.Name, v.TypeName))
		} else {
			ins = append(ins, fmt.Sprintf("%s %s", v.Name, v.TypeName))
		}

		if v.IsCloser() || v.IsRealElemCloser() {
			if v.IsPtr() {
				inValues = append(inValues, v.Name+".coreElemPtr()")
			} else {
				if !v.IsSlice() {
					inValues = append(inValues, v.Name+".coreElem()")
				} else {
					inValues = append(inValues, fmt.Sprintf("SliceToGoCVCloser(%s)", v.Name))
				}
			}
		} else {
			inValues = append(inValues, v.Name)
		}
	}
	for i, v := range m.Outs {
		if v.IsCloser() && (!v.IsPtr()) {
			outs = append(outs, "*"+v.TypeName)
		} else {
			outs = append(outs, v.TypeName)
		}
		name := fmt.Sprintf("_ov%d", i+1)
		outValues = append(outValues, name)

		resourceTrackerName := "rt"
		if isTypeMethod {
			resourceTrackerName = "ptr.ResourceTracker"
		}

		if v.IsCloser() || v.IsRealElemCloser() {
			rtName := v.RealTypeName()
			if v.IsPtr() {
				outValuesToReturn = append(outValuesToReturn, fmt.Sprintf("new%sFromPtr(%s, %s)", rtName, resourceTrackerName, name))
			} else {
				if !v.IsSlice() {
					if !ShouldNotTrackerMethod[m.Name] {
						outValuesToReturn = append(outValuesToReturn, fmt.Sprintf("new%sFromElem(%s, %s)", rtName, resourceTrackerName, name))
					} else {
						outValuesToReturn = append(outValuesToReturn, fmt.Sprintf("new%sFromElemNoTracker(%s, %s)", rtName, resourceTrackerName, name))
					}
				} else {
					outValuesToReturn = append(outValuesToReturn, fmt.Sprintf("GoCVCloserToSlice(%s, %s)", name, resourceTrackerName))
				}
			}
		} else {
			outValuesToReturn = append(outValuesToReturn, name)
		}
	}

	inStr := strings.Join(ins, ", ")
	outStr := strings.Join(outs, ", ")
	inValuesStr := strings.Join(inValues, ", ")
	outValueStr := strings.Join(outValues, ", ")
	outValueToReturnStr := strings.Join(outValuesToReturn, ", ")

	originName := "origin"
	if isTypeMethod {
		originName = "ptr." + m.ReceverTypeName
	}

	if len(outs) == 0 {
		return fmt.Sprintf(`%s(%s) {
%s.%s(%s)
}
`, m.Name, inStr, originName, m.Name, inValuesStr)
	} else if len(outs) == 1 {
		return fmt.Sprintf(`%s(%s) %s {
	%s := %s.%s(%s)
	return %s
}`, m.Name, inStr, outStr, outValueStr, originName, m.Name, inValuesStr, outValueToReturnStr)
	} else {
		return fmt.Sprintf(`%s(%s) (%s) {
	%s := %s.%s(%s)
	return %s
}`, m.Name, inStr, outStr, outValueStr, originName, m.Name, inValuesStr, outValueToReturnStr)
	}
}

var (
	goTmpl = template.New("go")
)

func (i Importer) Encode(w io.Writer) error {
	return goTmpl.Execute(w, i)
}

func init() {
	var err error
	var buff []byte
	buff, err = goCodeTmpl.ReadFile("tmpl/go.tmpl")
	if err != nil {
		panic(err)
	}
	str := string(buff)

	goTmpl.Funcs(template.FuncMap{
		"GoFuncFormat": GoFuncFormat,
	})
	goTmpl.Parse(str)
}
