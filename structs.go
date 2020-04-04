package gomap

import (
	"encoding/json"
)

//Object read element as struct
func (elt *Element) Object(obj interface{}) error {
	if elt.Value == nil {
		return NewWrongPathError(elt.Path)
	}
	content, err := json.Marshal(elt.Value)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, obj)

	return err
}
