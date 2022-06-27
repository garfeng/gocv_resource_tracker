package main

import (
	"github.com/wzshiming/gotype"
	"go/ast"
	"path/filepath"
	"strings"
)

type Importer struct {
	Values         []*Value
	Methods        []*Method
	Types          []*Type
	ToPtrTypes     []*Type
	SameTypes      []*Type
	WithCloseTypes []*Type

	SrcPath    string
	GoFilename string
}

var (
	VarMap = map[string]bool{
		"ErrEmptyByteSlice": true,
	}
	ShouldNotTrackerMethod = map[string]bool{
		"At": true,
	}
)

func NewImporter(srcPath string, dstPath string) *Importer {
	_, dst := filepath.Split(dstPath)
	return &Importer{
		Values:         []*Value{},
		Methods:        []*Method{},
		Types:          []*Type{},
		SameTypes:      []*Type{},
		ToPtrTypes:     []*Type{},
		WithCloseTypes: []*Type{},
		SrcPath:        srcPath,
		GoFilename:     dst,
	}
}

func (i *Importer) Import() error {
	imp := gotype.NewImporter()
	pkgInfo, err := imp.Import(i.SrcPath, "")

	if err != nil {
		return err
	}

	num := pkgInfo.NumChild()

	for id := 0; id < num; id++ {
		v := pkgInfo.Child(id)
		if ast.IsExported(v.Name()) {
			switch v.Kind() {
			case gotype.Declaration:
				decl := v.Declaration()
				switch decl.Kind() {
				case gotype.Func:
					i.ImportFunc(v)
				default:
					i.ImportValue(v)
				}
			case gotype.Scope, gotype.Invalid:
			default:
				i.ImportType(v)
			}
		}
	}

	return nil
}

func (i *Importer) ImportFunc(one gotype.Type) {
	if !ast.IsExported(one.Name()) {
		return
	}

	i.Methods = append(i.Methods, NewMethod(one))
}

func NewMethod(one gotype.Type) *Method {
	decl := one.Declaration()
	method := &Method{
		BaseDef: BaseDef{
			Type:    one,
			Name:    one.Name(),
			Comment: comment(one.Comment().Text()),
			Doc:     comment(one.Doc().Text()),
		},
		Ins:  []*Value{},
		Outs: []*Value{},
	}

	num := decl.NumIn()

	for id := 0; id < num; id++ {
		v := decl.In(id)
		method.Ins = append(method.Ins, NewValue(v))
	}

	num = decl.NumOut()
	for id := 0; id < num; id++ {
		v := decl.Out(id)
		method.Outs = append(method.Outs, NewValue(v))
	}

	return method
}

func NewValue(one gotype.Type) *Value {
	v := &Value{
		BaseDef: BaseDef{
			Type:     one.Declaration(),
			Name:     one.Name(),
			TypeName: one.Declaration().String(),
			Comment:  comment(one.Comment().Text()),
			Doc:      comment(one.Doc().Text()),
		},
		IsVar: VarMap[one.Name()],
	}
	return v
}

func NewType(one gotype.Type) *Type {
	t := &Type{
		BaseDef: BaseDef{
			Type:     one,
			Name:     one.Name(),
			TypeName: one.String(),
			Comment:  comment(one.Comment().Text()),
			Doc:      comment(one.Doc().Text()),
		},
		Kind:                 one.Kind().String(),
		Fields:               []*Value{},
		Methods:              []*Method{},
		MethodsReturnsCloser: []*Method{},
	}

	realType := GetRealElementOfPtr(one)

	if realType.Kind() == gotype.Struct {
		num := realType.NumField()
		for i := 0; i < num; i++ {
			f := realType.Field(i)
			if ast.IsExported(f.Name()) {
				t.Fields = append(t.Fields, NewField(f))
			}
		}
	}

	num := one.NumMethod()
	for i := 0; i < num; i++ {
		m := one.Method(i)
		if ast.IsExported(m.Name()) {
			method := NewMethod(m)
			method.ReceverTypeName = one.Name()
			t.Methods = append(t.Methods, method)
			for _, v := range method.Outs {
				if v.IsCloser() || v.IsRealElemCloser() {
					t.MethodsReturnsCloser = append(t.MethodsReturnsCloser, method)
				}
			}
		}
	}
	return t
}

func NewField(one gotype.Type) *Value {
	typeName := ""
	var tp gotype.Type
	if one.Kind() == gotype.Declaration {
		typeName = one.Declaration().String()
		tp = one.Declaration()
	} else {
		typeName = one.String()
		typeName = strings.Replace(typeName, one.Name(), "", 1)
		tp = one
	}
	return &Value{
		BaseDef: BaseDef{
			Type:     tp,
			Name:     one.Name(),
			TypeName: typeName,
			Comment:  comment(one.Comment().Text()),
			Doc:      comment(one.Doc().Text()),
		}, IsVar: true}
}

func (i *Importer) ImportType(v gotype.Type) {
	if !ast.IsExported(v.Name()) {
		return
	}
	tp := NewType(v)
	i.Types = append(i.Types, tp)

	if len(tp.Methods) > 0 || len(tp.Fields) > 0 {
		hasClose := false
		for _, method := range tp.Methods {
			if method.Name == "Close" {
				hasClose = true
				tp.IsCloser = true
				if len(method.Outs) > 0 {
					tp.IsCloserError = true
				}
			}
		}
		if hasClose {
			i.WithCloseTypes = append(i.WithCloseTypes, tp)

		} else {
			i.ToPtrTypes = append(i.ToPtrTypes, tp)
		}
	} else {
		i.SameTypes = append(i.SameTypes, tp)
	}

}

func (i *Importer) ImportValue(v gotype.Type) {
	if !ast.IsExported(v.Name()) {
		return
	}

	i.Values = append(i.Values, NewValue(v))
}

type BaseDef struct {
	Type     gotype.Type
	Name     string
	TypeName string
	Comment  string
	Doc      string
}

type Value struct {
	BaseDef
	IsVar bool
}

type Method struct {
	BaseDef
	Ins             []*Value
	Outs            []*Value
	ReceverTypeName string
}

type Type struct {
	BaseDef

	Kind          string
	IsCloser      bool
	IsCloserError bool

	Fields               []*Value
	Methods              []*Method
	MethodsReturnsCloser []*Method
}

func comment(s string) string {
	ss := []string{}
	dd := strings.Split(s, "\n")
	for i, v := range dd {
		if v == "" && i == len(dd)-1 {
			continue
		}
		v = strings.TrimSpace(v)
		if strings.Index(v, "//") != 0 {
			v = "// " + v
		}
		ss = append(ss, v)
	}

	return strings.Join(ss, "\n")
}
