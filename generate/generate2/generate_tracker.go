package generate2

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/garfeng/gocv_resource_tracker/generate/standard"

	"github.com/garfeng/gocv_resource_tracker/generate/gocv_tracker"
)

func GenerateTrackerPackage(core *gocv_tracker.GoCVTracker, typeCore *gocv_tracker.GoCVTrackerType, packagePath string, dstPath string) error {
	funcCoreGroup := reflect.TypeOf(core)
	os.MkdirAll(dstPath, 0777)

	typeCoreGroup := reflect.TypeOf(typeCore)

	names, err := standard.ScanDirs(packagePath)
	if err != nil {
		return err
	}

	for _, name := range names {
		name = strings.ToLower(name)
		ext := filepath.Ext(name)
		if ext == ".go" {
			fullName := filepath.Join(packagePath, name)
			funcs, types, typeFuncs, prefix, err := ReadFuncsFromGoFile(funcCoreGroup, typeCoreGroup, fullName)
			if err != nil {
				return err
			}

			err = WriteFuncsTo(funcs, types, typeFuncs, prefix, filepath.Join(dstPath, name))

			if err != nil {
				return err
			}
		}
	}

	return nil
}

var (
	regexpFunc   = regexp.MustCompile(`func \(g \*GoCVTracker\) ([^(]+)\(([^\)]*)\).+\{`)
	regexpOfType = regexp.MustCompile(`type (.+) = gocv\.(.+)`)
)

func WriteFuncsTo(funcs []FuncParam, types []string, methodOfTypes []string, prefix, dstName string) error {

	res := []string{}
	for _, v := range funcs {
		s := strings.ReplaceAll(v.Str, "GoCVTracker", "GoCVResourceTracker")
		gocvCallCmd := `gocv.` + v.FuncName + "("
		ins := []string{}

		for _, param := range v.In {
			if _, _, isCloser := GetCloseFunc(param.Type); isCloser {
				isPtr := param.Type.Kind() == reflect.Ptr

				//fmt.Println(shouldPtr, param.Name, param.Type, "||||", param.Type.Name())
				if !isPtr {
					ins = append(ins, fmt.Sprintf("%s.%s", param.Name, param.Type.Name()))
				} else {
					ins = append(ins, fmt.Sprintf("&(%s.%s)",
						param.Name, param.Type.Elem().Name()))
				}
			} else {
				isSlice := param.Type.Kind() == reflect.Slice
				if !isSlice {
					ins = append(ins, param.Name)
				} else {
					element := param.Type.Elem()
					_, _, isSliceCloser := GetCloseFunc(element)

					if !isSliceCloser {
						ins = append(ins, param.Name)
					} else {
						ins = append(ins, fmt.Sprintf("SliceToGoCVCloser(%s)", param.Name))
					}
				}
			}
		}

		gocvCallCmd += strings.Join(ins, ", ")
		gocvCallCmd += ")"

		if len(v.Out) == 0 {
			s += "\r\n    " + gocvCallCmd + "\r\n}\r\n"
			res = append(res, s)
		} else {
			returnData := []string{}
			for i := 0; i < len(v.Out); i++ {
				returnData = append(returnData, fmt.Sprintf("rs%d", i))
			}

			gocvReturnedNames := strings.Join(returnData, ", ")

			gocvCallCmd = gocvReturnedNames + " := " + gocvCallCmd
			s += "\r\n    " + gocvCallCmd + "\r\n"
			myReturnedNames := []string{}

			for i := 0; i < len(v.Out); i++ {
				closeFn, shouldPtr, find := GetCloseFunc(v.Out[i].Type)
				if find {
					toAddTracker := returnData[i]
					if shouldPtr {
						toAddTracker = "&" + toAddTracker
					}
					if closeFn.Type.NumOut() == 0 {
						s += "    " + "g.TrackCloser(" + toAddTracker + ")\r\n"
					} else {
						s += "    " + "g.TrackCloseError(" + toAddTracker + ")\r\n"
					}
					pkgDataName := fmt.Sprintf("pkg%d", i)
					myReturnedNames = append(myReturnedNames, pkgDataName)

					structName := v.Out[i].Type.Name()

					if v.Out[i].Type.Kind() == reflect.Ptr {
						structName = "&" + v.Out[i].Type.Elem().Name()
						returnData[i] = "*" + returnData[i]
					}

					s += "    " + pkgDataName + " := " + structName + fmt.Sprintf(`{
	    %s,
	    g,
    }
`, returnData[i])

					//s += "    " + returnData[i] + ".Close()\r\n"
				} else {

					isSlice := v.Out[i].Type.Kind() == reflect.Slice

					if isSlice {
						element := v.Out[i].Type.Elem()
						_, _, isSliceCloser := GetCloseFunc(element)

						if !isSliceCloser {
							myReturnedNames = append(myReturnedNames, returnData[i])
						} else {
							myReturnedNames = append(myReturnedNames, fmt.Sprintf("GoCVCloserToSlice(%s, g)", returnData[i]))
						}
					} else {
						myReturnedNames = append(myReturnedNames, returnData[i])
					}

				}
			}
			s += "    return " + strings.Join(myReturnedNames, ", ") + "\r\n"

			s += "}\r\n"
			res = append(res, s)
		}
	}

	code := strings.Join(res, "\r\n\r\n")

	code += strings.Join(methodOfTypes, "\r\n")

	code = strings.ReplaceAll(code, "interface##", "interface{}")

	prefix += `package gocv_resource_tracker
import (
`
	/*
		    "gocv.io/x/gocv"
		    "image"
		    "image/color"
		)
		`
	*/

	typeStr := strings.Join(types, "\r\n")
	code = typeStr + "\r\n\r\n" + code

	if strings.Contains(code, "gocv.") {
		prefix += `    "gocv.io/x/gocv"`
		prefix += "\r\n"
	}
	if strings.Contains(code, "image.") {
		prefix += `    "image"`
		prefix += "\r\n"

	}
	if strings.Contains(code, "color.") {
		prefix += `    "image/color"`
		prefix += "\r\n"
	}
	prefix += ")\r\n\r\n"

	//prefix += "\r\n\r\n"

	return ioutil.WriteFile(dstName, []byte(prefix+code), 0755)
}

