package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	jsonContentType = []string{"application/json; charset=utf-8"}
)

type H map[string]interface{}
type S struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
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

func (c *Context) Data(data S) (err error) {
	c.Status(200)
	writeContentType(c.Response, jsonContentType)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.Response.Write(jsonBytes)
	return err
}

func (c *Context) BindBody(obj interface{}) (err error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	return bindBody(body, obj)
}

func (c *Context) GetBody() (interface{}, error) {
	var obj interface{}
	body, err1 := ioutil.ReadAll(c.Request.Body)
	if err1 != nil {
		return nil, err1
	}
	err2 := bindBody(body, &obj)
	if err2 != nil {
		return nil, err2
	}
	return obj, nil
}
