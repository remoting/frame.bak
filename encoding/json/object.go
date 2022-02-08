package json

type Object struct {
	value map[string]interface{}
}

func (json *Object) GetObject(name string) *Object {
	if obj, ok1 := json.value[name]; ok1 {
		if ret, ok := obj.(map[string]interface{}); ok {
			jsonObject := &Object{
				value: ret,
			}
			return jsonObject
		}
	}
	return nil
}
func (json *Object) GetArray(name string) *Array {
	if obj, ok1 := json.value[name]; ok1 {
		if ret, ok := obj.([]interface{}); ok {
			jsonArray := &Array{
				value: ret,
			}
			return jsonArray
		}
	}
	return nil
}
func (json *Object) GetString(name, def string) string {
	val, ok := json.value[name]
	if ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return def
}
func (json *Object) GetInt(name string, def int) int {
	obj, ok := json.value[name]
	if ok {
		if ret, ok := obj.(int); ok {
			return ret
		}
	}
	return def
}
func (json *Object) GetInt64(name string, def int64) int64 {
	obj, ok := json.value[name]
	if ok {
		if ret, ok := obj.(int64); ok {
			return ret
		}
	}
	return def
}
func (json *Object) Contains(name string) bool {
	_, ok := json.value[name]
	return ok
}