func ReadFuncsFromGoFile(fnGroup reflect.Type,
	typeGroup reflect.Type, name string) ([]FuncParam, []string, []string, string, error) {
	buff, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, nil, nil, "", err
	}
	data := string(buff)

	data = strings.ReplaceAll(data, "interface{}", "interface##")
	data = strings.ReplaceAll(data, " ...", "... ")

	matched := regexpFunc.FindAllStringSubmatch(data, -1)

	res := []FuncParam{}

	_, filename := filepath.Split(name)

	for _, v := range matched {
		funcName := v[1]
		params := strings.Trim(v[2], " ")
		fnParam := FuncParam{
			In:       []ParamNameAndType{},
			Out:      []ParamNameAndType{},
			FuncName: funcName,
			Filename: filename,
			Str:      v[0],
		}

		if len(params) > 0 {
			paramsList := strings.Split(params, ",")
			for _, oneParam := range paramsList {
				oneParam = strings.Trim(oneParam, " ")
				if len(oneParam) > 0 {
					oneParamList := strings.Split(oneParam, " ")
					fnParam.In = append(fnParam.In, ParamNameAndType{
						Name: oneParamList[0],
						Type: nil,
					})
				}
			}
		}

		err = fnParam.ReadDetailsFromType(fnGroup)
		if err != nil && err != ErrMethodNotFind {
			return nil, nil, nil, "", fmt.Errorf("%s, func: %s, %d, %s, %s",
				err.Error(), fnParam.FuncName, len(fnParam.In), v[1], v[2])
		}
		if err == nil {
			res = append(res, fnParam)
		}
	}

	data2 := string(buff)
	types := regexpOfType.FindAllStringSubmatch(data2, -1)

	typesStr := []string{}

	methodOfTypes := []string{}

	for i := range types {
		typeName := types[i][1]
		types[i][0] = strings.Trim(types[i][0], "\r\n")
		types[i][0] = strings.Trim(types[i][0], "\n")

		typeFunc, find := typeGroup.MethodByName(typeName)
		if !find {
			continue
			//return nil, nil, "", errors.New("type not found:" + typeName)
		}
		typeInfo := typeFunc.Type.Out(0)
		_, _, isCloser := GetCloseFunc(typeInfo)
		if isCloser {
			typesStr = append(typesStr, fmt.Sprintf(`type %s struct {
    gocv.%s
	ResourceTracker *GoCVResourceTracker
}`, typeName, typeName))
		} else {
			typesStr = append(typesStr, types[i][0])
		}

		methods := ParseTypesMethod(typeInfo, nil)
		if len(methods) > 0 {
			methodOfTypes = append(methodOfTypes, strings.Join(methods, "\r\n\r\n"))
		}

	}

	prefixs := strings.Split(data2, "package")

	return res, typesStr, methodOfTypes, prefixs[0], nil
}

