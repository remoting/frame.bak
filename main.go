package main

import (
	"io"

	"github.com/remoting/frame/web"
)

type Demo struct {
}

func (demo *Demo) Index(c *web.Context) {
	io.WriteString(c.Response, "hello"+c.Request.URL.String())
}
func (demo *Demo) Home(c *web.Context) {
	io.WriteString(c.Response, "home"+c.Request.URL.String())
}
func xxx(h web.HandlerFunc, g *web.RouterGroup) {
	g.ALL("/test", h)
}
func main() {
	r := web.New()
	r.Group("/api", func(group *web.RouterGroup) {

		group.ALL("/service", &Demo{})
		group.GET("/xxx", func(c *web.Context) {
			panic(nil)
			io.WriteString(c.Response, "xxx"+c.Request.URL.String())
		})
		group.ALL("/test", func(c *web.Context) {
			io.WriteString(c.Response, "test"+c.Request.URL.String())
		})

		group.ALL("/test", func(c *web.Context) {
			io.WriteString(c.Response, "test"+c.Request.URL.String())
		})
		xxx(func(c *web.Context) {
			io.WriteString(c.Response, "test"+c.Request.URL.String())
		}, group)

	})
	r.Run(":8080")
}
