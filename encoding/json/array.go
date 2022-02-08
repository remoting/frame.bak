package json

type Array struct {
	value []interface{}
}

func (json *Array) Size() int {
	return len(json.value)
}
func (json *Array) GetArray(i int) *Array {
	obj := json.value[i]
	if ret, ok := obj.([]interface{}); ok {
		jsonArray := &Array{
			value: ret,
		}
		return jsonArray
	}
	return nil
}

func (json *Array) GetObject(i int) *Object {
	obj := json.value[i]
	if ret, ok := obj.(map[string]interface{}); ok {
		jsonObj := &Object{
			value: ret,
		}
		return jsonObj
	}
	return nil
}
func (json *Array) GetString(i int, def string) string {
	obj := json.value[i]
	if ret, ok := obj.(string); ok {
		return ret
	}
	return def
}
func (json *Array) GetInt(i int, def int) int {
	obj := json.value[i]
	if ret, ok := obj.(int); ok {
		return ret
	}
	return def
}
func (json *Array) GetInt64(i int, def int64) int64 {
	obj := json.value[i]
	if ret, ok := obj.(int64); ok {
		return ret
	}
	return def
}
