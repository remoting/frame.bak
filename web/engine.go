package web

import "net/http"

type HandlerFunc func(*Context)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Router   *Router
}

type Engine struct {
	RouterGroup
}

func New() *Engine {
	engine := &Engine{}
	engine.parent = nil
	engine.prefix = ""
	return engine
}

func (engine *Engine) Run(addr ...string) (err error) {
	engine.print()
	address := resolveAddress(addr)
	http.ListenAndServe(address, engine)
	return
}
