package web

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type RouterGroup struct {
	parent     *RouterGroup
	prefix     string         // Prefix for sub-route.
	children   []*RouterGroup // children group.
	middleware []*HandlerFunc // middleware array.
	router     []*Router
}

type Router struct {
	group   *RouterGroup
	Path    string
	Method  string
	handler HandlerFunc
}

func (router *Router) GetFullPath() string {
	path := router.group.GetPrefix()
	if strings.HasPrefix(router.Path, "/") {
		path = path + router.Path
	} else {
		path = path + "/" + router.Path
	}
	return path
}
func proxyHandlerFunc(method reflect.Method, object interface{}) HandlerFunc {
	return func(c *Context) {
		params := []reflect.Value{reflect.ValueOf(object), reflect.ValueOf(c)}
		method.Func.Call(params)
	}
}
func (group *RouterGroup) print() {
	for _, v := range group.children {
		v.print()
	}
	for _, v := range group.router {
		fmt.Printf("%s,%s,%v\n", v.GetFullPath(), v.Method, v.handler)
	}
}
func (group *RouterGroup) GetPrefix() string {
	if group.parent != nil {
		return group.parent.GetPrefix() + group.prefix
	} else {
		return group.prefix
	}
}
func (group *RouterGroup) Group(prefix string, groups ...func(group *RouterGroup)) *RouterGroup {
	if len(prefix) > 0 && prefix[0] != '/' {
		prefix = "/" + prefix
	}
	if prefix == "/" {
		prefix = ""
	}
	child := &RouterGroup{
		parent: group,
		prefix: prefix,
	}
	group.children = append(group.children, child)
	if len(groups) > 0 {
		for _, v := range groups {
			v(child)
		}
	}
	return child
}

//获取指针的原始结构体类型
func ttt(t reflect.Type) {
	tt := t
	for tt.Kind() == reflect.Ptr || tt.Kind() == reflect.Slice {
		tt = tt.Elem()
	}
}
func (group *RouterGroup) ALL(prefix string, controller interface{}) *RouterGroup {
	t := reflect.TypeOf(controller)
	kind := t.Kind()
	if kind == reflect.Func {
		switch v := controller.(type) {
		case HandlerFunc:
			group.addRouter(prefix, "*", controller.(HandlerFunc))
		case func(*Context):
			group.addRouter(prefix, "*", controller.(func(ctx *Context)))
		default:
			fmt.Printf("函数不符合约定:%v", v)
		}
	}
	if kind == reflect.Struct {

	}
	if kind == reflect.Ptr {
		// 结构体指针
		group.Group(prefix, func(g *RouterGroup) {
			for i := 0; i < t.NumMethod(); i++ {
				m := t.Method(i)
				// 判断是否为控制器方法
				if m.Type.NumIn() == 2 && m.Type.In(1) == reflect.TypeOf(&Context{}) {
					g.addRouter(m.Name, "*", proxyHandlerFunc(m, controller))
				}
			}
		})
	}
	return group
}
func (group *RouterGroup) addRouter(path, method string, handler HandlerFunc) {
	group.router = append(group.router, &Router{
		Path:    path,
		Method:  method,
		group:   group,
		handler: handler,
	})
}
func (group *RouterGroup) GET(path string, handler HandlerFunc) {
	group.addRouter(path, "GET", handler)
}

func (group *RouterGroup) POST(path string, handler HandlerFunc) {
	group.addRouter(path, "POST", handler)
}
func (engine *RouterGroup) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	isMatch := false
	path := req.URL.Path
	for _, v := range engine.children {
		pre := v.GetPrefix()
		if len(path) >= len(pre) && pre == path[0:len(pre)] {
			isMatch = true
			v.ServeHTTP(w, req)
			break
		}
	}
	if !isMatch {
		engine.handleHTTPRequest(w, req)
	}
}
func (group *RouterGroup) handleHTTPRequest(w http.ResponseWriter, req *http.Request) {
	c := &Context{
		Request:  req,
		Response: w,
	}
	for _, v := range group.router {
		pattern := v.GetFullPath()
		if pattern == req.URL.Path && (v.Method == "*" || v.Method == req.Method) {
			c.Router = v
			v.handler(c)
			break
		}
	}
}