func ShouldFuncRewrite(fn reflect.Type) bool {
	inNum := fn.NumIn()
	outNum := fn.NumOut()

	for i := 1; i < inNum; i++ {
		in := fn.In(i)
		if _, _, isCloser := GetCloseFunc(in); isCloser {
			return true
		}
	}

	for i := 0; i < outNum; i++ {
		out := fn.Out(i)
		if _, _, isCloser := GetCloseFunc(out); isCloser {
			return true
		}
	}
	return false
}

var (
	GoCVRoot = "../../vendor/gocv.io/x/gocv"
)

func GetFuncLineStrFromGoCVRoot(dataType, fnName string) ([]string, error) {
	names, err := standard.ScanDirs(GoCVRoot)
	if err != nil {
		return nil, err
	}
	regStr := fmt.Sprintf(`func \((.+ \**%s)\) %s\(([^)]*)\)(.*)\{`, dataType, fnName)
	reg, err := regexp.Compile(regStr)

	if err != nil {
		return nil, err
	}

	for _, v := range names {
		fullName := filepath.Join(GoCVRoot, v)
		buff, _ := ioutil.ReadFile(fullName)
		data := string(buff)
		data = strings.ReplaceAll(data, "interface{}", "interface##")

		matched := reg.FindAllStringSubmatch(data, -1)

		if len(matched) > 0 {
			return matched[0], nil
		}
	}
	return nil, errors.New("method not find: (p *" + dataType + ") " + fnName + "()")
}

