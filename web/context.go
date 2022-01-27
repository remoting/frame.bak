package web

import (
	"encoding/json"
	"net/http"
)

var (
	jsonContentType = []string{"application/json; charset=utf-8"}
)

type H map[string]interface{}
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Router   *Router
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

func (c *Context) Status(code int) {
	c.Response.WriteHeader(code)
}

func (c *Context) JSON(code int, data interface{}) (err error) {
	c.Status(code)
	writeContentType(c.Response, jsonContentType)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.Response.Write(jsonBytes)
	return err
}
