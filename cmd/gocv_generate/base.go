package main

import "github.com/wzshiming/gotype"

func GetRealDecl(v gotype.Type) gotype.Type {
	if v.Kind() == gotype.Declaration {
		return GetRealDecl(v.Declaration())
	}
	return v
}

func GetRealElement(v gotype.Type) gotype.Type {
	kind := v.Kind()
	if kind == gotype.Ptr || kind == gotype.Slice || kind == gotype.Map {
		return GetRealElement(v.Elem())
	}
	return v
}

func GetRealElementOfPtr(v gotype.Type) gotype.Type {
	kind := v.Kind()
	if kind == gotype.Ptr {
		return GetRealElementOfPtr(v.Elem())
	}
	return v
}