func ReWriteFunc(typeName string, fn reflect.Method) (string, bool, error) {
	if !ShouldFuncRewrite(fn.Type) {
		return "", false, nil
	}

	fnInfos, err := GetFuncLineStrFromGoCVRoot(typeName, fn.Name)

	if err != nil {
		return "", false, err
	}

	ptr := fnInfos[1]
	ptrs := strings.Split(ptr, " ")
	entryName := ptrs[0]

	fnPrefix := ""

	fnCore := entryName + "." + typeName + "." + fn.Name + "("

	paramList := strings.Trim(fnInfos[2], " ")

	paramNames := []string{}
	if len(paramList) > 0 {
		paramNames = strings.Split(paramList, ",")
		for i := range paramNames {
			tmp := strings.Trim(paramNames[i], " ")
			tmpL := strings.Split(tmp, " ")
			paramNames[i] = strings.Trim(tmpL[0], " ")
		}
	}
	if len(paramNames) != fn.Type.NumIn()-1 {
		return "", false, errors.New("param number not equal")
	}

	params := []string{}

	for i := 1; i < fn.Type.NumIn(); i++ {
		in := fn.Type.In(i)
		paramName := paramNames[i-1]
		inInfo := GetTypeNameAndParamName(paramName, in)

		if !inInfo.IsCloser {
			params = append(params, paramName)
		} else {
			if inInfo.IsPtr && inInfo.IsSlice {
				fnPrefix += fmt.Sprintf("param%d := SliceToGoCVCloser(*%s)", i, paramName)
				params = append(params, fmt.Sprintf("&param%d", i))
			} else if inInfo.IsPtr {
				params = append(params, "&("+paramName+"."+inInfo.TypeName+")")
			} else if inInfo.IsSlice {
				params = append(params, fmt.Sprintf("SliceToGoCVCloser(%s)", paramName))
			} else {
				params = append(params, paramName+"."+inInfo.TypeName)
			}
		}
	}

	fnCore += strings.Join(params, ", ")
	fnCore += ")"

	if fn.Type.NumOut() == 0 {
		fnCore = "    " + fnCore
	} else {
		outNames := []string{}
		for i := 0; i < fn.Type.NumOut(); i++ {
			outNames = append(outNames, fmt.Sprintf("rs%d", i))
		}

		outS := strings.Join(outNames, ", ")

		fnCore = outS + " := " + fnCore

		if len(fnPrefix) > 0 {
			fnCore = "    " + fnPrefix + "\r\n    " + fnCore
		} else {
			fnCore = "    " + fnCore
		}

		returnedData := []string{}

		trackers := []string{}

		for i := 0; i < fn.Type.NumOut(); i++ {
			outInfo := GetTypeNameAndParamName(outNames[i], fn.Type.Out(i))
			if !outInfo.IsCloser {
				returnedData = append(returnedData, outInfo.ParamName)
			} else {
				if outInfo.IsPtr {
					closeFn, shouldPtr, _ := GetCloseFunc(fn.Type.Out(i))
					at := ""
					if shouldPtr {
						at = "&"
					}
					if closeFn.Type.NumOut() == 0 {
						trackers = append(trackers, fmt.Sprintf("    %s.ResourceTracker.TrackCloser(%s%s)", entryName, at, outInfo.ParamName))
					} else {
						trackers = append(trackers, fmt.Sprintf("    %s.ResourceTracker.TrackCloseError(%s%s)", entryName, at, outInfo.ParamName))
					}

					returnedData = append(returnedData, fmt.Sprintf("&%s{%s, %s.ResourceTracker}", outInfo.TypeName, outInfo.ParamName, entryName))
				} else if outInfo.IsSlice {
					returnedData = append(returnedData, fmt.Sprintf("GoCVCloserToSlice(%s)", outInfo.ParamName))
				} else {

					closeFn, shouldPtr, _ := GetCloseFunc(fn.Type.Out(i))
					at := ""
					if shouldPtr {
						at = "&"
					}
					if closeFn.Type.NumOut() == 0 {
						trackers = append(trackers, fmt.Sprintf("    %s.ResourceTracker.TrackCloser(%s%s)", entryName, at, outInfo.ParamName))
					} else {
						trackers = append(trackers, fmt.Sprintf("    %s.ResourceTracker.TrackCloseError(%s%s)", entryName, at, outInfo.ParamName))
					}

					returnedData = append(returnedData, fmt.Sprintf("%s{%s, %s.ResourceTracker}", outInfo.TypeName, outInfo.ParamName, entryName))
				}
			}
		}

		if len(trackers) > 0 {
			fnCore += "\r\n\r\n" + strings.Join(trackers, "\r\n")
		}

		fnCore += "\r\n    return " + strings.Join(returnedData, ", ")
	}

	newFunc := fnInfos[0] + "\r\n" + fnCore + "\r\n}"

	return newFunc, true, nil
}

