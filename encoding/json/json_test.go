package json

import (
	"fmt"
	"testing"
)

const (
	jsonData1 = `
[{
"a":"1"
},{
"a":"2"
}]
`
	jsonData2 = `
 {
"a":["1",0]
} 
`
	jsonData3 = `
{
	"a":[{"s":"2"},["a",1]]
} 
`
)

func TestJSON001(t *testing.T) {
	obj1 := ParseString(jsonData3)
	if ok, array := obj1.IsArray(); ok {
		fmt.Printf("--:%v\n", array)
	}
	if ok, object := obj1.IsObject(); ok {
		fmt.Printf("==1:%v\n", object.GetArray("a").GetObject(0).GetString("s", ""))
		fmt.Printf("==2:%v\n", object.GetArray("a").GetArray(1).GetString(0, ""))
	}
}
