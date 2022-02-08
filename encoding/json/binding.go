package json

import (
	"bytes"
	"encoding/json"
	"io"
)

var EnableDecoderUseNumber = false
var EnableDecoderDisallowUnknownFields = false

func bindBody(body []byte, obj interface{}) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

func decodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	if EnableDecoderUseNumber {
		decoder.UseNumber()
	}
	if EnableDecoderDisallowUnknownFields {
		decoder.DisallowUnknownFields()
	}
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