func GetTypeNameAndParamName(name string, v reflect.Type) TypeNameAndParamName {
	inInfo := TypeNameAndParamName{
		TypeName:  "",
		ParamName: name,
	}
	_, _, isCloser := GetCloseFunc(v)
	if isCloser {
		inInfo.IsCloser = true
		isPtr := v.Kind() == reflect.Ptr
		if isPtr {
			inInfo.TypeName = v.Elem().Name()
			inInfo.IsPtr = true
		} else {
			inInfo.TypeName = v.Name()
			inInfo.IsPtr = false
		}
	} else {
		isSlice := v.Kind() == reflect.Slice
		if isSlice {
			inInfo.TypeName = v.Elem().Name()
			inInfo.IsSlice = true
			_, _, isCloser = GetCloseFunc(v.Elem())
			inInfo.IsCloser = isCloser
		} else {
			inInfo.TypeName = v.Name()
			inInfo.IsSlice = false
		}
	}
	if !inInfo.IsCloser {
		isPtr := v.Kind() == reflect.Ptr
		if isPtr {
			info2 := GetTypeNameAndParamName(name, v.Elem())
			if info2.IsCloser {
				info2.IsPtr = true
				return info2
			}
		}
	}

	return inInfo
}

type TypeNameAndParamName struct {
	TypeName  string
	ParamName string
	IsPtr     bool
	IsSlice   bool
	IsCloser  bool
}

func ParseTypesMethod(data reflect.Type, rewritedMap map[string]bool) []string {
	_, _, isCloser := GetCloseFunc(data)
	if !isCloser {
		return []string{}
	}

	res := []string{}
	funcNum := data.NumMethod()

	name := data.Name()
	if data.Kind() == reflect.Ptr {
		name = data.Elem().Name()
	}

	if rewritedMap == nil {
		rewritedMap = map[string]bool{}
	}

	for i := 0; i < funcNum; i++ {
		fn := data.Method(i)
		if !rewritedMap[fn.Name] {
			s, shouldRewrite, err := ReWriteFunc(name, fn)
			if err != nil {
				panic(err)
			}
			if shouldRewrite {
				res = append(res, s)
				rewritedMap[fn.Name] = true
			}
		}

	}

	if data.Kind() != reflect.Ptr {
		res2 := ParseTypesMethod(reflect.New(data).Type(), rewritedMap)
		res = append(res, res2...)
	}

	return res
}

type ParamNameAndType struct {
	Name string
	Type reflect.Type
}

type FuncParam struct {
	In       []ParamNameAndType
	Out      []ParamNameAndType
	FuncName string
	Filename string
	Str      string
}

var (
	ErrMethodNotFind = errors.New("method not found")
)

func (f *FuncParam) ReadDetailsFromType(tp reflect.Type) error {
	method, find := tp.MethodByName(f.FuncName)
	if !find {
		return ErrMethodNotFind
	}
	paramsNum := method.Type.NumIn()

	if paramsNum-1 != len(f.In) {
		return errors.New("paramsNumber from type != paramsNumber from code")
	}

	for i := 0; i < paramsNum-1; i++ {
		f.In[i].Type = method.Type.In(i + 1)
	}

	returnNum := method.Type.NumOut()

	for i := 0; i < returnNum; i++ {
		f.Out = append(f.Out, ParamNameAndType{
			Name: "",
			Type: method.Type.Out(i),
		})
	}

	return nil
}

func GetCloseFunc(v reflect.Type) (closeFn reflect.Method, shouldPtr bool, find bool) {
	closeFn1, find1 := v.MethodByName("Close")
	ptr := reflect.New(v)
	closeFn2, find2 := ptr.Type().MethodByName("Close")

	//var closeFn reflect.Method
	shouldPtr = false
	if find1 {
		closeFn = closeFn1
	} else if find2 {
		shouldPtr = true
		closeFn = closeFn2
	}

	return closeFn, shouldPtr, find1 || find2
}
