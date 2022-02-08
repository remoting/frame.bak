package json

import (
	"encoding/json"
)

type JSON struct {
	value interface{}
}

func ParseString(str string) *JSON {
	obj := JSON{}
	err := json.Unmarshal([]byte(str), &obj.value)
	if err != nil {
		panic(err)
	}
	/*
		obj := &JSON{}
		err2 := bindBody([]byte(str), obj.value)
		if err2 != nil {
			panic(err2)
		}
		return &obj
	*/
	return &obj
}

func (json *JSON) IsArray() (bool, *Array) {
	if obj, ok := json.value.([]interface{}); ok {
		array := Array{
			value: obj,
		}
		return true, &array
	}
	return false, nil
}
func (json *JSON) IsObject() (bool, *Object) {
	if obj, ok := json.value.(map[string]interface{}); ok {
		object := Object{
			value: obj,
		}
		return true, &object
	}
	return false, nil
}
