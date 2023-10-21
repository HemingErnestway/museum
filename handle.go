package main

import (
	"museum/api"
	"museum/config"
	"museum/engine"
	"net/http"
	"reflect"
	"strings"
)

var types map[string]bool
var hdl *api.Handler
var apiMap map[string]map[string]reflect.Value

func init() {
	cfg := config.Get()
	apiMap = make(map[string]map[string]reflect.Value)
	apiMap["POST"] = make(map[string]reflect.Value)
	apiMap["PUT"] = make(map[string]reflect.Value)
	apiMap["DELETE"] = make(map[string]reflect.Value)
	apiMap["GET"] = make(map[string]reflect.Value)
	maps := cfg.Api

	types = make(map[string]bool)
	types[".ico"] = true
	types[".png"] = true
	types[".html"] = true
	types[".js"] = true
	types[".svg"] = true

	hdl = &api.Handler{}
	services := reflect.ValueOf(hdl)
	_struct := reflect.TypeOf(hdl)

	for methodNum := 0; methodNum < _struct.NumMethod(); methodNum++ {
		method := _struct.Method(methodNum)
		val, ok := maps[method.Name]
		if !ok {
			continue
		}
		if _, ok := apiMap[val.Method]; !ok {

		}
		apiMap[val.Method][val.Url] = services.Method(methodNum)
	}

}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	ctx := engine.Context{
		Response: w,
		Request:  r,
	}

	url := r.URL
	path := url.Path[1:]
	pathArr := strings.Split(path, "/")
	pathName := pathArr[0]

	if pathArr[0] == "" {
		sendFile("index.html", ctx)
		return
	}

	//last := pathArr[(len(pathArr) - 1)]
	//str := last[strings.LastIndex(last, "."):]
	//if len(last) > 3 && types[str] {
	//	sendFile("tpl/index.html", ctx)
	//}

	maps, ok := apiMap[r.Method]
	if !ok {
		w.Write([]byte("No such method"))
	}

	if len(pathArr) > 1 {
		pathName += "/{id}"
	}

	if fun, ok := maps[pathName]; ok {
		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(&ctx)
		//fmt.Println(pathName)
		fun.Call(in)
	}
}

func sendFile(url string, ctx engine.Context) {
	ctx.Response.Write([]byte{})
}
