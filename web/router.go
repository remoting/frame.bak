package web

import (
	"fmt"
	"net/http"
	"reflect"
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
	method  *reflect.Method
	object  interface{}
	handler HandlerFunc
	otype   reflect.Type
}

func (group *RouterGroup) print() {
	for _, v := range group.children {
		v.print()
	}
	for _, v := range group.router {
		if v.method != nil {
			fmt.Printf("%s,%v\n", group.GetPrefix()+"/"+v.Path, v.method.Type)
		} else {
			fmt.Printf("%s,%v\n", group.GetPrefix()+v.Path, v.handler)
		}
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

func (rg *RouterGroup) ALL(prefix string, controller interface{}) *RouterGroup {
	rg.Group(prefix, func(group *RouterGroup) {
		t := reflect.TypeOf(controller)
		fmt.Println(t.Kind())
		if t.Kind() == reflect.Ptr {
			//指针
			fmt.Println(t.Elem().Name())
			fmt.Println(t.Elem().Kind())
			fmt.Println(t.Elem().PkgPath())
		} else {
			//结构体
			fmt.Println(t.Name())
			fmt.Println(t.Kind())
			fmt.Println(t.PkgPath())
		}

		for i := 0; i < t.NumMethod(); i++ {
			m := t.Method(i)
			if m.Type.NumIn() == 2 && m.Type.In(1).String() == "*web.Context" {
				group.router = append(group.router, &Router{
					Path:   m.Name,
					group:  group,
					object: controller,
					otype:  t,
					method: &m,
				})
			}
		}
	})
	return rg
}

func (group *RouterGroup) GET(prefix string, handler func(c *Context)) {

	group.router = append(group.router, &Router{
		Path:    prefix,
		method:  nil,
		group:   group,
		object:  nil,
		handler: (HandlerFunc)(handler),
	})
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
		//fmt.Printf("%s,%s,%s", v.group.GetPrefix(), v.Path, req.URL.Path)
		pattern := v.group.GetPrefix() + "/" + v.Path
		if v.method == nil {
			pattern = v.group.GetPrefix() + v.Path
		}
		if pattern == req.URL.Path {
			c.Router = v
			if v.method == nil {
				v.handler(c)
			} else {
				params := []reflect.Value{reflect.ValueOf(v.object), reflect.ValueOf(c)}
				v.method.Func.Call(params)
			}
		}
	}
}
